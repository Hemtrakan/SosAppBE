package restapi

import (
	"errors"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/control"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/db"
	"github.com/go-playground/validator"
	config "github.com/spf13/viper"
	"os"
)

type Controller struct {
	Properties db.Properties
	Access     db.Access
	Ctx        control.ConController
}

func Build() *db.Properties {
	var prop db.Properties
	if _, err := env.UnmarshalFromEnviron(&prop); err != nil {
		panic(err)
	}
	return &prop
}

func Initial(properties *db.Properties) *db.Access {
	return &db.Access{
		ENV:   properties,
		RDBMS: db.Create(properties),
		//GRPC: grpc.Create(properties),
	}
}

func ConController(db *db.Access) *control.ConController {
	res := control.ConController{
		Access: db,
	}
	return &res
}

func NewController() Controller {
	build := Build()
	access := Initial(build)
	ctx := ConController(access)
	return Controller{
		Properties: *build,
		Access:     *access,
		Ctx:        *ctx,
	}
}

func (ctrl Controller) LoadConfigFile() error {
	env := os.Getenv("ENV")
	env = "dev"
	if env == "" {
		env = os.Args[1]
	}

	//ctrl.Logger.Info(transID, fmt.Sprintf("Server start running on %s environment configuration", env))
	config.SetConfigName(env)
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")
	err := config.ReadInConfig()
	if err != nil {
		errMsg := fmt.Sprintf("Read config file %s.yml occur error: %s", env, err.Error())
		panic(errMsg)
		return err
	}
	return err
}

func ValidateStruct(dataStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("%s: %s", err.StructField(), err.Tag()))
		}
	} else {
		return nil
	}
	return err
}
