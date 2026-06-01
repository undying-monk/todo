import numpy as np

def T(v):
    w = np.zeros((3,1))
    w[0,0] = 3*v[0,0]
    w[2,0] = -2*v[1,0]
    
    return w

e1 = np.array([[1], [0]])
e2 = np.array([[0], [1]])


V = np.hstack((e1,e2))
print(V)
print(T(V))