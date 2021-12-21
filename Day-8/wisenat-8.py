from pydoc import importfile
f = importfile("../functions.py")

def sameChars(string, chars):
    return [char in string for char in chars]

def part1():
    with open("input.txt") as f:
        data = [[partition.split(" ") for partition in line.split(" | ")] for line in f.read().split("\n")]
        unique = 0
        for segment in data:
            for value in segment[1]:
                if len(value) in (2, 3, 4, 7):
                    unique += 1
    return unique

def part2():
    '''
    2 Segment: 1
    3 Segment: 7
    4 Segment: 4
    5 Segment: 2, 3, 5
    6 Segment: 0, 6, 9
    7 Segment: 8

    Unique: 1, 7, 4, 8
    Determine 9 using 4.. shares all segments
    Determine 2 using 4.. shares all but 2 segments
    Determine 5 using 4.. shares all but 1 segment
    Determine 3 using 1.. shares all segments
    Determine 0 using 1.. shares all segments
    Determine 6 using None.. last value
    '''
    with open("input.txt") as f:
        data = [[partition.split(" ") for partition in line.split(" | ")] for line in f.read().split("\n")]

    outputs = []
    for segment in data:
        digits = ["" for _ in range(10)]
        inp, out = segment[0], segment[1]
        inp.sort(key=len)

        # Determine 1, 7, 4, 8 using their unique lengths
        for key, ind in zip((1, 7, 4, 8), (0, 1, 2, -1)):
            digits[key] = inp[ind]

        del inp[:3]
        del inp[-1]

        # Determine 9(6) using 4.. shares all segments
        # Determine 2(5) using 4.. shares all but 2 segments
        # Determine 5(5) using 4.. shares all but 1 segment
            # AND doesn't share all of 1s segments
                # 3(5) is the value if it does
        for val in inp[:]:
            bool4 = sameChars(val, digits[4])
            bool4.sort()
            if not digits[9] and all(bool4):
                digits[9] = val
            elif not digits[2] and bool4 == [False, False, True, True]:
                digits[2] = val
            elif bool4 == [False, True, True, True]:
                bool1 = sameChars(val, digits[1])
                if not digits[3] and all(bool1):
                    digits[3] = val
                elif not digits[5]:
                    digits[5] = val
                else:
                    continue
            else:
                continue
            inp.remove(val)

        # Determine 0(6) using 1.. shares all segments
        # Determine 6(6) using None.. last value       
        bool1 = sameChars(inp[0], digits[1])
        bool1.sort()
 
        digits[0] = inp[1]
        digits[6] = inp[0]
        if bool1 == [True, True]:
            digits[0], digits[6] = digits[6], digits[0]

        # Find output values
        res = ""
        digits = list(map(lambda x: "".join(sorted(x)), digits))
        for val in out:
            res += str(digits.index("".join(sorted(val))))
        outputs.append(int(res))

    return sum(outputs)


if __name__ == "__main__":
    # s1 = f.average_time(part1, iter=1)
    s2 = f.average_time(part2, iter=1)
    
    # print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")