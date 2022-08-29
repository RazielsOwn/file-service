package amqprpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"

	"file-service/internal/entity"
	"file-service/internal/usecase"
	"file-service/pkg/rabbitmq/rmq_rpc/server"
)

type fileServiceRoutes struct {
	fileStore usecase.IFileStore
}

func newFileServiceRoutes(routes map[string]server.CallHandler, t usecase.IFileStore) {
	r := &fileServiceRoutes{t}
	{
		routes["downloadFile"] = r.downloadFile()
	}
}

type fileResponse struct {
	File entity.FileEntity `json:"file"`
}

func (r *fileServiceRoutes) downloadFile() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		request := struct {
			Id int `json:"id"`
		}{}

		err := json.Unmarshal(d.Body, &request)
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - fileServiceRoutes - downloadFile - json.Marshal : %w", err)
		}

		var id = request.Id
		file, err := r.fileStore.GetFileById(context.Background(), id)
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - fileServiceRoutes - downloadFile : %w", err)
		}

		response := fileResponse{file}
		response.File.Path = fmt.Sprintf("/files/%v", file.Id)

		return response, nil
	}
}
