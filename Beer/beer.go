package Beer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func Beer(option string, event *Struct.Message) bool {

	idOfBar := IndexOfBars(option)
	if idOfBar == "-1" {
		event.API.PostMessage(event.Channel, "Le bar n'existe pas ou n'est actuellement pas dans notre base de donnée", Struct.SlackParams)
		return false
	}
	res, err := http.Get("https://www.mistergoodbeer.com/map/" + idOfBar)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		fmt.Println(err)
		return false
	}
	// panel := doc.Find(".widget-pane .widget-pane-visible")
	fmt.Println("https://www.mistergoodbeer.com/map/" + idOfBar)
	params := Struct.SlackParams
	var fields []slack.AttachmentField
	for i := 0; i < 2; i++ {
		fields = append(fields, slack.AttachmentField{
			Title: doc.Find(".widget-pane-section-header-price").ChildrenFiltered(".Score").Eq(i).Find(".score-label").Text(),
			Value: doc.Find(".widget-pane-section-header-price").ChildrenFiltered(".Score").Eq(i).Find("score-base").Text(),
		})
	}
	fmt.Println(doc.Text())
	attachment := slack.Attachment{
		Color:      "ffcc00",
		AuthorLink: "https://www.mistergoodbeer.com/map/" + idOfBar,
		Title:      doc.Find(".widget-pane-section-header-hero-title").First().Text(),
		AuthorName: doc.Find(".widget-pane-section-header-hero-subtitle").First().Text(),
		Fields:     fields,
		Footer:     "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
	return true
}

func IndexOfBars(word string) string {
	for _, v := range Bars {
		if word == strings.ToLower(v.Name) {
			return v.ID
		}
		for _, a := range v.Abre {
			if word == strings.ToLower(a) {
				return v.ID
			}
		}
	}
	return "-1"
}
