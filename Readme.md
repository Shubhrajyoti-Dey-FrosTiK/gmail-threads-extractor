# Mails Label Extractor

This is a simple Google Cloud Function which extracts the threads of your email addresses and sends it back as an API Call response.

# Prerequisites

1. Golang installed on system
2. Make a GCP Account
3. Create OAuth2 Credentials from [here](https://console.cloud.google.com/apis/credentials/oauthclient)
4. Now you need a refresh token.
   The easiest way to get te refresh token is from [here](https://developers.google.com/oauthplayground/).
   Select `GMail API V1`.
   Select `https://mail.google.com/`
   Select `https://www.googleapis.com/auth/gmail.readonly`
   These will give you readonly rights
   Now select `Authorize` and choose your account.
   Then exchange the Authorization Code with a refresh token.
5. Install `Make` in the system. (Optional and only required for the last step in setup)
6. Install `direnv` in the system

# Setup

1. Clone the repo

```
git clone https://github.com/Shubhrajyoti-Dey-FrosTiK/gmail-labels-extractor.git
```

2. Go inside the directory

```
cd gmail-labels-extractor
```

3. Create a `.envrc` file

```
touch .envrc
```

4. Put the credentials in the `.env` file like this

```
export $OATH_CLIENT_ID=...
export $OATH_PROJECT_ID=...
export $OATH_CLIENT_SECRET=...
export $OATH_REDIRECT_URI=http://localhost
export $REFRESH_TOKEN=...
```

5. Start the server by running this commande

```
make
```

# Walkthrough

`cmd/main.go` is the starting point of the application where the function server will be initialized.

`makefile` will set an env variable named `TARGET_FUNCTION` to the name of the path which we are calling the function. Suppose the URL is `http://..../api` then the function will be `api` here. After setting this we are starting the server by running the main.go file

In `cmd/main.go` we also see a blank import.This will import `functions.go` and will look for the init function.

In `functions.go` we define all the functions and we call the `labels` functions.

After you have started the server with the `make` command you can hit `http://localhost8080/labels` and check the output.

$ Endpoints

`/threads` extracts the recent list of threads. Here optionally you can provide `thread_page` which is the `id` of the thread page and it will automatically detect it and fetch it.
`/threads/from` extracts the list of threads provided `from` in the header
`/threads/label` extracts the list of threads provided `label` in the header
`threads/id` extracts the thread provided `thread_id` in the header

# Deployment

In order to deploy just run this command

```
gcloud functions deploy gmail-labels-extractor --runtime go119 --trigger-http --allow-unauthenticated --region=asia-south1 --gen2 --entry-point labels
```

Or just run

```
bash deploy.sh
```
