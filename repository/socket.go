package repository

import (
	esl "github.com/cgrates/fsock"
	"EchoServer/configs"
	"EchoServer/adapters"
// 	"fmt"
// 	"time"
// 	"strings"
// 	"errors"
)

type eslAdapterRepository struct {
	config  *configs.Config
	eslConn *ESLsessions
}

// The freeswitch session manager type holding a buffer for the network connection
// and the active sessions
type ESLsessions struct {
	Cfg         *configs.Config
	Conns       map[string]*eslAdapterRepository // Keep the list here for connection management purposes
	SenderPools map[string]*esl.FSockPool
	RedisAdapter adapters.RedisAdapter// Keep sender pools here
}

type getHttpQuery struct {
	FromNumber string `url:"from" json:"from"`
	ToNumber string `url:"to" json:"to"`
	DidNumber string `url:"did_number" json:"did_number"`
	ALegCallUUID string `url:"a_leg_call_uuid" json:"a_leg_call_uuid"`
	AnswerURL string `url:"answer_url" json:"answer_url"`
}

func NewESLsessions(config *configs.Config) (eslPool *ESLsessions, redisAdapters adapters.RedisAdapter) {
	eslPool = &ESLsessions{
		Cfg:         config,
		Conns:       make(map[string]*eslAdapterRepository),
		SenderPools: make(map[string]*esl.FSockPool),
		RedisAdapter: redisAdapters,
	}
	return
}


