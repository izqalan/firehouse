package utils

// write a firbase client that can be used globally

import (
	"context"
	"encoding/json"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/izqalan/firehouse/cmd/config"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	Client  *firebase.App
	Context context.Context
}

func NewFirebaseClient() *FirebaseClient {
	ctx := context.Background()

	credentials, err := config.GetCredentials()
	if err != nil {
		log.Fatal(err)
	}

	// parse credentials to json byte
	bytes, err := json.Marshal(credentials)
	if err != nil {
		log.Fatal(err)
	}

	sa := option.WithCredentialsJSON(bytes)
	client, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatal(err)
	}

	return &FirebaseClient{Client: client, Context: ctx}
}

func (f *FirebaseClient) GetClient() *firebase.App {
	return f.Client
}

func (f *FirebaseClient) GetContext() context.Context {
	return f.Context
}
