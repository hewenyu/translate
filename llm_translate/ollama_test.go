package llmtranslate

import (
	"context"
	"testing"
)

func TestOllamaTranslate(t *testing.T) {
	provider, err := CreateOllamaProvider(ollamaEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	translate := NewOllamaTranslate(provider)

	text := "こんにちは、世界！"
	sourceLang := LangsCodeJa
	targetLang := LangsCodeZh
	result, err := translate.Translate(context.Background(), text, sourceLang, targetLang)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestOllamaTranslate2(t *testing.T) {
	provider, err := CreateOllamaProvider(ollamaEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	translate := NewOllamaTranslate(provider)

	text := "Hello, world!"
	sourceLang := LangsCodeEn
	targetLang := LangsCodeZh
	result, err := translate.Translate(context.Background(), text, sourceLang, targetLang)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
