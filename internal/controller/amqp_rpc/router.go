package amqprpc

import (
	"file-service/internal/usecase"
	"file-service/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.IFileStore) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newFileServiceRoutes(routes, t)
	}

	return routes
}
