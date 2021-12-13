from pydoc import importfile
f = importfile("../functions.py")

def fillGrid(grid, val):
    return [["x" if el == val else el for el in row] for row in grid]

def bingo(grid):
    for grid in [grid, list(zip(*grid))]:
        for row in grid:
            if row.count("x") == len(row):
                return True
    return False

def bingoResult(grid, num):
    return sum([int(i.replace("x", "0")) for row in grid for i in row]) * int(num)

def Step1():
    with open("input.txt") as f:
        data = {
            "draw_order": f.readline().rstrip().split(","),
            "grids": [],
        }
        f.readline()
        data["grids"] = [[list(filter(None, y.strip().split(" "))) for y in x.rstrip().split("\n")] for x in f.read().split("\n\n")]

    for draw in data["draw_order"]:
        for ind, grid in enumerate(data["grids"]):
            grid = fillGrid(grid, draw)
            data["grids"][ind] = grid 
            win = bingo(grid)
            if win:
                return bingoResult(grid, draw)

def Step2():
    with open("input.txt") as f:
        data = {
            "draw_order": f.readline().rstrip().split(","),
            "grids": [],
        }
        f.readline()
        data["grids"] = [[list(filter(None, y.strip().split(" "))) for y in x.rstrip().split("\n")] for x in f.read().split("\n\n")]

    for draw in data["draw_order"]:
        for ind, grid in enumerate(data["grids"]):
            data["grids"][ind] = fillGrid(grid, draw)
            win = bingo(data["grids"][ind])
            if win:
                if len(data["grids"]) == 1:
                    return bingoResult(data["grids"][0], draw)
                data["grids"][ind] = ""
        data["grids"] = list(filter(None, data["grids"]))
                
if __name__ == "__main__":
    s1 = f.average_time(Step1, iter=1)
    s2 = f.average_time(Step2, iter=1)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")
