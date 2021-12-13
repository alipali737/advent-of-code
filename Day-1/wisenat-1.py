from pydoc import importfile
f = importfile("../functions.py")

def increase_counter(matrix):
    counter = 0
    for i in range(len(matrix) - 1):
        if matrix[i] < matrix[i + 1]:
            counter += 1
    return counter

def part1():
    with open("input.txt") as f:
        data = [int(i) for i in f.read().split("\n")]
        return increase_counter(data)


def part2():
    with open("input.txt") as f:
        data = [int(i) for i in f.read().split("\n")]
        summed_data = [sum(data[i: i+3]) for i in range(len(data) - 2)]
        return increase_counter(summed_data)
    

if __name__ == "__main__":
    s1 = f.average_time(part1)
    s2 = f.average_time(part2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")