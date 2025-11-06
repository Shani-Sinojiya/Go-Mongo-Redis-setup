package databases

import (
	"context"

	"accounts.workflecks.com/constants"
)

func RedisFlush() error {
	info := RedisClient.Del(context.Background(), constants.DatabaseRedisPrefix+"*")
	return info.Err()
}
