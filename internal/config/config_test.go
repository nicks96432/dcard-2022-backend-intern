package config_test

import (
	"os"
	"path"
	"reflect"
	"testing"
)

import "dcard-2022-backend-intern/internal/config"

func createFile(t *testing.T, content string) string {
	temp := t.TempDir()
	p := path.Join(temp, "config.json")

	f, err := os.Create(p)
	if err != nil {
		t.Fatalf("createFile(): create file failed: %v", err)
	}

	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("createFile(): cannot write file: %v\n", err)
	}

	if err := f.Close(); err != nil {
		t.Fatalf("createFile(): cannot close file: %v\n", err)
	}

	return p
}

func TestFromValid(t *testing.T) {
	p := createFile(t, `{"port": 8763, "inMemory": true, "hostname": "example.com"}`)

	expected := &config.Config{
		Hostname: "example.com",
		Port:     8763,
		InMemory: true,
	}

	if got, err := config.From(p); err != nil {
		t.Fatalf("TestFromValid(): cannot read config: %v\n", err)
	} else if !reflect.DeepEqual(expected, got) {
		t.Fatalf("TestFromValid(): expected %v, got %v\n", expected, got)
	}
}

func TestFromNotExist(t *testing.T) {
	if _, err := config.From("/path/that/does/not/exist/config.json"); err == nil {
		t.Fatalf("TestFromNotExist(): expected error, got nil\n")
	}
}

func TestFromInvalidJson(t *testing.T) {
	p := createFile(t, "starburst stream")
	if _, err := config.From(p); err == nil {
		t.Fatalf("TestFromInvalidJson(): expected error, got nil\n")
	}
}

func TestDefault(t *testing.T) {
	expected := &config.Config{Port: 48763, InMemory: true, Hostname: "localhost"}
	got := config.Default()

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("TestDefault(): expected %v, got %v\n", expected, got)
	}

}
