package signup_test

import (
	"app/client"
	"app/shared"
	"net/http"
	"testing"
)

func TestSignupGet(t *testing.T) {
	client := client.New(t)
	defer client.Close()

	resp := client.Request(
		http.MethodGet,
		shared.RouterUserSignup,
		http.Header{},
		"",
	)

	got := resp.Header.Get("Content-Type")
	want := "text/html; charset=utf-8"
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
