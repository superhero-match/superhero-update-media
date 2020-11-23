prepare:
	go mod download

run:
	go build -o bin/main cmd/media/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/chat/main.go
	chmod +x bin/main

dkb:
	docker build -t superhero-update-media .

dkr:
	docker run --rm -p "7100:7100" -p "8260:8260" superhero-update-media

launch: dkb dkr

update-media-log:
	docker logs superhero-update-media -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

update-media-ssh:
	docker exec -it superhero-update-media /bin/bash

PHONY: prepare build dkb dkr launch update-media-log update-media-ssh rmc rmi clear