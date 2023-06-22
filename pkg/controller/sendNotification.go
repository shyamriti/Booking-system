package controller

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendBookingNotification(recipient, msg string) {
	accountSid := "AC6e5bbbde4e1e3c5266228e175c5d2335"
	authToken := "54d8b929e428d6bb619c42bf48089300"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(recipient)
	params.SetFrom("+12179552661")
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
