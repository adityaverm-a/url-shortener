# URL Shortener Microservice in Golang

## Introduction
This is a URL Shortener microservice implemented in Golang. It exposes a set of RESTful APIs to shorten long URLs, resolve short URLs to their original form, and manage custom short URLs with optional Time-to-Live (TTL) support.

The microservice uses `sync.Map` for thread-safe, in-memory storage and offers various features to ensure reliability, scalability, and ease of use.

## Table of Contents
- [URL Shortener Microservice in Golang](#url-shortener-microservice-in-golang)
  - [Introduction](#introduction)
  - [Table of Contents](#table-of-contents)
  - [Features Implemented](#features-implemented)
  - [Setting up Development Environment](#setting-up-development-environment)
    - [Configuring env values](#configuring-env-values)
    - [Running the Server](#running-the-server)
    - [Postman Collection](#postman-collection)
    - [Launching the Debugger](#launching-the-debugger)
    - [Unit Testing Guidelines](#unit-testing-guidelines)

## Features Implemented

### 1. **URL Shortening**
   - The service allows users to shorten long URLs.
   - If the long URL has been shortened before, it returns the same short URL.

### 2. **Custom Short URLs**
   - Users can optionally provide a custom short URL.
   - The system ensures the custom short URL is unique and not already in use.
   - If the custom URL already exists, the service returns an error: `"custom short URL already exists"`.

### 3. **Time-to-Live (TTL) for Short URLs**
   - Shortened URLs can have an expiration time defined in seconds (`TTL`).
   - Once the TTL expires, the short URL becomes invalid and is no longer accessible.
   - The service handles the expiration check before redirecting to the original URL.

### 4. **Concurrency Safe with `sync.Map`**
   - The service uses `sync.Map` for thread-safe, in-memory storage of short URLs.
   - `sync.Map` ensures that multiple requests can be handled concurrently without race conditions.

### 5. **Error Handling**
   - Proper error messages are returned for invalid inputs, expired URLs, and duplicate short URLs.
   - Detailed error logs are provided for repository failures and other internal errors.

### 6. **URL Redirection**
   - When accessing a short URL, the service redirects to the original URL.
   - The redirection is permanent (`301 Moved Permanently`), ensuring a smooth user experience.

### 7. **Unit Tests**
   - The microservice includes unit tests for key functionality like shortening, resolving URLs, and validating input data.

---

## Setting up Development Environment

### <a name="configure-env">Configuring env values</a>
Yaml syntax is used for injecting config values into the application environment. A valid YAML config snapshot is given below. The file name should be `[env].config.yaml`.

    `local.config.yaml`, `dev.config.yaml` are valid file names.

This file should be placed in the root of the source code.

```yaml
port: <:port>
#port: :5000
charset: <your_char_set>
#charset: abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
shortURLLength: <short_url_length>
#shortURLLength: 6

```

### <a name="running-the-server">Running the Server <a/>
Use below command to start your server

    go run cmd/main.go

Once the server is up & runnning, use the provided Postman Collection!

### <a name="running-the-server">Postman Collection <a/> 
I have attached the Postman collection for testing all the API endpoints. You can find the collection on my [GitHub repository](https://github.com/adityaverm-a/url-shortener) to quickly test the functionality of the service.

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
