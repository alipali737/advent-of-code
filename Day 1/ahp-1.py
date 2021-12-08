with open("input.txt") as f:
    lines = f.readlines()
    lines = [line.rstrip() for line in lines]
    counter = 0
    for i in range(len(lines)-1):
            if lines[i] < lines[i+1]:
                counter+=1
    print(counter)