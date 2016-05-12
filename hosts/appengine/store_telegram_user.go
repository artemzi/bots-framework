package gae_host

import (
	"github.com/strongo/bots-framework/platforms/telegram"
	"github.com/strongo/bots-framework/core"
	"google.golang.org/appengine/datastore"
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"time"
)

type GaeTelegramUserStore struct {
	GaeBotUserStore
}
var _ bots.BotUserStore = (*GaeTelegramUserStore)(nil) // Check for interface implementation at compile time

func NewGaeTelegramUserStore(log bots.Logger, r *http.Request, gaeAppUserStore GaeAppUserStore) GaeTelegramUserStore {
	return GaeTelegramUserStore{
		GaeBotUserStore: GaeBotUserStore{
			GaeBaseStore: NewGaeBaseStore(log, r, telegram_bot.TelegramUserKind),
			gaeAppUserStore: gaeAppUserStore,
			newBotUserEntity: func(apiUser bots.WebhookActor) bots.BotUser {
				if apiUser == nil {
					return &telegram_bot.TelegramUser{
					}
				} else {
					return &telegram_bot.TelegramUser{
						BotUserEntity: bots.BotUserEntity{
							BotEntity: bots.BotEntity{
								OwnedByUser: bots.OwnedByUser{
									DtCreated: time.Now(),
								},
							},
							FirstName: apiUser.GetFirstName(),
							LastName: apiUser.GetLastName(),
							UserName: apiUser.GetUserName(),
						},
					}
				}
			},
			validateBotUserEntityType: func(entity bots.BotUser) {
				if _, ok := entity.(*telegram_bot.TelegramUser); !ok {
					panic(fmt.Sprintf("Expected *telegram_bot.TelegramUser but received %T", entity))
				}
			},
			botUserKey: func(botUserId interface{}) *datastore.Key {
				if intID, ok := botUserId.(int); ok {
					return datastore.NewKey(appengine.NewContext(r), telegram_bot.TelegramUserKind, "", (int64)(intID), nil)
				} else {
					panic(fmt.Sprintf("Expected botUserId as int, got: %T", botUserId))
				}
			},
		},
	}
}