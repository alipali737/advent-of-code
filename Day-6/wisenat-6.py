from pydoc import importfile
import numpy as np
f = importfile("../functions.py")

def part1():
    '''
    Part 1 sucks but I'm not going to change it to the updated version
    '''
    with open("input.txt") as f:
        LanternFish = [int(i) for i in f.read().split(",")]
        print(f"Initial:\t{LanternFish}")
        for day in range(1, 80 + 1):
            LanternFish = [i - 1 for i in LanternFish]
            for ind, el in enumerate(LanternFish):
                if el == -1:
                    LanternFish[ind] = 6
                    LanternFish.append(8)
            # print(f"Day {day}:\t\t{LanternFish}")

    return len(LanternFish)


def part2():
    '''
    1. Create list with the index as fish lifespans, elements as amount of fish with said lifespan 
    2. Grab reproducing fish
    3. Reduce all fish lifespans down by 1 (reproducing moved to lifespan 8)
        [1, 2, 3, 4, 5, 6, 7, 8] -> [2, 3, 4, 5, 6, 7, 8, 1]
    4. Add all reproducing fish to lifespan 6
        [2, 3, 4, 5, 6, 7 + 1, 8, 1]
    ''' 
    with open("input.txt") as f:
        inp = [int(i) for i in f.read().split(",")]
        LanternFish = np.array([inp.count(ind) for ind in range(8 + 1)], dtype=np.int64)
        days = 256
        
        # Model time
        for _ in range(1, days + 1):
            LanternFish = np.roll(LanternFish, -1)
            LanternFish[6] += LanternFish[8]

    print(LanternFish)
    return sum(LanternFish)
    

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")