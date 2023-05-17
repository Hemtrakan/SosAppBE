package control

import (
	"emergency/constant"
	"emergency/control/structure"
	rdbmsstructure "emergency/db/structure"
	"emergency/db/structure/query"
	"emergency/restapi/model"
	"emergency/utility/encoding"
	"emergency/utility/pointer"
	"errors"
	"fmt"
	config "github.com/spf13/viper"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (ctrl Controller) GetInform(UserId uint, Token, role string) (res []model.InformResponse, Error error) {
	var resp []*query.InformInfoList
	var err error

	if role == constant.Admin {
		resp, err = ctrl.Access.RDBMS.GetAllInformListForAdmin()
		if err != nil {
			Error = err
			return
		}
	} else {
		resp, err = ctrl.Access.RDBMS.GetInformList(UserId)
		if err != nil {
			Error = err
			return
		}
	}

	for _, m1 := range resp {
		UserRes := new(structure.UserRes)
		UserNotiID := ""
		Username := ""
		PhoneNumber := ""
		if pointer.GetStringValue(m1.UserNotiID) != "0" {
			UserNotiID = pointer.GetStringValue(m1.UserNotiID)
		}

		if UserNotiID != "" {
			account := config.GetString("url.account")
			URL := ""

			if role == constant.Admin {
				URL = account + "admin/user/" + UserNotiID
			} else {
				URL = account + "user/" + UserNotiID
			}

			httpHeaderMap := map[string]string{}
			httpHeaderMap["Authorization"] = Token

			HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
			if err != nil {
				Error = err
				return
			}

			if HttpResponse.HttpStatusCode != 200 {
				Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v \n Msg : %#v", HttpResponse.HttpStatusCode, HttpResponse.ResponseMsg))
				return
			}

			err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
			if err != nil {
				Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
				return
			}

			if UserRes.Data.FirstName != "" && UserRes.Data.LastName != "" {
				Username = UserRes.Data.FirstName + " " + UserRes.Data.LastName
			}
			PhoneNumber = UserRes.Data.PhoneNumber
		}

		var status, _ = constant.Status(constant.StatusStep1).Status()
		if pointer.GetStringValue(m1.Status) != "" {
			s, _ := strconv.ParseInt(pointer.GetStringValue(m1.Status), 0, 0)
			status, _ = constant.Status(s).Status()
		}

		mapData := model.InformResponse{
			ID:                  pointer.GetStringValue(m1.ID),
			Description:         pointer.GetStringValue(m1.Description),
			PhoneNumberCallBack: pointer.GetStringValue(m1.CALLBack),
			Latitude:            pointer.GetStringValue(m1.Latitude),
			Longitude:           pointer.GetStringValue(m1.Longitude),
			UserName:            Username,
			PhoneNumber:         PhoneNumber,
			Workplace:           UserRes.Data.Workplace,
			SubTypeName:         pointer.GetStringValue(m1.SubTypeName),
			Date:                pointer.GetStringValue(m1.InformCreatedAt),
			UpdateDate:          pointer.GetStringValue(m1.InformUpdateAt),
			Status:              pointer.GetStringValue(status),
			StatusChat:          pointer.GetBooleanValue(m1.StatusChat),
		}
		res = append(res, mapData)
	}

	return
}

func (ctrl Controller) GetInformById(ReqInformId, Token, role string) (res model.InformResponse, Error error) {
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
	UserNotiID := ""
	if pointer.GetStringValue(resp.UserNotiID) != "0" {
		UserNotiID = pointer.GetStringValue(resp.UserNotiID)
	}

	UserRes := new(structure.UserRes)
	if UserNotiID != "" {
		account := config.GetString("url.account")
		URL := ""
		if role == constant.Admin {
			URL = account + "admin/user/" + UserNotiID
		} else {
			URL = account + "user/" + UserNotiID
		}
		httpHeaderMap := map[string]string{}
		httpHeaderMap["Authorization"] = Token

		HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
		if err != nil {
			Error = err
			return
		}

		if HttpResponse.HttpStatusCode != 200 {
			Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v \n Msg : %#v", HttpResponse.HttpStatusCode, HttpResponse.ResponseMsg))

			return
		}

		err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
		if err != nil {
			Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
			return
		}
	}

	Username := ""
	if UserRes.Data.FirstName != "" && UserRes.Data.LastName != "" {
		Username = UserRes.Data.FirstName + " " + UserRes.Data.LastName
	}
	var PhoneNumber = UserRes.Data.PhoneNumber
	var ImageInfoArr []model.ImageInfo

	for _, image := range resp.ImageInfo {
		ImageInfo := model.ImageInfo{
			ImageId: pointer.GetStringValue(image.ImageId),
			Image:   pointer.GetStringValue(image.Image),
		}
		ImageInfoArr = append(ImageInfoArr, ImageInfo)
	}

	var status, _ = constant.Status(constant.StatusStep1).Status()
	if pointer.GetStringValue(resp.Status) != "" {
		s, _ := strconv.ParseInt(pointer.GetStringValue(resp.Status), 0, 0)
		status, _ = constant.Status(s).Status()
	}

	mapData := model.InformResponse{
		ID:                  pointer.GetStringValue(resp.ID),
		Description:         pointer.GetStringValue(resp.Description),
		Image:               ImageInfoArr,
		PhoneNumberCallBack: pointer.GetStringValue(resp.CALLBack),
		Latitude:            pointer.GetStringValue(resp.Latitude),
		Longitude:           pointer.GetStringValue(resp.Longitude),
		UserName:            Username,
		PhoneNumber:         PhoneNumber,
		Workplace:           UserRes.Data.Workplace,
		SubTypeName:         pointer.GetStringValue(resp.SubTypeName),
		Date:                pointer.GetStringValue(resp.InformCreatedAt),
		UpdateDate:          pointer.GetStringValue(resp.InformUpdateAt),
		Status:              pointer.GetStringValue(status),
		StatusChat:          pointer.GetBooleanValue(resp.StatusChat),
	}
	res = mapData

	return
}

func (ctrl Controller) GetAllInformOps(Token string) (res []model.InformResponse, Error error) {
	resp, err := ctrl.Access.RDBMS.GetAllInformList()
	if err != nil {
		Error = err
		return
	}

	for _, m1 := range resp {
		Username := ""
		UserRes := new(structure.UserRes)
		UserInformID := ""
		//PhoneNumber := ""
		if pointer.GetStringValue(m1.UserInformID) != "0" {
			UserInformID = pointer.GetStringValue(m1.UserInformID)
		}

		if UserInformID != "" {
			account := config.GetString("url.account")
			URL := account + "user/" + UserInformID

			httpHeaderMap := map[string]string{}
			httpHeaderMap["Authorization"] = Token

			HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
			if err != nil {
				Error = err
				return
			}

			if HttpResponse.HttpStatusCode != 200 {
				Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v \n Msg : %#v", HttpResponse.HttpStatusCode, HttpResponse.ResponseMsg))
				return
			}

			err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
			if err != nil {
				Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
				return
			}

			if UserRes.Data.FirstName != "" && UserRes.Data.LastName != "" {
				Username = UserRes.Data.FirstName + " " + UserRes.Data.LastName
			}
			//PhoneNumber = UserRes.Data.PhoneNumber
		}

		var status, _ = constant.Status(constant.StatusStep1).Status()
		if pointer.GetStringValue(m1.Status) != "" {
			s, _ := strconv.ParseInt(pointer.GetStringValue(m1.Status), 0, 0)
			status, _ = constant.Status(s).Status()
		}

		mapData := model.InformResponse{
			ID:                  pointer.GetStringValue(m1.ID),
			Description:         pointer.GetStringValue(m1.Description),
			PhoneNumberCallBack: UserRes.Data.PhoneNumber,
			Latitude:            pointer.GetStringValue(m1.Latitude),
			Longitude:           pointer.GetStringValue(m1.Longitude),
			UserName:            Username,
			//PhoneNumber:         PhoneNumber,
			Workplace:   UserRes.Data.Workplace,
			SubTypeName: pointer.GetStringValue(m1.SubTypeName),
			Date:        pointer.GetStringValue(m1.InformCreatedAt),
			UpdateDate:  pointer.GetStringValue(m1.InformUpdateAt),
			Status:      pointer.GetStringValue(status),
			StatusChat:  pointer.GetBooleanValue(m1.StatusChat),
		}
		res = append(res, mapData)
	}
	return
}

func (ctrl Controller) GetInformOps(OpsId uint, Token, role string) (res []model.InformResponse, Error error) {
	resp, err := ctrl.Access.RDBMS.GetInformListByOpsId(OpsId)
	if err != nil {
		Error = err
		return
	}

	for _, m1 := range resp {
		UserInformID := ""
		Username := ""
		//PhoneNumber := ""
		UserRes := new(structure.UserRes)
		if pointer.GetStringValue(m1.UserInformID) != "0" {
			UserInformID = pointer.GetStringValue(m1.UserInformID)
		}

		if UserInformID != "" {
			account := config.GetString("url.account")
			URL := ""
			if role == constant.Admin {
				URL = account + "admin/user/" + UserInformID
			} else {
				URL = account + "user/" + UserInformID
			}

			httpHeaderMap := map[string]string{}
			httpHeaderMap["Authorization"] = Token

			HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
			if err != nil {
				Error = err
				return
			}

			if HttpResponse.HttpStatusCode != 200 {
				Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v \n Msg : %#v", HttpResponse.HttpStatusCode, HttpResponse.ResponseMsg))
				return
			}

			err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
			if err != nil {
				Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
				return
			}

			if UserRes.Data.FirstName != "" && UserRes.Data.LastName != "" {
				Username = UserRes.Data.FirstName + " " + UserRes.Data.LastName
			}
			//PhoneNumber = UserRes.Data.PhoneNumber

		}

		var status, _ = constant.Status(constant.StatusStep1).Status()
		if pointer.GetStringValue(m1.Status) != "" {
			s, _ := strconv.ParseInt(pointer.GetStringValue(m1.Status), 0, 0)
			status, _ = constant.Status(s).Status()
		}

		mapData := model.InformResponse{
			ID:                  pointer.GetStringValue(m1.ID),
			Description:         pointer.GetStringValue(m1.Description),
			PhoneNumberCallBack: UserRes.Data.PhoneNumber,
			Latitude:            pointer.GetStringValue(m1.Latitude),
			Longitude:           pointer.GetStringValue(m1.Longitude),
			UserName:            Username,
			Workplace:           UserRes.Data.Workplace,
			//PhoneNumber:         PhoneNumber,
			SubTypeName: pointer.GetStringValue(m1.SubTypeName),
			Date:        pointer.GetStringValue(m1.InformCreatedAt),
			UpdateDate:  pointer.GetStringValue(m1.InformUpdateAt),
			Status:      pointer.GetStringValue(status),
			StatusChat:  pointer.GetBooleanValue(m1.StatusChat),
		}
		res = append(res, mapData)
	}
	return
}

func (ctrl Controller) GetInformOpsById(ReqInformId, Token, role string) (res model.InformResponse, Error error) {
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
	UserID := ""
	if pointer.GetStringValue(resp.UserInformID) != "0" {
		UserID = pointer.GetStringValue(resp.UserInformID)
	}

	UserRes := new(structure.UserRes)
	if UserID != "" {
		account := config.GetString("url.account")
		URL := ""
		if role == constant.Admin {
			URL = account + "admin/user/" + UserID
		} else {
			URL = account + "user/" + UserID
		}

		httpHeaderMap := map[string]string{}
		httpHeaderMap["Authorization"] = Token

		HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
		if err != nil {
			Error = err
			return
		}

		if HttpResponse.HttpStatusCode != 200 {
			Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v \n Msg : %#v", HttpResponse.HttpStatusCode, HttpResponse.ResponseMsg))

			return
		}

		err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
		if err != nil {
			Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
			return
		}
	}

	Username := ""
	if UserRes.Data.FirstName != "" && UserRes.Data.LastName != "" {
		Username = UserRes.Data.FirstName + " " + UserRes.Data.LastName
	}
	//var PhoneNumber = UserRes.Data.PhoneNumber

	var ImageInfoArr []model.ImageInfo

	for _, image := range resp.ImageInfo {
		ImageInfo := model.ImageInfo{
			ImageId: pointer.GetStringValue(image.ImageId),
			Image:   pointer.GetStringValue(image.Image),
		}
		ImageInfoArr = append(ImageInfoArr, ImageInfo)
	}

	var status, _ = constant.Status(constant.StatusStep1).Status()
	if pointer.GetStringValue(resp.Status) != "" {
		s, _ := strconv.ParseInt(pointer.GetStringValue(resp.Status), 0, 0)
		status, _ = constant.Status(s).Status()
	}

	mapData := model.InformResponse{
		ID:                  pointer.GetStringValue(resp.ID),
		Description:         pointer.GetStringValue(resp.Description),
		Image:               ImageInfoArr,
		PhoneNumberCallBack: UserRes.Data.PhoneNumber,
		Latitude:            pointer.GetStringValue(resp.Latitude),
		Longitude:           pointer.GetStringValue(resp.Longitude),
		UserId:              UserID,
		UserName:            Username,
		//PhoneNumber:         PhoneNumber,
		Workplace:   UserRes.Data.Workplace,
		SubTypeName: pointer.GetStringValue(resp.SubTypeName),
		Date:        pointer.GetStringValue(resp.InformCreatedAt),
		UpdateDate:  pointer.GetStringValue(resp.InformUpdateAt),
		Status:      pointer.GetStringValue(status),
		StatusChat:  pointer.GetBooleanValue(resp.StatusChat),
	}
	res = mapData

	return
}

func (ctrl Controller) PostInform(req *model.InformRequest) (Error error) {

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
		Status:              strconv.Itoa(constant.StatusStep1),
		StatusChat:          false,
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

func (ctrl Controller) UpdateInform(req *model.UpdateInformRequest, token string, informId uint, role string) (Error error) {
	UserRes := new(structure.UserRes)
	if informId != 0 {
		account := config.GetString("url.account")
		URL := ""
		if role == constant.Admin {
			URL = account + "admin/"
		} else {
			URL = account + "ops/"
		}

		httpHeaderMap := map[string]string{}
		httpHeaderMap["Authorization"] = token

		HttpResponse, err := ctrl.HttpClient.Get(URL, httpHeaderMap)
		if err != nil {
			Error = err
			return
		}

		if HttpResponse.HttpStatusCode != 200 {
			Error = errors.New(fmt.Sprintf("Error HttpStatusCode : %#v \n Msg : %#v", HttpResponse.HttpStatusCode, HttpResponse.ResponseMsg))
			return
		}

		err = encoding.JsonToStruct(HttpResponse.ResponseMsg, UserRes)
		if err != nil {
			Error = errors.New(fmt.Sprintf("URL : %#v json response message invalid", err.Error()))
			return
		}
	}

	if UserRes.Code == constant.SuccessCode {

		OpsId, err := strconv.ParseUint(UserRes.Data.ID, 0, 0)

		newReqInform := rdbmsstructure.Inform{
			Model: gorm.Model{
				ID:        informId,
				UpdatedAt: time.Now().Add(time.Hour * 7),
			},
			Description:         pointer.GetStringValue(req.Description),
			PhoneNumberCallBack: pointer.GetStringValue(req.PhoneNumberCallBack),
			Latitude:            pointer.GetStringValue(req.Latitude),
			Longitude:           pointer.GetStringValue(req.Longitude),
			SubTypeID:           pointer.GetUintValue(req.SubTypeID),
			OpsID:               uint(OpsId),
			Status:              strconv.Itoa(pointer.GetIntValue(req.Status)),
			StatusChat:          pointer.GetBooleanValue(req.StatusChat),
		}

		err = ctrl.Access.RDBMS.PutInform(newReqInform)
		if err != nil {
			Error = err
			return
		}
	}

	return
}

func (ctrl Controller) DeleteInform(userId, informId uint) (Error error) {
	data := rdbmsstructure.Inform{
		Model: gorm.Model{
			ID:        informId,
			UpdatedAt: time.Now().Add(time.Hour * 7),
		},
		DeletedBy: userId,
	}

	err := ctrl.Access.RDBMS.PutInform(data)
	if err != nil {
		Error = err
		return
	}

	err = ctrl.Access.RDBMS.DeleteInform(data)
	if err != nil {
		Error = err
		return
	}

	return
}
