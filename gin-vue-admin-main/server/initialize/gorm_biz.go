package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/group"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(user.Users{}, group.Chats{})
	if err != nil {
		return err
	}
	return nil
}
