import threading

def compute():
    print("Thread running")  # Runs sequentially under GIL

threads = []
for _ in range(8):
    t = threading.Thread(target=compute)
    t.start()
    threads.append(t)

for t in threads:
    t.join()