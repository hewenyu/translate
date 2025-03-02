package llmtranslate

import (
	"context"
	"testing"
)

func TestOllamaTranslate(t *testing.T) {

	translate, err := NewOllamaTranslate(ollamaEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	text := "こんにちは、世界！"
	sourceLang := LangsCodeJa
	targetLang := LangsCodeZh
	result, err := translate.Translate(context.Background(), text, string(sourceLang), string(targetLang))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestOllamaTranslate2(t *testing.T) {

	translate, err := NewOllamaTranslate(ollamaEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	text := "Hello, world!"
	sourceLang := LangsCodeEn
	targetLang := LangsCodeZh
	result, err := translate.Translate(context.Background(), text, string(sourceLang), string(targetLang))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
