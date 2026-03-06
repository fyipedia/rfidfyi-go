// Package rfidfyi provides a Go client for the RFIDFYI API.
//
// RFIDFYI is a comprehensive RFID reference covering tags, readers, families,
// frequencies, standards, EPC schemes, and use cases. This client requires
// no authentication and has zero external dependencies.
//
// Usage:
//
//	client := rfidfyi.NewClient()
//	result, err := client.Search("uhf")
package rfidfyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the RFIDFYI API.
const DefaultBaseURL = "https://rfidfyi.com/api"

// Client is an RFIDFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new RFIDFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("rfidfyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("rfidfyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("rfidfyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across RFID tags, standards, and glossary terms.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Tag returns details for an RFID tag by slug.
func (c *Client) Tag(slug string) (*TagDetail, error) {
	var result TagDetail
	if err := c.get("/tag/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Reader returns details for an RFID reader by slug.
func (c *Client) Reader(slug string) (*ReaderDetail, error) {
	var result ReaderDetail
	if err := c.get("/reader/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Family returns details for an RFID family by slug.
func (c *Client) Family(slug string) (*FamilyDetail, error) {
	var result FamilyDetail
	if err := c.get("/family/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Frequency returns details for an RFID frequency band by slug.
func (c *Client) Frequency(slug string) (*FrequencyDetail, error) {
	var result FrequencyDetail
	if err := c.get("/frequency/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Standard returns details for an RFID standard by slug.
func (c *Client) Standard(slug string) (*StandardDetail, error) {
	var result StandardDetail
	if err := c.get("/standard/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EPC returns details for an EPC scheme by slug.
func (c *Client) EPC(slug string) (*EPCDetail, error) {
	var result EPCDetail
	if err := c.get("/epc/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UseCase returns details for an RFID use case by slug.
func (c *Client) UseCase(slug string) (*UseCaseDetail, error) {
	var result UseCaseDetail
	if err := c.get("/use-case/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Compare compares two RFID tags.
func (c *Client) Compare(slugA, slugB string) (*CompareResult, error) {
	var result CompareResult
	path := fmt.Sprintf("/compare/?a=%s&b=%s", url.QueryEscape(slugA), url.QueryEscape(slugB))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random RFID tag.
func (c *Client) Random() (*TagDetail, error) {
	var result TagDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
