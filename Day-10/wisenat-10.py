from pydoc import importfile
f = importfile("../functions.py")

def emptyFunc():
    pass

def scanLine(string, syntax=False):
    OPENERS = {"(": ")", "[": "]", "{": "}", "<": ">"}
    SYNTAX_SCORE = {")": 3, "]": 57, "}": 1197, ">": 25137}
    AUTO_SCORE = {")": 1, "]": 2, "}": 3, ">": 4}

    default = 0 if syntax != None else None
    opcode = []
    for opr in string:
        # Open chunk operator
        if opr in OPENERS.keys():
            opcode.append(opr)
        # Close chunk operator
        else:
            # Valid
            if len(opcode) != 0 and OPENERS[opcode[-1]] == opr:
                del opcode[-1]
            # Corrupt
            else:
                if syntax:
                    return SYNTAX_SCORE[opr]
                else:
                    return None
    
    # Incomplete
    if len(opcode) != 0 and not syntax:
        opcode.reverse()
        for missingCloser in opcode:
            default = default * 5 + AUTO_SCORE[OPENERS[missingCloser]]

    return default

def printGrid(grid):
    for row in grid:
        print(" ".join([i.replace("0", ".") for i in map(str, row)]))


def Step1():
    with open("input.txt") as f:
        data = [line for line in f.read().split("\n")]

    score = 0
    for line in data:
        score += scanLine(line, syntax=True)

    return score

def Step2():
    with open("input.txt") as f:
        data = [line for line in f.read().split("\n")]
    
    score = []
    for line in data:
        val = scanLine(line)
        if val != None:
            score.append(val)
    
    return sorted(score)[round(len(score) / 2)]

if __name__ == "__main__":
    s1 = f.average_time(Step1)
    s2 = f.average_time(Step2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")
