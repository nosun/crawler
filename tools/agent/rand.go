package agent

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomUserAgent() string {
	return UserAgents[r.Intn(len(UserAgents))]
}