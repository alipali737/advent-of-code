from pydoc import importfile
f = importfile("../functions.py")


def cleanInstructions(instructions):
    return [instruction.replace("fold along ", "") for instruction in instructions]


def fold(coords: set[tuple[int]], instuctions: list[str]) -> int:

    for instruction in instuctions: #x=5
        axis, ivalue = instruction.split("=")
        ivalue = int(ivalue)

        axisInd = 0 if axis == "x" else 1
        maxValue = max(list(zip(*coords))[axisInd])

        new_coords = set()
        for coord in coords:
            coord = list(coord)
            coordVal = coord[axisInd]

            # Value on fold line
            if coordVal == ivalue:
                continue
            
            # Value outside of fold line
            if coordVal > ivalue:
                coord[axisInd] = maxValue - coordVal
            
            # Adding coord to new coords
            new_coords.add(tuple(coord))       

        # Updating coords
        coords = new_coords

    return coords


def part1():
    with open("input-data.txt") as f:
        data = set(tuple(map(int, coord.split(","))) for coord in f.read().split("\n"))
    with open("input-instructions.txt") as f:
        instructions = cleanInstructions(instruction for instruction in f.read().split("\n"))
    
    return len(fold(data, instructions))



def part2():
    with open("input-data.txt") as f:
        data = set(tuple(map(int, coord.split(","))) for coord in f.read().split("\n"))
    with open("input-instructions.txt") as f:
        instructions = cleanInstructions(instruction for instruction in f.read().split("\n"))
    
    chars = fold(data, instructions)
    maxRow = max(list(zip(*chars))[1])
    maxCol = max(list(zip(*chars))[0])

    # Grid Generation
    grid = [["#" if (col, row) in chars else " " for col in range(maxCol + 1)] for row in range(maxRow + 1)]
    for row in grid:
        print("".join(map(str, row)))

if __name__ == "__main__":
    s1 = f.average_time(part1, iter=1)
    s2 = f.average_time(part2, iter=1)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")