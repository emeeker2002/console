package domains

import (
	"context"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
)

type (
	Repository interface {
		GetCount(context.Context, string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.Domain, error)
		GetDomainByDomainSuffix(ctx context.Context, domainSuffix, tenantID string) (*entity.Domain, error)
		GetByName(ctx context.Context, name, tenantID string) (*entity.Domain, error)
		Delete(ctx context.Context, name, tenantID string) (bool, error)
		Update(ctx context.Context, d *entity.Domain) (bool, error)
		Insert(ctx context.Context, d *entity.Domain) (string, error)
	}
	Feature interface {
		GetCount(context.Context, string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.Domain, error)
		GetDomainByDomainSuffix(ctx context.Context, domainSuffix, tenantID string) (*entity.Domain, error)
		GetByName(ctx context.Context, name, tenantID string) (*entity.Domain, error)
		Delete(ctx context.Context, name, tenantID string) error
		Update(ctx context.Context, d *entity.Domain) (*entity.Domain, error)
		Insert(ctx context.Context, d *entity.Domain) (*entity.Domain, error)
	}
)
