from pydoc import importfile
f = importfile("functions.py")

def Step1():
    horizontal = depth = 0
    with open("./Day-2/input.txt") as f:
        for line in f.readlines():
            line = line.split()
            if line[0] == "forward":
                horizontal += int(line[1])
            elif line[0] == "down":
                depth += int(line[1])
            else:
                depth -= int(line[1])
    return horizontal * depth

def Step2():
    horizontal = depth = aim = 0
    with open("./Day-2/input.txt") as f:
        for line in f.readlines():
            line = line.split()
            if line[0] == "forward":
                horizontal += int(line[1])
                depth += int(line[1])*aim
            elif line[0] == "down":
                aim += int(line[1])
            else:
                aim -= int(line[1])
    return horizontal * depth

    
if __name__ == "__main__":
    print(f.average_time(Step1))
    print(f.average_time(Step2))
