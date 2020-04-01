package redis

func (c Client) HMSet(key string, fields map[string]interface{}) error {
	err := c.baseClient.HMSet(c.CacheKey(key), fields).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	return c.baseClient.HMGet(c.CacheKey(key), fields...).Result()
}

func (c Client) HGet(key string, field string) (string, error) {
	return c.baseClient.HGet(c.CacheKey(key), field).Result()
}

func (c Client) HSet(key string, field string, value interface{}) (bool, error) {
	return c.baseClient.HSet(c.CacheKey(key), field, value).Result()
}

func (c Client) HGetAll(key string) (map[string]string, error) {
	return c.baseClient.HGetAll(c.CacheKey(key)).Result()
}
