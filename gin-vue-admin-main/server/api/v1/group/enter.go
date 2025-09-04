package group

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ChatsApi }

var chatsService = service.ServiceGroupApp.GroupServiceGroup.ChatsService
