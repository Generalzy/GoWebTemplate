package acount

type WebUserInfo struct {
	ID         uint   `gorm:"primarykey;column:id" json:"-"`
	Username   string `gorm:"column:username" json:"username"`
	Password   string `gorm:"column:password" json:"password"`
	Email      string `gorm:"column:email" json:"email"`
	Phone      string `gorm:"column:phone" json:"phone"`
	ProjectNum int    `gorm:"column:project_num" json:"project_num"`
}
