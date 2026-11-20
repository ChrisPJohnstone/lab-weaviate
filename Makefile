CONTAINER_ENGINE ?= podman

.PHONY: deploy generator destroy

deploy:
	$(CONTAINER_ENGINE) compose up -d

generator:
	make -C generator

destroy:
	make -C generator clean
	$(CONTAINER_ENGINE) compose down -v
