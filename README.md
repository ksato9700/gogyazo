# gogyazo
Package gogyazo proivdes structs and functions for accessing Gyazo API.

## Command

## Environment Variables

- `OAUTH2_GYAZO_CLIENT_ID`
- `OAUTH2_GYAZO_CLIENT_SECRET`
- `OAUTH2_GYAZO_REDIRECT_URI`

## Usage

### Initial setup

1. Register an application at https://gyazo.com/oauth/applications
2. Set the client_id, client_secret and return_uri as above environment variables
3. Run `go run cmd/getAccessToken/main.go` to get the authentication URL
4. Open the URL with a web browser, sign-in if you didn't, authorizes the app, and you'll be redirect to ${return_uri}?code=${authcode}
5. Run `go run cmd/getAccessToken/main.go --code=${authcode}, then access token will be shown.

Once you follow the steps, you'll have `cache.json` file, in which access token is stored.

### Get image list

Run `go run cmd/getImages/main.go`



