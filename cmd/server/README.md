# 翻译服务 API 文档

## 基本信息

- 服务启动命令：`./server -config config.yaml -addr :1188`
- 默认监听地址：`:1188`
- 默认配置文件：`config.yaml`

## 配置文件说明

配置文件使用 YAML 格式，结构如下：

```yaml
translate_type: <string>  # 翻译服务类型：qcloud/auzer/ollama
qcloud:        # 腾讯云翻译配置
  # 腾讯云相关配置
auzer:         # Auzer翻译配置
  # Auzer相关配置
ollama:        # Ollama翻译配置
  # Ollama相关配置
```

## API 接口

### 1. 文本翻译

**接口地址**：`/translate`  
**请求方法**：`POST`  
**Content-Type**：`application/json`

#### 请求参数

```json
{
    "text": "要翻译的文本",        // 必填
    "source_lang": "源语言代码",   // 必填
    "target_lang": "目标语言代码"  // 必填
}
```

#### 响应格式

成功响应：
```json
{
    "code": 200,                  // 状态码
    "data": "翻译后的文本",       // 翻译结果
    "id": 1234567890,            // 翻译任务ID
    "method": "翻译服务类型",     // qcloud/auzer/ollama
    "source_lang": "源语言代码",  // 大写格式
    "target_lang": "目标语言代码" // 大写格式
}
```

错误响应：
```json
{
    "code": 400/500,             // 错误状态码
    "message": "错误信息",        // 错误描述
    "data": null                 // 数据为空
}
```

#### 状态码说明

- 200: 请求成功
- 400: 请求参数错误
- 500: 服务器内部错误

## 启动参数说明

服务器支持以下命令行参数：

- `-config`: 指定配置文件路径，默认为 `config.yaml`
- `-addr`: 指定服务器监听地址，默认为 `:1188`

示例：
```bash
./server -config custom_config.yaml -addr :8080
```

## 注意事项

1. 语言代码在请求时不区分大小写，但响应中会统一返回大写格式
2. 每个翻译请求会生成一个唯一的任务 ID
3. 服务使用 Gin 框架提供 HTTP 服务
4. 服务内置了异常恢复中间件 