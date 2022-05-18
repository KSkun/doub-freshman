# doub-freshman

Unique Hackday Project. The web backend of a game named 胡闹大学, presenting the daily lives of college students.

## Features

This is a HTTP backend server for webpage game. Functions are provided by RESTful HTTP APIs. See [API document](doc/api.md).

## Build & Run

The project is configured for go mod & Docker. You also need [MongoDB](https://www.mongodb.com/) and [Redis](https://redis.io/) for its databases.

### go mod

- `cd src`
- `go build`
- Get the binary file in current directory.

### Docker

- `sudo docker-compose build`
- `sudo docker-compose up -d`
- Check the container info by `sudo docker ps | grep doub`.
