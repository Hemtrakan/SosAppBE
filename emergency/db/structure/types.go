package structure

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name      string
	DeletedBy uint
}

type SubType struct {
	gorm.Model
	Name      string
	Type      Type
	TypeID    uint
	DeletedBy uint
}
