from time import perf_counter

def average_time(func, iter = 10):
    total = 0
    for i in range(iter):
        start = perf_counter()
        func()
        end = perf_counter()
        total += end - start

    print(f"Time Taken: {total/iter}")