package repos_redis

import (
	"encoding/json"

	"github.com/go-redis/redis/v7"
)

// Redis driver
type Redis struct {
	Client   *redis.Client
	Addr     string
	Password string
	DB       int
}

var driver *Redis

//初始化
func InitRedis() {
	driver = newRedisDriver()
	_, err := driver.Connect()
	if err != nil {
		panic(err.Error())
	}
}

// LoadFromJSON load object from json
func LoadFromJSON(jsonObj string, vPtr interface{}) error {
	err := json.Unmarshal([]byte(jsonObj), vPtr)
	if err != nil {
		return err
	}
	return nil
}

// ConvertToJSON converts object to json
func ConvertToJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
