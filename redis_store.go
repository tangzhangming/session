package session

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type redis_store struct {
	id      string
	name    string
	context context.Context
	redis   *redis.Client
	Values  map[string]string
}


func (r *redis_store) ID() string {
	return r.id
}

func (r *redis_store) Name() string {
	return r.name
}

func (r *redis_store) StorageID() string {
	return "session-" + r.ID()
}

func (r *redis_store) Get(name string) string {
	if val, ok := r.All()[name]; ok {
		return val
	} else {
		return ""
	}
}

func (r *redis_store) DefaultGet(name string, def string) string {
	if val := r.Get(name); val != "" {
		return val
	} else {
		return def
	}
}

func (r *redis_store) All() map[string]string {
	attributes := make(map[string]string)

	//读取Session
	session, err := r.redis.Get(r.context, r.StorageID()).Result()
	if err != nil {
		return attributes
	}

	//反序列化
	json.Unmarshal([]byte(session), &attributes)
	return attributes
}

func (r *redis_store) Set(name string, value string) {
	sessionMap := r.All()
	sessionMap[name] = value
	str, _ := json.Marshal(sessionMap)
	r.redis.Set(r.context, r.StorageID(), str, time.Hour)
}

func (r *redis_store) Del(name string) bool {
	attributes := r.All()
	delete(attributes, name)

	str, _ := json.Marshal(attributes)
	if err := r.redis.Set(r.context, r.StorageID(), str, time.Hour).Err(); err == nil {
		return true
	} else {
		return false
	}
}

func (r *redis_store) Clear() bool {
	if err := r.redis.Del(r.context, r.StorageID()).Err(); err == nil {
		return true
	} else {
		return false
	}
}
