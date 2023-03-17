package db

import (
	"emergency/db/structure"
	"emergency/db/structure/responsedb"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {
	GetImageByInformId(informId uint) (response *responsedb.InformInfoById, Error error)
	GetInformList(UserId uint) (response []*responsedb.InformInfoList, Error error)
	PostInform(imageArr []structure.InformImage, inform structure.Inform) (Error error)
}

func Create(env *Properties) FactoryInterface {
	once.Do(func() {
		factory = gormInstance(env)
	})
	return factory
}

type GORMFactory struct {
	env    *Properties
	client *gorm.DB
}

func gormInstance(env *Properties) GORMFactory {
	databaseSet := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		env.GormHost, env.GormPort, env.GormUser, env.GormName, env.GormPass, "disable", "Asia/Bangkok")

	db, err := gorm.Open(postgres.Open(databaseSet), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
	}

	if env.Flavor != Production {
		db = db.Debug()
	}

	_ = db.AutoMigrate(

		//Inform
		structure.Type{},
		structure.SubType{},
		structure.Inform{},
		structure.InformImage{},
		structure.InformNotification{},
	)

	var typeStructureArr []structure.Type
	db.Find(&typeStructureArr)
	if len(typeStructureArr) == 0 {
		Type1 := structure.Type{
			Name:      "โรงพยาบาล",
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Type1)

		Type2 := structure.Type{
			Name:      "ปอเต็กตึ๊ง",
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Type2)

		Type3 := structure.Type{
			Name:      "สถานีดับเพลิง",
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Type3)

		Type4 := structure.Type{
			Name:      "สถานีตำรวจ",
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Type4)

		subType1 := structure.SubType{
			Name:      "เจ็บป่วย",
			TypeID:    Type1.ID,
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&subType1)

		subType2 := structure.SubType{
			Name:      "อุบัติเหตุ",
			TypeID:    Type2.ID,
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&subType2)

		subType3 := structure.SubType{
			Name:      "อาคาร/สถานที่",
			TypeID:    Type3.ID,
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&subType3)

		subType4 := structure.SubType{
			Name:      "อื่น",
			TypeID:    Type4.ID,
			DeletedBy: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&subType4)

	}
	return GORMFactory{env: env, client: db}
}

type Access struct {
	ENV   *Properties
	RDBMS FactoryInterface
}

type Flavor string
type URL string

const (
	Develop    Flavor = "DEVELOP"
	Devspace   Flavor = "DEVSPACE"
	Production Flavor = "PRODUCTION"
)

type Properties struct {
	// -- core
	Flavor Flavor `env:"FLAVOR,default=DEVELOP"`
	// --

	// -- Gorm
	//GormHost string `env:"GORM_HOST,default=emergency-rdbms"`
	GormHost string `env:"GORM_HOST,default=localhost"`
	GormPort string `env:"GORM_PORT,default=5433"`
	GormName string `env:"GORM_NAME,default=postgresdb"`
	GormUser string `env:"GORM_USER,default=postgres"`
	GormPass string `env:"GORM_PASS,default=pgpassword"`
}
