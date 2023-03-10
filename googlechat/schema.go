package googlechat

type Template struct {
	CardsV2 []CardV2 `json:"cardsV2"`
}

type CardV2 struct {
	CardId string `json:"cardId"`
	Card   Card   `json:"card"`
}

type Card struct {
	Header   Header    `json:"header"`
	Sections []Section `json:"sections"`
}

type Header struct {
	Title        string `json:"title,omitempty"`
	Subtitle     string `json:"subtitle,omitempty"`
	ImageUrl     string `json:"imageUrl,omitempty"`
	ImageType    string `json:"imageType,omitempty"`
	ImageAltText string `json:"imageAltText,omitempty"`
}

type Section struct {
	Header                    string    `json:"header,omitempty"`
	Collapsible               bool      `json:"collapsible,omitempty"`
	UncollapsibleWidgetsCount int       `json:"uncollapsibleWidgetsCount,omitempty"`
	Widgets                   []*Widget `json:"widgets,omitempty"`
}

type Widget struct {
	DecoratedText *DecoratedText `json:"decoratedText,omitempty"`
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`
	Divider       *Divider       `json:"divider,omitempty"`
	ButtonList    *ButtonList    `json:"buttonList,omitempty"`
}

type DecoratedText struct {
	StartIcon struct {
		KnownIcon string `json:"knownIcon"`
	} `json:"startIcon,omitempty"`
	Text string `json:"text,omitempty"`
}

type TextParagraph struct {
	Text string `json:"text"`
}

type Divider struct {
}

type ButtonList struct {
	Buttons []*Button `json:"buttons,omitempty"`
}

type Button struct {
	Text    string  `json:"text"`
	OnClick OnClick `json:"onClick"`
}

type OnClick struct {
	OpenLink OpenLink `json:"openLink"`
}

type OpenLink struct {
	Url string `json:"url"`
}

func (c *Card) SetHeader(header Header) {
	c.Header = header
}

func (c *Card) AddSection(section Section) {
	c.Sections = append(c.Sections, section)
}

func (s *Section) AddWidget(widget *Widget) {
	s.Widgets = append(s.Widgets, widget)
}

func (w *Widget) SetTextParagraph(textParagraph *TextParagraph) {
	w.TextParagraph = textParagraph
}

func (w *Widget) SetDecoratedText(decoratedText *DecoratedText) {
	w.DecoratedText = decoratedText
}

func (w *Widget) SetDivider(divider *Divider) {
	w.Divider = divider
}

func (w *Widget) AddButton(buttonList *ButtonList) {
	w.ButtonList = buttonList
}

func (bList *ButtonList) AddButton(b *Button) {
	bList.Buttons = append(bList.Buttons, b)
}

func (h *Header) SetTitle(title string) {
	h.Title = title
}

func (h *Header) SetSubTitle(subTitle string) {
	h.Subtitle = subTitle
}
