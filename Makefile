DOCKER_HUB_ID?=ossifrage
NAME?=hellologdriver
DOCKER_PLUGIN_NAME?=${DOCKER_HUB_ID}/${NAME}
DOCKER_PLUGIN_VERSION?=0.0.1
DOCKER_PLUGIN=${DOCKER_PLUGIN_NAME}:${DOCKER_PLUGIN_VERSION}
CONTAINER_NAME=${NAME}_container

all: build install

.PHONY: build
build: clean
	docker build --target builder -t rootfsimage-build .
	docker build --target final -t rootfsimage .
	mkdir rootfs
	docker create --name ${CONTAINER_NAME} rootfsimage true
	docker export ${CONTAINER_NAME} | tar -x -C rootfs; \
	docker rm -vf ${CONTAINER_NAME}
	docker rmi rootfsimage
	docker rmi rootfsimage-build

.PHONY: install
install:
	docker plugin create ${DOCKER_PLUGIN} .
	docker plugin enable ${DOCKER_PLUGIN}

.PHONY: clean
clean:
	rm -rf rootfs
	-docker plugin disable ${DOCKER_PLUGIN}
	-docker plugin rm ${DOCKER_PLUGIN}
