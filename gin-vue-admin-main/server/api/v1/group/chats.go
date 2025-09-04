package group

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/group"
    groupReq "github.com/flipped-aurora/gin-vue-admin/server/model/group/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ChatsApi struct {}



// CreateChats 创建chats表
// @Tags Chats
// @Summary 创建chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body group.Chats true "创建chats表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /chats/createChats [post]
func (chatsApi *ChatsApi) CreateChats(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var chats group.Chats
	err := c.ShouldBindJSON(&chats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = chatsService.CreateChats(ctx,&chats)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteChats 删除chats表
// @Tags Chats
// @Summary 删除chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body group.Chats true "删除chats表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /chats/deleteChats [delete]
func (chatsApi *ChatsApi) DeleteChats(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := chatsService.DeleteChats(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteChatsByIds 批量删除chats表
// @Tags Chats
// @Summary 批量删除chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /chats/deleteChatsByIds [delete]
func (chatsApi *ChatsApi) DeleteChatsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := chatsService.DeleteChatsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateChats 更新chats表
// @Tags Chats
// @Summary 更新chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body group.Chats true "更新chats表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /chats/updateChats [put]
func (chatsApi *ChatsApi) UpdateChats(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var chats group.Chats
	err := c.ShouldBindJSON(&chats)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = chatsService.UpdateChats(ctx,chats)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindChats 用id查询chats表
// @Tags Chats
// @Summary 用id查询chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询chats表"
// @Success 200 {object} response.Response{data=group.Chats,msg=string} "查询成功"
// @Router /chats/findChats [get]
func (chatsApi *ChatsApi) FindChats(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	rechats, err := chatsService.GetChats(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rechats, c)
}
// GetChatsList 分页获取chats表列表
// @Tags Chats
// @Summary 分页获取chats表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query groupReq.ChatsSearch true "分页获取chats表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /chats/getChatsList [get]
func (chatsApi *ChatsApi) GetChatsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo groupReq.ChatsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := chatsService.GetChatsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetChatsPublic 不需要鉴权的chats表接口
// @Tags Chats
// @Summary 不需要鉴权的chats表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /chats/getChatsPublic [get]
func (chatsApi *ChatsApi) GetChatsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    chatsService.GetChatsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的chats表接口信息",
    }, "获取成功", c)
}
