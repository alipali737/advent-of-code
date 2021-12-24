from pydoc import importfile
import pandas as pd
import numpy as np
f = importfile("../functions.py")

def gridToCsv(grid):
    df = pd.DataFrame(np.array(grid))
    df.to_csv("grid.csv", index=False, header=False)

def getRelativeElement(pos, posInc, matrix):
    row = pos[0] + posInc[0]
    ind = pos[1] + posInc[1]
    if not isValidPos(matrix, (row, ind)):
        return None, None

    return matrix[row][ind], (row, ind)

def isValidPos(matrix, pos):
    row = pos[0]
    ind = pos[1]
    if 0 <= row < len(matrix) and 0 <= ind < len(matrix[row]):
        return True
    return False

def floodFill(matrix, pos):
    toFill = {pos}
    total = 0

    while toFill:
        (row, ind) = toFill.pop()
        
        if not isValidPos(matrix, (row, ind)):
            continue
        
        el = matrix[row][ind]
        # print(f"({row}, {ind}) -> {el}")

        if el in [9, "."]:
            continue
        
        total += 1
        matrix[row][ind] = "."

        toFill.add((row - 1, ind))
        toFill.add((row + 1, ind))
        toFill.add((row, ind - 1))
        toFill.add((row, ind + 1))
    
    return total

def part1():
    '''
    1. Generate num grid and marked grid
    2. Start from top left (0, 0); iterating left to right, top to bottom
    3. Check whether the cardinal elements are smaller than the target element
        a. Cardinal element smaller? Mark target element
        b. Cardinal element bigger? Mark cardinal element
    4. Target element marked?
        a. No -> add it's value + 1 to the total (var)
    5. Move on to next element that isn't marked
    6. Continue until you reach the end of grid  
    '''
    
    # Read in grid data
    with open("input.txt") as f:
        numberGrid = [list(map(int, list(i))) for i in f.read().split("\n")]
        markedGrid = [["" for x in range(len(numberGrid[0]))] for y in range(len(numberGrid))]
    
    MARKED = "X"
    total = 0

    # Iterate through all elements in numberGrid
    for rowInd, row  in enumerate(numberGrid):
        for elInd, el in enumerate(row):
            # Skip target element if marked
            if markedGrid[rowInd][elInd] == MARKED:
                continue
            
            # Pull all cardinal elements relative to target element [Left, Top, Right, Bottom] and check them 
            for pos in ((0, -1), (-1, 0), (0, 1), (1, 0)):
                cardinalEl, cardinalPos = getRelativeElement((rowInd, elInd), pos, numberGrid)
                if cardinalEl == None:
                    continue

                if cardinalEl > el:
                    markedGrid[cardinalPos[0]][cardinalPos[1]] = MARKED
                else:
                    markedGrid[rowInd][elInd] = MARKED
            
            # Checking if Target Element is marked (low point)
            if markedGrid[rowInd][elInd] != MARKED:
                total += el + 1
            
    return total


def part2():
    # Read in grid data
    with open("input.txt") as f:
        numberGrid = [list(map(int, list(i))) for i in f.read().split("\n")]
        markedGrid = [["" for x in range(len(numberGrid[0]))] for y in range(len(numberGrid))]
    
    MARKED = "X"
    basins = set()

    # Iterate through all elements in numberGrid
    for rowInd, row  in enumerate(numberGrid):
        for elInd, el in enumerate(row):
            # Skip target element if marked
            if markedGrid[rowInd][elInd] == MARKED:
                continue
            
            # Pull all cardinal elements relative to target element [Left, Top, Right, Bottom] and check them 
            for pos in ((0, -1), (-1, 0), (0, 1), (1, 0)):
                cardinalEl, cardinalPos = getRelativeElement((rowInd, elInd), pos, numberGrid)
                if cardinalEl == None:
                    continue

                if cardinalEl > el:
                    markedGrid[cardinalPos[0]][cardinalPos[1]] = MARKED
                else:
                    markedGrid[rowInd][elInd] = MARKED
            
            # Checking if Target Element is marked (low point)
            if markedGrid[rowInd][elInd] != MARKED:
                basins.add((rowInd, elInd))

    markedGrid = numberGrid[:]
    result = np.product(np.array(sorted([floodFill(markedGrid, basin) for basin in basins], reverse=True)[:3]))
    return result
    


if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")