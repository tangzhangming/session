package session

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Session struct {
}

var session *Session

var session_name string

func init() {
	session_name = "s_id"
}

func New(ctx *gin.Context, client *redis.Client) store {

	var session_id string
	session_id, _ = ctx.Cookie(session_name)

	if session_id == "" {
		session_id = uuid.New().String()
		ctx.SetCookie(session_name, session_id, 3600, "/", "poorobject.net", false, true)
	}

	store := &redis_store{
		id:      session_id,
		name:    session_name,
		context: context.Background(),
		redis:   client,
		Values:  make(map[string]string),
	}

	store.Set("uuid_version", uuid.New().Version().String())

	return store
}
