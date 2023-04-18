package control

import (
	"errors"
	"fmt"
	config "github.com/spf13/viper"
	"messenger/control/structure"
	rdbmsstructure "messenger/db/structure"
	"messenger/restapi/model/chat/request"
	"messenger/utility/encoding"
)

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

func (ctrl Controller) JoinChat(req request.GroupChat, Token string) (msg string, Error error) {
	checkAlready := false
	msg1 := ""
	var arrUser []uint
	UserRes := new(structure.UserRes)
	for _, m1 := range req.UserID {
		res, _ := ctrl.Access.RDBMS.CheckRoomChatForUser(req.RoomChatID, m1)
		if m1 != res.UserID {
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

			Username := UserRes.Data.FirstName + " " + UserRes.Data.LastName

			msg1 = msg1 + fmt.Sprintf(" %v, ", Username)
		}
	}

	if checkAlready {
		msg = fmt.Sprintf("User already exists in this chat in ( %v ) .", msg1)
	}

	for _, userId := range arrUser {
		reqGroupChat := rdbmsstructure.GroupChat{
			UserID:     userId,
			RoomChatID: req.RoomChatID,
		}

		err := ctrl.Access.RDBMS.JoinChat(reqGroupChat)
		if err != nil {
			Error = err
			return
		}
	}

	return
}
