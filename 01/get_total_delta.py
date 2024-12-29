with open('input') as input:
  left, right = zip(*([int(n) for n in line.split('   ')] for line in input))
  deltas = [abs(l-r) for l, r in zip(sorted(left), sorted(right))]
  print(sum(deltas))