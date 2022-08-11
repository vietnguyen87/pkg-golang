package schema

type AccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
}

type SendZNSResp struct {
	Data    Data   `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type Data struct {
	SentTime string `json:"sent_time"`
	MsgId    string `json:"msg_id"`
}
