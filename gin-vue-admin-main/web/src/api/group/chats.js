import service from '@/utils/request'
// @Tags Chats
// @Summary 创建chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Chats true "创建chats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chats/createChats [post]
export const createChats = (data) => {
  return service({
    url: '/chats/createChats',
    method: 'post',
    data
  })
}

// @Tags Chats
// @Summary 删除chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Chats true "删除chats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chats/deleteChats [delete]
export const deleteChats = (params) => {
  return service({
    url: '/chats/deleteChats',
    method: 'delete',
    params
  })
}

// @Tags Chats
// @Summary 批量删除chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除chats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chats/deleteChats [delete]
export const deleteChatsByIds = (params) => {
  return service({
    url: '/chats/deleteChatsByIds',
    method: 'delete',
    params
  })
}

// @Tags Chats
// @Summary 更新chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Chats true "更新chats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chats/updateChats [put]
export const updateChats = (data) => {
  return service({
    url: '/chats/updateChats',
    method: 'put',
    data
  })
}

// @Tags Chats
// @Summary 用id查询chats表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Chats true "用id查询chats表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chats/findChats [get]
export const findChats = (params) => {
  return service({
    url: '/chats/findChats',
    method: 'get',
    params
  })
}

// @Tags Chats
// @Summary 分页获取chats表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取chats表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chats/getChatsList [get]
export const getChatsList = (params) => {
  return service({
    url: '/chats/getChatsList',
    method: 'get',
    params
  })
}

// @Tags Chats
// @Summary 不需要鉴权的chats表接口
// @Accept application/json
// @Produce application/json
// @Param data query groupReq.ChatsSearch true "分页获取chats表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /chats/getChatsPublic [get]
export const getChatsPublic = () => {
  return service({
    url: '/chats/getChatsPublic',
    method: 'get',
  })
}
