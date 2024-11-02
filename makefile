IMAGE_NAME =? <USERNAME>/<IMAGE-NAME>:<IMAGE-TAG>
docker-build:
	docker build -t ${IMAGE_NAME} -f .\build\Dockerfile .

docker-push:
	docker push -t ${IMAGE_NAME} -f .\build\Dockerfile .