package services

import (
	"net/http"

	"github.com/mky-cyber/bookstore_users-api/domain/users"
	"github.com/mky-cyber/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return nil, nil

	return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusBadRequest,
	}
}
