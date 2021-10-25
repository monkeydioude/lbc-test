.PHONY: all

IMAGE_NAME = lbc-test
CONTAINER_NAME = lbc-test_app
# test_container_name = lbc-test_test

all:
	IMAGE_NAME=$(image_name) CONTAINER_NAME=$(CONTAINER_NAME) ./script/docker-run.sh

# build:
# 	docker build -t $(image_name) -f ./build/Dockerfile .

# stop:
# 	docker stop $(container_name)

# rm: stop
# 	docker rm $(container_name)

# re: rm run

# reall: rm build run

# test:
# 	docker build --target build -t build-lbc-test -f ./build/Dockerfile .
# 	docker run --name $(test_container_name) build-lbc-test go test -v ./...