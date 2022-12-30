# googlechat 

Google Chat supply feature send notification. 

## Install

`go get github.com/vietnguyen87/pkg-golang`

## Usage

```go 
package main

import (   dsgjksdlgjskdlg 
	"github.com/vietnguyen87/pkg-golang/googlechat"
)

func main() {
	//Notification to group chat
	data := builder.BuildTicketNotificationTemplate(ticket)
	chatNotiUrl = config.GetAppConfig().GoogleTicketNotificationHost
	go googlechat.SendMessageTemplate(context.Background(), chatNotiUrl, data)
}
```        

## Options

### Build Ticket Notification Template

Set Header 

```go
var card googlechat.Card
//Set Header
card.SetHeader(googlechat.Header{
    Title: fmt.Sprintf("%s - Ticket owner: %s", ticketProp.Subject, ticketProp.HubspotOwnerId),
})
```

Add Widgets: DecoratedText, TextParagraph, Divider, ButtonList

```go
//Define widget
var widget googlechat.Widget
//TextParagraph widget
widget.SetTextParagraph(&googlechat.TextParagraph{
    Text: ticketProp.Content,
})
//Button widget
var btnWidget googlechat.Widget
var btnList googlechat.ButtonList
btnList.AddButton(&googlechat.Button{
    Text: fmt.Sprintf("Xem chi tiết: %s", ticket.Id),
    OnClick: googlechat.OnClick{
        OpenLink: googlechat.OpenLink{
            Url: fmt.Sprintf("https://app.hubspot.com/contacts/%v/ticket/%v", portalId, ticket.Id),
        },
    },
})
btnWidget.AddButton(&btnList)

//Section
var section googlechat.Section
section.AddWidget(&widget)
section.AddWidget(&googlechat.Widget{
    Divider: &googlechat.Divider{},
})
section.AddWidget(&btnWidget)
// Set Section
card.AddSection(section)
```
Return Template 

```go 
//Template
return googlechat.Template{
    CardsV2: []googlechat.CardV2{
        {
            CardId: "ticket-new-notification",
            Card:   card,
        },
    },
}
```