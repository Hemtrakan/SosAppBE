package control

import (
	"errors"
	"fmt"
	config "github.com/spf13/viper"
	"messenger/control/structure"
	rdbmsstructure "messenger/db/structure"
	"messenger/restapi/model/chat/request"
	chatRes "messenger/restapi/model/chat/response"
	"messenger/utility/encoding"
)

func (ctrl Controller) GetChatList(userId uint) (res []chatRes.GetChatList, Error error) {
	resDB, err := ctrl.Access.RDBMS.GetRoomChatListByUserId(userId)
	if err != nil {
		Error = err
		return
	}
	var dataArr []chatRes.GetChatList
	for _, m1 := range resDB {
		data := chatRes.GetChatList{
			RoomChatID: fmt.Sprintf("%v", m1.RoomChat.ID),
			RoomName:   m1.RoomChat.Name,
			OwnerId:    fmt.Sprintf("%v", m1.RoomChat.UserOwnerId),
			CreatedAt:  m1.RoomChat.CreatedAt,
			UpdatedAt:  m1.RoomChat.UpdatedAt,
			DeletedAT:  m1.RoomChat.DeletedAt,
			DeleteBy:   fmt.Sprintf("%v", m1.RoomChat.DeletedBy),
		}
		dataArr = append(dataArr, data)
	}
	res = dataArr
	return
}

func (ctrl Controller) RoomChat(userId uint, req request.RoomChatReq, Token string) (Error error) {
	reqGroupChat := rdbmsstructure.GroupChat{
		UserID: userId,
		RoomChat: rdbmsstructure.RoomChat{
			Name:        req.RoomName,
			UserOwnerId: userId,
		},
	}

	res, err := ctrl.Access.RDBMS.RoomChat(reqGroupChat)
	if err != nil {
		Error = err
		return
	}

	requestGroupChat := request.GroupChat{
		RoomChatID: res.RoomChatID,
		UserID:     req.GroupChat.UserID,
	}

	_, err = ctrl.JoinChat(requestGroupChat, Token)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl Controller) JoinChat(req request.GroupChat, Token string) (res chatRes.JoinChatRes, Error error) {

	roomChat, err := ctrl.Access.RDBMS.GetRoomChatById(req.RoomChatID)
	if err != nil {
		Error = errors.New("roomChat not found.")
		return
	}

	checkAlready := false
	var arrUser []uint
	var UserAlready []string

	UserRes := new(structure.UserRes)
	for _, m1 := range req.UserID {
		resDB, _ := ctrl.Access.RDBMS.CheckRoomChatUser(req.RoomChatID, m1)
		if m1 != resDB.UserID {
			arrUser = append(arrUser, m1)
		} else {
			checkAlready = true
			account := config.GetString("url.account")

			URL := account + "user/" + fmt.Sprintf("%v", m1)

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

			UserAlready = append(UserAlready, UserRes.Data.FirstName+" "+UserRes.Data.LastName)
		}
	}

	if checkAlready {
		res.Mag = fmt.Sprintf("User already exists in this chat in.")
		res.Username = UserAlready
	}

	for _, userId := range arrUser {
		reqGroupChat := rdbmsstructure.GroupChat{
			UserID:     userId,
			RoomChatID: roomChat.ID,
		}

		err := ctrl.Access.RDBMS.JoinChat(reqGroupChat)
		if err != nil {
			Error = err
			return
		}
	}

	return
}
