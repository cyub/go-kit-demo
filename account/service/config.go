package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config define config struct
type Config struct {
	confCenterHost  string
	confCenterPort  int
	appName         string
	appEnv          string
	refreshInternal int
	viper           *viper.Viper
}

// C is global pointer to Config
var C *Config

// NewConfig generage config
func NewConfig(confCenterHost string, confCenterPort int, appName string, appEnv string, refreshInternal int) *Config {
	C = &Config{
		confCenterHost,
		confCenterPort,
		appName,
		appEnv,
		refreshInternal,
		viper.New(),
	}
	return C
}

// Load load Config
func (c *Config) Load() {
	c.loadConfFromConfCenter()

	// @todo 配置更新之后，数据库连接池需要重连等
	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Info("config refresh")
			c.loadConfFromConfCenter()
		}
	}()
}

// Reload reload config
func (c *Config) Reload() {
	c.loadConfFromConfCenter()
}

func (c *Config) loadConfFromConfCenter() {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://" + c.confCenterHost + ":" + strconv.Itoa(c.confCenterPort)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("connect consul error " + err.Error())
		panic(err)
	}

	kv := consulClient.KV()
	keyPrefix := fmt.Sprintf("%s/%s,%s/", "config", c.appName, c.appEnv)
	if pairs, _, err := kv.List(keyPrefix, nil); err != nil {
		panic(err)
	} else {
		for _, pair := range pairs {
			k := strings.TrimLeft(pair.Key, keyPrefix)
			if len(k) <= 0 {
				continue
			}
			c.viper.Set(k, string(pair.Value))
		}
	}
}

// IsSet checks to see if the key has been set in any of the data locations.
func (c *Config) IsSet(key string) bool {
	return c.viper.IsSet(key)
}

// Get can retrieve any value given the key to use.
func (c *Config) Get(key string) interface{} {
	return c.viper.Get(key)
}

// GetViper get the  Viper instance.
func (c *Config) GetViper() *viper.Viper {
	return c.viper
}

// Set sets the value for the key in the override register.
func (c *Config) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

// GetString returns the value associated with the key as a string.
func (c *Config) GetString(key string) string {
	return c.viper.GetString(key)
}

// GetInt returns the value associated with the key as an integer
func (c *Config) GetInt(key string) int {
	return c.viper.GetInt(key)
}
