package repositories

import "errors"

var ErrorOrderNotFound = errors.New("Order not found with the given id.")

var ErrorOrdersNotFound = errors.New("Orders not found with the given filters")

var ErrorOrderStatusAlreadyUpdated = errors.New("Order status already updated.")

var ErrorOrderByInvalid = errors.New("malformed order direction, should be asc or desc")

var ErrorInvalidSortBy = errors.New("unknown field in sortBy query parameter")
