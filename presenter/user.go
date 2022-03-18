package presenter

import "my-arch/domain"

type userJson struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func User(user *domain.User) *userJson {
	return &userJson{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func Users(users []*domain.User) []*userJson {
	usersJson := make([]*userJson, len(users))
	for i, v := range users {
		usersJson[i] = User(v)
	}

	return usersJson
}
