package hubspot

type EngagementsRequest struct {
	EngagementProperties
}

type EngagementProperties struct {
	Engagement   Engagement             `json:"engagement"`
	Associations EngagementAssociations `json:"associations"`
	Metadata     Metadata               `json:"metadata"`
}

type EngagementsResponse struct {
	EngagementProperties
}

type Engagement struct {
	Active bool `json:"active"`
	//Optional long, corresponding to an Owner. Task engagements use the ownerId to populate the Assigned to field.
	OwnerId   int64  `json:"ownerId,omitempty"`
	Type      string `json:"type"`      //Required. One of: EMAIL, CALL, MEETING, TASK, NOTE
	Timestamp int64  `json:"timestamp"` //Optional timestamp (in milliseconds).
}

type EngagementAssociations struct {
	ContactIds []int64 `json:"contactIds,omitempty"`
	CompanyIds []int64 `json:"companyIds,omitempty"`
	DealIds    []int64 `json:"dealIds,omitempty"`
	OwnerIds   []int64 `json:"ownerIds,omitempty"`
	TicketIds  []int64 `json:"ticketIds,omitempty"`
}

type Metadata struct {
	Body          string   `json:"body"`
	Subject       string   `json:"subject"`
	Status        string   `json:"status"`              // String; One of NOT_STARTED, COMPLETED, IN_PROGRESS, WAITING, or DEFERRED
	ForObjectType string   `json:"forObjectType"`       // String; One of CONTACT or COMPANY, what object type the task is for.
	Reminders     []uint64 `json:"reminders,omitempty"` // Optional timestamp (in milliseconds). Reminder by time.
	TaskType      string   `json:"taskType,omitempty"`  // String; One of CALL, EMAIL, TODO
}
