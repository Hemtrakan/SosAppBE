package logs

import (
	"log"
	"time"
)

//var times = time.Now().Add(time.Minute * 3).Add(time.Hour * 7)

func LogStart(nameAPI string) {
	log.Println("---------------------- Start ---------------------")
	log.Printf("Start API : %s , Time : %v", nameAPI, time.Now().Add(time.Hour*7))
}

func LogRequest(req interface{}) {
	log.Println("")
	log.Println("--------- Start Log Request -----------")
	log.Printf("Request : %v", req)
	log.Println(" -------- End Log Request -------------")
}

func LogResponse(res interface{}) {
	log.Println("")
	log.Println("--------- Start Log Response ---------")
	log.Printf("Response : %v", res)
	log.Printf("Time : %v", time.Now().Add(time.Hour*7))
	log.Println("--------- End Log Response -----------")
	log.Println("------------------------ End ----------------------")
}

func LogError(err error) {
	log.Println("")
	log.Println(" -------- Start Log Error ---------")
	log.Printf("Error : %v", err.Error())
	log.Printf("Time : %v", time.Now().Add(time.Hour*7))
	log.Println("--------- End Log Error -----------")
	log.Println("------------------------ End ----------------------")
}
