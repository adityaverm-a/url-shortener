package services

import "errors"

var ErrorURLAlreadyExists = errors.New("custom short URL already exists")

var ErrorURLNotFound = errors.New("the requested short URL does not exist")

var ErrorURLExpired = errors.New("short URL has expired")
