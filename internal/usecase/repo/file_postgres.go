package repo

import (
	"context"
	"fmt"

	"file-service/internal/entity"
	"file-service/pkg/postgres"

	"github.com/jackc/pgx/v4"
)

// FilesRepo -.
type FilesRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *FilesRepo {
	return &FilesRepo{pg}
}

// GetFileById -.
func (r *FilesRepo) GetFileEntityById(ctx context.Context, id int) (entity.FileEntity, error) {
	query, _, err := r.Builder.
		Select("id, name, description, path").
		From("files").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return entity.FileEntity{}, fmt.Errorf("FilesRepo - GetFileById - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, query, id)

	e := entity.FileEntity{}

	err = row.Scan(&e.Id, &e.Name, &e.Description, &e.Path)
	if err == pgx.ErrNoRows {
		return e, nil
	} else if err != nil {
		return entity.FileEntity{}, fmt.Errorf("FilesRepo - GetFileById - rows.Scan: %w", err)
	}

	return e, nil
}

// SaveFile -.
func (r *FilesRepo) SaveFileEntity(ctx context.Context, t entity.FileEntity) error {
	query, args, err := r.Builder.
		Insert("files").
		Columns("name, description, path").
		Values(t.Name, t.Description, t.Path).
		ToSql()
	if err != nil {
		return fmt.Errorf("FilesRepo - SaveFile - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("FilesRepo - SaveFile - r.Pool.Exec: %w", err)
	}

	return nil
}
