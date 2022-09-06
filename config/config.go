package config

import "github.com/jinzhu/configor"

type Config struct {
	AccountConfig `toml:"account_config"`

}

type AccountConfig struct {
	UIN uint64 `toml:"uin"`
	Password string `toml:"password"`

}

func (c *Config) Load(filePath string) (err error) {
	err = configor.Load(c, filePath)
	return err
}
