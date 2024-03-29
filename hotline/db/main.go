package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hotline/db/structure"
	"sync"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {
	GetHotLine() (response []structure.HotlineNumber, Error error)
	GetHotLineById(id uint) (response structure.HotlineNumber, Error error)
	PostHotLine(req structure.HotlineNumber) (Error error)
	PutHotLine(req structure.HotlineNumber) (Error error)
	DeleteHotLine(req structure.HotlineNumber) (Error error)

	PostHistory(req structure.History) (Error error)
	GetHistory() (res []structure.History, Error error)
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
	databaseSet := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		env.GormHost, env.GormPort, env.GormUser, env.GormName, env.GormPass, "disable")

	db, err := gorm.Open(postgres.Open(databaseSet), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
		//panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
	}

	if env.Flavor != Production {
		db = db.Debug()
	}

	_ = db.AutoMigrate(

		// hotline.go
		&structure.HotlineNumber{},
		&structure.History{},
	)

	var CheckHotlineNumber []structure.HotlineNumber
	db.Find(&CheckHotlineNumber)
	if len(CheckHotlineNumber) == 0 {
		hotline := structure.HotlineNumber{
			Number:           "191",
			Description:      "ตำรวจ",
			UserIDLogUpdated: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&hotline)
		hotline = structure.HotlineNumber{
			Number:           "1193",
			Description:      "ตำรวจทางหลวง",
			UserIDLogUpdated: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&hotline)
		hotline = structure.HotlineNumber{
			Number:           "1192",
			Description:      "ศูนย์รับแจ้งรถหาย",
			UserIDLogUpdated: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&hotline)
		hotline = structure.HotlineNumber{
			Number:           "1691",
			Description:      "โรงพยาบาลตำรวจ",
			UserIDLogUpdated: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&hotline)
		hotline = structure.HotlineNumber{
			Number:           "1155",
			Description:      "ตำรวจท่องเที่ยว",
			UserIDLogUpdated: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&hotline)
		hotline = structure.HotlineNumber{
			Number:           "1197",
			Description:      "ตำรวจจราจร",
			UserIDLogUpdated: 0,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&hotline)

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
	//GormHost string `env:"GORM_HOST,default=access"`
	GormHost string `env:"GORM_HOST,default=localhost"`
	//GormHost string `env:"GORM_HOST,default=access"`
	GormPort string `env:"GORM_PORT,default=5434"`
	GormName string `env:"GORM_NAME,default=postgresdb"`
	GormUser string `env:"GORM_USER,default=postgres"`
	GormPass string `env:"GORM_PASS,default=pgpassword"`
}
