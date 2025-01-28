package server

// TranslateRequest 翻译请求
type TranslateRequest struct {
	Text       string `json:"text" binding:"required"`
	SourceLang string `json:"source_lang" binding:"required"`
	TargetLang string `json:"target_lang" binding:"required"`
}

// TranslateResponse 翻译响应
type TranslateResponse struct {
	Code       int    `json:"code"`
	Data       string `json:"data"`
	ID         int64  `json:"id"`
	Method     string `json:"method"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

// Response 通用响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
