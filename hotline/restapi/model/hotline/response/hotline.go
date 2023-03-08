package response

type HotlineNumber struct {
	Id          uint   `json:"id"`
	Number      string `json:"number"`
	Description string `json:"description"`
}
