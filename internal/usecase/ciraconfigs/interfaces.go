package ciraconfigs

import (
	"context"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
)

type (
	Repository interface {
		GetCount(ctx context.Context, tenantID string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.CIRAConfig, error)
		GetByName(ctx context.Context, configName, tenantID string) (*entity.CIRAConfig, error)
		Delete(ctx context.Context, profileName, tenantID string) (bool, error)
		Update(ctx context.Context, p *entity.CIRAConfig) (bool, error)
		Insert(ctx context.Context, p *entity.CIRAConfig) (string, error)
	}
	Feature interface {
		GetCount(ctx context.Context, tenantID string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.CIRAConfig, error)
		GetByName(ctx context.Context, configName, tenantID string) (*entity.CIRAConfig, error)
		Delete(ctx context.Context, profileName, tenantID string) error
		Update(ctx context.Context, p *entity.CIRAConfig) (*entity.CIRAConfig, error)
		Insert(ctx context.Context, p *entity.CIRAConfig) (*entity.CIRAConfig, error)
	}
)
