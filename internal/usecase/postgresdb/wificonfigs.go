package postgresdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
	"github.com/open-amt-cloud-toolkit/console/pkg/postgres"
)

// WirelessRepo -.
type WirelessRepo struct {
	*postgres.DB
}

// New -.
func NewWirelessRepo(pg *postgres.DB) *WirelessRepo {
	return &WirelessRepo{pg}
}

// CheckProfileExits -.
func (r *WirelessRepo) CheckProfileExists(ctx context.Context, profileName, tenantID string) (bool, error) {
	sqlQuery, _, err := r.Builder.
		Select("COUNT(*) OVER() AS total_count").
		From("wirelessconfigs").
		Where("wireless_profile_name and tenant_id = ?", profileName, tenantID).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("WirelessRepo - CheckProfileExists - r.Builder: %w", err)
	}

	var count int

	err = r.Pool.QueryRow(ctx, sqlQuery, tenantID).Scan(&count)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("WirelessRepo - CheckProfileExists - r.Pool.QueryRow: %w", err)
	}

	return true, nil
}

// GetCount -.
func (r *WirelessRepo) GetCount(ctx context.Context, tenantID string) (int, error) {
	sqlQuery, _, err := r.Builder.
		Select("COUNT(*) OVER() AS total_count").
		From("wirelessconfigs").
		Where("tenant_id = ?", tenantID).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("WirelessRepo - GetCount - r.Builder: %w", err)
	}

	var count int

	err = r.Pool.QueryRow(ctx, sqlQuery, tenantID).Scan(&count)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}

		return 0, fmt.Errorf("WirelessRepo - GetCount - r.Pool.QueryRow: %w", err)
	}

	return count, nil
}

// Get -.
func (r *WirelessRepo) Get(ctx context.Context, top, skip int, tenantID string) ([]entity.WirelessConfig, error) {
	if top == 0 {
		top = 100
	}

	sqlQuery, _, err := r.Builder.
		Select(`
			wireless_profile_name,
			authentication_method,
			encryption_method,
			ssid,
			psk_value,
			psk_passphrase,
			link_policy,
			tenant_id,
			ieee8021x_profile_name,
      		CAST(xmin as text) as xmin
			`).
		From("wirelessconfigs").
		Where("tenant_id = ?", tenantID).
		OrderBy("wireless_profile_name").
		Limit(uint64(top)).
		Offset(uint64(skip)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("WirelessRepo - Get - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sqlQuery, tenantID)
	if err != nil {
		return nil, fmt.Errorf("WirelessRepo - Get - r.Pool.Query: %w", err)
	}

	defer rows.Close()

	wirelessConfigs := make([]entity.WirelessConfig, 0)

	for rows.Next() {
		p := entity.WirelessConfig{}

		err = rows.Scan(&p.ProfileName, &p.AuthenticationMethod, &p.EncryptionMethod, &p.SSID, &p.PSKValue, &p.PSKPassphrase, &p.LinkPolicy, &p.TenantID, &p.IEEE8021xProfileName, &p.Version)
		if err != nil {
			return nil, fmt.Errorf("WirelessRepo - Get - rows.Scan: %w", err)
		}

		wirelessConfigs = append(wirelessConfigs, p)
	}

	return wirelessConfigs, nil
}

// GetByName -.
func (r *WirelessRepo) GetByName(ctx context.Context, profileName, tenantID string) (entity.WirelessConfig, error) {
	sqlQuery, _, err := r.Builder.
		Select(`
			wireless_profile_name,
			authentication_method,
			encryption_method,
			ssid,
			psk_value,
			psk_passphrase,
			link_policy,
			tenant_id,
			ieee8021x_profile_name,
			CAST(xmin as text) as xmin
			`).
		From("wirelessconfigs").
		Where("wireless_profile_name = ? and tenant_id = ?", profileName, tenantID).
		ToSql()
	if err != nil {
		return entity.WirelessConfig{}, fmt.Errorf("WirelessRepo - GetByName - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sqlQuery, profileName, tenantID)
	if err != nil {
		return entity.WirelessConfig{}, fmt.Errorf("WirelessRepo - GetByName - r.Pool.Query: %w", err)
	}

	defer rows.Close()

	wirelessConfigs := make([]entity.WirelessConfig, 0)

	for rows.Next() {
		p := entity.WirelessConfig{}

		err = rows.Scan(&p.ProfileName, &p.AuthenticationMethod, &p.EncryptionMethod, &p.SSID, &p.PSKValue, &p.PSKPassphrase, &p.LinkPolicy, &p.TenantID, &p.IEEE8021xProfileName, &p.Version)
		if err != nil {
			return p, fmt.Errorf("WirelessRepo - GetByName - rows.Scan: %w", err)
		}

		wirelessConfigs = append(wirelessConfigs, p)
	}

	if len(wirelessConfigs) == 0 {
		return entity.WirelessConfig{}, errors.New(postgres.NotFound)
	}

	return wirelessConfigs[0], nil
}

// Delete -.
func (r *WirelessRepo) Delete(ctx context.Context, profileName, tenantID string) (bool, error) {
	sqlQuery, args, err := r.Builder.
		Delete("wirelessconfigs").
		Where("wireless_profile_name = ? AND tenant_id = ?", profileName, tenantID).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("WirelessRepo - Delete - r.Builder: %w", err)
	}

	res, err := r.Pool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return false, fmt.Errorf("WirelessRepo - Delete - r.Pool.Exec: %w", err)
	}

	return res.RowsAffected() > 0, nil
}

// Update -.
func (r *WirelessRepo) Update(ctx context.Context, p *entity.WirelessConfig) (bool, error) {
	sqlQuery, args, err := r.Builder.
		Update("wirelessconfigs").
		Set("authentication_method", p.AuthenticationMethod).
		Set("encryption_method", p.EncryptionMethod).
		Set("ssid", p.SSID).
		Set("psk_value", p.PSKValue).
		Set("psk_passphrase", p.PSKPassphrase).
		Set("link_policy", p.LinkPolicy).
		Set("ieee8021x_profile_name", p.IEEE8021xProfileName).
		Where("wireless_profile_name = ? AND tenant_id = ?", p.ProfileName, p.TenantID).
		Suffix("AND xmin::text = ?", p.Version).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("WirelessRepo - Update - r.Builder: %w", err)
	}

	res, err := r.Pool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return false, fmt.Errorf("WirelessRepo - Update - r.Pool.Exec: %w", err)
	}

	return res.RowsAffected() > 0, nil
}

// Insert -.
func (r *WirelessRepo) Insert(ctx context.Context, p *entity.WirelessConfig) (string, error) {
	date := time.Now().Format("2006-01-02 15:04:05")

	ieeeProfileName := p.IEEE8021xProfileName

	if p.IEEE8021xProfileName != nil {
		if *p.IEEE8021xProfileName == "" {
			ieeeProfileName = nil
		}
	}

	sqlQuery, args, err := r.Builder.
		Insert("wirelessconfigs").
		Columns("wireless_profile_name", "authentication_method", "encryption_method", "ssid", "psk_value", "psk_passphrase", "link_policy", "creation_date", "tenant_id", "ieee8021x_profile_name").
		Values(p.ProfileName, p.AuthenticationMethod, p.EncryptionMethod, p.SSID, p.PSKValue, p.PSKPassphrase, p.LinkPolicy, date, p.TenantID, ieeeProfileName).
		Suffix("RETURNING xmin::text").
		ToSql()
	if err != nil {
		return "", fmt.Errorf("WirelessRepo - Insert - r.Builder: %w", err)
	}

	var version string

	err = r.Pool.QueryRow(ctx, sqlQuery, args...).Scan(&version)
	if err != nil {
		return "", fmt.Errorf("WirelessRepo - Insert - r.Pool.QueryRow: %w", err)
	}

	return version, nil
}
