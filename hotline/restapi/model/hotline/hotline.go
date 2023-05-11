package hotline

type HotlineNumber struct {
	Id          uint   `json:"id"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

type HotlineReq struct {
	Number      string `json:"number,omitempty"`
	Description string `json:"description,omitempty"`
}

type HistoryReq struct {
	HotlineNumberID int   `json:"hotlineNumberID,omitempty"`
	UserId          *uint `json:"userId,omitempty"`
}
