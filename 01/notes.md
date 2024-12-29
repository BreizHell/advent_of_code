# Day 1: Historian Hysteria

## Résumé global

- `Chief historian` à trouver
- Recherche de `Chief historian` à faire, en fonction de la proba qu'il soit à tel ou tel point
- Chaque point observé sera marqué d'une `star`
- `Chief historian` est dans l'une des 50 première places
- Pour réussir, il faut _50_ `star` avant le 25/12, 1 puzzle résolu donne une `star`
- 2 Puzzles/jours, le premier débloque le second

## Résumé p1d1

- La liste est vide
- Une liste de `location ID` pour des localisation possible est donné
- Deux liste indépendantes (l'input du puzzle)
- Il faut réconcilier les deux listes
  - Pairer les n-ièmes plus gros nombres de gauche et droite
  - Obtenir le delta absolue entre chaque éléments d'une paire
  - renvoyer la somme des écarts

## Résumé p2d1

- On ne sait pas quelle est la bonne manière de lire l'écriture de `Chief historian`
- On multiplie chaque nb de gauche par le nb de fois qu'il apparait à droite
- On renvoie la somme du tout
