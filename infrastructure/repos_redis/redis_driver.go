package repos_redis

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

// NewRedisDriver create a new instance
func newRedisDriver() *Redis {
	return &Redis{
		Addr:     viper.GetString("REDIS_ADDR"),
		Password: viper.GetString("REDIS_PSW"),
		DB:       viper.GetInt("REDIS_DB"),
	}
}

// Connect establish a redis connection
func (r *Redis) Connect() (bool, error) {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})

	_, err := r.Ping()

	if err != nil {
		return false, err
	}

	return true, nil
}

// Ping checks the redis connection
func (r *Redis) Ping() (bool, error) {
	pong, err := r.Client.Ping().Result()

	if err != nil {
		return false, err
	}
	return pong == "PONG", nil
}

// Set sets a record
func (r *Redis) Set(key, value string, expiration time.Duration) (bool, error) {
	result := r.Client.Set(key, value, expiration)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val() == "OK", nil
}

// Get gets a record value
func (r *Redis) Get(key string) (string, error) {
	result := r.Client.Get(key)

	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Val(), nil
}

// Exists deletes a record
func (r *Redis) Exists(key string) (bool, error) {
	result := r.Client.Exists(key)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val() > 0, nil
}

// Del deletes a record
func (r *Redis) Del(key string) (int64, error) {
	result := r.Client.Del(key)

	if result.Err() != nil {
		return 0, result.Err()
	}

	return result.Val(), nil
}

// HGet gets a record from hash
func (r *Redis) HGet(key, field string) (string, error) {
	result := r.Client.HGet(key, field)

	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Val(), nil
}

// HSet sets a record in hash
func (r *Redis) HSet(key, field, value string) (int64, error) {
	result := r.Client.HSet(key, field, value)

	if result.Err() != nil {
		return -1, result.Err()
	}

	return result.Val(), nil
}

// HExists checks if key exists on a hash
func (r *Redis) HExists(key, field string) (bool, error) {
	result := r.Client.HExists(key, field)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val(), nil
}

// HDel deletes a hash record
func (r *Redis) HDel(key, field string) (int64, error) {
	result := r.Client.HDel(key, field)

	if result.Err() != nil {
		return 0, result.Err()
	}

	return result.Val(), nil
}

// HLen count hash records
func (r *Redis) HLen(key string) (int64, error) {
	result := r.Client.HLen(key)

	if result.Err() != nil {
		return 0, result.Err()
	}

	return result.Val(), nil
}

// HTruncate deletes a hash
func (r *Redis) HTruncate(key string) (int64, error) {
	result := r.Client.Del(key)

	if result.Err() != nil {
		return 0, result.Err()
	}

	return result.Val(), nil
}

// HScan return an iterative obj for a hash
func (r *Redis) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.Client.HScan(key, cursor, match, count)
}

// RPush
func (r *Redis) RPush(key string, values ...interface{}) (int64, error) {
	result := r.Client.RPush(key, values...)

	if result.Err() != nil {
		return 0, result.Err()
	}

	return result.Val(), nil
}
