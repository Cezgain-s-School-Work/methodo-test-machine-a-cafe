package main

import (
	"fmt"
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

// 4. Insertion d’un montant supérieur au prix (1€)
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

// 5. Insertion d’un montant non multiple de 50cts (2€)
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

// 6. Machine défaillante à tout moment
func Test_MachineDefaillanteAnyCoin(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = true
	driver.On("ReturnChange", 5).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(5)
	driver.AssertNotCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 5)
}

// 7. Insertion d'une pièce de 20 centimes
func Test_Piece20Centimes(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("ReturnChange", 20).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(20)
	driver.AssertNotCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 20)
}

// 8. Pas de café disponible
func Test_PasDeCafeDisponible(t *testing.T) {
	driver := new(MockDriver)
	driver.defective = false
	driver.On("BrewCoffee").Return(fmt.Errorf("Plus de café")).Once()
	driver.On("ReturnChange", 50).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 50)
}

// 0. Stock de café vide
func Test_Feature_StockCafeVide(t *testing.T) {
	driver := new(MockDriver)
	driver.On("BrewCoffee").Return(fmt.Errorf("Plus de café")).Once()
	driver.On("ReturnChange", 50).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 50)
}

// 2. Pas de gobelet disponible (détecté par absence de récipient)
func Test_Feature_AbsenceRecipient(t *testing.T) {
	driver := new(MockDriver)
	driver.On("BrewCoffee").Return(fmt.Errorf("Aucun récipient détecté")).Once()
	driver.On("ReturnChange", 50).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 50)
}

// 2b. Prix d’un gobelet (vérification du calcul)
func Test_Feature_PrixGobelet(t *testing.T) {
	machine := NewCoffeeMachine(new(MockDriver))
	prixSansTasse := machine.PrixCommande(false)
	prixAvecTasse := machine.PrixCommande(true)
	if prixSansTasse != 60 {
		t.Errorf("Le prix avec gobelet doit être 60, obtenu %d", prixSansTasse)
	}
	if prixAvecTasse != 50 {
		t.Errorf("Le prix avec tasse doit être 50, obtenu %d", prixAvecTasse)
	}
}

// 3. Récipient déjà présent (café distribué normalement)
func Test_Feature_RecipientPresent(t *testing.T) {
	driver := new(MockDriver)
	driver.On("BrewCoffee").Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
}

// 4. Distribution automatique de touillette (vérification indirecte)
func Test_Feature_TouilletteDistribuee(t *testing.T) {
	// Ce test serait à faire côté intégration/hardware
}

// 5. Dosage café normal (1 dose d'eau, café boolean)
func Test_Feature_DosageCafeNormal(t *testing.T) {
	machine := NewCoffeeMachine(new(MockDriver))
	machine.SetDosage(false)
	if machine.doseEau != 1 || !machine.doseCafe {
		t.Errorf("Le dosage normal doit être 1 dose d'eau et café true, obtenu %d/%v", machine.doseEau, machine.doseCafe)
	}
}

// 6. Dosage café allongé (2 doses d'eau, café boolean)
func Test_Feature_DosageCafeAllonge(t *testing.T) {
	machine := NewCoffeeMachine(new(MockDriver))
	machine.SetDosage(true)
	if machine.doseEau != 2 || !machine.doseCafe {
		t.Errorf("Le dosage allongé doit être 2 doses d'eau et café true, obtenu %d/%v", machine.doseEau, machine.doseCafe)
	}
}

// 7. Sélection automatique de la source d’eau
func Test_Feature_SelectionSourceEau(t *testing.T) {
	machine := NewCoffeeMachine(new(MockDriver))
	machine.SetSourceEau(false, true)
	if machine.sourceEau != Reseau {
		t.Errorf("La source d'eau doit être le réseau si dispo")
	}
	machine.SetSourceEau(false, false)
	if machine.sourceEau != Bonbonne {
		t.Errorf("La source d'eau doit être la bonbonne si réseau indisponible et bonbonne pleine")
	}
	machine.SetSourceEau(true, false)
	if machine.sourceEau != Aucune {
		t.Errorf("La source d'eau doit être aucune si tout est vide")
	}
}

// 8. Bonbonne vide, réseau disponible
func Test_Feature_BonbonneVideReseauDispo(t *testing.T) {
	machine := NewCoffeeMachine(new(MockDriver))
	machine.SetSourceEau(true, true)
	if machine.sourceEau != Reseau {
		t.Errorf("La machine doit utiliser le réseau si bonbonne vide et réseau dispo")
	}
}

// 9. Aucune source d’eau disponible
func Test_Feature_AucuneSourceEau(t *testing.T) {
	driver := new(MockDriver)
	driver.On("BrewCoffee").Return(fmt.Errorf("Pas d'eau disponible")).Once()
	driver.On("ReturnChange", 50).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 50)
}

// 10. Remboursement en cas de dysfonctionnement
func Test_Feature_Dysfonctionnement(t *testing.T) {
	driver := new(MockDriver)
	driver.On("BrewCoffee").Return(fmt.Errorf("Erreur interne")).Once()
	driver.On("ReturnChange", 50).Return(nil).Once()
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	driver.AssertCalled(t, "BrewCoffee")
	driver.AssertCalled(t, "ReturnChange", 50)
}
