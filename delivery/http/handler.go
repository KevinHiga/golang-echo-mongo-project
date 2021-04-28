package http

import (
	"echo-mongo-project2/config/dbiface"
)

type Handler struct {
	Col dbiface.CollectionAPI
}
