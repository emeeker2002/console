package dto

type BootSetting struct {
	Action int  `json:"action" binding:"required" example:"8"`
	UseSOL bool `json:"useSOL" binding:"required" example:"true"`
}
