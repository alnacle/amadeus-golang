package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Client is a core class
type Client struct {
	ClientID     string
	ClientSecret string
	UserAgent    string
	BaseURL      *url.URL
	httpClient   http.Client
	Debug        bool
}

// NewClient creates a new ClIent instance
func NewClient(clientID string, clientSecret string, host string) *Client {

	baseURL, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		UserAgent:    "Amadeus Golang SDK",
		BaseURL:      baseURL,
		Debug:        false,
	}
}

// Get makes an authenticated GET API call.
func (c *Client) Get(path string, params ...string) ([]byte, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	// query parameters
	v := url.Values{}

	for _, param := range params {
		sp := strings.Split(param, "=")
		v.Add(sp[0], sp[1])
	}

	u.RawQuery = v.Encode()
	fmt.Println(u)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer USXj0NokGqCP38JFidHrnMGdSvZi")

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&users)
	return b, nil
}
