# Translate Service

一个支持多种翻译服务的 API 服务器，目前支持腾讯云翻译和 Azure 翻译服务。

## 功能特性

- 支持多种翻译服务
  - 腾讯云翻译 (qcloud)
  - Azure 翻译服务 (auzer)
- RESTful API 接口
- 统一的配置管理
- 支持自动语言检测
- 支持 SRT 字幕文件翻译

## 快速开始

1. 复制配置文件模板并修改：
```bash
cp config.yaml.template config.yaml
```

2. 编辑 `config.yaml` 文件，填入你的翻译服务凭证：
```yaml
translate_type: qcloud  # 或 auzer

qcloud:
  secret_id: "your-secret-id"
  secret_key: "your-secret-key"
  region: "ap-guangzhou"
  project_id: 0

auzer:
  resource_key: "your-azure-resource-key"
  endpoint: "https://api.cognitive.microsofttranslator.com/translate"
  region: "eastasia"
```

3. 启动服务器：
```bash
go run cmd/server/main.go -addr :1188
```

## API 使用

### 文本翻译

**请求：**
```bash
curl -X POST http://localhost:1188/translate \
-H "Content-Type: application/json" \
-d '{
    "text": "你好，世界",
    "source_lang": "ZH",
    "target_lang": "EN"
}'
```

**响应：**
```json
{
  "code": 200,
  "data": "Hello, world!",
  "id": 8356681003,
  "method": "qcloud",
  "source_lang": "ZH",
  "target_lang": "EN"
}
```

## 配置说明

- `translate_type`: 选择使用的翻译服务 (`qcloud` 或 `auzer`)
- `qcloud`: 腾讯云翻译配置
  - `secret_id`: 腾讯云 SecretId
  - `secret_key`: 腾讯云 SecretKey
  - `region`: 服务地区
  - `project_id`: 项目ID
- `auzer`: Azure 翻译配置
  - `resource_key`: Azure 资源密钥
  - `endpoint`: 服务端点
  - `region`: 服务地区

## 开发

### 构建
```bash
make build
```

### 测试
```bash
make test
```

### 运行
```bash
make run
```

## 注意事项

1. 请妥善保管你的翻译服务凭证
2. 配置文件 `config.yaml` 已添加到 `.gitignore`
3. 建议在生产环境使用 HTTPS
