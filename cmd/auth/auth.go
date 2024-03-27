package auth

import (
	"log"

	fauth "firebase.google.com/go/v4/auth"
	"github.com/izqalan/firehouse/models"
	"github.com/izqalan/firehouse/utils"
)

// Update user details Firebase Auth
// --uid string
// --email string
// --password string
// --display-name string
// --photo-url string
// --phone-number string
// --disabled bool
// --email-verified bool
// TODO: --custom-claims
// --custom-claims string expect key value pairs separated by comma

func UpdateUserDetails(a models.AuthOptions) error {

	log.Printf("Updating user details for %s", a.UID)
	// update user details

	// get firebase client
	client := utils.NewFirebaseClient()

	// get firebase auth client
	auth, err := client.Client.Auth(client.GetContext())
	if err != nil {
		return err
	}

	// only update the fields that are not empty
	params := map[string]interface{}{}

	// -e, --email
	if a.Email != "" {
		params["email"] = a.Email
	}
	// -p, --password
	if a.Password != "" {
		params["password"] = a.Password
	}
	// -d, --display-name
	if a.DisplayName != "" {
		params["displayName"] = a.DisplayName
	}
	// -u, --photo-url
	if a.PhotoURL != "" {
		params["photoURL"] = a.PhotoURL
	}
	// -n, --phone-number
	if a.PhoneNumber != "" {
		params["phoneNumber"] = a.PhoneNumber
	}
	// -d, --disabled
	if a.Disabled {
		params["disabled"] = a.Disabled
	}
	// -v, --email-verified
	if a.EmailVerified {
		params["emailVerified"] = a.EmailVerified
	}

	// convert params to UserToUpdate type
	userToUpdate := &fauth.UserToUpdate{}
	for key, value := range params {
		switch key {
		case "email":
			userToUpdate.Email(value.(string))
		case "password":
			userToUpdate.Password(value.(string))
		case "displayName":
			userToUpdate.DisplayName(value.(string))
		case "photoURL":
			userToUpdate.PhotoURL(value.(string))
		case "phoneNumber":
			userToUpdate.PhoneNumber(value.(string))
		case "disabled":
			userToUpdate.Disabled(value.(bool))
		case "emailVerified":
			userToUpdate.EmailVerified(value.(bool))
		}
	}

	// update user
	_, err = auth.UpdateUser(client.GetContext(), a.UID, userToUpdate)
	if err != nil {
		return err
	}

	log.Print("User details updated")

	return nil
}

func DeleteUser(uid string) error {
	log.Printf("Deleting user %s", uid)
	// get firebase client
	client := utils.NewFirebaseClient()

	// get firebase auth client
	auth, err := client.Client.Auth(client.GetContext())
	if err != nil {
		return err
	}

	// delete user
	err = auth.DeleteUser(client.GetContext(), uid)
	if err != nil {
		return err
	}

	log.Print("User deleted")

	return nil
}

func CreateUser(a models.AuthOptions) error {
	log.Printf("Creating user %s", a.Email)
	// get firebase client
	client := utils.NewFirebaseClient()

	// get firebase auth client
	auth, err := client.Client.Auth(client.GetContext())
	if err != nil {
		return err
	}

	// create user
	params := &fauth.UserToCreate{}

	if a.Email != "" {
		params.Email(a.Email)
	}

	if a.Password != "" {
		params.Password(a.Password)
	}

	if a.DisplayName != "" {
		params.DisplayName(a.DisplayName)
	}

	if a.PhotoURL != "" {
		params.PhotoURL(a.PhotoURL)
	}

	if a.PhoneNumber != "" {
		params.PhoneNumber(a.PhoneNumber)
	}

	params.Disabled(a.Disabled)
	params.EmailVerified(a.EmailVerified)

	_, err = auth.CreateUser(client.GetContext(), params)
	if err != nil {
		return err
	}

	log.Print("User created")

	return nil
}
