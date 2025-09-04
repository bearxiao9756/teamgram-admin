
// 自动生成模板Chats
package group
import (
	"time"
)

// chats表 结构体  Chats
type Chats struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //id字段
  CreatorUserId  *int64 `json:"creatorUserId" form:"creatorUserId" gorm:"column:creator_user_id;"`  //creatorUserId字段
  AccessHash  *int64 `json:"accessHash" form:"accessHash" gorm:"column:access_hash;"`  //accessHash字段
  RandomId  *int64 `json:"randomId" form:"randomId" gorm:"column:random_id;"`  //randomId字段
  ParticipantCount  *int32 `json:"participantCount" form:"participantCount" gorm:"column:participant_count;"`  //participantCount字段
  Title  *string `json:"title" form:"title" gorm:"column:title;size:255;"`  //title字段
  About  *string `json:"about" form:"about" gorm:"column:about;size:255;"`  //about字段
  PhotoId  *int64 `json:"photoId" form:"photoId" gorm:"column:photo_id;"`  //photoId字段
  DefaultBannedRights  *int64 `json:"defaultBannedRights" form:"defaultBannedRights" gorm:"column:default_banned_rights;"`  //defaultBannedRights字段
  MigratedToId  *int64 `json:"migratedToId" form:"migratedToId" gorm:"column:migrated_to_id;"`  //migratedToId字段
  MigratedToAccessHash  *int64 `json:"migratedToAccessHash" form:"migratedToAccessHash" gorm:"column:migrated_to_access_hash;"`  //migratedToAccessHash字段
  AvailableReactionsType  *int32 `json:"availableReactionsType" form:"availableReactionsType" gorm:"column:available_reactions_type;"`  //availableReactionsType字段
  AvailableReactions  *string `json:"availableReactions" form:"availableReactions" gorm:"column:available_reactions;size:128;"`  //availableReactions字段
  Deactivated  *bool `json:"deactivated" form:"deactivated" gorm:"column:deactivated;"`  //deactivated字段
  Noforwards  *bool `json:"noforwards" form:"noforwards" gorm:"column:noforwards;"`  //noforwards字段
  TtlPeriod  *int32 `json:"ttlPeriod" form:"ttlPeriod" gorm:"column:ttl_period;"`  //ttlPeriod字段
  Version  *int32 `json:"version" form:"version" gorm:"column:version;"`  //version字段
  Date  *int64 `json:"date" form:"date" gorm:"column:date;"`  //date字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
}


// TableName chats表 Chats自定义表名 chats
func (Chats) TableName() string {
    return "chats"
}





