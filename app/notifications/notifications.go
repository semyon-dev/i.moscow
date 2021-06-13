package notifications

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
	"i-moscow-backend/app/config"
	"log"
)

var App *firebase.App

func Init() {
	var err error
	opt := option.WithCredentialsFile(config.FireBaseFileName)
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
	}
}

func Send(registrationToken, title, body string) {
	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n\n", err)
	}
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: registrationToken,
	}
	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Println("Unsuccessfully sent message:", err)
		return
	}
	log.Println("Successfully sent message:", response)
}
