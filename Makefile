#absolute path is expected e.g.: /somefolder/
SERVICE_BINDING_ROOT ?= bindings

IMAGE_TAG ?= quay.io/myeung/crdb-go-quickstart:v0.0.1

.PHONY: binary
binary:
	go build -o crdb-go-quickstart ./cmd/.

.PHONY: run-binary
run-binary: binary
	./crdb-go-quickstart

.PHONY: docker-build
docker-build:
	docker build -t ${IMAGE_TAG} .

.PHONY: docker-push
docker-push:
	docker push ${IMAGE_TAG}
