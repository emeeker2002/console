package domains

import (
	"context"
	"fmt"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
	"github.com/open-amt-cloud-toolkit/console/pkg/consoleerrors"
	"github.com/open-amt-cloud-toolkit/console/pkg/logger"
)

// UseCase -.
type UseCase struct {
	repo Repository
	log  logger.Interface
}

// New -.
func New(r Repository, log logger.Interface) *UseCase {
	return &UseCase{
		repo: r,
		log:  log,
	}
}

// History - getting translate history from store.
func (uc *UseCase) GetCount(ctx context.Context, tenantID string) (int, error) {
	count, err := uc.repo.GetCount(ctx, tenantID)
	if err != nil {
		return 0, fmt.Errorf("DomainsUseCase - Count - uc.repo.GetCount: %w", err)
	}

	return count, nil
}

func (uc *UseCase) Get(ctx context.Context, top, skip int, tenantID string) ([]entity.Domain, error) {
	data, err := uc.repo.Get(ctx, top, skip, tenantID)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - Get - uc.repo.Get: %w", err)
	}

	return data, nil
}

func (uc *UseCase) GetDomainByDomainSuffix(ctx context.Context, domainSuffix, tenantID string) (*entity.Domain, error) {
	data, err := uc.repo.GetDomainByDomainSuffix(ctx, domainSuffix, tenantID)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - GetDomainByDomainSuffix - uc.repo.GetDomainByDomainSuffix: %w", err)
	}

	return data, nil
}

func (uc *UseCase) GetByName(ctx context.Context, domainName, tenantID string) (*entity.Domain, error) {
	data, err := uc.repo.GetByName(ctx, domainName, tenantID)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - GetByName - uc.repo.GetByName: %w", err)
	}

	if data.DomainSuffix == "" {
		return nil, consoleerrors.ErrNotFound
		//return nil, fmt.Errorf("DomainsUseCase - GetByName - uc.repo.GetByName: %w", consoleerrors.ErrNotFound)
	}

	return data, nil
}

func (uc *UseCase) Delete(ctx context.Context, domainName, tenantID string) error {
	isSuccessful, err := uc.repo.Delete(ctx, domainName, tenantID)
	if err != nil {
		return fmt.Errorf("DomainsUseCase - Delete - uc.repo.Delete: %w", err)
	}
	if !isSuccessful {
		return consoleerrors.ErrNotFound
		// return fmt.Errorf("DomainsUseCase - Delete - uc.repo.Delete: %w", consoleerrors.ErrNotFound)
	}

	return nil
}

func (uc *UseCase) Update(ctx context.Context, d *entity.Domain) (*entity.Domain, error) {
	_, err := uc.repo.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - Update - uc.repo.Update: %w", err)
	}

	updateDomain, err := uc.repo.GetByName(ctx, d.ProfileName, "")
	if err != nil {
		return nil, err
	}

	return updateDomain, nil
}

func (uc *UseCase) Insert(ctx context.Context, d *entity.Domain) (*entity.Domain, error) {
	_, err := uc.repo.Insert(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - Insert - uc.repo.Insert: %w", err)
	}

	newDomain, err := uc.repo.GetByName(ctx, d.ProfileName, "")
	if err != nil {
		return nil, err
	}

	return newDomain, nil
}
