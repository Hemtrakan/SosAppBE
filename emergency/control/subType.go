package control

import (
	rdbmsstructure "emergency/db/structure"
	structuer "emergency/restapi/model"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func (ctrl Controller) GetSubType() (res []structuer.SubTypeRes, Error error) {
	subTypes, err := ctrl.Access.RDBMS.GetSubType()
	if err != nil {
		Error = err
	}

	for _, m1 := range subTypes {
		types := m1.Type
		obj := structuer.SubTypeRes{
			ID:           fmt.Sprintf("%v", m1.ID),
			CreatedAt:    m1.CreatedAt,
			UpdatedAt:    m1.UpdatedAt,
			NameSubType:  m1.Name,
			ImageSubType: m1.ImageSubType,
			DeletedBy:    fmt.Sprintf("%v", m1.DeletedBy),
			TypeRes: &structuer.TypeRes{
				ID:        fmt.Sprintf("%v", types.ID),
				CreatedAt: types.CreatedAt,
				UpdatedAt: types.UpdatedAt,
				NameType:  types.Name,
				ImageType: types.ImageType,
				DeletedBy: fmt.Sprintf("%v", types.DeletedBy),
			},
		}

		res = append(res, obj)
	}
	return
}

func (ctrl Controller) GetSubTypeById(id uint) (res structuer.SubTypeRes, Error error) {
	subTypes, err := ctrl.Access.RDBMS.GetSubTypeById(id)
	if err != nil {
		Error = err
	}

	types := subTypes.Type

	res = structuer.SubTypeRes{
		ID:           fmt.Sprintf("%v", subTypes.ID),
		CreatedAt:    subTypes.CreatedAt,
		UpdatedAt:    subTypes.UpdatedAt,
		NameSubType:  subTypes.Name,
		ImageSubType: subTypes.ImageSubType,
		DeletedBy:    fmt.Sprintf("%v", subTypes.DeletedBy),
		TypeRes: &structuer.TypeRes{
			ID:        fmt.Sprintf("%v", types.ID),
			CreatedAt: types.CreatedAt,
			UpdatedAt: types.UpdatedAt,
			NameType:  types.Name,
			ImageType: types.ImageType,
			DeletedBy: fmt.Sprintf("%v", types.DeletedBy),
		},
	}

	return
}

func (ctrl Controller) PostSubType(req structuer.SubTypeReq) (Error error) {
	typeId, err := strconv.ParseUint(req.TypeId, 0, 0)
	if err != nil {
		Error = err
		return
	}

	data := rdbmsstructure.SubType{
		TypeID:       uint(typeId),
		Name:         req.NameSubType,
		ImageSubType: req.ImageSubType,
	}

	err = ctrl.Access.RDBMS.PostSubType(data)
	if err != nil {
		Error = err
	}

	return
}

func (ctrl Controller) PutSubType(id uint, req structuer.SubTypeReq) (Error error) {

	typeId, err := strconv.ParseUint(req.TypeId, 0, 0)
	if err != nil {
		Error = err
		return
	}
	data := rdbmsstructure.SubType{
		Model: gorm.Model{
			ID: id,
		},
		Name:         req.NameSubType,
		ImageSubType: req.ImageSubType,
		TypeID:       uint(typeId),
	}

	err = ctrl.Access.RDBMS.PutSubType(data)
	if err != nil {
		Error = err
		return
	}

	return
}

func (ctrl Controller) DeleteSubType(id, userId uint) (Error error) {
	data := rdbmsstructure.SubType{
		Model: gorm.Model{
			ID: id,
		},
		DeletedBy: userId,
	}

	err := ctrl.Access.RDBMS.PutSubType(data)
	if err != nil {
		Error = err
		return
	}

	err = ctrl.Access.RDBMS.DeleteSubType(id)
	if err != nil {
		Error = err
		return
	}

	return
}
