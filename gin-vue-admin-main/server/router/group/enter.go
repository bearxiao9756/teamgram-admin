package group

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ChatsRouter }

var chatsApi = api.ApiGroupApp.GroupApiGroup.ChatsApi
