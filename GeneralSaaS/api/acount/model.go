package acount

type WebUserInfo struct {
	ID         uint   `gorm:"primarykey;column:id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
	ProjectNum int    `gorm:"column:project_num"`
}
