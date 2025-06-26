# Plan de test – Machine à café (MVP)

## Objectif
Valider le comportement de la machine à café selon les exigences du MVP, en couvrant le happy path et les edge cases.

## Scénarios Gherkin

### 1. Achat réussi d’un café (Happy Path)
```gherkin
ÉTANT DONNÉ une machine à café fonctionnelle
QUAND on insère une pièce de 50cts ou plus
ALORS le brewer reçoit l’ordre de faire un café
ET la monnaie excédentaire est restituée
```

### 2. Pièce insuffisante
```gherkin
ÉTANT DONNÉ une machine à café fonctionnelle
QUAND on insère une pièce de moins de 50cts (1cts, 2cts, 5cts, 10cts, 20cts)
ALORS le brewer ne reçoit pas d’ordre
ET l’argent est restitué
```

### 3. Machine défaillante
```gherkin
ÉTANT DONNÉ une machine à café défaillante
QUAND on insère une pièce de 50cts ou plus
ALORS l’argent est restitué
ET le brewer ne reçoit pas d’ordre
```

### 4. Insertion de plusieurs pièces
```gherkin
ÉTANT DONNÉ une machine à café fonctionnelle
QUAND on insère deux fois une pièce de 50cts
ALORS le brewer reçoit deux fois l’ordre de faire un café
ET aucune monnaie n’est restituée
```

### 5. Insertion d’un montant supérieur au prix
```gherkin
ÉTANT DONNÉ une machine à café fonctionnelle
QUAND on insère une pièce de 1€
ALORS le brewer reçoit l’ordre de faire un café
ET 50cts sont restitués
```

### 6. Insertion d’un montant non multiple de 50cts
```gherkin
ÉTANT DONNÉ une machine à café fonctionnelle
QUAND on insère une pièce de 2€
ALORS le brewer reçoit l’ordre de faire un café
ET 1,50€ sont restitués
```

### 7. Machine défaillante à tout moment
```gherkin
ÉTANT DONNÉ une machine à café défaillante
QUAND on insère n’importe quelle pièce
ALORS l’argent est restitué
ET le brewer ne reçoit pas d’ordre
```

## Remarques
- La machine n’a pas d’affichage.
- Elle accepte uniquement les pièces de 1cts à 2€.
- Elle rend la monnaie si nécessaire.
- Pas d’annulation ni de gestion de stock pour le MVP.
- Le driver gère la restitution de monnaie, l’envoi de commandes au brewer, et la détection des pannes.
