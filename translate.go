package translate

import "context"

// 实现翻译器需要实现的方法
type Translator interface {
	Translate(ctx context.Context, text string, sourceLang string, targetLang string) (string, error)
}
