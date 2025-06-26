package main

// Double de test pour CoffeeMachineDriver (pas de mock, comportement contrôlé)
type TestDriver struct {
	BrewCoffeeFunc   func() error
	ReturnChangeFunc func(amount int) error
	IsDefectiveFunc  func() bool
}

func (d *TestDriver) BrewCoffee() error {
	if d.BrewCoffeeFunc != nil {
		return d.BrewCoffeeFunc()
	}
	return nil
}

func (d *TestDriver) ReturnChange(amount int) error {
	if d.ReturnChangeFunc != nil {
		return d.ReturnChangeFunc(amount)
	}
	return nil
}

func (d *TestDriver) IsDefective() bool {
	if d.IsDefectiveFunc != nil {
		return d.IsDefectiveFunc()
	}
	return false
}
