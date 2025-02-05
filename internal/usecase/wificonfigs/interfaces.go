package wificonfigs

import (
	"context"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
)

type (
	Repository interface {
		CheckProfileExists(ctx context.Context, profileName, tenantID string) (bool, error)
		GetCount(ctx context.Context, tenantID string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.WirelessConfig, error)
		GetByName(ctx context.Context, guid, tenantID string) (*entity.WirelessConfig, error)
		Delete(ctx context.Context, profileName, tenantID string) (bool, error)
		Update(ctx context.Context, p *entity.WirelessConfig) (bool, error)
		Insert(ctx context.Context, p *entity.WirelessConfig) (string, error)
	}

	Feature interface {
		CheckProfileExists(ctx context.Context, profileName, tenantID string) (bool, error)
		GetCount(ctx context.Context, tenantID string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.WirelessConfig, error)
		GetByName(ctx context.Context, guid, tenantID string) (*entity.WirelessConfig, error)
		Delete(ctx context.Context, profileName, tenantID string) error
		Update(ctx context.Context, p *entity.WirelessConfig) (*entity.WirelessConfig, error)
		Insert(ctx context.Context, p *entity.WirelessConfig) (*entity.WirelessConfig, error)
	}
)
