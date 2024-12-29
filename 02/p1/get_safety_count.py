from functools import reduce
with open('../input') as input:
  reports = [[int(level) for level in report.split(' ')] for report in input]
  consecutives = [reduce(lambda acc,v: [*acc, (acc[-1][1], v)], r[2:], [tuple(r[0:2])]) for r in reports]

  reports_within_range = {i for i, p in enumerate(consecutives) if all(1 <= abs(a - b) <= 3 for a,b in p)}
  reports_with_directional_consistency = {i for i, pairs in enumerate(consecutives) if all((p[1] - p[0]) > 0 for p in pairs) or all((p[1] - p[0]) < 0 for p in pairs)}

  print(len(reports_within_range & reports_with_directional_consistency))