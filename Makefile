SHELL := /bin/bash

execute-tests:
	docker-compose --profile tests up --abort-on-container-exit
	docker-compose --profile tests down

# Worth to mention that when this is started there is no need to start the emulators and also
# when editing the code the changes are reflected in the container without the need to rebuild it.
# Seems that the local app-engine is behaving like hot-reload.
start-local-emulators:
	docker-compose --profile appengine up --abort-on-container-exit

stop-local-emulators:
	docker-compose --profile appengine down

start-datastore-emulator:
	docker-compose --profile datastore-emulator up -d

stop-datastore-emulator:
	docker-compose --profile datastore-emulator down

