package uptime

import (
	"context"
	"fmt"
	"net/http"
)

type TagService service

// Tag represents a check tag in Uptime.com.
type Tag struct {
	PK       int    `json:"pk,omitempty"`
	URL      string `json:"url,omitempty"`
	Tag      string `json:"tag,omitempty"`
	ColorHex string `json:"color_hex,omitempty"`
}

// TagListOptions specifies the optional parameters to the TagService.List method.
type TagListResponse struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []*Tag `json:"results,omitempty"`
}

type TagListOptions struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type TagResponse struct {
	Messages map[string]interface{} `json:"messages,omitempty"`
	Results  Tag                    `json:"results,omitempty"`
}

// List retrieves a list of Uptime.com check tags.
func (s *TagService) List(ctx context.Context, opt *TagListOptions) ([]*Tag, *http.Response, error) {
	u := "check-tags"
	return s.listTags(ctx, u, opt)
}

func (s *TagService) listTags(ctx context.Context, url string, opt *TagListOptions) ([]*Tag, *http.Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)

	var tags TagListResponse
	resp, err := s.client.Do(ctx, req, &tags)
	if err != nil {
		return nil, nil, err
	}
	return tags.Results, resp, nil
}

// Create creates a new check tag in Uptime.com.
func (s *TagService) Create(ctx context.Context, tag *Tag) (*Tag, *http.Response, error) {
	u := "check-tags"

	req, err := s.client.NewRequest("POST", u, tag)
	if err != nil {
		return nil, nil, err
	}

	tr := &TagResponse{}
	resp, err := s.client.Do(ctx, req, tr)
	if err != nil {
		return nil, resp, err
	}

	return &tr.Results, resp, nil
}

// Get retrieves a specific tag from Uptime.com.
func (s *TagService) Get(ctx context.Context, pk int) (*Tag, *http.Response, error) {
	u := fmt.Sprintf("check-tags/%v", pk)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	t := &Tag{}
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

// Update updates an Uptime.com tag in-place.
func (s *TagService) Update(ctx context.Context, tag *Tag) (*Tag, *http.Response, error) {
	u := fmt.Sprintf("check-tags/%v", tag.PK)
	if tag.PK == 0 {
		return nil, nil, fmt.Errorf("Error updating tag with empty PK")
	}

	req, err := s.client.NewRequest("PATCH", u, tag)
	if err != nil {
		return nil, nil, err
	}

	tr := &TagResponse{}
	resp, err := s.client.Do(ctx, req, tr)
	if err != nil {
		return nil, resp, err
	}

	return &tr.Results, resp, nil
}

// Delete removes an Uptime.com tag.
func (s *TagService) Delete(ctx context.Context, pk int) (*http.Response, error) {
	u := fmt.Sprintf("check-tags/%v", pk)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
