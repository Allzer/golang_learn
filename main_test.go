package main

import "testing"

func TestIsValidEmailTable(t *testing.T) {
	table := []struct {
		email string
		want bool
	}{
		{"example@mail.com", true},
		{"miss@tld", false},
	}

	for _, data := range table {
		result := IsValidEmail(data.email)
		if result != data.want {
			t.Errorf("%v: %tm want: %t", data.email, result, data.want)
		}
	}
}