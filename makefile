build:
	sh ./build.sh
docker:
	docker build -t transgo/transgo:$(version) -t transgo/transgo:latest --build-arg version=$(version) .