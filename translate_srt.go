package translate

import "context"

// translate srt file

type SrtTranslator struct {
	t Translator // 翻译器
}

// NewSrtTranslator 创建一个新的 Srt 翻译器
func NewSrtTranslator(t Translator) *SrtTranslator {
	return &SrtTranslator{
		t: t,
	}
}

// Translate 翻译字幕文件
func (t *SrtTranslator) Translate(ctx context.Context, srt string, sourceLang string, targetLang string) (string, error) {
	// reed srt file
	
	return "", nil
}
