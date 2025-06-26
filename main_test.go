package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// Mock du driver matériel
// Permet de vérifier les appels à BrewCoffee et ReturnChange
// et de simuler une machine défaillante

type MockDriver struct {
	mock.Mock
	defective bool
}

func (m *MockDriver) BrewCoffee() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDriver) ReturnChange(amount int) error {
	args := m.Called(amount)
	return args.Error(0)
}

func (m *MockDriver) IsDefective() bool {
	return m.defective
}

// 1. Achat réussi d’un café (Happy Path)
func Test_AchatReussiCafe(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("BrewCoffee").Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
}

// 2. Pièce insuffisante
func Test_PieceInsuffisante(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("ReturnChange", 20).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(20)
	driver.AssertNotCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 20)
}

// 3. Machine défaillante
func Test_MachineDefaillante(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = true
	driver.On("ReturnChange", 100).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(100)
	driver.AssertNotCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 100)
}

// 4. Insertion de plusieurs pièces (2 cafés)
func Test_InsertionDeuxFois50cts(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("BrewCoffee").Return(nil).Twice()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	machine.InsertCoin(50)
	driver.AssertNumberOfCalls(t, "BrewCoffee", 2)
}

// 5. Insertion d’un montant supérieur au prix (1€)
func Test_Insertion1Euro(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("BrewCoffee").Return(nil).Once()
	driver.On("ReturnChange", 50).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(100)
	driver.AssertNumberOfCalls(t, "BrewCoffee", 1)
	driver.AssertCalled(t, "ReturnChange", 50)
}

// 6. Insertion d’un montant non multiple de 50cts (2€)
func Test_Insertion2Euros(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("BrewCoffee").Return(nil).Once()
	driver.On("ReturnChange", 150).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(200)
	driver.AssertNumberOfCalls(t, "BrewCoffee", 1)
	driver.AssertCalled(t, "ReturnChange", 150)
}

// 7. Machine défaillante à tout moment
func Test_MachineDefaillanteAnyCoin(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = true
	driver.On("ReturnChange", 5).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(5)
	driver.AssertNotCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 5)
}

// 8. Insertion d'une pièce de 20 centimes
func Test_Piece20Centimes(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("ReturnChange", 20).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(20)
	driver.AssertNotCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 20)
}
