package base62_test

import "testing"

import "dcard-2022-backend-intern/internal/base62"

func TestBase62_Encode(t *testing.T) {
	got := base62.Encode(4876348763)
	expected := "000005k0F4v"

	if got != expected {
		t.Fatalf("base62.Encode(): expect %s, got %s\n", expected, got)
	}
}

func TestBase62_Decode(t *testing.T) {
	got, err := base62.Decode("5k0F4v")
	var expected uint64 = 4876348763

	if err != nil {
		t.Fatalf("error happened: %v\n", err)
	}

	if got != expected {
		t.Fatalf("base62.Decode(): expect %d, got %d\n", expected, got)
	}
}
