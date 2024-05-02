package profiles

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
		return 0, fmt.Errorf("ProfilesUseCase - Count - uc.repo.GetCount: %w", err)
	}

	return count, nil
}

func (uc *UseCase) Get(ctx context.Context, top, skip int, tenantID string) ([]entity.Profile, error) {
	data, err := uc.repo.Get(ctx, top, skip, tenantID)
	if err != nil {
		return nil, fmt.Errorf("ProfilesUseCase - Get - uc.repo.Get: %w", err)
	}

	return data, nil
}

func (uc *UseCase) GetByName(ctx context.Context, profileName, tenantID string) (*entity.Profile, error) {
	data, err := uc.repo.GetByName(ctx, profileName, tenantID)
	if err != nil {
		return nil, fmt.Errorf("ProfilesUseCase - GetByName - uc.repo.GetByName: %w", err)
	}

	return data, nil
}

func (uc *UseCase) Delete(ctx context.Context, profileName, tenantID string) error {
	isSuccessful, err := uc.repo.Delete(ctx, profileName, tenantID)
	if err != nil {
		return fmt.Errorf("ProfilesUseCase - Delete - uc.repo.Delete: %w", err)
	}

	if !isSuccessful {
		return consoleerrors.ErrNotFound
	}

	return nil
}

func (uc *UseCase) Update(ctx context.Context, d *entity.Profile) (*entity.Profile, error) {
	_, err := uc.repo.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("ProfilesUseCase - Update - uc.repo.Update: %w", err)
	}

	updatedProfile, err := uc.repo.GetByName(ctx, d.ProfileName, "")
	if err != nil {
		return nil, err
	}

	return updatedProfile, nil
}

func (uc *UseCase) Insert(ctx context.Context, d *entity.Profile) (*entity.Profile, error) {
	_, err := uc.repo.Insert(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("ProfilesUseCase - Insert - uc.repo.Insert: %w", err)
	}

	newProfile, err := uc.repo.GetByName(ctx, d.ProfileName, "")
	if err != nil {
		return nil, err
	}

	return newProfile, nil
}
