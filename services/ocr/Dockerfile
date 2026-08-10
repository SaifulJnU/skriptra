# Tesseract and poppler are system packages, which is the entire reason this
# service exists as a separate container rather than as more Go.
FROM python:3.12-slim

# tesseract-ocr-deu alongside eng: a German university corpus mixes both, often
# within one paper.
# poppler-utils provides pdftoppm, which pdf2image shells out to.
RUN apt-get update && apt-get install -y --no-install-recommends \
        tesseract-ocr \
        tesseract-ocr-eng \
        tesseract-ocr-deu \
        poppler-utils \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY main.py .

EXPOSE 50052

# One worker. OCR is CPU-bound and already uses what it is given; more workers
# on the same cores would just add contention and memory.
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "50052", "--workers", "1"]
