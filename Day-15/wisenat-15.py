from heapq import heapify, heappop, heappush
from itertools import chain
import numpy as np
from pydoc import importfile
f = importfile("../functions.py")

class Node:
    def __init__(self, risk, minPath):
        self.risk = risk
        self.minPath = minPath
    
    def __add__(self, obj):
        self.risk += obj
        return self
    
    def __radd__(self, other):
        if other == 0:
            return self
        else:
            return self.__add__(other)

    def __mod__(self, other):
        self.risk = self.risk % other
        return self

def getNeighbours(node: list[int], grid: list[list]) -> list[Node]:
    neighbours = []
    # y,x -> L, T, R, B
    for incX, incY in ((0, -1), (-1, 0), (0, 1), (1, 0)):
        totalPos = (node[0] + incX, node[1] + incY)

        if 0 <= totalPos[0] < len(grid) and 0 <= totalPos[1] < len(grid[totalPos[0]]):
            neighbours.append(totalPos)
    
    return neighbours

def posInQueue(pos: list[int], queue: list[tuple]) -> bool:
    if len(queue) != 0 and pos in list(zip(*queue))[1]:
        return True
    return False

def getEl(pos: tuple[int], grid: list[list]) -> Node:
    return grid[pos[1]][pos[0]]

def dijkstras(grid: list[list[Node]], source, sink) -> int:
    # Initialise Visited, Visiting and Source values
    # Add source to Visiting
    # Iterate through Visiting while Nodes exist within
    #   Grab Node (cNode) with smallest path to it
    #   Add cNode to Visited
    #   Grab the Neighbours of the cNode and update their minPaths if necessary
    #   Add Neighbours to Visiting if they haven't been Visited and aren't already in Visiting
    # Keep going till Visiting is empty
    visited = set()

    visiting = [(0, source)]
    heapify(visiting)
    
    sourceNode = getEl(source, grid)
    sourceNode.minPath = sourceNode.risk 
    
    while visiting:
        cPos = heappop(visiting)[1]
        cNode = getEl(cPos, grid)

        visited.add(cPos)

        # Iterate through neighbour nodes..
        neighbours = getNeighbours(cPos, grid)
        for nPos in neighbours:
            nNode = getEl(nPos, grid)
            if (path := cNode.minPath + nNode.risk) < nNode.minPath:
                nNode.minPath = path
            
            if not posInQueue(nPos, visiting) and nPos not in visited:
                heappush(visiting, (nNode.minPath, nPos))
    
    return getEl(sink, grid).minPath


def part1():
    with open("input.txt") as f:
        data = f.read().split("\n")
        maxVal = len(data) * len(data[0]) * 9
        grid = np.array([[Node(int(i), maxVal) for i in row] for row in data])

    source = (0, 0)
    sink = (len(grid) - 1, len(grid[0]) - 1)

    return dijkstras(grid, source, sink) - getEl(source, grid).risk

def part2():
    '''
    My part 2 is extremely slow. Not because of the grid generation logic below,
    that's surprisingly quick. Dijkstras just kinda says no with how I've done
    it.

    Most of AoC is just learning for me anyway - not necesssarily trying to do
    everything fast. Here's a good example I found of how to do this well:
    https://todd.ginsberg.com/post/advent-of-code/2021/day15/
    '''

    with open("input.txt") as f:
        data = f.read().split("\n")
    
    # Initial Array logic - increment and original
    initGrid = np.array([[int(i) for i in row] for row in data])
    incGrid = {val: np.full((len(initGrid), len(initGrid[0])), val) for val in range(8 + 1)}
    maxVal = len(data) * len(data[0]) * 9

    # Generating larger grid
    buffer = [[[] for _ in range(5)] for _ in range(5)]
    for row in range(5):
        for col in range(5):
            buffer[row][col] = initGrid + incGrid[row + col]
    buffer = np.where((v := np.mod(buffer, 9)) == 0, 9, v)

    grid = [list(chain(*combined)) for row in buffer for combined in zip(*row)]
    grid = [[Node(i, maxVal) for i in row] for row in grid]
    
    # Final value decs for source/sink Nodes
    source = (0, 0)
    sink = (len(grid) - 1, len(grid[0]) - 1)

    return dijkstras(grid, source, sink) - getEl(source, grid).risk
   

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")