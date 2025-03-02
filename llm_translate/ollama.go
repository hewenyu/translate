package llmtranslate

import (
	"context"
	"fmt"

	"github.com/hewenyu/llm"
)

var (
	ollamaEndpoint = "http://localhost:11434"
)

type OllamaConfig struct {
	Endpoint string `yaml:"endpoint"`
}

const (
	modelJapaneseToChinese = "yuebanlaosiji/sakura-1.5b-qwen2.5-v1.0:latest"
	modelOther             = "qwen2.5:latest"
)

// CreateOllamaProvider 创建一个OllamaProvider
func CreateOllamaProvider(endpoint string) (llm.Provider, error) {
	return llm.NewOllamaProvider(endpoint)
}

type OllamaTranslate struct {
	llm.Provider
}

// NewOllamaTranslate 创建一个OllamaTranslate
func NewOllamaTranslate(provider llm.Provider) *OllamaTranslate {
	return &OllamaTranslate{Provider: provider}
}

// Translate 翻译
func (t *OllamaTranslate) Translate(ctx context.Context, text string, sourceLang LangsCode, targetLang LangsCode) (string, error) {
	// check model
	exist, err := t.checkModel(ctx)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", fmt.Errorf("model is not exist, please download it first")
	}
	// 日文转中文使用 yuebanlaosiji/sakura-1.5b-qwen2.5-v1.0
	// 其他使用 qwen2.5
	var model string
	var req llm.ChatRequest
	if sourceLang == LangsCodeJa && targetLang == LangsCodeZh {
		model = modelJapaneseToChinese
		req = llm.ChatRequest{
			Messages: []llm.Message{
				{Role: "user", Content: text},
			},
		}
	} else {
		model = modelOther
		// build prompt
		prompt := t.buildPrompt(text, sourceLang, targetLang)
		req = llm.ChatRequest{
			Messages: []llm.Message{
				{Role: "user", Content: prompt},
			},
		}
	}
	response, err := t.Provider.Chat(ctx, model, req)
	if err != nil {
		return "", err
	}
	return response.Message.Content, nil
}

func (t *OllamaTranslate) checkModel(ctx context.Context) (bool, error) {
	// check modelJapaneseToChinese, modelOther is exist
	models, err := t.Provider.ListModels(ctx)
	if err != nil {
		return false, err
	}
	existModelJapaneseToChinese := false
	existModelOther := false
	for _, m := range models {
		if m.Name == modelJapaneseToChinese {
			existModelJapaneseToChinese = true
		}
		if m.Name == modelOther {
			existModelOther = true
		}
	}
	if !existModelJapaneseToChinese {
		return false, fmt.Errorf("modelJapaneseToChinese is not exist, please download it first,run `ollama pull %s`", modelJapaneseToChinese)
	}
	if !existModelOther {
		return false, fmt.Errorf("modelOther is not exist, please download it first,run `ollama pull %s`", modelOther)
	}
	return true, nil
}

// build prompt
func (t *OllamaTranslate) buildPrompt(text string, sourceLang LangsCode, targetLang LangsCode) string {
	prompt := fmt.Sprintf("Translate the following text from %s to %s:\n%s", sourceLang, targetLang, text)
	return prompt
}
