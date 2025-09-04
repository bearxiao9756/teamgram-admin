
package group

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/group"
    groupReq "github.com/flipped-aurora/gin-vue-admin/server/model/group/request"
)

type ChatsService struct {}
// CreateChats 创建chats表记录
// Author [yourname](https://github.com/yourname)
func (chatsService *ChatsService) CreateChats(ctx context.Context, chats *group.Chats) (err error) {
	err = global.GVA_DB.Create(chats).Error
	return err
}

// DeleteChats 删除chats表记录
// Author [yourname](https://github.com/yourname)
func (chatsService *ChatsService)DeleteChats(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&group.Chats{},"id = ?",id).Error
	return err
}

// DeleteChatsByIds 批量删除chats表记录
// Author [yourname](https://github.com/yourname)
func (chatsService *ChatsService)DeleteChatsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]group.Chats{},"id in ?",ids).Error
	return err
}

// UpdateChats 更新chats表记录
// Author [yourname](https://github.com/yourname)
func (chatsService *ChatsService)UpdateChats(ctx context.Context, chats group.Chats) (err error) {
	err = global.GVA_DB.Model(&group.Chats{}).Where("id = ?",chats.Id).Updates(&chats).Error
	return err
}

// GetChats 根据id获取chats表记录
// Author [yourname](https://github.com/yourname)
func (chatsService *ChatsService)GetChats(ctx context.Context, id string) (chats group.Chats, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chats).Error
	return
}
// GetChatsInfoList 分页获取chats表记录
// Author [yourname](https://github.com/yourname)
func (chatsService *ChatsService)GetChatsInfoList(ctx context.Context, info groupReq.ChatsSearch) (list []group.Chats, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&group.Chats{})
    var chatss []group.Chats
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&chatss).Error
	return  chatss, total, err
}
func (chatsService *ChatsService)GetChatsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
