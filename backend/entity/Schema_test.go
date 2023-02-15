package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerValidate (t *testing.T){
	g := NewGomegaWithT(t)

	user := Customer{
		Name: "",
		DOB: time.Now(),
		Email: "aha0037@gmail.com",
		Password: "B2312312",
	}

	ok, err := govalidator.ValidateStruct(user)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}


func TestCustomerPasswordValidate (t *testing.T){
	g := NewGomegaWithT(t)

	user := Customer{
		Name: "Natthawat",
		DOB: time.Now(),
		Email: "aha0037@gmail.com",
		Password: "V1234567",
	}

	ok, err := govalidator.ValidateStruct(user)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("อักษรตัวแรกต้องเป็น B ต่อมาอีก 7 ตัวเป็นตัวเลข 0-9"))
}