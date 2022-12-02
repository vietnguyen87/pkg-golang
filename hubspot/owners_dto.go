package hubspot

type OwnersResponse struct {
	ErrorResponse
	Id        string                `json:"id"`
	Email     string                `json:"email" example:"bcooper@biglytics.net"`
	Firstname string                `json:"firstname" example:"Bryan"`
	Lastname  string                `json:"lastname" example:"Cooper"`
	UserId    string                `json:"userId" example:"27047699"`
	Teams     []OwnersResponseTeams `json:"teams"`
	CreatedAt string                `json:"createdAt"`
	UpdatedAt string                `json:"updatedAt"`
	Archived  bool                  `json:"archived"`
}

type OwnersResponseTeams struct {
	Id      string `json:"id"`
	Name    string `json:"name" example:"sale"`
	Primary bool   `json:"primary"`
}

type GetListOwnerResponse struct {
	Results []OwnersResponse `json:"results"`
	Paging  OwnerPagination  `json:"paging"`
}

type OwnerPagination struct {
	Next OwnerPaginationNext `json:"next"`
}
type OwnerPaginationNext struct {
	After string `json:"after"`
	Link  string `json:"link"`
}
