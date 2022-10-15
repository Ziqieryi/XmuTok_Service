package loginImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
)

func Login(username *string, password *string, data *module.UserTable) string {
	// Judge whether the username and password is it false
	err := dao.Db.Where("user_name = ? AND account_password = ? ", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "The username or password is incorrect"
		}
		return err.Error()
	}
	return ""
}
