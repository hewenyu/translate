package common

import (
	"fmt"
	"os"

	"github.com/hewenyu/translate/auzer"
	llmtranslate "github.com/hewenyu/translate/llm_translate"
	"github.com/hewenyu/translate/qcloud"
	"gopkg.in/yaml.v3"
)

// TranslateType 翻译服务类型
type TranslateType string

const (
	QCloudTranslate TranslateType = "qcloud" // 腾讯云翻译
	AuzerTranslate  TranslateType = "auzer"  // Auzer翻译
	OllamaTranslate TranslateType = "ollama" // Ollama翻译
)

// Config 配置结构
type Config struct {
	TranslateType TranslateType             `yaml:"translate_type"` // 使用的翻译服务
	QCloud        qcloud.QCloudConfig       `yaml:"qcloud"`         // 腾讯云配置
	Auzer         auzer.AuzerConfig         `yaml:"auzer"`          // Auzer配置
	Ollama        llmtranslate.OllamaConfig `yaml:"ollama"`         // Ollama配置
}

// LoadConfig 从文件加载配置
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}

	return &cfg, nil
}
