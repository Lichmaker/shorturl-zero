package shorten

import (
	"context"
	"fmt"

	"golang.org/x/sync/singleflight"

	"github.com/lichmaker/short-url-micro/model/shorts"
	"github.com/lichmaker/short-url-micro/pkg/cachex"
	"github.com/lichmaker/short-url-micro/pkg/hash"
	"github.com/lichmaker/short-url-micro/pkg/helpers"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"gorm.io/gorm"
)

type Shorten struct {
	Ctx      context.Context
	GormDB   *gorm.DB
	Redis    *redis.Redis
	BloomKey string
	Sg       *singleflight.Group
}

func (s *Shorten) Make(long string) (shorts.Short, error) {
	// 补充协议头
	long = helpers.FillHttpScheme(long)

	// 哈希
	short := hash.Make(long)

	var shortModel shorts.Short

	// myBloom := bloom.New(s.Redis, s.BloomKey, 1024)
	// exists, err := myBloom.Exists([]byte(short))
	// if err != nil {
	// 	return shortModel, err
	// }

	cacheInstance := &cachex.ShortenCache{
		RedisClient: s.Redis,
	}
	// 访问缓存
	shortModel, err := cacheInstance.Get(short)
	if err != nil {
		return shortModel, err
	}
	if shortModel.Id > 0 {
		return shortModel, nil
	}

	// 查库
	if shortModel.Id == 0 {
		logx.WithContext(s.Ctx).Infof("%s 无缓存,开始查库", short)
		shortModel, err = s.getFromDb(short)
		if err != nil {
			return shortModel, err
		}
		if shortModel.Id > 0 {
			return shortModel, nil
		}
	}

	// 写库
	shortModel = shorts.Short{
		Long:      long,
		Short:     short,
		ExpiredAt: 0,
	}
	res := s.GormDB.Create(&shortModel)
	if res.Error != nil {
		return shortModel, res.Error
	}
	// myBloom.Add([]byte(short))
	cacheInstance.Set(shortModel)
	return shortModel, nil
}

func (s *Shorten) Get(shortStr string) (shorts.Short, error) {
	cacheInstance := &cachex.ShortenCache{
		RedisClient: s.Redis,
	}
	// 访问缓存
	shortModel, err := cacheInstance.Get(shortStr)
	if err != nil {
		return shortModel, err
	}
	if shortModel.Id > 0 {
		return shortModel, nil
	}
	// 查库
	shortModel, err = s.getFromDb(shortStr)
	if err != nil {
		return shortModel, err
	}
	return shortModel, nil
}

func (s *Shorten) getFromDb(shortStr string) (shorts.Short, error) {
	cacheInstance := &cachex.ShortenCache{
		RedisClient: s.Redis,
	}

	shortModelRes, err, fromShared := s.Sg.Do(fmt.Sprintf("shorten_query:%s", shortStr), func() (interface{}, error) {
		var shortModel shorts.Short
		shortModel, err := shorts.GetByShort(s.Ctx, s.GormDB, shortStr)
		if err != nil {
			return shortModel, err
		}
		cacheInstance.Set(shortModel)
		return shortModel, nil
	})
	logx.WithContext(s.Ctx).Infof("%s 查库完成，是否使用了共享数据%b", shortStr, fromShared)
	return shortModelRes.(shorts.Short), err
}
