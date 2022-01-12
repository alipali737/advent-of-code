from copy import deepcopy
from pydoc import importfile
f = importfile("../functions.py")

class Node:
    def __init__(self, reference, connections = None):
        self.connections = connections if connections else set() 
        self.reference = reference
    
    def addCon(self, node):
        self.connections.add(node)

    def addCons(self, nodes):
        for node in nodes:
            self.addCon(node)

    def __repr__(self):
        return self.reference

    def __str__(self):
        strConnections = [con.reference for con in self.connections]
        return f"{self.reference} -> {', '.join(strConnections)}"
    
class Graph:
    def __init__(self, start = Node("start"), end = Node("end")):
        self.nodes = {"start": start, "end": end}
    
    def generateGraph(self, data):
        for headRef, tailRef in data:
            
            def genNode(ref):
                node = Node(ref)
                self.nodes[ref] = node
                return node

            if (headNode := self.nodes.get(headRef)) is None:
                headNode = genNode(headRef)
            if (tailNode := self.nodes.get(tailRef)) is None:
                tailNode = genNode(tailRef)

            # Add node connections
            headNode.addCon(tailNode)
            tailNode.addCon(headNode)

    
 
    def DFS(self, canVisit):
        '''
        Modified Depth-first Search algorithm to dictate routes.
        1. Traverse each adjacent node
        2. If the current node (not the adjacent) is a small cave, pass itself through visited
        3. If the node is "end", add one to routes
        '''
        def DFSUtil(node: Node, routes, visited, canVisit):
            if node.reference == "end":
                routes[0] += 1
                return
            
            visited.append(node)

            # Recur for all the connected nodes
            for neighbour in node.connections:
                if canVisit(neighbour, visited):
                    DFSUtil(neighbour, routes, visited[:], canVisit)

        # routes is a list to mimic pass by reference.
        # It's mutable object and will update its value without needing a return in DFSUtil
        routes = [0]
        visited = []
        DFSUtil(self.nodes["start"], routes, visited, canVisit)

        return routes[0]
    
def part1():
    with open("input.txt") as f:
        data = [i.split("-") for i in f.read().split("\n")]
    cave = Graph()
    cave.generateGraph(data)

    def canVisit(node: Node, visited):
        return node.reference.isupper() or node not in visited

    return cave.DFS(canVisit)

def part2():
    with open("input.txt") as f:
        data = [i.split("-") for i in f.read().split("\n")]
    cave = Graph()
    cave.generateGraph(data)

    def canVisit(node: Node, visited: list):
        if node.reference == "start":
            return False
        if node.reference.isupper():
            return True

        visitedSet = set([node for node in visited if not node.reference.isupper()])
        visitedCount = [val for el in visitedSet if (val := visited.count(el)) > 1]
        if len(visitedCount) == 0:
            return True
        if len(visitedCount) == 1 and visitedCount[0] == 2:
            return True
        return False

        
    return cave.DFS(canVisit)
    

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")