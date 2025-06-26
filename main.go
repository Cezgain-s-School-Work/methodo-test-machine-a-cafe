package main

import "fmt"

const prixCafe = 50
const prixGobelet = 10

// Interface du driver matériel (mockable)
type CoffeeMachineDriver interface {
	BrewCoffee() error
	ReturnChange(amount int) error
	IsDefective() bool
}

type SourceEau int

const (
	Aucune SourceEau = iota
	Bonbonne
	Reseau
)

type CoffeeMachine struct {
	driver    CoffeeMachineDriver
	doseEau   int
	doseCafe  bool
	sourceEau SourceEau
}

func NewCoffeeMachine(driver CoffeeMachineDriver) *CoffeeMachine {
	return &CoffeeMachine{driver: driver, doseEau: 1, doseCafe: true, sourceEau: Reseau}
}

// Calcule le prix total selon la présence d'une tasse
func (m *CoffeeMachine) PrixCommande(avecTasse bool) int {
	if avecTasse {
		return prixCafe
	}
	return prixCafe + prixGobelet
}

// Définit le dosage pour café normal ou allongé
func (m *CoffeeMachine) SetDosage(allonge bool) {
	if allonge {
		m.doseEau = 2
	} else {
		m.doseEau = 1
	}
	m.doseCafe = true
}

// Définit la source d'eau à utiliser
func (m *CoffeeMachine) SetSourceEau(bonbonneVide, reseauDispo bool) {
	if reseauDispo {
		m.sourceEau = Reseau
	} else if !bonbonneVide {
		m.sourceEau = Bonbonne
	} else {
		m.sourceEau = Aucune
	}
}

// Insertion d'une pièce (en centimes)
func (m *CoffeeMachine) InsertCoin(amount int) {
	if m.driver.IsDefective() {
		_ = m.driver.ReturnChange(amount)
		return
	}
	if amount >= prixCafe {
		err := m.driver.BrewCoffee()
		if err != nil {
			_ = m.driver.ReturnChange(amount)
			return
		}
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
