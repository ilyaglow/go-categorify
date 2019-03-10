// Package categorify simplifies interaction with Categorify API.
// More about Categorify: https://categorify.org.
package categorify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiLocation = "https://categorify.org/api?website=%s"

// Categorify defines configuration to make queries.
type Categorify struct {
	Client *http.Client
}

// New bootstraps categorify configuration with a default http client.
func New() *Categorify {
	return &Categorify{
		Client: http.DefaultClient,
	}
}

// NewWithClient bootstraps categorify with a specified http client.
func NewWithClient(client *http.Client) *Categorify {
	return &Categorify{
		Client: client,
	}
}

// Report represents categorify's website report.
type Report struct {
	Domain      string `json:"domain,omitempty"`
	IP          string `json:"ip,omitempty"`
	CountryCode string `json:"country-code,omitempty"`
	Country     string `json:"country,omitempty"`
	Rating      struct {
		Language    bool   `json:"language,omitempty"`
		Violence    bool   `json:"violence,omitempty"`
		Nudity      bool   `json:"nudity,omitempty"`
		Adult       bool   `json:"adult,omitempty"`
		Value       string `json:"value,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"rating"`
	Confidence     string         `json:"confidence_level,omitempty"`
	Categories     []string       `json:"category,omitempty"`
	KeywordHeatmap map[string]int `json:"keyword_heatmap,omitempty"`
}

type respError struct {
	Result string `json:"result,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type getFunc func(string) (*http.Response, error)

func lookup(location string, f getFunc) (*Report, error) {
	resp, err := f(location)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var e respError
		err = dec.Decode(&e)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("result: %s, reason: %s", e.Result, e.Reason)
	}

	var r Report
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Lookup gets categorify report with default http client.
func Lookup(domain string) (*Report, error) {
	return lookup(fmt.Sprintf(apiLocation, domain), http.Get)
}

// Lookup gets categorify report with a customised configuration.
func (c *Categorify) Lookup(domain string) (*Report, error) {
	return lookup(fmt.Sprintf(apiLocation, domain), c.Client.Get)
}
