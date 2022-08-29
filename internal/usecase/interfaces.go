// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"file-service/internal/entity"
)

type (
	// IFileStore -.
	IFileStore interface {
		SaveFile(ctx context.Context, file entity.FileEntity) (entity.FileEntity, error)
		GetFileById(ctx context.Context, id int) (entity.FileEntity, error)
	}

	// IFileStoreRepo -.
	IFileStoreRepo interface {
		SaveFileEntity(ctx context.Context, file entity.FileEntity) error
		GetFileEntityById(ctx context.Context, id int) (entity.FileEntity, error)
	}
)
