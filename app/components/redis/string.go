package redis

import "time"

// GetString, error is redis.Nil when key does not exist.
func (c Client) GetString(key string) (string, error) {
	return c.baseClient.Get(c.CacheKey(key)).Result()
}

func (c Client) GetInt(key string) (int, error) {
	return c.baseClient.Get(c.CacheKey(key)).Int()
}

// Mget, 空数据会顺序返回 nil
func (c Client) MGet(keys []string) ([]interface{}, error) {
	pKeys := []string{}
	for _, key := range keys {
		pKeys = append(pKeys, c.CacheKey(key))
	}
	return c.baseClient.MGet(pKeys...).Result()
}

// expiration=0不过期
func (c Client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.baseClient.Set(c.CacheKey(key), value, expiration).Err()
}

func (c Client) IncrBy(key string, value int, expiration time.Duration) error {
	key = c.CacheKey(key)
	_, err := c.baseClient.IncrBy(key, int64(value)).Result()
	if err != nil {
		return err
	}
	c.baseClient.Expire(key, expiration)
	return nil
}
