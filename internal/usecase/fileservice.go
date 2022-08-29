package usecase

import (
	"context"
	"fmt"

	"file-service/internal/entity"
)

// FileUseCase -.
type FileUseCase struct {
	repo IFileStoreRepo
}

// New -.
func New(r IFileStoreRepo) *FileUseCase {
	return &FileUseCase{
		repo: r,
	}
}

// Files - getting file info from store.
func (uc *FileUseCase) GetFileById(ctx context.Context, id int) (entity.FileEntity, error) {
	fileFound, err := uc.repo.GetFileEntityById(ctx, id)
	if err != nil {
		return entity.FileEntity{}, fmt.Errorf("FileServiceUseCase - GetFileById - s.repo.GetFileById: %w", err)
	}

	return fileFound, nil
}

// Add file -.
func (uc *FileUseCase) SaveFile(ctx context.Context, t entity.FileEntity) (entity.FileEntity, error) {
	err := uc.repo.SaveFileEntity(context.Background(), t)
	if err != nil {
		return entity.FileEntity{}, fmt.Errorf("FileServiceUseCase - SaveFile - s.repo.SaveFile: %w", err)
	}

	return t, nil
}
