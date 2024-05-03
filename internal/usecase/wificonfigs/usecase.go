package wificonfigs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
	"github.com/open-amt-cloud-toolkit/console/internal/entity/dto"
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
func (uc *UseCase) CheckProfileExists(ctx context.Context, profileName, tenantID string) (bool, error) {
	data, err := uc.repo.CheckProfileExists(ctx, profileName, tenantID)
	if err != nil {
		return false, fmt.Errorf("WificonfigsUseCase - Count - s.repo.GetCount: %w", err)
	}

	return data, nil
}

func (uc *UseCase) GetCount(ctx context.Context, tenantID string) (int, error) {
	count, err := uc.repo.GetCount(ctx, tenantID)
	if err != nil {
		return 0, fmt.Errorf("WificonfigsUseCase - Count - s.repo.GetCount: %w", err)
	}

	return count, nil
}

func (uc *UseCase) Get(ctx context.Context, top, skip int, tenantID string) ([]dto.WirelessConfig, error) {
	data, err := uc.repo.Get(ctx, top, skip, tenantID)

	if err != nil {
		return nil, fmt.Errorf("WificonfigsUseCase - Get - s.repo.Get: %w", err)
	}

	// iterate over the data and convert each entity to dto
	d1 := make([]dto.WirelessConfig, len(data))
	for i, v := range data {
		d1[i] = *uc.entityToDTO(&v)
	}

	return d1, nil
}

func (uc *UseCase) GetByName(ctx context.Context, profileName, tenantID string) (dto.WirelessConfig, error) {
	data, err := uc.repo.GetByName(ctx, profileName, tenantID)
	if err != nil {
		return dto.WirelessConfig{}, fmt.Errorf("WificonfigsUseCase - GetByName - s.repo.GetByName: %w", err)
	}

	d1 := uc.entityToDTO(&data)

	return *d1, nil
}

func (uc *UseCase) Delete(ctx context.Context, profileName, tenantID string) (bool, error) {
	data, err := uc.repo.Delete(ctx, profileName, tenantID)
	if err != nil {
		return false, fmt.Errorf("WificonfigsUseCase - Delete - s.repo.Delete: %w", err)
	}

	return data, nil
}

func (uc *UseCase) Update(ctx context.Context, d *dto.WirelessConfig) (bool, error) {
	d1 := uc.dtoToEntity(d)

	data, err := uc.repo.Update(ctx, d1)
	if err != nil {
		return false, fmt.Errorf("WificonfigsUseCase - Update - s.repo.Update: %w", err)
	}

	return data, nil
}

func (uc *UseCase) Insert(ctx context.Context, d *dto.WirelessConfig) (string, error) {
	d1 := uc.dtoToEntity(d)

	data, err := uc.repo.Insert(ctx, d1)
	if err != nil {
		return "", fmt.Errorf("WificonfigsUseCase - Insert - s.repo.Insert: %w", err)
	}

	return data, nil
}

// convert dto.WirelessConfig to entity.WirelessConfig
func (uc *UseCase) dtoToEntity(d *dto.WirelessConfig) *entity.WirelessConfig {
	// convert []int to comma separated string
	linkPolicy := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(d.LinkPolicy)), ","), "[]")

	d1 := &entity.WirelessConfig{
		ProfileName:          d.ProfileName,
		AuthenticationMethod: d.AuthenticationMethod,
		EncryptionMethod:     d.EncryptionMethod,
		SSID:                 d.SSID,
		PSKValue:             d.PSKValue,
		PSKPassphrase:        d.PSKPassphrase,
		LinkPolicy:           &linkPolicy,
		TenantID:             d.TenantID,
		IEEE8021xProfileName: d.IEEE8021xProfileName,
		Version:              d.Version,
	}
	return d1
}

// convert entity.WirelessConfig to dto.WirelessConfig
func (uc *UseCase) entityToDTO(d *entity.WirelessConfig) *dto.WirelessConfig {
	// convert comma separated string to []int
	linkPolicyInt := []int{}
	if d.LinkPolicy != nil {
		linkPolicy := strings.Split(*d.LinkPolicy, ",")
		// convert []string to []int
		linkPolicyInt := make([]int, len(linkPolicy))
		for i, v := range linkPolicy {
			val, err := strconv.Atoi(v)
			if err != nil {
				// handle the error, e.g. log or return an error
				uc.log.Error("error converting string to int")
			}
			linkPolicyInt[i] = val
		}
	}

	d1 := &dto.WirelessConfig{
		ProfileName:          d.ProfileName,
		AuthenticationMethod: d.AuthenticationMethod,
		EncryptionMethod:     d.EncryptionMethod,
		SSID:                 d.SSID,
		PSKValue:             d.PSKValue,
		PSKPassphrase:        d.PSKPassphrase,
		LinkPolicy:           linkPolicyInt,
		TenantID:             d.TenantID,
		IEEE8021xProfileName: d.IEEE8021xProfileName,
		Version:              d.Version,
	}
	return d1
}
