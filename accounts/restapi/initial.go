package restapi

import (
	"accounts/control"
	"accounts/db"
	"accounts/httpclient"
	"errors"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/go-playground/validator"
	config "github.com/spf13/viper"
	"os"
	"time"
)

type Controller struct {
	Properties db.Properties
	Access     db.Access
	Ctx        control.Controller
	HttpClient httpclient.HttpClient
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

func ConController(db *db.Access) *control.Controller {
	res := control.Controller{
		Access: db,
	}
	return &res
}

func NewHttpClient() (httpClient httpclient.HttpClient) {
	httpClient = httpclient.HttpClient{
		Charset:        "utf-8",
		CertSkipVerify: true,
		Timeout:        10 * time.Second,
	}
	return
}

func NewController() Controller {
	build := Build()
	access := Initial(build)
	ctx := ConController(access)
	httpclient := NewHttpClient()
	return Controller{
		Properties: *build,
		Access:     *access,
		Ctx:        *ctx,
		HttpClient: httpclient,
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
