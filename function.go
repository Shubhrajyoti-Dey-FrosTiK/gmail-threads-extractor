package function

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

/* ------- Interfaces --------- */
type InstalledOAuthClient struct {
	Client_id						string	`json:"client_id"`
	Project_id						string	`json:"project_id"`
	Auth_uri						string	`json:"auth_uri"`
	Token_uri						string	`json:"token_uri"`
	Auth_provider_x509_cert_url		string	`json:"auth_provider_x509_cert_url"`
	Client_secret					string	`json:"client_secret"`
	Redirect_uris					[]string `json:"redirect_uris"`
}

type OAthCredentials struct {
	Installed	InstalledOAuthClient	`json:"installed"`
}

/*------------Helper Functions -----------*/
// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) http.Client {
	godotenv.Load()
	var token oauth2.Token

	token.TokenType = "authorized_user"
	token.RefreshToken = os.Getenv("REFRESH_TOKEN")

	return *config.Client(context.Background(), &token)
}


func getCredentialBytes() []byte {
	godotenv.Load()
	var data  = OAthCredentials{
		Installed:	 InstalledOAuthClient{
			Client_id: os.Getenv("OATH_CLIENT_ID"),
			Project_id: os.Getenv("OATH_PROJECT_ID"),
			Auth_uri: "https://accounts.google.com/o/oauth2/auth",
			Token_uri: "https://oauth2.googleapis.com/token",
			Auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs",
			Client_secret: os.Getenv("OATH_CLIENT_SECRET"),
			Redirect_uris: []string{os.Getenv("OATH_REDIRECT_URI")},
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
 
	return file
}

/*------------- Initialization ---------*/
func init() {
	// functions.HTTP("HelloWorld", helloWorld)
	functions.HTTP("threads",fetchMailLabels)
}

/*------------- Global Declarations ---------*/
var ctx = context.Background()
var config, _ = google.ConfigFromJSON(getCredentialBytes(), gmail.GmailReadonlyScope)
var client = getClient(config)
var gmailClient, _ = gmail.NewService(ctx, option.WithHTTPClient(&client))

/*-------------Handler Functions-------*/
func fetchMailLabels(writer http.ResponseWriter, request *http.Request){
	user := "me"
	
	path := request.URL.Path
	fmt.Print(path)

	switch path {
	case "/threads" : 
		threadPageId := request.Header.Get("thread_page")
		threadsList, _ := gmailClient.Users.Threads.List(user).PageToken(threadPageId).Do()
		out, _ := json.Marshal(threadsList)
		fmt.Fprintf(writer,string(out))

	case "/threads/from" :
		from := request.Header.Get("from")
		fmt.Print("from" + from)
		threadsList, _ := gmailClient.Users.Threads.List(user).Q("from: " + from).Do()
		out, _ := json.Marshal(threadsList)
		fmt.Fprintf(writer,string(out))

	case "/threads/label" :
		label := request.Header.Get("label")
		fmt.Print("from" + label)
		threadsList, _ := gmailClient.Users.Threads.List(user).Q("label: " + label).Do()
		out, _ := json.Marshal(threadsList)
		fmt.Fprintf(writer,string(out))

	case "/threads/id" :
		threadId := request.Header.Get("thread_id")
		thread, _ := gmailClient.Users.Threads.Get("me",threadId).Do()
		out, _ := json.Marshal(thread)
		fmt.Fprintf(writer,string(out))
	}
}