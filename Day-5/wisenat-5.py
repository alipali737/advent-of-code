from os import linesep
from pydoc import importfile
import pandas as pd
import numpy
f = importfile("../functions.py")

def printGrid(grid):
    for row in grid:
        print(" ".join(map(str, row)))

def gridToCsv(grid):
    df = pd.DataFrame(numpy.array(grid).transpose())
    df.to_csv("grid.csv", index=False)

def generateBoundingGrid(coords):
    bounding = numpy.max(coords, 0)
    return [[0 for x in range(bounding[0] + 1)] for y in range(bounding[1] + 1)]

def addLines(grid, lines):
    # 1 Iteration = 1 Line
    for head, tail in lines:
        def genIterSel(c1, c2, val):
            a = c1[1 - val]
            b = c2[1 - val]
            iteration = range(min(a, b), max(a, b) + 1)
            selection = [c1[val] for _ in iteration]
            return iteration, selection

        # X values are equal... iterate through Y in matrix selection of X
        if head[0] == tail[0]:
            y, x = genIterSel(head, tail, 0)
        # Y values are equal... iterate through X in matrix selection of Y
        elif head[1] == tail[1]:
            x, y = genIterSel(head, tail, 1)
        else:
            # print("Not a straight line..")
            continue
        
        for x, y in zip(*[x, y]):
            grid[y][x] += 1

    return grid

def part1():
    '''
    1. Read in all the data in a nice format: [[head_coord1, tail_coord1], [head_coord2, tail_coord2], ...]
    2. Find biggest value for X and Y in data, generate an empty grid using these values
    3. Iterate through data, appending lines to the grid
    4. Count values of 0 and 1 and subtract them from the length*width of the grid'''

    with open("input.txt") as f:
        lines = numpy.array([[list(map(int, coords.split(","))) for coords in line.split(" -> ")] for line in f.read().split("\n")])
    
    coords = lines.reshape(-1, lines.shape[-1])
    grid = generateBoundingGrid(coords)
    grid = addLines(grid, lines)
    # gridToCsv(grid)
    
    result = len(grid) * len(grid[0]) - sum(row.count(0) + row.count(1) for row in grid)
    
    return result
    

def part2():
    return None


if __name__ == "__main__":
    s1 = f.average_time(part1, iter=1)
    # s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    # print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")