package loggers

import "fmt"

func LogProvider(httpStatus, APIName string, message interface{}) {
	fmt.Printf("\t<====== Open APIName : %#v ======>\n", APIName)
	fmt.Printf("HttpStatus : %#v \nMessage : %#v \n", httpStatus, message)
	fmt.Printf("\t<====== End APIName : %#v ======>\n", APIName)
}
