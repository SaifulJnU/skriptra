"""OCR sidecar.

The one place in Skriptra where Python earns its keep. Everything else, the
API, retrieval, embeddings, generation, PDF text extraction, is Go, because it
is HTTP, SQL or a Go library. Reading text off a photograph is not: Tesseract
and the modern OCR stack have no Go equivalent worth using.

It registers as one more DocumentParser behind the Chain in
backend/internal/ingest. The Go side declares what a document needs and picks
the cheapest parser that can handle it, so a digital PDF never reaches this
service and a photograph always does.

HTTP rather than gRPC. The payload is one file and one JSON reply, the service
is called once per document rather than per request, and a plain endpoint keeps
the sidecar something anyone can run and curl.
"""

import io
import logging
import os

import pytesseract
from fastapi import FastAPI, File, Form, HTTPException, UploadFile
from pdf2image import convert_from_bytes
from PIL import Image, ImageOps

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")
log = logging.getLogger("ocr")

app = FastAPI(title="Skriptra OCR", version="1.0.0")

# German and English together: a Dortmund exam paper routinely mixes both, and
# running one language over the other produces confident nonsense rather than a
# visible failure.
DEFAULT_LANGS = os.getenv("OCR_LANGS", "eng+deu")

# Above this, rendering every page at 300 DPI would exhaust memory before it
# produced anything useful. A textbook is not what this service is for.
MAX_PAGES = int(os.getenv("OCR_MAX_PAGES", "40"))

# Tesseract's default page segmentation assumes a single uniform block. Exam
# papers are multi-column and heavily indented, so 3 (fully automatic with
# layout analysis) recovers question numbering that the default flattens.
TESS_CONFIG = os.getenv("OCR_TESSERACT_CONFIG", "--oem 1 --psm 3")


@app.get("/healthz")
def healthz():
    try:
        version = str(pytesseract.get_tesseract_version())
    except Exception as exc:  # tesseract missing from the image
        raise HTTPException(status_code=503, detail=f"tesseract unavailable: {exc}")
    return {"status": "ok", "tesseract": version, "langs": DEFAULT_LANGS}


def _prepare(image: Image.Image) -> Image.Image:
    """Normalise an image before OCR.

    Grayscale and autocontrast, not thresholding. A phone photo of an exam paper
    has uneven lighting, and a global threshold turns the shadowed half of the
    page into a solid block. Tesseract does its own binarisation and does it
    better with the greyscale it was given.
    """
    image = image.convert("L")
    image = ImageOps.autocontrast(image)
    # EXIF orientation: a photo taken in portrait is often stored rotated, and
    # OCR on a sideways page returns nothing at all.
    return ImageOps.exif_transpose(image)


def _ocr_page(image: Image.Image, langs: str) -> str:
    return pytesseract.image_to_string(_prepare(image), lang=langs, config=TESS_CONFIG)


@app.post("/ocr")
async def ocr(file: UploadFile = File(...), langs: str = Form(DEFAULT_LANGS)):
    """Extract text from a photograph or a scanned PDF.

    Returns one entry per page so the Go side keeps page numbers, which every
    citation depends on.
    """
    content = await file.read()
    if not content:
        raise HTTPException(status_code=422, detail="empty file")

    name = (file.filename or "").lower()
    is_pdf = content[:4] == b"%PDF" or name.endswith(".pdf")

    try:
        if is_pdf:
            # 300 DPI is the point where Tesseract's accuracy stops improving
            # for printed text; higher just costs memory.
            images = convert_from_bytes(content, dpi=300)
            if len(images) > MAX_PAGES:
                raise HTTPException(
                    status_code=422,
                    detail=f"{len(images)} pages exceeds the {MAX_PAGES} page limit",
                )
        else:
            images = [Image.open(io.BytesIO(content))]
    except HTTPException:
        raise
    except Exception as exc:
        log.exception("could not decode %s", file.filename)
        raise HTTPException(status_code=422, detail=f"could not read file: {exc}")

    pages = []
    for i, image in enumerate(images, start=1):
        try:
            text = _ocr_page(image, langs)
        except Exception as exc:
            # One unreadable page must not lose the other twenty. Record it
            # empty; the Go side decides whether the document as a whole failed.
            log.warning("page %d failed: %s", i, exc)
            text = ""
        pages.append({"number": i, "text": text, "width": image.width, "height": image.height})

    extracted = sum(len(p["text"].strip()) for p in pages)
    log.info("ocr %s: %d pages, %d chars", file.filename, len(pages), extracted)

    return {
        "pages": pages,
        "pageCount": len(pages),
        "parsedBy": "tesseract",
        # Reported rather than judged. The Go side owns the decision about
        # whether an extraction is good enough to index.
        "charactersExtracted": extracted,
    }
