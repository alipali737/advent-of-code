def Step1():
    with open("input.txt") as f:
        lines = [int(line.rstrip()) for line in f.readlines()]
        counter = 0
        for i in range(len(lines)-1):
                if lines[i] < lines[i+1]:
                    counter+=1
        print(counter)

def Step2():
    with open("input.txt") as f:
        lines = [int(line.rstrip()) for line in f.readlines()]
        counter = 0
        for i in range(len(lines)-3):
            if sum(lines[i:i+3]) < sum(lines[i+1:i+4]):
                counter+=1
        print(counter)

if __name__ == "__main__":
    Step1()
    Step2()