from pydoc import importfile
f = importfile("functions.py")

def Step1():
    # code
    pass

def Step2():
    # code
    pass
    
if __name__ == "__main__":
    s1 = f.average_time(Step1)
    s2 = f.average_time(Step2)
    
    print(f"Part 1 - Average Time: {s1[0]}, Result: {s1[1]}")
    print(f"Part 2 - Average Time: {s2[0]}, Result: {s2[1]}")
