from collections import Counter
from pydoc import importfile
f = importfile("../functions.py")

# Original
def pairInsertion(polyTemp, pairs, iter=1):
    for _ in range(iter):

        head = 1
        while head <= len(polyTemp) - 1:
            insertion =  pairs[polyTemp[head-1:head+1]]
            polyTemp = f"{polyTemp[:head]}{insertion}{polyTemp[head:]}"

            head += 2
    
    return polyTemp

# New Version (counter for the pairs in input.txt. No need to generate a string)
def pairCounter(polyTemp, pairs, iter=1):
    polyTempList = [polyTemp[i] + polyTemp[i+1] for i in range(len(polyTemp)-1)]
    counter = Counter(polyTempList)

    for _ in range(iter):
        counterAdd = Counter()

        # Add new pair entries and remove old ones
        for pair in counter:
            newPairs = pairs[pair]
            val = 1 * counter[pair]

            counterAdd[pair] -= val
            counterAdd[newPairs[0]] += val
            counterAdd[newPairs[1]] += val

        counter.update(counterAdd)
        del counterAdd

        # Remove 0 value keys
        counter = Counter({key:val for key, val in counter.items() if val > 0})

    return counter

def pairCalculate(counter: Counter, removal: tuple) -> int:
    res = Counter()
    for key, val in counter.items():
        res[key[0]] += val
        res[key[1]] += val
    
    res[removal[0]] += 1
    res[removal[1]] += 1

    res = sorted(res.values())
    return int((res[-1] - res[0])/2)

def part1():
    with open("input.txt") as f:
        polyTemp = f.readline().replace("\n", "")
        f.readline()
        pairs = {i[0]: i[1] for i in [pair.split(" -> ") for pair in f.read().split("\n")]}
    polyTemp = pairInsertion(polyTemp, pairs, iter=10)
    common = sorted(polyTemp.count(char) for char in set(polyTemp))
    return common[-1] - common[0]


def part2():
    with open("input.txt") as f:
        polyTemp = f.readline().replace("\n", "")
        f.readline()
        pairs = {i[0]: [i[0][0] + i[1], i[1] + i[0][1]] for i in [pair.split(" -> ") for pair in f.read().split("\n")]}
    counter = pairCounter(polyTemp, pairs, iter=40)
    return pairCalculate(counter, (polyTemp[0], polyTemp[-1]))

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")