package helper

import (
	"github.com/google/go-github/v75/github"

	"github.com/abgeo/follytics/internal/model"
)

func MapGitHubUserToModel(user *github.User, userType model.UserType) *model.User {
	return &model.User{
		Type:     userType,
		GHID:     user.GetID(),
		Username: user.GetLogin(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Avatar:   user.GetAvatarURL(),
	}
}

func MapGitHubUsersToModels(users []*github.User, userType model.UserType) []*model.User {
	models := make([]*model.User, len(users))
	for i, follower := range users {
		models[i] = MapGitHubUserToModel(follower, userType)
	}

	return models
}
