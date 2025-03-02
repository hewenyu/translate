package translate

import (
	"context"
)

type TranslateSrt interface {
	Translate(ctx context.Context, srtPath string, sourceLang string, targetLang string) (string, error)
}
