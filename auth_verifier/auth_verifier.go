package auth_verifier

import (
	"github.com/bborbe/auth/model"
	"github.com/bborbe/http/header"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type check func(authToken model.AuthToken, requiredGroups []model.GroupName) (*model.UserName, error)

type auth struct {
	check          check
	requiredGroups []model.GroupName
}

func New(check check, requiredGroups ...model.GroupName) *auth {
	a := new(auth)
	a.check = check
	a.requiredGroups = requiredGroups
	return a
}

func (a *auth) Verify(username string, password string) (bool, error) {
	logger.Debugf("verify user %s has groups %v", username, a.requiredGroups)
	token := header.CreateAuthorizationToken(username, password)
	user, err := a.check(model.AuthToken(token), a.requiredGroups)
	if err != nil {
		logger.Debugf("verify failed: %v", err)
		return false, err
	}
	result := len(*user) > 0
	logger.Debugf("verify user %s => %v", username, result)
	return result, nil
}
