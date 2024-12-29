from collections import Counter

with open('input') as input:
  left, right = zip(*([int(n) for n in line.split('   ')] for line in input))
  left_set, right_counts = set(left), Counter(right)
  print(sum(l * right_counts[l] for l in left_set))