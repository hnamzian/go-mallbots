package domain

type BasketCheckedOut struct {
	Basket *Basket
}

func (e *BasketCheckedOut) EventName() string {
	return "basket.BasketCheckedOut"
}
