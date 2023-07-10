package trawlingwebgo

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// TrwArticle API Structure
type TrwArticle struct {
	ID                   string  `json:"id"`
	Hash                 string  `json:"hash"`
	Published            string  `json:"published"`
	Crawled              int64   `json:"crawled"`
	Updated              int64   `json:"updated"`
	PostID               string  `json:"post_id"`
	URL                  string  `json:"url"`
	Text                 string  `json:"text"`
	Lang                 string  `json:"lang"`
	RetweetCount         int64   `json:"retweet_count"`
	ReplyCount           int64   `json:"reply_count"`
	FavoriteCount        int64   `json:"favorite_count"`
	ReproductionsCount   int64   `json:"reproductions_count"`
	EntitiesURL          string  `json:"entities_url"`
	URLImage             string  `json:"url_image"`
	Hashtags             string  `json:"hashtags"`
	UserMentions         string  `json:"user_mentions"`
	TimeDistance         float64 `json:"time_distance"`
	Reply                bool    `json:"reply"`
	UserName             string  `json:"user_name"`
	UserScreenName       string  `json:"user_screen_name"`
	UserCreationDate     string  `json:"user_creation_date"`
	UserURL              string  `json:"user_url"`
	UserProfileImageURL  string  `json:"user_profile_image_url"`
	UserProfileBannerURL string  `json:"user_profile_banner_url"`
	UserDescription      string  `json:"user_description"`
	UserExternalURL      string  `json:"user_external_url"`
	UserLocation         string  `json:"user_location"`
	UserFollowerCount    int64   `json:"user_follower_count"`
	UserFollowingCount   int64   `json:"user_following_count"`
	UserFavouritesCount  int64   `json:"user_favourites_count"`
	UserIsPrivate        bool    `json:"user_is_private"`
	UserIsVerified       bool    `json:"user_is_verified"`
	UserIsBlueVerified   bool    `json:"user_is_blue_verified"`
	UserNumberOfTweets   int64   `json:"user_number_of_tweets"`
}

// TrwResponse API structure
type TrwResponse struct {
	Response struct {
		Data         []TrwArticle `json:"data"`
		RequestLeft  int          `json:"requestLeft"`
		TotalResults int          `json:"totalResults"`
		RestResults  int          `json:"restResults"`
		Next         string       `json:"next"`
	} `json:"response"`
}

// TrwRequest API structure
type TrwRequest struct {
	Token string
	Query string
	Ts    string
	Tsi   string
	Sort  string
	Order string
}

// TrwError for get problems
type TrwError struct {
	Response struct {
		Error string `json:"error"`
	} `json:"response"`
}

// Request to https service
func Request(url string) (TrwResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	var res TrwResponse
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", "trawlingweb-cli.go 1.2")
	resp, err2 := client.Do(req)
	if err2 != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	err3 := json.NewDecoder(resp.Body).Decode(&res)
	return res, err3
}

// Query Initial function
// Query Initial function
func Query(params TrwRequest) (TrwResponse, error) {
	values := reflect.ValueOf(params)
	twurl := "https://twitter.trawlingweb.com/posts_full/?"
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() != "" {
			if i != 0 {
				twurl += "&"
			}
			if values.Type().Field(i).Name == "Query" {
				sturl := values.Field(i).String()
				encodedPath := url.QueryEscape(sturl)
				twurl += "q=" + encodedPath
			} else {
				twurl += strings.ToLower(values.Type().Field(i).Name) + "=" + values.Field(i).String()
			}
		}
	}

	return Request(twurl)
}

// Next query function
func Next(twurl string) (TrwResponse, error) {
	return Request(twurl)
}
