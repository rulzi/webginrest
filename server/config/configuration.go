package config

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defCfg      map[string]string
	initialized = false
)

// initialize this configuration
func initialize() {
	viper.SetConfigFile(`config.yml`)
	viper.ReadInConfig()

	defCfg = make(map[string]string)

	defCfg["name"] = "Gogin Web"
	defCfg["debug"] = "true"
	defCfg["appurl"] = "http://localhost"
	defCfg["secretkey"] = "secretkey"

	defCfg["server.address"] = "8080"
	defCfg["server.log.level"] = "error"

	defCfg["context.timeout"] = "2"

	defCfg["database.host"] = "localhost"
	defCfg["database.port"] = "3306"
	defCfg["database.user"] = "devuser"
	defCfg["database.password"] = "devpassword"
	defCfg["database.name"] = "devdb"

	defCfg["redis.host"] = "localhost"
	defCfg["redis.port"] = "6379"

	defCfg["token.lifespan"] = "10"
	defCfg["token.api_secret"] = "secret"

	for k := range defCfg {
		err := viper.BindEnv(k)
		if err != nil {
			log.Errorf("Failed to bind env \"%s\" into configuration. Got %s", k, err)
		}
	}

	initialized = true
}

// SetConfig put configuration key value
func SetConfig(key, value string) {
	viper.Set(key, value)
}

// Get fetch configuration as string value
func Get(key string) string {
	if !initialized {
		initialize()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

// GetBoolean fetch configuration as boolean value
func GetBoolean(key string) bool {
	if len(Get(key)) == 0 {
		return false
	}
	b, err := strconv.ParseBool(Get(key))
	if err != nil {
		panic(err)
	}
	return b
}

// GetInt fetch configuration as integer value
func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// GetFloat fetch configuration as float value
func GetFloat(key string) float64 {
	if len(Get(key)) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(Get(key), 64)
	if err != nil {
		panic(err)
	}
	return f
}

// Set configuration key value
func Set(key, value string) {
	defCfg[key] = value
}
