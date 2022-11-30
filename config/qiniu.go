package config

type Qiniu struct {
	AccessKey   string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SercetKey   string `mapstructure:"sercet-key" json:"sercet-key" yaml:"sercet-key"`
	Bucket      string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	QiniuServer string `mapstructure:"qiniu-server" json:"qiniu-server" yaml:"qiniu-server"`
}
