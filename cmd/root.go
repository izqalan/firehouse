package cmd

import (
	"fmt"
	"log"

	"github.com/izqalan/firehouse/cmd/auth"
	"github.com/izqalan/firehouse/cmd/config"
	"github.com/izqalan/firehouse/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "firehouse",
	Short: "A CLI tool for interacting with Firebase",
}

func init() {
	initService()
	initAuth()
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func loadFirebaseCredentials(cmd *cobra.Command, args []string) error {
	serviceAccountPath, err := cmd.Flags().GetString("service-account")
	if err != nil {
		return err
	}

	log.Print("Loading Firebase service account credentials from ", serviceAccountPath)
	opt, err := config.LoadCredentials(serviceAccountPath)
	if err != nil {
		return err
	}

	// use opt to save to config gile using viper
	viper.Set("service-account", opt)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Println("Error saving configuration:", err)
		return err
	}

	fmt.Println("Successfully loaded Firebase service account credentials")
	return nil
}

func getFirebaseCredentials(cmd *cobra.Command, args []string) error {
	credentials, err := config.GetCredentials()
	if err != nil {
		return err
	}

	// print credentials
	log.Println("Service Account Credentials")
	log.Println("Type:", credentials.AccountType)
	log.Println("Project ID:", credentials.ProjectID)
	log.Println("Client Email:", credentials.ClientEmail)
	log.Println("Client ID:", credentials.ClientID)
	log.Println("Auth URI:", credentials.AuthURI)
	log.Println("Token URI:", credentials.TokenURI)
	log.Println("Auth Provider X509 Cert URL:", credentials.AuthProviderX509CertURL)
	log.Println("Client X509 Cert URL:", credentials.ClientX509CertURL)
	log.Println("Universe Domain:", credentials.UniverseDomain)
	// log.Println("Private Key:", credentials.PrivateKey)

	return nil
}

func initService() {
	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Load Firebase service account credentials",
		// if flag --get run loadFirebaseCredentials
		RunE: loadFirebaseCredentials,
	}

	rootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().StringP("service-account", "s", "", "Path to the Firebase service account JSON file")
	serviceCmd.MarkFlagRequired("service-account")

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get Firebase service account credentials",
		Args:  cobra.ExactArgs(0),
		RunE:  getFirebaseCredentials,
	}

	serviceCmd.AddCommand(getCmd)

}

func initAuth() {

	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Interact with Firebase Authentication",
	}

	authCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// get all the flags
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			displayName, _ := cmd.Flags().GetString("display-name")
			photoURL, _ := cmd.Flags().GetString("photo-url")
			phoneNumber, _ := cmd.Flags().GetString("phone-number")
			disabled, _ := cmd.Flags().GetBool("disabled")
			emailVerified, _ := cmd.Flags().GetBool("email-verified")

			a := models.AuthOptions{
				Email:         email,
				Password:      password,
				DisplayName:   displayName,
				PhotoURL:      photoURL,
				PhoneNumber:   phoneNumber,
				Disabled:      disabled,
				EmailVerified: emailVerified,
			}

			// call the CreateUser function
			err := auth.CreateUser(a)
			if err != nil {
				return err
			}

			return nil
		},
	}

	authUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update user details",
		// take in 9 flags, --uid is compulsory
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// get all the flags
			uid, _ := cmd.Flags().GetString("uid")
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			displayName, _ := cmd.Flags().GetString("display-name")
			photoURL, _ := cmd.Flags().GetString("photo-url")
			phoneNumber, _ := cmd.Flags().GetString("phone-number")
			disabled, _ := cmd.Flags().GetBool("disabled")
			emailVerified, _ := cmd.Flags().GetBool("email-verified")

			a := models.AuthOptions{
				UID:           uid,
				Email:         email,
				Password:      password,
				DisplayName:   displayName,
				PhotoURL:      photoURL,
				PhoneNumber:   phoneNumber,
				Disabled:      disabled,
				EmailVerified: emailVerified,
			}

			// call the UpdateUserDetails function
			err := auth.UpdateUserDetails(a)
			if err != nil {
				return err
			}

			return nil
		},
	}

	authDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete user",
		Long:  "Delete a user from Firebase Authentication",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			uid, _ := cmd.Flags().GetString("uid")
			// prompt confirmation
			fmt.Printf("Are you sure you want to delete user %s? (y/n): ", uid)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" {
				fmt.Println("Operation cancelled")
				return nil
			}

			err := auth.DeleteUser(uid)
			if err != nil {
				return err
			}
			return nil
		},
	}

	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(authCreateCmd)
	authCreateCmd.Flags().StringP("email", "e", "", "Email")
	authCreateCmd.Flags().StringP("password", "p", "", "Password")
	authCreateCmd.Flags().StringP("display-name", "d", "", "Display Name")
	authCreateCmd.Flags().StringP("photo-url", "o", "", "Photo URL")
	authCreateCmd.Flags().StringP("phone-number", "n", "", "Phone Number")
	authCreateCmd.Flags().BoolP("disabled", "s", false, "Disabled")
	authCreateCmd.Flags().BoolP("email-verified", "v", false, "Email Verified")

	authCmd.AddCommand(authUpdateCmd)
	authUpdateCmd.Flags().StringP("uid", "u", "", "User ID")
	authUpdateCmd.Flags().StringP("email", "e", "", "Email")
	authUpdateCmd.Flags().StringP("password", "p", "", "Password")
	authUpdateCmd.Flags().StringP("display-name", "d", "", "Display Name")
	authUpdateCmd.Flags().StringP("photo-url", "o", "", "Photo URL")
	authUpdateCmd.Flags().StringP("phone-number", "n", "", "Phone Number")
	authUpdateCmd.Flags().BoolP("disabled", "s", false, "Disabled")
	authUpdateCmd.Flags().BoolP("email-verified", "v", false, "Email Verified")

	authCmd.AddCommand(authDeleteCmd)
	authDeleteCmd.Flags().StringP("uid", "u", "", "User ID")

}
