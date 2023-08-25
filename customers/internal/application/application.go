package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/internal/domain"
	"github.com/stackus/errors"
)

type (
	Application struct{
		customers domain.CustomersRepository
	}

	RegisterCustomer struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		SmsNumber string `json:"sms_number"`
	}

	AuthorizeCustomer struct {
		ID string `json:"id"`
    }

	GetCustomer struct {
		ID string `json:"id"`
	}

	EnableCustomer struct {
		ID string `json:"id"`
    }

	DisableCustomer struct {
        ID string `json:"id"`
    }
)

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) RegisterCustomer(ctx context.Context, register *RegisterCustomer) error{
	customer, err := domain.RegisterCustomer(register.ID, register.Name, register.SmsNumber)
	if err!= nil {
        return err
    }
	return a.customers.Save(ctx, customer)
}

func (a *Application) AuthorizeCustomer(ctx context.Context, id string) error {
	customer, err := a.customers.Get(ctx, id)
	if err!= nil {
        return err
    }
	if !customer.Enabled {
		return errors.Wrap(errors.ErrUnauthorized, "Customer is not enabled")
	}
	return nil
}

func (a *Application) GetCustomer(ctx context.Context, id string) (*domain.Customer, error) {
	customer, err := a.customers.Get(ctx, id)
	if err!= nil {
        return nil, err
    }
	return customer, nil
}

func (a *Application) EnableCustomer(ctx context.Context, id string) error {
	customer, err := a.customers.Get(ctx, id)
	if err!= nil {
        return err
    }
	if err = customer.EnableCustomer(); err!= nil {
		return err
	}
	if err = a.customers.Update(ctx, customer); err!= nil {
		return err
	}
	return nil
}

func (a *Application) DisableCustomer(ctx context.Context, id string) error {
	customer, err := a.customers.Get(ctx, id)
	if err!= nil {
        return err
    }
	if err = customer.DisableCustomer(); err!= nil {
        return err
    }
	if err = a.customers.Update(ctx, customer); err!= nil {
        return err
    }
	return nil
}
