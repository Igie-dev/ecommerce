package dto

import "github.com/laptop-shop_api/model"

func ToUser(c *model.User) *model.User {
	return &model.User{
		ID:         c.ID,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		Email:      c.Email,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		UserId: c.UserId,
		Contact:    c.Contact,
	}
}

func ToUsers(users []*model.User) []*model.User {
	res := make([]*model.User, len(users))
	for i, user := range users {
		res[i] = ToUser(user)
	}
	return res
}
