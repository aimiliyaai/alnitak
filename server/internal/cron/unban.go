package cron

import (
	"time"

	"interastral-peace.com/alnitak/internal/global"
	"interastral-peace.com/alnitak/utils"

	"interastral-peace.com/alnitak/internal/domain/model"
)

// 解除封禁
func UnbanUser() {
	// 查找封禁表中需要解封的用户
	var userBans []model.UserBan
	if err := global.Mysql.Where("end_time <= ? and status = ?", time.Now(), 0).Find(&userBans).Error; err != nil {
		utils.ErrorLog("查询封禁表失败", "cron", err.Error())
		return
	}

	for _, v := range userBans {
		tx := global.Mysql.Begin()

		// 更新用户表状态
		if err := tx.Model(&model.User{}).Where("id = ?", v.Uid).Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
			utils.ErrorLog("更新用户表失败", "cron", err.Error())
			tx.Rollback()
			return
		}

		// 更新封禁状态
		if err := tx.Model(&model.UserBan{}).Where("id = ?", v.ID).Updates(map[string]interface{}{
			"status": 2,
		}).Error; err != nil {
			utils.ErrorLog("更新用户封禁表失败", "cron", err.Error())
			tx.Rollback()
			return
		}

		tx.Commit()
	}
}
