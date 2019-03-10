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
	Domain      string `json:"domain"`
	IP          string `json:"ip,omitempty"`
	CountryCode string `json:"country-code,omitempty"`
	Country     string `json:"country,omitempty"`
	Rating      struct {
		Language    bool   `json:"language"`
		Violence    bool   `json:"violence"`
		Nudity      bool   `json:"nudity"`
		Adult       bool   `json:"adult"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"rating"`
	Confidence     string         `json:"confidence_level,omitempty"`
	Category       []string       `json:"category"`
	KeywordHeatmap map[string]int `json:"keyword_heatmap"`
}

// Lookup gets categorify report with default http client.
func Lookup(domain string) (*Report, error) {
	resp, err := http.Get(fmt.Sprintf(apiLocation, domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Report
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Lookup gets categorify report with a customised configuration.
func (c *Categorify) Lookup(domain string) (*Report, error) {
	resp, err := c.Client.Get(fmt.Sprintf(apiLocation, domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Report
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
