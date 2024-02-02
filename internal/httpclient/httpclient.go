package httpclient

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/newUser1337/task-news/internal/newserror"
)

type Client struct {
	client *http.Client
}

type Option func(*Client)

func WithTimeout(timeout time.Duration) Option {
	return func(cl *Client) {
		cl.client.Timeout = timeout
	}
}

func NewClient(opts ...Option) *Client {
	client := &Client{
		client: &http.Client{},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *Client) FetchData(url string, data any) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return newserror.NewErrorNews(
			newserror.ExternalErr,
			"http-client: failed to get response",
			err.Error(),
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return newserror.NewErrorNews(
			newserror.ExternalErr,
			fmt.Sprintf("http-client: failed to get response, status is %d", resp.StatusCode),
			err.Error(),
		)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return newserror.NewErrorNews(
			newserror.ExternalErr,
			"http-client: failed to read response",
			err.Error(),
		)
	}

	if err := xml.Unmarshal(body, data); err != nil {
		return newserror.NewErrorNews(
			newserror.ExternalErr,
			"http-client: failed to unmarshal response",
			err.Error(),
		)
	}
	return nil
}
