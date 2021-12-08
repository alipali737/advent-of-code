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
        print(f"Result: {increase_counter(data)}")


def part2():
    with open("input.txt") as f:
        data = [int(i) for i in f.read().split("\n")]
        summed_data = [sum(data[i: i+3]) for i in range(len(data) - 2)]
        print(f"Result: {increase_counter(summed_data)}")
    

if __name__ == "__main__":
    time1 = f.average_time(part1)
    time2 = f.average_time(part2)
    
    print(f"Part 1 Average Time: {time1}")
    print(f"Part 2 Average Time: {time2}")