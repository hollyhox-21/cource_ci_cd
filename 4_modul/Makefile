CONTAINER_NAME=my_app
APP_VERSION=1.0

build:
# сборка базового образа из Dockerfile
	docker build -t $(CONTAINER_NAME):$(APP_VERSION) .
run:
# запуск контейнера в фоновом режиме с именем с удалением после остановки и подключением к порту 8080
	docker run \
	-d \
	--rm \
	--name $(CONTAINER_NAME) \
	-p 8080:8080 \
	$(CONTAINER_NAME):$(APP_VERSION)
stop:
# остановка контейнера
	docker stop $(CONTAINER_NAME)
