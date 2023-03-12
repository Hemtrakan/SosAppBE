package control

import (
	rdbmsstructure "emergency/db/structure"
	"emergency/restapi/model/inform"
	"strconv"
)

func (ctrl Controller) PostInform(req *inform.InformRequest) (Error error) {

	userId, err := strconv.ParseUint(req.UserID, 10, 32)
	if err != nil {
		Error = err
		return
	}
	newReqInform := rdbmsstructure.InformImage{
		Inform: rdbmsstructure.Inform{
			Description:         req.Description,
			PhoneNumberCallBack: req.PhoneNumberCallBack,
			Latitude:            req.Latitude,
			Longitude:           req.Longitude,
			UserID:              uint(userId),
			SubTypeID:           req.SubTypeID,
		},
		Image: req.Images,
	}

	err = ctrl.Access.RDBMS.PostInform(newReqInform)
	if err != nil {
		Error = err
		return
	}
	return
}
