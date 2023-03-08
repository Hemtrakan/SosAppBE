package control

import (
	"hotline/restapi/model/hotline/response"
)

func (ctrl Controller) GetHotLine() (res []response.HotlineNumber, Error error) {
	data, err := ctrl.Access.RDBMS.GetHotLine()
	if err != nil {
		Error = err
		return
	}

	arrHotline := []response.HotlineNumber{}

	for _, hotline := range data {
		reshotline := response.HotlineNumber{
			Id:          hotline.ID,
			Number:      hotline.Number,
			Description: hotline.Description,
		}
		arrHotline = append(arrHotline, reshotline)
	}

	res = arrHotline

	return
}
