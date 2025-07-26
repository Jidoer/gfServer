package cache

import (
	"context"
	"gfAdmin/internal/cache/file"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
)

var redisConfig = &gredis.Config{
	Address: "47.93.100.223:6379",
	Db:      0,
	Pass:    "26514849bca771f4",
}
var redis *gredis.Redis

// cache 缓存驱动
var cache *gcache.Cache
var err error

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter(ctx context.Context) {
	var cacheAdapter gcache.Adapter
	// adapter := g.Cfg().MustGet(ctx, "cache.adapter").String()
	// fileDir := g.Cfg().MustGet(ctx, "cache.fileDir").String()
	adapter := "redis"
	fileDir := ""

	switch adapter {
	case "redis":
		redis, err = gredis.New(redisConfig)
		if err != nil {
			panic(err)
		}
		cacheAdapter = gcache.NewAdapterRedis(redis) //(g.Redis())
	case "file":
		if fileDir == "" {
			g.Log().Fatal(ctx, "file path must be configured for file caching.")
			return
		}
		if !gfile.Exists(fileDir) {
			if err := gfile.Mkdir(fileDir); err != nil {
				g.Log().Fatalf(ctx, "failed to create the cache directory. procedure, err:%+v", err)
				return
			}
		}
		cacheAdapter = file.NewAdapterFile(fileDir)
	default:
		cacheAdapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(cacheAdapter)
	// g.DB().GetCache().SetTTL(60 * time.Second)
	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(cacheAdapter)
}
