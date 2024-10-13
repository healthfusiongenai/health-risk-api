VERSION ?= $(shell cat VERSION)
APP_NAME ?= energy-usage-simulator

ifeq ($(shell command -v podman 2> /dev/null),)
    CMD=docker
else
    CMD=podman
endif

CMD=docker

# Build the docker image
build:
#	go build -o /energy-usage-simulator .
	$(CMD) build \
	 -t $(APP_NAME):$(VERSION) \
	 --build-arg APP_NAME=$(APP_NAME) \
	 .

# Run the docker container
run:
	$(CMD) run -d -p 8000:8000 --name $(APP_NAME) $(APP_NAME):$(VERSION)

exec:
	$(CMD) exec -it $(APP_NAME) /bin/sh

# Push the docker image to the registry
push:
	$(CMD) push $(APP_NAME):$(VERSION)

clean:
	$(CMD) system prune -a --filter label=org.opencontainers.image.title="$(APP_NAME)" -f
#	docker rm -f $(APP_NAME)
#	docker rmi -f $(APP_NAME):$(VERSION)


