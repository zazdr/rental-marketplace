package client

import (
	"app/bootstrap"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Client struct {
	t      *testing.T
	server *httptest.Server
}

func New(t *testing.T) *Client {
	server, err := bootstrap.New()
	if err != nil {
		t.Fatal(err)
	}

	return &Client{
		t:      t,
		server: httptest.NewServer(server.Echo),
	}
}

func (c *Client) Close() {
	c.server.Close()
}

type Response struct {
	Code   int
	Header http.Header
	Body   string
}

func (c *Client) Request(
	method string,
	url string,
	header http.Header,
	body string,
) *Response {
	c.t.Helper()

	req, err := http.NewRequest(
		method,
		c.server.URL+url,
		bytes.NewReader([]byte(body)),
	)
	if err != nil {
		c.t.Fatal("create request: ", err)
	}

	for k, v := range header {
		req.Header[k] = v
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.t.Fatal("do request: ", err)
	}
	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		c.t.Fatal("read body: ", err)
	}
	bodyString := string(bodyByte)

	return &Response{
		Code:   resp.StatusCode,
		Header: resp.Header,
		Body:   bodyString,
	}
}
