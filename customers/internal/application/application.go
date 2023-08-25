package application

import (
	"context"

	"github.com/hnamzian/go-mallbots/customers/internal/domain"
	"github.com/stackus/errors"
)

type (
	App interface {
		RegisterCustomer(ctx context.Context, register RegisterCustomer) error
		AuthorizeCustomer(ctx context.Context, authorize AuthorizeCustomer) error
		GetCustomer(ctx context.Context, get GetCustomer) (*domain.Customer, error)
		EnableCustomer(ctx context.Context, enable EnableCustomer) error
		DisableCustomer(ctx context.Context, disable DisableCustomer) error
	}
	
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

func NewApplication(customers domain.CustomersRepository) Application {
	return Application{customers}
}

func (a Application) RegisterCustomer(ctx context.Context, register RegisterCustomer) error{
	customer, err := domain.RegisterCustomer(register.ID, register.Name, register.SmsNumber)
	if err!= nil {
        return err
    }
	return a.customers.Save(ctx, customer)
}

func (a Application) AuthorizeCustomer(ctx context.Context, authorize AuthorizeCustomer) error {
	customer, err := a.customers.Get(ctx, authorize.ID)
	if err!= nil {
        return err
    }
	if !customer.Enabled {
		return errors.Wrap(errors.ErrUnauthorized, "Customer is not enabled")
	}
	return nil
}

func (a Application) GetCustomer(ctx context.Context, get GetCustomer) (*domain.Customer, error) {
	customer, err := a.customers.Get(ctx, get.ID)
	if err!= nil {
        return nil, err
    }
	return customer, nil
}

func (a Application) EnableCustomer(ctx context.Context, enable EnableCustomer) error {
	customer, err := a.customers.Get(ctx, enable.ID)
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

func (a Application) DisableCustomer(ctx context.Context, disable DisableCustomer) error {
	customer, err := a.customers.Get(ctx, disable.ID)
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
