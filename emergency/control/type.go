package control

import (
	rdbmsstructure "emergency/db/structure"
	structuer "emergency/restapi/model"
	"fmt"
	"gorm.io/gorm"
)

func (ctrl Controller) GetType() (res []structuer.TypeRes, Error error) {
	data, err := ctrl.Access.RDBMS.GetType()
	if err != nil {
		Error = err
	}

	for _, m1 := range data {
		obj := structuer.TypeRes{
			ID:        fmt.Sprintf("%v", m1.ID),
			CreatedAt: m1.CreatedAt,
			UpdatedAt: m1.UpdatedAt,
			NameType:  m1.Name,
			ImageType: m1.ImageType,
			DeletedBy: fmt.Sprintf("%v", m1.DeletedBy),
		}
		res = append(res, obj)
	}
	return
}

func (ctrl Controller) GetTypeById(id uint) (res structuer.TypeRes, Error error) {
	types, err := ctrl.Access.RDBMS.GetTypeById(id)
	if err != nil {
		Error = err
	}
	subTypes, err := ctrl.Access.RDBMS.GetSubTypeByTypeId(id)

	var SubTypeRes []*structuer.SubTypeRes
	for _, subtype := range subTypes {
		obj := &structuer.SubTypeRes{
			ID:           fmt.Sprintf("%v", subtype.ID),
			CreatedAt:    subtype.CreatedAt,
			UpdatedAt:    subtype.UpdatedAt,
			NameSubType:  subtype.Name,
			ImageSubType: subtype.ImageSubType,
			DeletedBy:    fmt.Sprintf("%v", subtype.DeletedBy),
		}
		SubTypeRes = append(SubTypeRes, obj)
	}

	res = structuer.TypeRes{
		ID:         fmt.Sprintf("%v", types.ID),
		CreatedAt:  types.CreatedAt,
		UpdatedAt:  types.UpdatedAt,
		NameType:   types.Name,
		ImageType:  types.ImageType,
		DeletedBy:  fmt.Sprintf("%v", types.DeletedBy),
		SubTypeRes: SubTypeRes,
	}

	return
}

func (ctrl Controller) PostType(req structuer.TypeReq) (Error error) {
	data := rdbmsstructure.Type{
		Name:      req.NameType,
		ImageType: req.ImageType,
	}

	err := ctrl.Access.RDBMS.PostType(data)
	if err != nil {
		Error = err
	}

	return
}

func (ctrl Controller) PutType(id uint, req structuer.TypeReq) (Error error) {
	data := rdbmsstructure.Type{
		Model: gorm.Model{
			ID: id,
		},
		Name:      req.NameType,
		ImageType: req.ImageType,
	}

	err := ctrl.Access.RDBMS.PutType(data)
	if err != nil {
		Error = err
		return
	}

	return
}

func (ctrl Controller) DeleteType(id, userId uint) (Error error) {
	data := rdbmsstructure.Type{
		Model: gorm.Model{
			ID: id,
		},
		DeletedBy: userId,
	}

	err := ctrl.Access.RDBMS.PutType(data)
	if err != nil {
		Error = err
		return
	}

	err = ctrl.Access.RDBMS.DeleteType(id)
	if err != nil {
		Error = err
		return
	}

	return
}
