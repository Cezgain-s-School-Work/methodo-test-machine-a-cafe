package main

import (
	"fmt"
	"testing"
)

// 1. Achat réussi d’un café (Happy Path)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on insère une pièce de 50cts ou plus
// ALORS le brewer reçoit l’ordre de faire un café
// ET la monnaie excédentaire est restituée
func Test_AchatReussiCafe(t *testing.T) {
	calledBrew := false
	driver := &TestDriver{
		IsDefectiveFunc: func() bool { return false },
		BrewCoffeeFunc:  func() error { calledBrew = true; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	if !calledBrew {
		t.Error("BrewCoffee n'a pas été appelé")
	}
}

// 2. Pièce insuffisante
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on insère une pièce de moins de 50cts (1cts, 2cts, 5cts, 10cts, 20cts)
// ALORS le brewer ne reçoit pas d’ordre
// ET l’argent est restitué
func Test_PieceInsuffisante(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return nil },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(20)
	if calledBrew {
		t.Error("BrewCoffee ne doit pas être appelé")
	}
	if !calledReturn || amountReturned != 20 {
		t.Errorf("ReturnChange doit être appelé avec 20, obtenu %d", amountReturned)
	}
}

// 3. Machine défaillante
// Gherkin :
// ÉTANT DONNÉ une machine à café défaillante
// QUAND on insère une pièce de 50cts ou plus
// ALORS l’argent est restitué
// ET le brewer ne reçoit pas d’ordre
func Test_MachineDefaillante(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return true },
		BrewCoffeeFunc:   func() error { calledBrew = true; return nil },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(100)
	if calledBrew {
		t.Error("BrewCoffee ne doit pas être appelé si défaillant")
	}
	if !calledReturn || amountReturned != 100 {
		t.Errorf("ReturnChange doit être appelé avec 100, obtenu %d", amountReturned)
	}
}

// 4. Insertion d’un montant supérieur au prix (1€)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on insère une pièce de 1€
// ALORS le brewer reçoit l’ordre de faire un café
// ET 50cts sont restitués
func Test_Insertion1Euro(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return nil },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(100)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 50 {
		t.Errorf("ReturnChange doit être appelé avec 50, obtenu %d", amountReturned)
	}
}

// 5. Insertion d’un montant non multiple de 50cts (2€)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on insère une pièce de 2€
// ALORS le brewer reçoit l’ordre de faire un café
// ET 1,50€ sont restitués
func Test_Insertion2Euros(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return nil },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(200)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 150 {
		t.Errorf("ReturnChange doit être appelé avec 150, obtenu %d", amountReturned)
	}
}

// 6. Machine défaillante à tout moment
// Gherkin :
// ÉTANT DONNÉ une machine à café défaillante
// QUAND on insère n'importe quelle pièce
// ALORS l’argent est restitué
// ET le brewer ne reçoit pas d’ordre
func Test_MachineDefaillanteAnyCoin(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return true },
		BrewCoffeeFunc:   func() error { calledBrew = true; return nil },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(5)
	if calledBrew {
		t.Error("BrewCoffee ne doit pas être appelé si défaillant")
	}
	if !calledReturn || amountReturned != 5 {
		t.Errorf("ReturnChange doit être appelé avec 5, obtenu %d", amountReturned)
	}
}

// 7. Insertion d'une pièce de 20 centimes
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on insère une pièce de 20cts
// ALORS l’argent est restitué
// ET le brewer ne reçoit pas d’ordre
func Test_Piece20Centimes(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return nil },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(20)
	if calledBrew {
		t.Error("BrewCoffee ne doit pas être appelé")
	}
	if !calledReturn || amountReturned != 20 {
		t.Errorf("ReturnChange doit être appelé avec 20, obtenu %d", amountReturned)
	}
}

// 8. Pas de café disponible
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// ET le réservoir de café est vide
// QUAND on insère une pièce de 50cts
// ALORS un message d'erreur est affiché
// ET l’argent est restitué
func Test_PasDeCafeDisponible(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return fmt.Errorf("Plus de café") },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 50 {
		t.Errorf("ReturnChange doit être appelé avec 50, obtenu %d", amountReturned)
	}
}

// 0. Stock de café vide
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// ET le réservoir de café est vide
// QUAND on insère une pièce de 50cts
// ALORS un message d'erreur est affiché
// ET l’argent est restitué
func Test_Feature_StockCafeVide(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return fmt.Errorf("Plus de café") },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 50 {
		t.Errorf("ReturnChange doit être appelé avec 50, obtenu %d", amountReturned)
	}
}

// 2. Pas de gobelet disponible (détecté par absence de récipient)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// ET aucun récipient n'est détecté
// QUAND on insère une pièce de 50cts
// ALORS un message d'erreur est affiché
// ET l’argent est restitué
func Test_Feature_AbsenceRecipient(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return fmt.Errorf("Aucun récipient détecté") },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 50 {
		t.Errorf("ReturnChange doit être appelé avec 50, obtenu %d", amountReturned)
	}
}

// 2b. Prix d’un gobelet (vérification du calcul)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on consulte le prix d'un café avec gobelet
// ALORS le prix doit être supérieur de 10cts au prix sans gobelet
func Test_Feature_PrixGobelet(t *testing.T) {
	machine := NewCoffeeMachine(&TestDriver{})
	prixSansTasse := machine.PrixCommande(false)
	prixAvecTasse := machine.PrixCommande(true)
	if prixSansTasse != 60 {
		t.Errorf("Le prix avec gobelet doit être 60, obtenu %d", prixSansTasse)
	}
	if prixAvecTasse != 50 {
		t.Errorf("Le prix avec tasse doit être 50, obtenu %d", prixAvecTasse)
	}
}

// 5. Dosage café normal (1 dose d'eau, café boolean)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on sélectionne le dosage normal
// ALORS la machine doit utiliser 1 dose d'eau et du café
func Test_Feature_DosageCafeNormal(t *testing.T) {
	machine := NewCoffeeMachine(&TestDriver{})
	machine.SetDosage(false)
	if machine.doseEau != 1 || !machine.doseCafe {
		t.Errorf("Le dosage normal doit être 1 dose d'eau et café true, obtenu %d/%v", machine.doseEau, machine.doseCafe)
	}
}

// 6. Dosage café allongé (2 doses d'eau, café boolean)
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on sélectionne le dosage allongé
// ALORS la machine doit utiliser 2 doses d'eau et du café
func Test_Feature_DosageCafeAllonge(t *testing.T) {
	machine := NewCoffeeMachine(&TestDriver{})
	machine.SetDosage(true)
	if machine.doseEau != 2 || !machine.doseCafe {
		t.Errorf("Le dosage allongé doit être 2 doses d'eau et café true, obtenu %d/%v", machine.doseEau, machine.doseCafe)
	}
}

// 7. Sélection automatique de la source d’eau
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// QUAND on vérifie la source d'eau disponible
// ALORS la machine doit sélectionner le réseau s'il est disponible,
// sinon la bonbonne si elle est pleine,
// sinon aucune source
func Test_Feature_SelectionSourceEau(t *testing.T) {
	machine := NewCoffeeMachine(&TestDriver{})
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
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// ET la bonbonne est vide
// QUAND on insère une pièce
// ALORS la machine doit utiliser le réseau pour faire le café
func Test_Feature_BonbonneVideReseauDispo(t *testing.T) {
	machine := NewCoffeeMachine(&TestDriver{})
	machine.SetSourceEau(true, true)
	if machine.sourceEau != Reseau {
		t.Errorf("La machine doit utiliser le réseau si bonbonne vide et réseau dispo")
	}
}

// 9. Aucune source d’eau disponible
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// ET aucune source d'eau n'est disponible
// QUAND on insère une pièce de 50cts
// ALORS un message d'erreur est affiché
// ET l’argent est restitué
func Test_Feature_AucuneSourceEau(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return fmt.Errorf("Pas d'eau disponible") },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 50 {
		t.Errorf("ReturnChange doit être appelé avec 50, obtenu %d", amountReturned)
	}
}

// 10. Remboursement en cas de dysfonctionnement
// Gherkin :
// ÉTANT DONNÉ une machine à café fonctionnelle
// ET un dysfonctionnement est détecté
// QUAND on insère une pièce de 50cts
// ALORS un message d'erreur est affiché
// ET l’argent est restitué
func Test_Feature_Dysfonctionnement(t *testing.T) {
	calledBrew := false
	calledReturn := false
	amountReturned := 0
	driver := &TestDriver{
		IsDefectiveFunc:  func() bool { return false },
		BrewCoffeeFunc:   func() error { calledBrew = true; return fmt.Errorf("Erreur interne") },
		ReturnChangeFunc: func(amount int) error { calledReturn = true; amountReturned = amount; return nil },
	}
	machine := NewCoffeeMachine(driver)
	machine.InsertCoin(50)
	if !calledBrew {
		t.Error("BrewCoffee doit être appelé")
	}
	if !calledReturn || amountReturned != 50 {
		t.Errorf("ReturnChange doit être appelé avec 50, obtenu %d", amountReturned)
	}
}
