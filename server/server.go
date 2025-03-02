package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hewenyu/translate"
	"github.com/hewenyu/translate/common"
)

type Server struct {
	translator translate.Translator
	config     *common.Config
}

// NewServer 创建新的服务器实例
func NewServer(cfg *common.Config) (*Server, error) {
	translator, err := translate.CreateTranslator(cfg)
	if err != nil {
		return nil, fmt.Errorf("create translator: %w", err)
	}

	return &Server{
		translator: translator,
		config:     cfg,
	}, nil
}

// Run 运行服务器
func (s *Server) Run(addr string) error {
	r := gin.Default()

	// 中间件
	r.Use(gin.Recovery())

	// 路由
	r.POST("/translate", s.handleTranslate)

	return r.Run(addr)
}

// handleTranslate 处理翻译请求
func (s *Server) handleTranslate(c *gin.Context) {
	var req TranslateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "invalid request: " + err.Error(),
		})
		return
	}

	// 标准化语言代码
	sourceLang := strings.ToLower(req.SourceLang)
	targetLang := strings.ToLower(req.TargetLang)

	// 执行翻译
	translated, err := s.translator.Translate(c.Request.Context(), req.Text, sourceLang, targetLang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "translation failed: " + err.Error(),
		})
		return
	}

	// 生成随机ID
	rand.New(rand.NewSource(time.Now().UnixNano()))

	id := rand.Int63n(9999999999-1000000000) + 1000000000

	// 返回响应
	c.JSON(http.StatusOK, TranslateResponse{
		Code:       http.StatusOK,
		Data:       translated,
		ID:         id,
		Method:     strings.ToLower(string(s.config.TranslateType)),
		SourceLang: strings.ToUpper(sourceLang),
		TargetLang: strings.ToUpper(targetLang),
	})
}
