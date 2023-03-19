package loggers

import (
	"accounts/constant"
	"fmt"
)

func LogProvider(httpStatus, APIName string, message interface{}) {
	fmt.Printf("\t<====== Response %#v API  ======>\n", APIName)
	fmt.Printf("Servicename : %#v\n", constant.ServiceName)
	fmt.Printf("HttpStatus : %#v \nMessage : %#v \n", httpStatus, message)
	fmt.Printf("\t<====== End APIName : %#v ======>\n", APIName)
}

func LogStart(APIName string) {
	fmt.Printf("\t<====== Open APIName : %#v ======>\n", APIName)
}
