CONTAINER_ENGINE ?= podman

.PHONY: all deploy generator destroy

all: deploy generator

deploy:
	$(CONTAINER_ENGINE) compose up -d

generator:
	make -C generator

destroy:
	make -C generator clean
	$(CONTAINER_ENGINE) compose down -v
