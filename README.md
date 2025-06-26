# Methodo Test Machine à Café

Ce projet implémente la logique logicielle d'une machine à café connectée, avec gestion des pièces, du hardware, des cas d'erreur et des tests automatisés.

## Lancer le projet

1. **Installer les dépendances**

```sh
go mod tidy
```

2. **Lancer les tests**

```sh
go test -v
```

3. **Compiler et exécuter (exemple)**

```sh
go build -o cafe
./cafe
```

## Structure
- `main.go` : logique principale
- `main_test.go` : tests unitaires
- `PLAN_DE_TEST.md` : plan de test MVP
- `PLAN_DE_TEST_FEATURES.md` : plan de test des features avancées

## Prérequis
- Go 1.18 ou supérieur

## Auteur
- Projet pédagogique Ynov
