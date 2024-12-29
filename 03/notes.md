# Day 3

## Partie 1

- Les données en input sont corrompues
  - La corruption se manifeste par des caractères non-chiffre dans les valeur, et par des instructions invalides
- Le but du programme est de multiplier des nombres
  - On a des instructions comme `mul(x,y)`
    - x/y sont des nombres entre 1 et 3 chiffres
- Il faut ignorer toutes les sections invalides
- Le résultat final est la somme de toute les multiplications

## Partie 2

- 2 nouvelles instructions:
  - `do()`: autorise toute les prochaines instructions
  - `don't()`: annule toute les prochaines instructions
