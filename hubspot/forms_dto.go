package hubspot

type SubmitFormRequest struct {
	Firstname              string `json:"firstname"`
	Phone                  string `json:"phone"`
	Grade                  string `json:"grade"`
	InterestOnSubjects     string `json:"interest_on_subjects"`
	MktCampaignEngagement  string `json:"mkt_campaign_engagement"`
	HsAnalyticsSourceData1 string `json:"hs_analytics_source_data_1"`
	HsAnalyticsSourceData2 string `json:"hs_analytics_source_data_2"`
}

type SubmitRequest struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
