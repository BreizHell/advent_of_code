# Day 2: Red-Nosed Reports

## Résumé

- Deux groupes séparé
- On a une matrice de `report`
  - chaque ligne est un `report`
  - chaque colonne est un nombre correspondant à un `level`
- Un report peut être `safe` ou `unsafe`
  - `safe`: les `level` augmente/baisse graduellement
    - Tout les `level` augmente ou diminue
    - Deux `level` adjacent ne peuvent avoir un delta qu'entre 1 et 3 (inclus)

## Partie 2 (en Go !)

- Un report est maintenant considéré comme safe s'il n'y a qu'une seule transition d'unsafe dans toute la série
  - unsafe par delta trop elevé
  - unsafe par changement de direction
