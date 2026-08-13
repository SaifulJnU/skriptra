# Multi-stage, with a dev target used by docker-compose and a prod target that
# ships a static binary.
#
# The dev target deliberately keeps the toolchain in the image and runs from
# mounted source, so a code change is a restart rather than a rebuild. The prod
# target carries no toolchain, no shell and no package manager, which is both
# smaller and a smaller attack surface.

# ---------------------------------------------------------------- base -----
FROM golang:1.26-alpine AS base

WORKDIR /app

# Dependencies are copied and downloaded before the source, so editing code does
# not invalidate the module cache layer.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ----------------------------------------------------------------- dev -----
FROM base AS dev

# Source is bind-mounted over /app by compose; this is only the default command.
EXPOSE 8080
CMD ["go", "run", "./cmd/api"]

# --------------------------------------------------------------- build -----
FROM base AS build

# CGO_ENABLED=0 is what makes the binary static and lets the final stage be
# scratch-like. It is also why PDF extraction is pure Go: a cgo dependency here
# would force a distro base image and undo this.
ENV CGO_ENABLED=0 GOOS=linux

ARG VERSION=dev
RUN go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o /out/api    ./cmd/api  && \
    go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o /out/worker ./cmd/worker

# ------------------------------------------------------------------ api -----
FROM gcr.io/distroless/static-debian12:nonroot AS api

COPY --from=build /out/api /api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/api"]

# --------------------------------------------------------------- worker -----
FROM gcr.io/distroless/static-debian12:nonroot AS worker

COPY --from=build /out/worker /worker
USER nonroot:nonroot
ENTRYPOINT ["/worker"]
