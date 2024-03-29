package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"messenger/db/structure"
	"sync"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {
	GetMessage(roomChatId uint) (res []structure.Message, Error error)
	GetImageByMessageId(messageId uint) (res structure.Message, Error error)
	PostChat(message structure.Message) (Error error)
	PutChat(message structure.Message) (Error error)
	DeleteChat(messageId uint) (Error error)

	GetRoomChatListByUserId(UserID uint) (res []structure.GroupChat, Error error)
	GetRoomChatById(roomChatId uint) (res structure.RoomChat, Error error)

	GetMembersRoomChat(RoomChatID uint) (res []structure.GroupChat, Error error)
	CheckRoomChatUser(RoomChatID, UserID uint) (res structure.GroupChat, Error error)
	RoomChat(groupChat structure.GroupChat) (res structure.GroupChat, Error error)
	JoinChat(groupChat structure.GroupChat) (Error error)
	PutRoomChat(groupChat structure.RoomChat) (Error error)
	DeleteRoomChatById(roomChatId uint) (Error error)

	GetAllForAdminChatList() (res []structure.GroupChat, Error error)
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
		//chat
		&structure.RoomChat{},
		&structure.GroupChat{},
		&structure.Message{},
	)

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
	GormPort string `env:"GORM_PORT,default=5435"`
	GormName string `env:"GORM_NAME,default=postgresdb"`
	GormUser string `env:"GORM_USER,default=postgres"`
	GormPass string `env:"GORM_PASS,default=pgpassword"`
}
