package configs

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/ChangSZ/mall-go/pkg/env"
	"github.com/ChangSZ/mall-go/pkg/file"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	MySQL struct {
		Read struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"read"`
		Write struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"mysql"`

	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
		Database     string `toml:"database"`
		Key          struct {
			Admin        string `toml:"admin"`
			ResourceList string `toml:"resourceList"`
			AuthCode     string `toml:"authCode"`
			OrderId      string `toml:"orderId"`
			Member       string `toml:"member"`
		} `toml:"key"`
		Expire struct {
			Common   int64 `toml:"common"`
			AuthCode int64 `toml:"authCode"`
		} `toml:"expire"`
	} `toml:"redis"`

	Mail struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		To   string `toml:"to"`
	} `toml:"mail"`

	HashIds struct {
		Secret string `toml:"secret"`
		Length int    `toml:"length"`
	} `toml:"hashids"`

	Language struct {
		Local string `toml:"local"`
	} `toml:"language"`

	Jwt struct {
		TokenHeader string `toml:"tokenHeader"`
		Secret      string `toml:"secret"`
		Expiration  int64  `toml:"expiration"`
		TokenHead   string `toml:"tokenHead"`
	} `toml:"jwt"`

	Minio struct {
		Endpoint   string `toml:"endpoint"`
		BucketName string `toml:"bucketName"`
		AccessKey  string `toml:"accessKey"`
		SecretKey  string `toml:"secretKey"`
	} `toml:"minio"`

	Rabbitmq struct {
		Host        string `toml:"host"`
		Port        int64  `toml:"port"`
		VirtualHost string `toml:"virtualHost"`
		Username    string `toml:"username"`
		Password    string `toml:"password"`
	} `toml:"rabbitmq"`

	Alipay struct {
		GatewayUrl      string `toml:"gatewayUrl"`
		AppId           string `toml:"appId"`
		AlipayPublicKey string `toml:"alipayPublicKey"`
		AppPrivateKey   string `toml:"appPrivateKey"`
		ApiAESKey       string `toml:"apiAESKey"`
		ReturnUrl       string `toml:"returnUrl"`
		NotifyUrl       string `toml:"notifyUrl"`
	}
}

var (
	//go:embed dev_configs.toml
	devConfigs []byte

	//go:embed fat_configs.toml
	fatConfigs []byte

	//go:embed uat_configs.toml
	uatConfigs []byte

	//go:embed pro_configs.toml
	proConfigs []byte
)

func init() {
	var r io.Reader

	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfigs)
	case "fat":
		r = bytes.NewReader(fatConfigs)
	case "uat":
		r = bytes.NewReader(uatConfigs)
	case "pro":
		r = bytes.NewReader(proConfigs)
	default:
		r = bytes.NewReader(fatConfigs)
	}

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.AddConfigPath("./configs")

	configFile := "./configs/" + env.Active().Value() + "_configs.toml"
	_, ok := file.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}
