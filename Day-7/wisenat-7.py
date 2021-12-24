from pydoc import importfile
import numpy as np
from math import ceil, floor

f = importfile("../functions.py")

def part1():
    with open("input.txt") as f:
        arr = np.array([int(i) for i in f.read().split(",")])
    med = int(np.median(arr))

    return sum(abs(i - med) for i in arr)


def part2():
    with open("input.txt") as f:
        arr = np.array([int(i) for i in f.read().split(",")])

    av = np.average(arr)
    avgFloor = floor(av)
    avgCeil = ceil(av)
    res = [0, 0]
    for pos in arr:
        for i, avg in ((0, avgFloor), (1, avgCeil)):
            n = abs(pos - avg)
            res[i] += int((n/2) * (n + 1))
            
    return min(res)

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")