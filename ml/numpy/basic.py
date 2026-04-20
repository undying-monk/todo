import numpy as np

x = np.array([1,2,3,4])
print(x)
print(x.shape) # depth, number of rows, number of columns
print(type(x)) # class ndarray - multi-dimension array
print(x.ndim) # print number of dimensions
print(np.array('A').ndim) # zero dimension
y = np.array([['A','B','C'],['D','E','F']])
print(y.ndim) # 2 dimension
print(y.shape) # 2 row, 3 columns
print(y[0][0]) # A: chain indexing
print(y[0, 0]) # A: multi-dimensional indexing faster than chain indexing
print(y[0,0] + y[1,0]) # AD

# slice
print(y[0:1:1]) # A,B,C
print(y[-1::]) # D, E, F
print(y[-2::]) # [['A','B','C'],['D','E','F']]

# row, column
print(y[:,0]) # A,D column 0
print(y[:,1]) # B,E column 1
print(y[:,0:2]) # 2 columns A,B and D,E
print(y[:,1:]) # from column 1 B,C and E,F