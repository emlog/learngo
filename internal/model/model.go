package model

type Model struct {
	ID         uint32 `gorm: "primary_key" json:"id"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreateTime uint32 `json:"create_time"`
}
