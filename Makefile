SHELL := /bin/bash

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  help            to show this help message"
	@echo "  tests           to run tests in a consistent environment with all required dependencies"
	@echo "  start           to start the application in a consistent environment with all required dependencies"

tests:
	docker-compose --profile tests up --abort-on-container-exit

# Worth to mention that when this is started there is no need to start the emulators and also
# when editing the code the changes are reflected in the container without the need to rebuild it.
# Seems that the local app-engine is behaving like hot-reload.
start:
	docker-compose --profile appengine up

.PHONY: help tests start