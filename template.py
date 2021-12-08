from pydoc import importfile
f = importfile("../functions.py")

def Step1():
    # code

def Step2():
    # code
    
if __name__ == "__main__":
    f.average_time(Step1)
    f.average_time(Step2)
