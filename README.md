# New Go Service Template

A template for creating new Go services.

#### Table of contents

* [Introduction](#introduction)
    * [Project structure](#project-structure)
* [Development](#development)
    * [1. Prerequisites](#1-prerequisites)
    * [2. Setup](#2-setup)
    * [3. Running locally](#3-running-locally)
* [Testing](#testing)
    * [1. Running tests](#1-running-tests)
* [Deployment](#deployment)
    * [1. Building the image](#1-building-the-image)
    * [2. Running the image](#2-running-the-image)

## Introduction

### Project structure

Below is a quick outline of the structure of the app:

```text
.
├── cmd                        # Go apps.
|   ├── server
|   │   └── main.go            # Entry point for the app.
│   └── ...
└── pkg                        # A collection of reusable modules.
    ├── application
    |   └── application.go     # Sets up the configuration for the app.
    └── cleanup
    |   └── cleanup.go         # Gracefully cleans up the app.
    ├── config
    |   └── config.go          # Struct that defines the application config.
    └── error
    |   └── code.go            # All the error codes and the descriptions.
    ├── middleware             # Custom middleware.
    |   ├── somemiddleware.go
    |   └── ...
    ├── handler                # Handles HTTP requests and responses.
    |   ├── somehandler.go
    |   └── ...
    └── router
    |   └── router.go          # Creates the router and sets up the handlers.
    └── server
    |   ├── request.go         # Request structs.
    |   ├── response.go        # Response structs.
    |   └── server.go          # Starts the server.
    └── ...
```

## Development

### 1. Prerequisites

* Install [Go 1.14+](https://golang.org/dl/).
* Install [Docker](https://docs.docker.com/get-docker/).

### 2. Setup

1. Install the dependencies:
```bash
go get
```

### 3. Running locally

1. Simply run:
```bash
./scripts/start.sh
```

2. You can check the API is running using the following cURL command:
```shell script
curl -X GET http://localhost:${PORT}/healthcheck
```

## Testing

### 1. Running tests

1. Simply run the command:
```bash
./scripts/test.sh
```

## Deployment

### 1. Building the image

When building the Docker image, we want to inject env vars at build time, as the [`Dockerfile`](./Dockerfile) injects the build args as env vars into the container.
```bash
docker build \
-t kieranroneill/new_go_service_template \
--build-arg env=production \
--build-arg port=3000 \
--build-arg service_name=new-go-service-template \
--build-arg version=$(<VERSION) \
.
```

#### 2. Running the image

```bash
docker run \
--name new_go_service_template \
-it \
-p 1337:${PORT} \
kieranroneill/new_go_service_template:latest
```
