package group

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChatsRouter struct {}

// InitChatsRouter 初始化 chats表 路由信息
func (s *ChatsRouter) InitChatsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	chatsRouter := Router.Group("chats").Use(middleware.OperationRecord())
	chatsRouterWithoutRecord := Router.Group("chats")
	chatsRouterWithoutAuth := PublicRouter.Group("chats")
	{
		chatsRouter.POST("createChats", chatsApi.CreateChats)   // 新建chats表
		chatsRouter.DELETE("deleteChats", chatsApi.DeleteChats) // 删除chats表
		chatsRouter.DELETE("deleteChatsByIds", chatsApi.DeleteChatsByIds) // 批量删除chats表
		chatsRouter.PUT("updateChats", chatsApi.UpdateChats)    // 更新chats表
	}
	{
		chatsRouterWithoutRecord.GET("findChats", chatsApi.FindChats)        // 根据ID获取chats表
		chatsRouterWithoutRecord.GET("getChatsList", chatsApi.GetChatsList)  // 获取chats表列表
	}
	{
	    chatsRouterWithoutAuth.GET("getChatsPublic", chatsApi.GetChatsPublic)  // chats表开放接口
	}
}
