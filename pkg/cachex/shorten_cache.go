package cachex

import (
	"encoding/json"
	"fmt"

	"github.com/lichmaker/short-url-micro/model/shorts"
	"github.com/lichmaker/short-url-micro/pkg/helpers"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ShortenCache struct {
	RedisClient *redis.Redis
}

func (s *ShortenCache) Set(model shorts.Short) error {
	jsonByte, err := json.Marshal(model)
	if err != nil {
		return err
	}
	err = s.RedisClient.Setex(shortsCacheKey(model.Short), string(jsonByte), 86400)
	if err != nil {
		return err
	}
	_, err = s.RedisClient.Zadd(shortsSortedSetKey(), helpers.GetTimestamp(), model.Short)
	return err
}

func (s *ShortenCache) Get(shortStr string) (shorts.Short, error) {
	var model shorts.Short
	res, err := s.RedisClient.Get(shortsCacheKey(shortStr))
	if err != nil {
		return model, err
	}
	if len(res) == 0 {
		return model, nil
	}
	err = json.Unmarshal([]byte(res), &model)
	if err != nil {
		return model, err
	}
	return model, nil
}

func shortsCacheKey(shortStr string) string {
	return fmt.Sprintf("short_url_micro_cache:%s", shortStr)
}

func shortsSortedSetKey() string {
	return "short_url_micro_sorted_set"
}
