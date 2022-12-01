from os import remove, sep
from pydoc import importfile
from copy import deepcopy

f = importfile("../functions.py")

def invertBin(val):
    return val.replace("0", "a").replace("1", "0").replace("a", "1")

def removeElements(matrix, pos, val = "0"):
    return list(filter(lambda x: x[pos] != val, matrix))

def Step1():
    with open("input.txt") as f:
        # [0010, 1010, 1111] -> [[0, 1, 1], [0, 0, 1], [1, 1, 1], [0, 0, 1]]
        data = list(zip(*[i for i in f.read().split("\n")]))
        gamma = ""
        for column in data:
            if column.count("0") > column.count("1"):
                gamma += "0"
            else:
                gamma += "1"

        return int(gamma, 2) * int(invertBin(gamma), 2)

def Step2():
    with open("input.txt") as f:
        # [0010, 1010, 1111] -> [[0, 1, 1], [0, 0, 1], [1, 1, 1], [0, 0, 1]]
        buffer_data = [i for i in f.read().split("\n")]
        
        vals = []
        for bit in range(2):
            data = deepcopy(buffer_data)
            for pos in range(len(data[0])):
                if len(data) == 1:
                    break

                sep_data = list(zip(*data))
                column = sep_data[pos]

                zeroCount = column.count("0")
                oneCount = len(column) - zeroCount

                if zeroCount == oneCount:
                    data = removeElements(data, pos, val=str(bit))
                else:
                    higher = str(bit if zeroCount < oneCount else 1 - bit)
                    data = removeElements(data, pos, val=higher)

            vals.append(int(data[0], 2))

        return vals[0] * vals[1]
    
if __name__ == "__main__":
    s1 = f.average_time(Step1)
    s2 = f.average_time(Step2)

    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")
