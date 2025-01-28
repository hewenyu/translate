package qcloud

type QCloudConfig struct {
	SecretId  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
	Region    string `yaml:"region"`
	ProjectId int64  `yaml:"project_id"`
}
