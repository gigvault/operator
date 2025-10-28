FROM golang:1.23-bullseye AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/operator ./cmd/operator

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /out/operator .
USER 65532:65532
ENTRYPOINT ["/operator"]

