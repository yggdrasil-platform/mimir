package service

import (
	"fmt"
  "github.com/gomodule/redigo/redis"
	"github.com/kieranroneill/mimir/pkg/logger"
	"github.com/kieranroneill/mimir/pkg/model"
)

const (
  UserTokenExpiresIn = 2630000 // One month in seconds.
  ClientTokenExpiresIn = 10 // Ten seconds in... seconds!
)

type TokenService struct {
	store *redis.Pool
}

func(s *TokenService) Create(tkn model.Token) (*model.Token, error) {
  conn := s.store.Get()

  defer func() {
    if err := conn.Close(); err != nil {
      logger.Error.Printf("failed to close store connection: %v", err)
    }
  }()

  // If the Id has not been set, increment the id.
  if tkn.Id < 1 {
   id, err := s.IncrementId()
   if err != nil {
     return nil, err
   }

   tkn.Id = id
  }

	key := fmt.Sprintf("token:%d", tkn.Id)

	_, err := conn.Do(
		"HMSET",
		key,
		"access_token", tkn.AccessToken,
		"expires_in", tkn.ExpiresIn,
		"token_type", tkn.TokenType,
	)
	if err != nil {
		logger.Error.Print(err)
		return nil, err
	}

	// Add a TTL to expire.
	_, err = conn.Do("EXPIRE", key, tkn.ExpiresIn)
	if err != nil {
		logger.Error.Print(err)
		return nil, err
	}

	return &tkn, nil
}

func(s *TokenService) DeleteById(id int) error {
  conn := s.store.Get()

  defer func() {
    if err := conn.Close(); err != nil {
      logger.Error.Printf("failed to close store connection: %v", err)
    }
  }()

	key := fmt.Sprintf("token:%d", id)
	count, err := redis.Int(conn.Do("DEL", key))
	if err != nil {
		logger.Error.Print(err)
		return err
	}

	if count > 0 {
		logger.Info.Printf("deleted token: %d", id)
	}

	return nil
}

func(s *TokenService) GetById(id int) *model.Token {
	var tkn model.Token
  conn := s.store.Get()

  defer func() {
    if err := conn.Close(); err != nil {
      logger.Error.Printf("failed to close store connection: %v", err)
    }
  }()

	// Get the values for the specified key
	key := fmt.Sprintf("token:%d", id)
	values, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		logger.Error.Print(err)
		return nil
	}

	// Add to the struct
	if redis.ScanStruct(values, &tkn) != nil {
		logger.Error.Print(err)
		return nil
	}

	// Update the ID.
	tkn.Id = id

	return &tkn
}

func(s *TokenService) IncrementId() (int, error) {
  conn := s.store.Get()

  defer func() {
    if err := conn.Close(); err != nil {
      logger.Error.Printf("failed to close store connection: %v", err)
    }
  }()

	// Increment the latest ID, or 0.
	id, err := redis.Int(conn.Do("INCR", "token"))
	if err != nil {
		logger.Error.Print(err)
		return 0, err
	}

	return id, nil
}

func NewTokenService(p *redis.Pool) *TokenService {
	return &TokenService{store: p}
}
