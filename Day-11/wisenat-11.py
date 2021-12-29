from pydoc import importfile
f = importfile("../functions.py")

def printGrid(grid):
    print("\n".join(["".join(map(str, row)) for row in grid]))

def isValidPos(matrix, pos):
    row = pos[0]
    ind = pos[1]
    if 0 <= row < len(matrix) and 0 <= ind < len(matrix[row]):
        return True
    return False

def getRelativeElement(pos, posInc, matrix):
    row = pos[0] + posInc[0]
    ind = pos[1] + posInc[1]
    if not isValidPos(matrix, (row, ind)):
        return None, None

    return matrix[row][ind], (row, ind)

def part1():
    with open("input.txt") as f:
        data = [list(map(int, list(i))) for i in f.read().split("\n")]
    
    flashes = 0
    for steps in range(100):
        # Step 1: Increment, find flashing
        data = list(map(lambda row: list(map(lambda el: el + 1, row)), data))
        flashing = {(indRow, indEl) for indRow, row in enumerate(data) for indEl, octopus in enumerate(row) if octopus == 10}

        # Flash: increment all octopus around (unless they equal 0), if an octopus becomes > 9 add it to the set
        while flashing:
            flashes += 1
            indRow, indEl = flashing.pop()
            data[indRow][indEl] = 0

            # Left, Top, Right, Bottom, TL, TR, BL, BR
            for inc in ((0, -1), (-1, 0), (0, 1), (1, 0), (-1, -1), (-1, 1), (1, -1), (1, 1)):
                cardinalEl, cardinalPos = getRelativeElement((indRow, indEl), inc, data)
                if cardinalEl in (None, 0):
                    continue

                iRow, iEl = cardinalPos

                data[iRow][iEl] += 1
                if data[iRow][iEl] == 10:
                    flashing.add(cardinalPos)

    return flashes

def part2():
    with open("input.txt") as f:
        data = [list(map(int, list(i))) for i in f.read().split("\n")]
    
    step = 0
    while True:
        step += 1
        # Step 1: Increment, find flashing
        data = list(map(lambda row: list(map(lambda el: el + 1, row)), data))
        flashing = {(indRow, indEl) for indRow, row in enumerate(data) for indEl, octopus in enumerate(row) if octopus == 10}
        totalFlashed = len(flashing)

        # Flash: increment all octopus around (unless they equal 0), if an octopus becomes > 9 add it to the set
        while flashing:
            indRow, indEl = flashing.pop()
            data[indRow][indEl] = 0

            # Left, Top, Right, Bottom, TL, TR, BL, BR
            for inc in ((0, -1), (-1, 0), (0, 1), (1, 0), (-1, -1), (-1, 1), (1, -1), (1, 1)):
                cardinalEl, cardinalPos = getRelativeElement((indRow, indEl), inc, data)
                if cardinalEl in (None, 0):
                    continue

                iRow, iEl = cardinalPos

                data[iRow][iEl] += 1
                if data[iRow][iEl] == 10:
                    flashing.add(cardinalPos)
                    totalFlashed += 1
        
        if totalFlashed == len(data) * len(data[0]):
            return step

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")