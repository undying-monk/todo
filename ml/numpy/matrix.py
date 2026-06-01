import numpy as np

x = np.array([-1,3,8])
y = np.array([-4,2,1])

# Expect the same dot behaviour on 1D array
dot = np.matmul(x,y)
dot2 = np.dot(x,y)
print(dot, dot2) 

# Expect the maxtrix multiplication on 2D array
x = np.array([[-1,3,8], [-2,1,6]])
y= np.array([[-4,2,1], [-1,3,8]])
dot = np.matmul(x.transpose(),y)
print(dot)

# while dot, matrix multiplication unsupport for broastcasting
x = np.array([[-1,3,8], [-2,1,6]])
y= np.array([[-4,2,1]])
try:
    dot = np.matmul(x.transpose(),y)
    print(dot)
except ValueError as err:
    print(err) # err tell mismatch dimension

x = np.array([[3,0], [0,0], [0,-2]]) # 3,2
y = np.array([[3], [5]]) # 2,1

dot = np.matmul(x,y)
print(dot)