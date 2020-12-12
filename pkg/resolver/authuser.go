package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
  "context"
  "errors"
  "strconv"
  "time"

  "github.com/kieranroneill/mimir/pkg/encryption"
  _graphql "github.com/kieranroneill/mimir/pkg/graphql"
  "github.com/kieranroneill/mimir/pkg/model"
  "github.com/kieranroneill/mimir/pkg/service"
  "github.com/kieranroneill/mimir/pkg/utils"
)

func (r *mutationResolver) AuthenticateUser(ctx context.Context, input model.AuthenticateUserInput) (*model.Token, error) {
	ausrv := service.NewAuthUserService(r.Database)
	rslt := ausrv.GetByUserId(input.UserID) // Get user.
	if rslt == nil {
		return nil, errors.New("user doesn't exist")
	}

	// Decrypt the password.
	depsswd, err := encryption.Decrypt(rslt.Password, r.Config.EncryptionKey)
	if err != nil {
		return nil, err
	}

	// Check the password.
	if input.Password != depsswd {
		return nil, errors.New("unauthorized")
	}

	tsrv := service.NewTokenService(r.Store)
	tid, err := tsrv.IncrementId()
	if err != nil {
		return nil, err
	}

	// Create a JWT.
	iat := time.Now()
	exp := time.Now().Add(time.Duration(service.UserTokenExpiresIn) * time.Second)
	gty := "password"
	accstkn, err := utils.CreateJWT(strconv.Itoa(tid), strconv.Itoa(rslt.UserID), iat, exp, gty, r.Config.UserJWTSecretKey)
	if err != nil {
		return nil, err
	}

	// Save the token.
	tkn, err := tsrv.Create(model.Token{
		AccessToken: accstkn,
		ExpiresIn:   service.UserTokenExpiresIn,
		Id: tid,
		TokenType:   "Bearer",
	})
	if err != nil {
		return nil, err
	}

	return tkn, nil
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.AuthenticateUserInput) (*model.Token, error) {
	ausrv := service.NewAuthUserService(r.Database)

	// Check the user exists.
	if rslt := ausrv.GetByUserId(input.UserID); rslt != nil {
		return nil, errors.New("user already exists")
	}

	tsrv := service.NewTokenService(r.Store)
	tid, err := tsrv.IncrementId()
	if err != nil {
		return nil, err
	}

	// Create a JWT.
	iat := time.Now()
	exp := time.Now().Add(time.Duration(service.UserTokenExpiresIn) * time.Second)
	gty := "password"
	accstkn, err := utils.CreateJWT(strconv.Itoa(tid), strconv.Itoa(input.UserID), iat, exp, gty, r.Config.UserJWTSecretKey)
	if err != nil {
		return nil, err
	}

	// Save the token.
	tkn, err := tsrv.Create(model.Token{
		AccessToken: accstkn,
		ExpiresIn:   service.UserTokenExpiresIn,
    Id: tid,
		TokenType:   "Bearer",
	})
	if err != nil {
		return nil, err
	}

	// Encrypt the password.
	enpsswd, err := encryption.Encrypt(input.Password, r.Config.EncryptionKey)
	if err != nil {
		return nil, err
	}

	// Save a new AuthUser with the encrypted password.
	if _, err = ausrv.Create(model.AuthUser{
		Password: enpsswd,
		UserID:   input.UserID,
	}); err != nil {
		return nil, err
	}

	return tkn, nil
}

func (r *queryResolver) Verify(ctx context.Context, accessToken string) (bool, error) {
	return true, nil
}

// Mutation returns _graphql.MutationResolver implementation.
func (r *Resolver) Mutation() _graphql.MutationResolver { return &mutationResolver{r} }

// Query returns _graphql.QueryResolver implementation.
func (r *Resolver) Query() _graphql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
