package entity

import (
	// "testing"
	"time"
	// "github.com/asaskevich/govalidator"
	// "github.com/onsi/gomega"
)

type Customer struct{
	Name	string `valid:"required~Name cannot be blank"`
	DOB		time.Time
	Email	string
	Password	string `valid:"matches(^[B][0-9]{7}$)~อักษรตัวแรกต้องเป็น B ต่อมาอีก 7 ตัวเป็นตัวเลข 0-9"`
}