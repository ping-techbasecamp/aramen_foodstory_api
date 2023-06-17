package mailbit

import (
	"aramen-api/pkg/config"
	"fmt"
	"math/rand"
	"net/http"
)

const otpChars = "1234567890"
const refChars = "ABCDEFGHJKLMNOPQRSTUVWXYZ"

func Generate(length int, ref bool) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	var CharsLength int
	if ref {
		CharsLength = 26
	}
	CharsLength = 10
	for i := 0; i < length; i++ {
		if ref {
			buffer[i] = refChars[rand.Intn(100)%CharsLength]
		} else {
			buffer[i] = otpChars[int(buffer[i])%CharsLength]
		}
	}
	return string(buffer), nil
}

func SendOtp(ref, otp, phoneNumber string, appConfig *config.AppConfig) error {
	message := fmt.Sprintf("%v+คือรหัส+OTP+ของคุณ+(Ref:%v)", otp, ref)
	statement := fmt.Sprintf("http://sms.mailbit.co.th/vendorsms/pushsms.aspx?user=%v&password=%v&msisdn=%v&sid=%v&msg=%v&fl=0&dc=8", appConfig.Mailbit.Username, appConfig.Mailbit.Password, phoneNumber, appConfig.Mailbit.SenderName, message)
	fmt.Println(statement)
	_, err := http.Get(statement)
	if err != nil {
		return err
	}
	return nil
}
