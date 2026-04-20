import numpy as np

x = np.array([10,20,50,100,95])
print(x[x<60], x < 60)
print(x[(x>20) & (x<60)])
print(x[x % 2 != 0])

print(np.where(x > 20, x, 0)) # preserve the original value just replace value