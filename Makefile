prepare:
	go mod download

run:
	go build -o bin/main cmd/media/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/chat/main.go
	chmod +x bin/main

dkb:
	docker build -t superhero-register-media .

dkr:
	docker run --rm -p "7000:7000" -p "8230:8230" superhero-register-media

launch: dkb dkr

register-media-log:
	docker logs superhero-register-media -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

register-media-ssh:
	docker exec -it superhero-register-media /bin/bash

PHONY: prepare build dkb dkr launch register-media-log register-media-ssh rmc rmi clear