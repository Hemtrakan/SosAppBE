package constant

const (
	ServiceName             string = "emergency"
	SuccessMsg              string = "Success"
	SuccessCode             string = "0"
	ErrorCode               string = "400"
	ErrorStatusUnauthorized string = "401"
)

const (
	Required string = "required"
	Min      string = "min"
	Max      string = "max"
	Length   string = "len"
	Numeric  string = "numeric"
)

const (
	User  string = "user"
	Admin string = "admin"
	Ops   string = "ops"
)

type Status int

const (
	StatusStep1 = 1
	StatusStep2 = 2
	StatusStep3 = 3
	StatusStep4 = 4
)

var StatusData = []Status{
	StatusStep1,
	StatusStep2,
	StatusStep3,
	StatusStep4,
}

func (status Status) Status() (result *string, Errors error) {
	switch status {
	case StatusStep1:
		fullName := "แจ้งเหตุเรียบร้อย"
		result = &fullName
	case StatusStep2:
		fullName := "รับเรื่องการแจ้งเหตุแล้ว"
		result = &fullName
	case StatusStep3:
		fullName := "กำลังดำเนินงาน"
		result = &fullName
	case StatusStep4:
		fullName := "ได้รับการแก้ไขเรียบร้อนแล้ว"
		result = &fullName
	default:
		fullName := "ยังไม่ได้รับการดำเนินการ"
		result = &fullName
	}
	return
}
