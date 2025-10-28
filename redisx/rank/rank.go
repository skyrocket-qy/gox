package rank

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var _ TopNWithUpdate = (*TopNCache)(nil)

type TopNWithUpdate interface {
	SetScore(c context.Context, id string, newScore int) error
	GetTopN(c context.Context) ([]RankEntry, error)
	Clear(c context.Context) error
}

type RankEntry struct {
	Id    string
	Score int
}

type TopNCache struct {
	cli *redis.Client
	key string
	cap int // Capacity of the rank list
}

func NewTopNCache(redisClient *redis.Client, key string, cap int) *TopNCache {
	return &TopNCache{
		cli: redisClient,
		key: key,
		cap: cap,
	}
}

func (c *TopNCache) SetScore(ctx context.Context, id string, newScore int) error {
	_, err := c.cli.ZAdd(ctx, c.key, redis.Z{Score: float64(newScore), Member: id}).Result()
	if err != nil {
		return err
	}

	if c.cap > 0 {
		_, err = c.cli.ZRemRangeByRank(ctx, c.key, 0, int64(-c.cap-1)).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *TopNCache) GetTopN(ctx context.Context) ([]RankEntry, error) {
	cmd := c.cli.ZRevRangeWithScores(ctx, c.key, 0, int64(c.cap-1))
	zs, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	var result []RankEntry
	for _, z := range zs {
		result = append(result, RankEntry{Id: z.Member.(string), Score: int(z.Score)})
	}
	return result, nil
}

func (c *TopNCache) Clear(ctx context.Context) error {
	_, err := c.cli.Del(ctx, c.key).Result()
	return err
}
