package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// 设置User的表名为``
func (User) TableName() string {
	return "users"
}
