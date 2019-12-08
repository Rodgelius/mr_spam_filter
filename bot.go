package main
import (
  "github.com/Syfaro/telegram-bot-api"
  "log"
  "fmt"
  "strings"
)

func main() {
  bot, err := tgbotapi.NewBotAPI("ENTER_BOT_KEY_HERE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	// initialize channel for updates from API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, ok := bot.GetUpdatesChan(ucfg)
	// get updates from channel
	log.Printf("%v", ok)
	for update := range upd {
		if  update.Message != nil {
			fmt.Println("update is ", ok)
			UserName := update.Message.From.UserName
			Text := update.Message.Text
			if Text == "" {
				Text = update.Message.Caption
			}
			ChatName := update.Message.Chat.UserName
			var bws [41]string
			bws[0] = "investor"
			bws[1] = "Altcoin"
			bws[2] = "hydrạ2węb.com"
			bws[3] = "HYcheck_bot"
			bws[4] = "HydraoniondeepTGbot"
			bws[5] = "Yobit"
			bws[6] = "тєℓєgяαм"
			bws[7] = "hudra2web.com"
			bws[8] = "официальный список зеркал:"
			bws[9] = "gashish"
			bws[10] = "gazhish"
			bws[11] = "инутка рекламы"
			bws[12] = "微信"
			bws[13] = "电"
			bws[14] = "Ворованные вещи"
			bws[15] = "Хватит кормить магазины"
			bws[16] = "мфетамин"
			bws[17] = "Binance"
			bws[18] = "joinchat"
			bws[19] = "件"
			bws[20] = "Crypto Channel"
			bws[21] = "трейдер"
			bws[22] = "Трейдер"
			bws[23] = "ИНВЕСТ ФОНДЫ"
			bws[24] = "торгует на канале"
			bws[25] = "لم"
			bws[26] = "月"
			bws[27] = "аполните анкету"
			bws[28] = "щу  специалиста"
			bws[29] = "щу специалиста"
			bws[30] = "оводим набор на свободные вакансии"
                        bws[31] = "инвестируем с умом"
			bws[32] = "看"
			bws[33] = "#вакансия"
			bws[34] = "слуги по взлому"
			bws[35] = "CONTABO"
			bws[36] = "contabo"
			bws[37] = "тр3йд"
			bws[38] = "tart investing"
			bws[39] = "БОМБИТ ПУКАН"
			bws[40] = "бомбит пукан"
                        // Ban for Text or Capture
			for _, bw := range bws {
				// Compare Text variable with one of BanWordS
				if (strings.Contains(Text, bw)) {
					cmc := tgbotapi.ChatMemberConfig{
					ChatID:	update.Message.Chat.ID,
					UserID: update.Message.From.ID,
					}
					kcmc := tgbotapi.KickChatMemberConfig{
					cmc,
					CHATID, // put here chatid wherebot should check messages
					}
					bot.KickChatMember(kcmc)
					dmc := tgbotapi.DeleteMessageConfig{
					update.Message.Chat.ID,
					update.Message.MessageID,
					}
					bot.DeleteMessage(dmc)
					reply := "Banned:\nUser: @" + UserName + "\nin Chat: @" + ChatName + "\nfor: " + Text
					msg := tgbotapi.NewMessage(CHATID, reply) // chatid for logging messages of banned users
					bot.Send(msg)
				}
			}
                        // Only delete for this strings
			var dws [29]string
			dws[0] = "ебаная"
			dws[1] = "t.me/joinchat"
			dws[2] = "Вinаncе"
			dws[3] = "trade"
			dws[4] = "signals"
			dws[5] = "Реклама"
			dws[6] = "реклама"
			dws[7] = "shishki"
			dws[8] = "trading"
			dws[9] = "Путин"
			dws[10] = "ПУТИН"
			dws[11] = "Kryptomarkt"
			dws[12] = "د"
			dws[13] = "ТРЕЙДИНГ"
                        dws[14] = "наш канал"
			dws[15] = "пyтин"
			dws[16] = "реферал"
			dws[17] = "хуй"
			dws[18] = "издец"
			dws[19] = "ебаный"
			dws[20] = "хуев"
			dws[21] = " блять"
			dws[22] = "н а х у й"
			dws[23] = "блядс"
			dws[24] = "П о х у й"
			dws[25] = "п о х у й"
			dws[26] = "Ебучи"
			dws[27] = "ебучи"
			dws[28] = "издил"
			for _, dw := range dws {
				if strings.Contains(Text, dw) {
					dmc := tgbotapi.DeleteMessageConfig{
					update.Message.Chat.ID,
                                        update.Message.MessageID,
                                        }
					bot.DeleteMessage(dmc)
					reply := "Deleted message:\nUser: @" + UserName + "\nin Chat: @" + ChatName + "\nfor: " + Text
					msg := tgbotapi.NewMessage(CHAT_ID, reply) // chatid for logging deleted messages
					bot.Send(msg)
				}
			}
		}
	}
}
