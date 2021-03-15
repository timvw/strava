package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	OauthAuthorizeUrl    = "https://www.strava.com/oauth/authorize"
	OauthTokenUrl        = "https://www.strava.com/oauth/token"
	AthleteActivitiesUrl = "https://www.strava.com/api/v3/athlete/activities"
)

func main() {

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

	publicHost, ok := os.LookupEnv("PUBLIC_HOST")
	if !ok {
		publicHost = fmt.Sprintf("http://localhost:%v", port)
	}

	requestTokenCallbackPath := "/callback"
	requestTokenUrlCallbackUrl := fmt.Sprintf("%v%v", publicHost, requestTokenCallbackPath)

	http.HandleFunc(requestTokenCallbackPath, logged(makeRequestTokenCallbackHandler(clientId, clientSecret)))
	http.HandleFunc("/", logged(makeIndexHandler(clientId, requestTokenUrlCallbackUrl)))

	fmt.Printf("Visit %v/ to view the demo\n", publicHost)
	fmt.Printf("ctrl-c to exit\n")
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func logged(h func(http.ResponseWriter,*http.Request))  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL)
		h(w, r)
	}
}

func makeIndexHandler(clientId int, requestTokenCallbackUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := "activity:read_all"
		requestCodeUrl := fmt.Sprintf("%v?client_id=%v&redirect_uri=%v&response_type=code&scope=%v", OauthAuthorizeUrl, clientId, requestTokenCallbackUrl, scope)
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

	resp, err := http.Post(OauthTokenUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("Post to %v failed: %v", OauthTokenUrl, err)
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

		//activities, err := getActivitiesPage(accessToken, 10, 1)
		activities, err := getActivities(accessToken)
		if err != nil {
			fmt.Fprintf(w, "Failed to get activities: %v", err)
		}
		activitiesAsString := renderActivitiesAsStringArrays(activities)

		w.Header().Add("Content-Type", "application/CSV")
		w.Header().Add("Content-Disposition", "attachment; filename=\"strava.csv\"")

		csvWriter := csv.NewWriter(w)
		defer csvWriter.Flush()
		err = csvWriter.WriteAll(activitiesAsString)

		if err != nil {
			fmt.Fprintf(w, "Failed to write csv: %v", err)
		}
	}
}

/*
GET
/athlete/activities
Parameters
before: Integer, in query	An epoch timestamp to use for filtering activities that have taken place before a certain time.
after: Integer, in query	An epoch timestamp to use for filtering activities that have taken place after a certain time.
page: Integer, in query	Page number. Defaults to 1.
per_page: Integer, in query	Number of items per page. Defaults to 30.
*/

func getActivities(accessToken string) ([]map[string]interface{}, error) {

	var activities []map[string]interface{}
	var page = 1

	for {
		actitivitiesPage, err := getActivitiesPage(accessToken, 200, page)
		if err != nil {
			return nil, fmt.Errorf("Failed to get page: %v", err)
		}
		if len(actitivitiesPage) == 0 {
			break
		}
		activities = append(activities, actitivitiesPage...)
		page = page + 1
	}

	return activities, nil
}

func getActivitiesPage(accessToken string, pageSize int, page int) ([]map[string]interface{}, error) {

	url := fmt.Sprintf("%v?per_page=%v&page=%v", AthleteActivitiesUrl, pageSize, page)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create http request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", accessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get activities page: %v", err)
	}

	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}

func renderActivitiesAsStringArrays(activities []map[string]interface{}) [][]string {

	fields := []string{
		"resource_state",
		"name",
		"distance",
		"moving_time",
		"elapsed_time",
		"total_elevation_gain",
		"type",
		"workout_type",
		"id",
		"external_id",
		"upload_id",
		"start_date",
		"start_date_local",
		"timezone",
		"utc_offset",
		"start_latlng",
		"end_latlng",
		"location_city",
		"location_state",
		"location_country",
		"achievement_count",
		"kudos_count",
		"comment_count",
		"athlete_count",
		"photo_count",
		"trainer",
		"commute",
		"manual",
		"private",
		"flagged",
		"gear_id",
		"from_accepted_tag",
		"average_speed",
		"max_speed",
		"average_cadence",
		"average_watts",
		"weighted_average_watts",
		"kilojoules",
		"device_watts",
		"has_heartrate",
		"average_heartrate",
		"max_heartrate",
		"max_watts",
		"pr_count",
		"total_photo_count",
		"has_kudoed",
		"suffer_score",
	}

	fieldAsString := func(item interface{}) string {
		if item != nil {
			return fmt.Sprintf("%v", item)
		} else {
			return "N/A"
		}
	}

	activityAsStringArray := func(activity map[string]interface{}) []string {
		var result []string
		for _, field := range fields {
			result = append(result, fieldAsString(activity[field]))
		}
		return result
	}

	var activitiesAsString [][]string
	activitiesAsString = append(activitiesAsString, fields)
	for _, activity := range activities {
		activitiesAsString = append(activitiesAsString, activityAsStringArray(activity))
	}

	return activitiesAsString
}
