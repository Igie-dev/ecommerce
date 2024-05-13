package dto

import "laptop-shop/model"

func ToCustomer(c *model.Customer) *model.Customer {
	return &model.Customer{
		ID:         c.ID,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		Email:      c.Email,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		CustomerId: c.CustomerId,
		Contact:    c.Contact,
	}
}

func ToCustomers(customers []*model.Customer) []*model.Customer {
	res := make([]*model.Customer, len(customers))
	for i, customer := range customers {
		res[i] = ToCustomer(customer)
	}
	return res
}
