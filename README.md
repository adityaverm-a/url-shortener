# Order Management Microservice in Golang

## Introduction
This webservice has been implemented in golang and the set of exposed APIs are RESTful.

## Table of Content
- [Order Management Microservice in Golang](#order-management-microservice-in-golang)
  - [Introduction](#introduction)
  - [Table of Content](#table-of-content)
  - [Setting up Development Environment](#setting-up-development-environment)
    - [Configuring env values ](#configuring-env-values-)
    - [Running the Server ](#running-the-server-)
    - [Migrations ](#migrations-)
    - [Launching the Debugger ](#launching-the-debugger-)
    - [Unit Testing Guidelines ](#unit-testing-guidelines-)
        - [To run unit tests recursively, run -\> go test ./...](#to-run-unit-tests-recursively-run---go-test-)

## Setting up Development Environment

### <a name="configure-env">Configuring env values <a/>

Yaml syntax is used for injecting config values in application environment. A valid yaml config snapshot is given below. File name should be `[env].config.yaml`.

    `local.config.yaml`, `dev.config.yaml` are valid file names.

This file should on the root of source code.

    sql: &sqlConfig
    driverName: "mysql"
    dataSourceNameFormat: "%s:%s@tcp(%s)/%s"
    connConfig :
          user: <username>
          password: <password>
          host: <host>
          db: <db_name>
    port: <:port>
    # port: :5000

### <a name="running-the-server">Running the Server <a/>
Use below command to start your server

    go run cmd/main.go

Once the server is up & runnning, use the provided Postman Collection!

### <a name="starting-debugger">Migrations <a/>
Run Database Migrations

* For running up :

    `go run ./migrations/up/main.go`

* For running down :

    `go run ./migrations/down/main.go`

* For initializing tables with some data run :

    `go run ./migrations/init/main.go`

### <a name="starting-debugger">Launching the Debugger <a/>

Create a file `launch.json` in `.vscode` directory ( .vscode directory should be on the root) with the following content.

    {
        "version": "1.0.0",
        "configurations": [
            {
                "name": "Launch Package",
                "type": "go",
                "request": "launch",
                "mode": "debug",
                "program": "${workspaceFolder}/cmd",
                "cwd": "${workspaceFolder}"
            }
        ]
    }

Click on the play button, debugging server will start

### <a name="unit-testing-guidelines">Unit Testing Guidelines <a/>
##### To run unit tests recursively, run -> go test ./...
