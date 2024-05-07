package entity

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Profile struct {
	ProfileName                string               `json:"profileName,omitempty" binding:"required" example:"My Profile"`
	AMTPassword                string               `json:"amtPassword,omitempty" binding:"required_without=GenerateRandomPassword,len=0|min=8,max=32,genpasswordwone" example:"my_password"`
	CreationDate               string               `json:"creationDate,omitempty" example:"2021-07-01T00:00:00Z"`
	CreatedBy                  string               `json:"created_by,omitempty" example:"admin"`
	GenerateRandomPassword     bool                 `json:"generateRandomPassword,omitempty" example:"true"`
	CIRAConfigName             *string              `json:"ciraConfigName,omitempty" binding:"omitempty,ciraortls" example:"My CIRA Config"`
	Activation                 string               `json:"activation" binding:"required,oneof=ccmactivate acmactivate" example:"activate"`
	MEBXPassword               string               `json:"mebxPassword,omitempty" binding:"ccmactivation,len=0|min=8,max=32" example:"my_password"`
	GenerateRandomMEBxPassword bool                 `json:"generateRandomMEBxPassword,omitempty" example:"true"`
	CIRAConfigObject           *CIRAConfig          `json:"ciraConfigObject,omitempty"`
	Tags                       []string             `json:"tags,omitempty" example:"tag1,tag2"`
	DhcpEnabled                bool                 `json:"dhcpEnabled,omitempty" example:"true"`
	IPSyncEnabled              bool                 `json:"ipSyncEnabled,omitempty" example:"true"`
	LocalWifiSyncEnabled       bool                 `json:"localWifiSyncEnabled,omitempty" example:"true"`
	WifiConfigs                []ProfileWifiConfigs `json:"wifiConfigs,omitempty"`
	TenantID                   string               `json:"tenantId" example:"abc123"`
	TLSMode                    int                  `json:"tlsMode,omitempty" binding:"omitempty,min=1,max=4" example:"1"`
	TLSCerts                   *TLSCerts            `json:"tlsCerts,omitempty"`
	TLSSigningAuthority        string               `json:"tlsSigningAuthority,omitempty" binding:"omitempty,oneof=SelfSigned MicrosoftCA" example:"SelfSigned"`
	UserConsent                string               `json:"userConsent,omitempty" example:"All"`
	IDEREnabled                bool                 `json:"iderEnabled,omitempty" example:"true"`
	KVMEnabled                 bool                 `json:"kvmEnabled,omitempty" binding:"oneof=true false" example:"true"`
	SOLEnabled                 bool                 `json:"solEnabled,omitempty" binding:"oneof=true false" example:"true"`
	Ieee8021xProfileName       *string              `json:"ieee8021xProfileName,omitempty" example:"My Profile"`
	Ieee8021xProfileObject     *IEEE8021xConfig     `json:"ieee8021xProfileObject,omitempty"`
	Version                    string               `json:"version,omitempty" example:"1.0.0"`
}

const (
	TLSModeNone int = iota
	TLSModeServerOnly
	TLSModeServerAllowNonTLS
	TLSModeMutualOnly
	TLSModeMutualAllowNonTLS
)

const (
	TLSSigningAuthoritySelfSigned  string = "SelfSigned"
	TLSSigningAuthorityMicrosoftCA string = "MicrosoftCA"
)

const (
	UserConsentNone    string = "None"
	UserConsentAll     string = "All"
	UserConsentKVMOnly string = "KVM"
)

var ValidateCIRAOrTLS validator.Func = func(fl validator.FieldLevel) bool {
	tlsMode := fl.Parent().FieldByName("TLSMode").Interface().(int)
	if tlsMode != 0 {
		return false

	}
	return true
}

var GenPasswordWithOne validator.Func = func(fl validator.FieldLevel) bool {
	randPassword := fl.Parent().FieldByName("GenerateRandomPassword").Interface().(bool)
	password := fl.Parent().FieldByName("AMTPassword").Interface().(string)
	fmt.Println(randPassword)
	fmt.Println(password)
	if randPassword && password != "" {
		fmt.Println("it was passed")
		return false

	}
	return true
}

var ValidateCCMActivation validator.Func = func(fl validator.FieldLevel) bool {
	activation := fl.Parent().FieldByName("Activation").String()

	if activation == "ccmactivate" {
		return true
	}

	generateRandom := fl.Parent().FieldByName("GenerateRandomMEBxPassword").Bool()
	mebXPassword := fl.Parent().FieldByName("MEBXPassword").String()

	if !generateRandom && mebXPassword == "" {
		return false
	}

	return true
}
