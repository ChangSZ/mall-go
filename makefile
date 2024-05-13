IMAGE_VERSION=v1.0

docker-build-all: docker-build-admin docker-build-portal

docker-build-admin:
	docker build -f deploy/docker/Dockerfile_admin -t mall-go/admin:$(IMAGE_VERSION) .

docker-build-portal:
	docker build -f deploy/docker/Dockerfile_portal -t mall-go/portal:$(IMAGE_VERSION) .
