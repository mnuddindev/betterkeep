package utils

import (
	"crypto/rand"
	"log"
	"math/big"
	"reflect"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
)

func IsEmpty(data interface{}) bool {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface())
	}
}

func IsEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func CheckError(err error, message string) {
	if err != nil {
		log.Println(err)
	}
}

func HashPassword(password string) (string, error) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	CheckError(err, "failed to hash password")
	return string(hpass), err
}

func ComparePass(hashedPass, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	CheckError(err, "failed to compare passwords")
	return err
}

func GenerateOTP() (int64, error) {
	max := big.NewInt(999999)
	numb, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return numb.Int64(), nil
}

func ActiveUser(code, email, username string) {
	link := "http://localhost:4000/active-user"
	messBody := "Hello " + username + ", \n Your Activation code is " + code + " \n\n Active your Account By clicking on <a href='" + link + "'>this</a> link"

	mail := gomail.NewMessage()
	mail.SetHeader("From", "support@inadislam.com")
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Activate your account")
	mail.SetBody("text/html", messBody)
	dialer := gomail.NewDialer("0.0.0.0", 1025, "", "")
	if err := dialer.DialAndSend(mail); err != nil {
		panic(err)
	}
}
