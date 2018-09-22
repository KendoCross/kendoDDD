package tokens

import (
	"time"

	infra "../../ddd_infrastructure"
	"github.com/google/uuid"
)

// Tokens aggregate root
type Tokens struct {
}

//CreatToken 创建长连接Token，并缓存。
func (p *Tokens) CreatToken(userID string) (token string, err error) {
	token = uuid.New().String()
	err = infra.MemoryCache.Put(userID, token, 5*time.Minute)
	return
}

// New creates
func New() *Tokens {
	return &Tokens{}
}
