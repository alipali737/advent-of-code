if __name__ == "__main__":
    with open("input.txt") as f:
        counter = 0
        data = [int(i) for i in f.read().split("\n")]
        for i in range(len(data) - 1):
            if data[i] < data[i + 1]:
                counter += 1

    print(f"Result: {counter}")

