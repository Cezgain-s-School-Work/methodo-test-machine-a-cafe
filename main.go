package main

import "fmt"

const prixCafe = 50

// Interface du driver matériel (mockable)
type CoffeeMachineDriver interface {
	BrewCoffee() error
	ReturnChange(amount int) error
	IsDefective() bool
}

type CoffeeMachine struct {
	driver CoffeeMachineDriver
}

func NewCoffeeMachine(driver CoffeeMachineDriver) *CoffeeMachine {
	return &CoffeeMachine{driver: driver}
}

// Insertion d'une pièce (en centimes)
func (m *CoffeeMachine) InsertCoin(amount int) {
	if m.driver.IsDefective() {
		_ = m.driver.ReturnChange(amount)
		return
	}
	if amount >= prixCafe {
		_ = m.driver.BrewCoffee()
		if amount > prixCafe {
			_ = m.driver.ReturnChange(amount - prixCafe)
		}
	} else {
		_ = m.driver.ReturnChange(amount)
	}
}

func main() {
	fmt.Println("Machine à café prête.")
}
