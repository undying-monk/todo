import numpy as np

x = np.array([1,2,3])
print(x + 2)
print(x - 2)
print(x / 2)
print(x * 2)
print(x ** 2)

# vectorized with math
print("vector")
y = np.array([4,16,9]) # vector is 1 dimension array
print(np.sqrt(y))
print(np.round(np.array([1.2,2.02]))) # round  1, 2
print(np.floor(np.array([1.2,2.02]))) # round down 1, 2
print(np.ceil(np.array([1.2,2.02]))) # round up 2, 3
print(np.pi)