package set2

import (
	"net/url"
	"testing"
)

func TestProfileFor(t *testing.T) {
	email := "foo@bar.com&role=admin"
	rawOutput := profileFor(email)
	output, err := url.ParseQuery(rawOutput)
	if err != nil {
		t.Fatal(err)
	}
	if output.Get("email") != email {
		t.Fatalf("Expected '%v' but got '%v'", email, output.Get("email"))
	}
	if output.Get("role") != "user" {
		t.Fatalf("Expected 'user' but got '%v'", output.Get("role"))
	}
}

func TestCookieECBOracle(t *testing.T) {
	input := []byte("foo@bar.com")

	encrypt, decrypt := cookieECBOracle()
	output := decrypt(encrypt(input))

	if output.Get("email") != string(input) {
		t.Fatalf("Expected '%v' but got '%v'", input, output.Get("email"))
	}

	if output.Get("role") != "user" {
		t.Fatalf("Expected 'user' but got '%v'", output.Get("role"))
	}
}

func TestBreakECBCookie(t *testing.T) {
	encrypt, decrypt := cookieECBOracle()
	output := BreakECBCookie(encrypt)

	values := decrypt(output)
	if values.Get("role") != "admin" {
		t.Fatalf("Expected 'admin' but got '%v'", values.Get("role"))
	}
}
