package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const (
	RequestCodeUrl        = "https://www.strava.com/oauth/authorize"
	RequestAccessTokenUrl = "https://www.strava.com/oauth/token"
)

func main() {
	fmt.Println("Hello, world.")

	port := 9000

	clientIdString, ok := os.LookupEnv("STRAVA_CLIENT_ID")
	if !ok {
		fmt.Println("Failed to find 'STRAVA_CLIENT_ID' environment variable")
		return
	}
	clientId64, err := strconv.ParseInt(clientIdString, 10, 32)
	clientId := int(clientId64)

	if err != nil {
		fmt.Println("Provided 'STRAVA_CLIENT_ID' environment variable, '%v' is not an int: %v", clientIdString, err)
		return
	}

	clientSecret, ok := os.LookupEnv("STRAVA_CLIENT_SECRET")
	if !ok {
		fmt.Println("Failed to find 'STRAVA_CLIENT_SECRET' environment variable")
		return
	}

	requestTokenCallbackPath := "/callback"
	requestTokenUrlCallbackUrl := fmt.Sprintf("http://localhost:%d%v", port, requestTokenCallbackPath)

	http.HandleFunc(requestTokenCallbackPath, makeRequestTokenCallbackHandler(clientId, clientSecret))
	http.HandleFunc("/", makeIndexHandler(clientId, requestTokenUrlCallbackUrl))

	fmt.Printf("Visit http://localhost:%d/ to view the demo\n", port)
	fmt.Printf("ctrl-c to exit")
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func requestCodeUrl(clientId int, redirectUri string) string {
	scope := "activity:read_all"
	return fmt.Sprintf("%v?client_id=%v&redirect_uri=%v&response_type=code&scope=%v", RequestCodeUrl, clientId, redirectUri, scope)
}

func makeIndexHandler(clientId int, requestTokenCallbackUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := "activity:read_all"
		requestCodeUrl := fmt.Sprintf("%v?client_id=%v&redirect_uri=%v&response_type=code&scope=%v", RequestCodeUrl, clientId, requestTokenCallbackUrl, scope)
		fmt.Fprintf(w, `<a href="%s">`, requestCodeUrl)
		fmt.Fprint(w, `<img src="http://strava.github.io/api/images/ConnectWithStrava.png" />`)
		fmt.Fprint(w, `export activities`)
		fmt.Fprint(w, `</a>`)
	}
}

func requestAccessToken(clientId int, clientSecret string, code string) (string, error) {

	payload, err := json.Marshal(map[string]string{
		"client_id":     strconv.Itoa(clientId),
		"client_secret": clientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"f":             "json",
	})

	if err != nil {
		return "", fmt.Errorf("Failed to marshal payload: %v", err)
	}

	resp, err := http.Post(RequestAccessTokenUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("Post to %v failed: %v", RequestAccessTokenUrl, err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	accessToken := fmt.Sprintf("%v", result["access_token"])
	return accessToken, nil
}

func makeRequestTokenCallbackHandler(clientId int, clientSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			fmt.Fprintf(w, "Did not recieve code parameter")
		}

		accessToken, err := requestAccessToken(clientId, clientSecret, code)
		fmt.Fprintf(w, "received access token: %v , err: %v", accessToken, err)

	}
}
