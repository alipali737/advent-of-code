from pydoc import importfile

f = importfile("../functions.py")

def Step1():
    with open("input.txt") as f:
        data = [(dir, int(val)) for dir, val in map(str.split, f.read().split("\n"))]
        pos = {"x": 0, "y": 0}
        
        for dir, val in data:
            match dir.upper():
                case "FORWARD":
                    pos["x"] += val
                case "DOWN":
                    pos["y"] += val
                case "UP":
                    pos["y"] -= val

        return pos["x"] * pos["y"]

def Step2():
    with open("input.txt") as f:
        data = [(dir, int(val)) for dir, val in map(str.split, f.read().split("\n"))]
        pos = {"x": 0, "y": 0}
        aim = 0
        
        for dir, val in data:
            match dir.upper():
                case "FORWARD":
                    pos["x"] += val
                    pos["y"] += val * aim
                case "DOWN":
                    aim += val
                case "UP":
                    aim -= val

        return pos["x"] * pos["y"]
    
if __name__ == "__main__":
    s1 = f.average_time(Step1)
    s2 = f.average_time(Step2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")
