package llmtranslate

// Lang 语言
type Lang struct {
	Code string
	Name string
}

// Langs 语言列表

var (
	Langs = map[string]Lang{
		"zh": {Code: "zh", Name: "中文"},
		"en": {Code: "en", Name: "英文"},
		"ja": {Code: "ja", Name: "日文"},
		"ko": {Code: "ko", Name: "韩文"},
	}
)

type LangsCode string

const (
	LangsCodeZh LangsCode = "zh"
	LangsCodeEn LangsCode = "en"
	LangsCodeJa LangsCode = "ja"
	LangsCodeKo LangsCode = "ko"
)

// GetLang 获取语言
func GetLang(code string) *Lang {
	if lang, ok := Langs[code]; ok {
		return &lang
	}
	return nil
}

// GetLangs 获取所有语言
func GetLangs() []Lang {
	langs := make([]Lang, len(Langs))
	i := 0
	for _, lang := range Langs {
		langs[i] = lang
		i++
	}
	return langs
}
