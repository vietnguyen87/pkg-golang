package xredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"sort"
	"sync"
)

type Cmdable interface {
	redis.UniversalClient
	//AllKeys(ctx context.Context, matchKey string) *StringSliceCmd
}

type clusterClient struct {
	*redis.ClusterClient
}

func NewClusterClient(client *redis.ClusterClient) Cmdable {
	return &clusterClient{
		client,
	}
}

func (c *clusterClient) AllKeys(ctx context.Context, matchKey string) (result *StringSliceCmd) {
	keys := make([]string, 0, maxKeyCount)
	mutex := sync.Mutex{}
	err := c.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
		sliceResult := c.scanKeys(ctx, client, matchKey)
		if sliceResult.err != nil {
			return sliceResult.err
		}
		mutex.Lock()
		defer mutex.Unlock()
		keys = append(keys, sliceResult.val...)

		return nil
	})

	if err != nil {
		return &StringSliceCmd{
			err: err,
		}
	}
	sort.Strings(keys)
	return &StringSliceCmd{
		val: keys,
	}
}

func (c *clusterClient) scanKeys(ctx context.Context, client *redis.Client, matchKey string) (result StringSliceCmd) {
	var count int64 = maxKeyCount
	var err error
	keys := make([]string, 0, count)
	var scanResult []string
	var cursor uint64
	for {
		scanResult, cursor, err = client.Scan(ctx, cursor, matchKey, count).Result()
		if err != nil {
			result.err = err
			return
		}
		keys = append(keys, scanResult...)
		if cursor == 0 {
			break
		}
	}

	result.val = keys

	return
}

func (c *clusterClient) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	logrus.Warn("Warning: Keys method is return keys in one node only, please use AllKeys instead")
	return c.ClusterClient.Keys(ctx, pattern)
}

func (c *clusterClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	logrus.Warnf("Warning: Scan method is return keys in one node only, please use AllKeys instead")
	return c.ClusterClient.Scan(ctx, cursor, match, count)
}
