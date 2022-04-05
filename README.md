**Table of Contents**

- [Golang REST API](#golang-rest-api)
  - [About the project](#about-the-project)
    - [Status](#status)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Golang REST API

## About the project

This project is a simple rest API used for knowledge purposes.
Some packages used in this project:

- [Mux](https://github.com/gorilla/mux) - HTTP router and URL matcher for Go web servers
- [Gorm](https://github.com/go-gorm/gorm) - ORM library for Go

### Status

The project is complete, I don't want to modify anything, it's just a way to understand more about Go, Gorilla/Mux, and Gorm.

## Getting started

> First you **MUST** have [Go](https://go.dev/doc/install), [Docker](https://docs.docker.com/engine/install/), [Docker Compose](https://docs.docker.com/compose/install/) and [Git](https://git-scm.com/downloads) installed

1. Clone the repository:

```sh
git clone git@github.com:BikutaDesu/go-rest-api.git
```

2. Open the project:

```sh
cd go-rest-api
```

3. Start the containers:

```sh
docker-compose up
```

4. Run the project:

```sh
go run main.go
```

### Layout

```tree
├── .gitignore
├── docker-compose.yml
├── README.md
├── main.go
├── controllers
│   └── controllers.go
├── database
│   └── db.go
├── middleware
│   └── middleware.go
├── migration
│   └── docker-database-initial.sql
├── models
│   └── people.go
├── routes
│   └── routes.go
```

A brief description of the layout:

- `.gitignore` specifies intentionally untracked files that Git should ignore.
- `docker-compose.yml` file defining services, networks, and volumes for a Docker application.
- `README.md` description of the project.
- `main.go` main file.
- `controllers` api controllers.
- `database` location responsible for database services.
- `middleware` api middlewares.
- `migration` database migrations.
- `models` api models.
- `routes` api routes.

## Notes

- Before running the project you **MUST** build de containers in `docker-compose.yml` file.
