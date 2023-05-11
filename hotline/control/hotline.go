package control

import (
	"gorm.io/gorm"
	rdbmstructure "hotline/db/structure"
	model "hotline/restapi/model/hotline"
	"time"
)

func (ctrl Controller) PostHistory(req *model.HistoryReq) (Error error) {
	data := rdbmstructure.History{
		HotlineNumberID: req.HotlineNumberID,
		UserID:          *req.UserId,
		Time:            time.Now().Local(),
	}

	err := ctrl.Access.RDBMS.PostHistory(data)
	if err != nil {
		Error = err
		return
	}

	return
}

func (ctrl Controller) GetHotLine() (res []model.HotlineNumber, Error error) {
	data, err := ctrl.Access.RDBMS.GetHotLine()
	if err != nil {
		Error = err
		return
	}

	var arrHotline []model.HotlineNumber

	for _, hotline := range data {
		reshotline := model.HotlineNumber{
			Id:          hotline.ID,
			Number:      hotline.Number,
			Description: hotline.Description,
		}
		arrHotline = append(arrHotline, reshotline)
	}

	res = arrHotline

	return
}

func (ctrl Controller) GetHotLineById(id uint) (res model.HotlineNumber, Error error) {
	data, err := ctrl.Access.RDBMS.GetHotLineById(id)
	if err != nil {
		Error = err
		return
	}
	resHotline := model.HotlineNumber{
		Id:          data.ID,
		Number:      data.Number,
		Description: data.Description,
	}

	res = resHotline

	return
}

func (ctrl Controller) PostHotLine(req *model.HotlineReq, userId uint) (Error error) {
	data := rdbmstructure.HotlineNumber{
		Number:           req.Number,
		Description:      req.Description,
		UserIDLogUpdated: userId,
	}

	err := ctrl.Access.RDBMS.PostHotLine(data)
	if err != nil {
		Error = err
		return
	}

	return
}

func (ctrl Controller) PutHotLine(req *model.HotlineReq, HotLineId, userId uint) (Error error) {
	data := rdbmstructure.HotlineNumber{
		Model: gorm.Model{
			ID: HotLineId,
		},
		Number:           req.Number,
		Description:      req.Description,
		UserIDLogUpdated: userId,
	}

	err := ctrl.Access.RDBMS.PutHotLine(data)
	if err != nil {
		Error = err
		return
	}

	return
}

func (ctrl Controller) DeleteHotLine(HotLineId, userId uint) (Error error) {
	data := rdbmstructure.HotlineNumber{
		Model: gorm.Model{
			ID: HotLineId,
		},
		UserIDLogUpdated: userId,
		DeletedBy:        &userId,
	}

	err := ctrl.Access.RDBMS.PutHotLine(data)
	if err != nil {
		Error = err
		return
	}

	err = ctrl.Access.RDBMS.DeleteHotLine(data)
	if err != nil {
		Error = err
		return
	}

	return
}
