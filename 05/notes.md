# Advent of code J5

## P1

- Un ensemble de pages est à imprimer
- Notre problème:
  - On a un ensemble de lignes définissants des contrainte d'antériorité sur des pages, lorsqu'elles sont impliqués dans une MAJ
  - `X|Y` veut dire que la page `X` est à imprimer antérieurement à `Y`
  - Une MAJ est une ligne sous la forme `X,Y,Z,W` (longueur variable)
  - Une MAJ est considérée valide si l'ordre des pages qu'elle énumère ne viole aucune des règles d'ordonnancement
  - La réponse est la somme de la page du milieu de chaque MAJ valide

### Idées

- Dictionnaire de règles ?

  - K = une page
  - V = Précédance avec autres pages ?

- Algo
  - Pour chaque élément dans la liste, regarder s'il y'a au moins un élément sur sa gauche qui devrait lui être postérieur. Puis regarder sur sa droit s'il y a un élément qui devrait lui être antérieur
    - Si aucune violation, `somme += valDeLelemDuMilieu`
  - Divide & conquer ?
  - Faire un sort préalable des rêgles de prio ?
    => Non, on n'a aucune garantie que certaines rêgles ne soit pas contradictoire (ie. Certaine combinaison de pages seraient toujours invalides, peut importe leur ordonnancement)
