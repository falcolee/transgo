build:
	sh ./build.sh
docker:
	docker build -t xiaogouxo/transgo:$(version) -t xiaogouxo/transgo:latest --build-arg version=$(version) .