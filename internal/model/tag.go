package model

// Tag 标签模型
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TableName 获取表名
func (t Tag) TableName() string {
	return "blog_tag"
}
