# Microservice Clean Architecture

- Clean Architecture
- Microservice Architecture

## Additional resources

- Youtube channels `The Art of Development`, `Maksim Zhashkevych`
- Robert C. Martin `Clean Architecture` (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- RabbitMQ guides
- Minio guides

Service accepts the incoming picture in .jpeg or .png format and then sends it to the channel to optimize it and save it in the storage in formats 25, 50, 75 from 100 quality including original quality.
Service also have a API for download an optimized image from storage by image id and image quality.

> The service must be responsive and respond quickly to and accept incoming requests.
> Each request should be answered depending on the result, including http status.

## Tech

#### 1. Producer microservice

This service accepts incoming requests from the client
and sends them to the photo consumer.


- `Julien Schmidt HTTP` - http rest handler!
- `RabbitMQ` - message broker to queue with consumer image service
- `MINIO` - storage to download/upload file
- `Cleanenv` - greate env framework for environments

POSTMAN link: https://www.getpostman.com/collections/45677b9e5b12ef5ddc08

#### 2. Consumer microservice

This service accepts incoming images from the producer service
optimizing and save into a storage.

- `RabbitMQ` - message broker to queue with producer service!
- `MINIO` - storage to download/upload file
- `Cleanenv` - greate env framework for environments


## Local Run

Configure a storage and message broker in docker container.

```sh
docker compose up rabbitmq minio
or
make local
```
Next step a run producer and consumer..

PRODUCER:
```sh
cd producer
go mod tidy
go run cmd/app.go
```

CONSUMER:
```sh
cd consumer
go mod tidy
go run cmd/app.go
```


## Docker

Run all services and dependencies together

```sh
make up
or
docker-compose down && docker-compose up --build
```