package Bars

import (
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Cdt(option string, event *Struct.Message) bool {
	open := "Fermé !"
	color := "danger"
	if IsCdtOpen() {
		open = "Ouvert !"
		color = "good"
	}
	params := Struct.SlackParams
	attachment := slack.Attachment{
		Color: color,
		Title: "Café du Théatre",
		Text:  open,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Horaire",
				Value: "8-00h tous les jours!",
			},
			slack.AttachmentField{
				Title: "Mell's Brau :beer:",
				Value: "Pinte: 4.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Pelforth :beer:",
				Value: "Pinte: 5.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "La Chouffe :beer: :gnome:",
				Value: "Pinte: 6.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cuvée des trolls/Vedett/Affligem :beer:",
				Value: "Pinte: 6.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Vin :wine_glass:",
				Value: "4.80€",
				Short: true,
			},
		},
		Footer: "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
	return true
}

func IsCdtOpen() bool {
	now := time.Now()
	if now.Hour() >= 8 && now.Hour() < 24 {
		return true
	}
	return false
}
