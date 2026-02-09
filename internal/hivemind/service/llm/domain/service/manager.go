package service

import (
	"context"

	entity2 "github.com/kiosk404/eidolon/internal/hivemind/service/llm/domain/entity"
)

type ModelManager interface {
	CreateLLMModel(ctx context.Context, modelClass entity2.ModelClass, modelShowName string, conn *entity2.Connection, extra *entity2.ModelExtra) (int64, error)
	GetModelByID(ctx context.Context, id int64) (*entity2.ModelInstance, error)
	GetDefaultModel(ctx context.Context) (*entity2.ModelInstance, error)
	SetDefaultModel(ctx context.Context, id int64) error
	ListModelByType(ctx context.Context, modelType entity2.ModelType, limit int) ([]*entity2.ModelInstance, error)
	ListAllModelList(ctx context.Context) ([]*entity2.ModelInstance, error)
}
