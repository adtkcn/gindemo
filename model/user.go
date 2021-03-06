package model

// User 用户
type User struct {
	ID   uint   `gorm:"primaryKey"  form:"id" json:"id"`
	Name string `gorm:"type:varchar(20);uniqueIndex" form:"name" json:"name"`
	Pwd  string `gorm:"type:varchar(100)" form:"pwd" json:"pwd"`

	Updated int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created int64 `gorm:"autoCreateTime:milli"` // 使用时间戳秒数填充创建时间

	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
