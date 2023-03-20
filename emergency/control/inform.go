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
	Username := ""
	UserRes := new(structure.UserRes)
	for _, m1 := range resp {
		if pointer.GetStringValue(m1.UserNotiID) != "" {
			UserNotiID := pointer.GetStringValue(m1.UserNotiID)
			URL := "http://127.0.0.1:80/SosApp/accounts/user/" + UserNotiID
			fmt.Printf("Url : %#v \n", URL)
			httpHeaderMap := map[string]string{}
			httpHeaderMap["Authorization"] = Token

			HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
			if err != nil {
				Error = err
				fmt.Printf("Error : , %#v\n", Error)
				return
			}

			if HttpResponse.HttpStatusCode != 200 {
				Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v", HttpResponse.HttpStatusCode))
				fmt.Printf("%#v\n", Error)
				return
			}

			err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
			if err != nil {
				Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
				fmt.Printf("%#v\n", Error)
				return
			}

			if UserRes.FirstName != "" && UserRes.LastName != "" {
				Username = UserRes.FirstName + " " + UserRes.LastName
			}
		}

		mapData := inform.InformResponse{
			ID:                  pointer.GetStringValue(m1.ID),
			Description:         pointer.GetStringValue(m1.Description),
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

	return
}

func (ctrl Controller) GetInformById(ReqInformId, Token string) (res inform.InformResponse, Error error) {
	InformId, err := strconv.Atoi(ReqInformId)
	if err != nil {
		Error = err
		return
	}
	resp, err := ctrl.Access.RDBMS.GetImageByInformId(uint(InformId))
	if err != nil {
		Error = err
		return
	}
	URL := "http://127.0.0.1:80/SosApp/accounts/user/" + pointer.GetStringValue(resp.UserNotiID)
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
	var ImageInfoArr []inform.ImageInfo

	for _, image := range resp.ImageInfo {
		ImageInfo := inform.ImageInfo{
			ImageId: pointer.GetStringValue(image.ImageId),
			Image:   pointer.GetStringValue(image.Image),
		}
		ImageInfoArr = append(ImageInfoArr, ImageInfo)
	}

	mapData := inform.InformResponse{
		ID:                  pointer.GetStringValue(resp.ID),
		Description:         pointer.GetStringValue(resp.Description),
		Image:               ImageInfoArr,
		PhoneNumberCallBack: pointer.GetStringValue(resp.CALLBack),
		Latitude:            pointer.GetStringValue(resp.Latitude),
		Longitude:           pointer.GetStringValue(resp.Longitude),
		UserName:            Username,
		Workplace:           UserRes.Workplace,
		SubTypeName:         pointer.GetStringValue(resp.SubTypeName),
		Date:                pointer.GetStringValue(resp.InformCreatedAt),
		Status:              pointer.GetStringValue(resp.Status),
	}
	res = mapData

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
