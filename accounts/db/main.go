package db

import (
	"accounts/db/structure"
	"accounts/utility/verify"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	factory FactoryInterface
	once    sync.Once
)

type FactoryInterface interface {

	// LogLogin
	LogLogin(req structure.Users) (Error error)
	//Role
	GetRoleList() (response []structure.Role, Error error)
	GetRoleByName(req structure.Role) (response structure.Role, Error error)
	GetRoleById(req structure.Role) (response structure.Role, Error error)
	AddRole(req structure.Role) (Error error)

	// OTP
	SendOTPDB(req structure.OTP) (Error error)
	GetOTPDb(req structure.OTP) (response *structure.OTP, Error error)
	UpdateOTPDB(req structure.OTP) (Error error)

	// Users
	GetUserByPhone(req structure.Users) (response *structure.Users, Error error)
	GetUserByID(req structure.Users) (response *structure.Users, Error error)
	GetUserList() (response []*structure.Users, Error error)
	PostUser(req structure.Users) (Error error)
	PutUser(user *structure.Users, address *structure.Address, idCard *structure.IDCard) (Error []error)
	ChangePassword(req *structure.Users) (Error error)
	DeleteUser(req structure.Users) (Error error)
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
		// Account
		structure.Users{},
		structure.Role{},
		structure.IDCard{},
		structure.Address{},
		structure.OTP{},
		structure.LogLogin{},
	)

	var CheckRole []structure.Role
	db.Find(&CheckRole)
	if len(CheckRole) == 0 {
		dataAdmin := structure.Role{
			Name: "admin",
		}
		dataUser := structure.Role{
			Name: "user",
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dataAdmin)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dataUser)

		role := structure.Role{}
		address := structure.Address{
			Address:     "",
			SubDistrict: "",
			District:    "",
			Province:    "",
			PostalCode:  "",
			Country:     "",
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&address)
		db.Where("name = ?", "admin").Take(&role)

		IDCard := structure.IDCard{
			TextIDCard: "13xxxxxxxx347",
			PathImage:  "",
			DeletedBy:  nil,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&IDCard)

		Password, _ := verify.Hash("admin")
		data := structure.Users{
			PhoneNumber:  "admin",
			Password:     string(Password),
			Firstname:    "FirstNameAdmin",
			Lastname:     "LastNameAdmin",
			Email:        "",
			Birthday:     time.Now(),
			Gender:       "M",
			IDCardID:     IDCard.ID,
			ImageProfile: nil,
			DeletedBy:    nil,
			Workplace:    nil,
			AddressID:    address.ID,
			RoleID:       role.ID,
		}
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data)
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
	GormPort string `env:"GORM_PORT,default=5432"`
	GormName string `env:"GORM_NAME,default=postgresdb"`
	GormUser string `env:"GORM_USER,default=postgres"`
	GormPass string `env:"GORM_PASS,default=pgpassword"`
}
