package auzer

/*
	resourceKey string       // Azure 资源密钥
	endpoint    string       // Azure 翻译服务的终结点
	region      string       // Azure 翻译服务的区域
*/

// AuzerConfig Auzer配置
type AuzerConfig struct {
	ResourceKey string `yaml:"resource_key"` // Azure 资源密钥
	Endpoint    string `yaml:"endpoint"`     // Azure 翻译服务的终结点
	Region      string `yaml:"region"`       // Azure 翻译服务的区域
}

// new config
