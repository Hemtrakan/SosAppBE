package control

import (
	"emergency/control/structure"
	rdbmsstructure "emergency/db/structure"
	"emergency/restapi/model/inform"
	"emergency/utility/encoding"
	"emergency/utility/pointer"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (ctrl Controller) GetInform(UserId uint, Token string) (res []inform.InformResponse, Error error) {
	resp, err := ctrl.Access.RDBMS.GetInformList(UserId)
	if err != nil {
		Error = err
		return
	}
	for _, m1 := range resp {
		URL := "http://127.0.0.1:80/SosApp/accounts/user/" + pointer.GetStringValue(m1.UserNotiID)
		//URL := "localhost:80/SosApp/accounts/user/" + pointer.GetStringValue(m1.UserNotiID)
		httpHeaderMap := map[string]string{}
		httpHeaderMap["Authorization"] = Token

		HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
		if err != nil {
			Error = err
			return
		}

		if HttpResponse.HttpStatusCode != 200 {
			Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v", HttpResponse.HttpStatusCode))
			return
		}

		UserRes := new(structure.UserRes)
		err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
		if err != nil {
			Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
			return
		}

		Username := ""
		if UserRes.FirstName != "" && UserRes.LastName != "" {
			Username = UserRes.FirstName + " " + UserRes.LastName
		}

		mapData := inform.InformResponse{
			ID:                  pointer.GetStringValue(m1.ID),
			Description:         pointer.GetStringValue(m1.Description),
			Image:               pointer.GetStringValue(m1.Image),
			PhoneNumberCallBack: pointer.GetStringValue(m1.CALLBack),
			Latitude:            pointer.GetStringValue(m1.Latitude),
			Longitude:           pointer.GetStringValue(m1.Longitude),
			UserName:            Username,
			Workplace:           UserRes.Workplace,
			SubTypeName:         pointer.GetStringValue(m1.SubTypeName),
			Date:                pointer.GetStringValue(m1.InformCreatedAt),
			Status:              pointer.GetStringValue(m1.Status),
		}
		res = append(res, mapData)
	}
	//ctrl.HttpProxy.GetUser()
	//err = ctrl.Access.RDBMS.PostInform(newReqInform)
	//if err != nil {
	//	Error = err
	//	return
	//}
	return
}

func (ctrl Controller) PostInform(req *inform.InformRequest) (Error error) {

	userId, err := strconv.ParseUint(req.UserID, 10, 32)
	if err != nil {
		Error = err
		return
	}

	newReqInform := rdbmsstructure.Inform{
		Model: gorm.Model{
			CreatedAt: time.Now().Add(time.Hour * 7),
			UpdatedAt: time.Now().Add(time.Hour * 7),
		},
		Description:         req.Description,
		PhoneNumberCallBack: req.PhoneNumberCallBack,
		Latitude:            req.Latitude,
		Longitude:           req.Longitude,
		UserID:              uint(userId),
		SubTypeID:           req.SubTypeID,
	}

	var newReqImageArr []rdbmsstructure.InformImage
	for _, m1 := range req.Images {
		newReqImage := rdbmsstructure.InformImage{
			Model: gorm.Model{
				CreatedAt: time.Now().Add(time.Hour * 7),
				UpdatedAt: time.Now().Add(time.Hour * 7),
			},
			Image: m1,
		}
		newReqImageArr = append(newReqImageArr, newReqImage)
	}

	err = ctrl.Access.RDBMS.PostInform(newReqImageArr, newReqInform)
	if err != nil {
		Error = err
		return
	}
	return
}
