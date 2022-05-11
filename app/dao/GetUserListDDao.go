package dao

import (
	"poetry/app/model"
	"poetry/global"
)

var users []model.User

func GetUserListDao(page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(page_size).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0{
		return 0, userList
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	// result.Offset(offset).Limit(page_size).Find(&users)
	//
	for _, useSingle := range users {
		userItemMap := map[string]interface{}{
			"id":        useSingle.ID,
			"password":  useSingle.Password,
			"phone": useSingle.Phone,
		}
		userList = append(userList, userItemMap)
	}
	return total, userList
}

