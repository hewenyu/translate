package translate

import (
	"fmt"

	"github.com/hewenyu/translate/auzer"
	com "github.com/hewenyu/translate/common"
	llmtranslate "github.com/hewenyu/translate/llm_translate"
	"github.com/hewenyu/translate/qcloud"
)

// CreateTranslator 根据配置创建翻译器
func CreateTranslator(cfg *com.Config) (Translator, error) {
	switch cfg.TranslateType {
	case com.QCloudTranslate:
		return qcloud.NewTranslator(&cfg.QCloud)
	case com.AuzerTranslate:
		return auzer.NewTranslator(&cfg.Auzer)
	case com.OllamaTranslate:
		return llmtranslate.NewOllamaTranslate(cfg.Ollama.Endpoint)
	default:
		return nil, fmt.Errorf("unsupported translate type: %s", cfg.TranslateType)
	}
}
