CONTAINER_ENGINE ?= podman

.PHONY: deploy destroy

deploy:
	$(CONTAINER_ENGINE) compose up -d

destroy:
	$(CONTAINER_ENGINE) compose down -v
