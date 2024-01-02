# Golang Docker Demo

This repo outlines a basic Docker setup for a go app that can be used for development and built and deployed to production.

## Start dev
```BASH
docker-compose up
```
This will build a new `dev` container and start the watcher, when a new change is saved the app is re-built and loaded

## Build for production
These instructions outline how to build your final application for producting into a single small container image than can be deployed anywhere that will run docker containers

### Build production images
```BASH
docker build -t [image-repo]:[tag] -f Dockerfile --target prod .
```
This will build a new image tagged with the values placed ni the placeholders for `[image-repo]:[tag]` e.g: `myapp:latest`

### Run production image
```BASH
docker run --rm -it -p 8080:8080 [image-repo]:[tag]
```

### Docker Compose
To use your production image with docker compose you will need to put it into a registry like docker hub and then you will be able to start it as follows:

`docker-compose.yml`
```YAML
version: '3.9'

networks:
  golang:

services:
  app:
    image: [image-repo]:[tag]
    ports:
      - 8080:8080
    networks:
      - golang
```