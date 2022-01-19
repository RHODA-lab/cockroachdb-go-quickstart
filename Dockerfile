# Build the sample app
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

#copy all go files
COPY cmd/ cmd/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o crdb-go-quickstart ./cmd/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /

COPY public/ public/
COPY --from=builder /workspace/crdb-go-quickstart .
USER 65532:65532

EXPOSE 8080
ENTRYPOINT ["/crdb-go-quickstart"]
