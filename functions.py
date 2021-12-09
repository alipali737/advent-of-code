from time import perf_counter

def average_time(func, iter = 10):
    total = 0
    for _ in range(iter):
        start = perf_counter()
        funcReturnValue = func()
        end = perf_counter()
        total += end - start

    return total/iter, funcReturnValue