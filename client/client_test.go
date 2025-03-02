package client

import (
	"context"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := NewClient("http://127.0.0.1:1188")
	if err != nil {
		t.Fatal(err)
	}
	data, err := client.Translate(context.Background(), "Hello, world!", "en", "zh")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
