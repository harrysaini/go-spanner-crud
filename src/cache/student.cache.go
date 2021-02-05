package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-spanner-crud/src/libs"
	"go-spanner-crud/src/models"

	"github.com/go-redis/redis/v8"
)

// StudentCache - handle student cache ops
type StudentCache struct {
	cache     *libs.RedisCache
	keyPrefix string
}

// NewStudentCache - create new insatnce
func NewStudentCache(cache *libs.RedisCache) *StudentCache {
	return &StudentCache{
		cache:     cache,
		keyPrefix: "GSC-student",
	}
}

func (studentCache *StudentCache) buildCacheKey(key string) string {
	return fmt.Sprintf("%s-%s", studentCache.keyPrefix, key)
}

// AddNew - add student to redis cache
func (studentCache *StudentCache) AddNew(ctx context.Context, student models.Student) error {
	val, err := json.Marshal(student)
	if err != nil {
		return err
	}

	key := studentCache.buildCacheKey(student.UUID)
	return studentCache.cache.Set(ctx, key, val)
}

// Get - get from redis cache
func (studentCache *StudentCache) Get(ctx context.Context, uuid string) (*models.Student, error) {

	key := studentCache.buildCacheKey(uuid)
	val, err := studentCache.cache.Get(ctx, key)
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var student models.Student
	err = json.Unmarshal(val, &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// Purge deletes user from cache
func (studentCache *StudentCache) Purge(ctx context.Context, uuid string) error {
	key := studentCache.buildCacheKey(uuid)
	return studentCache.cache.Delete(ctx, key)
}
