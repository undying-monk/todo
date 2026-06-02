import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from matplotlib import colors
from nn_model import *


# Dataset
m = 30

X = np.random.randint(0, 2, (2, m))
Y = np.logical_and(X[0] == 0, X[1] == 1).astype(int).reshape((1, m))

print('Training dataset X containing (x1, x2) coordinates in the columns:')
print(X)
print('Training dataset Y containing labels of two classes (0: blue, 1: red)')
print(Y)

print ('The shape of X is: ' + str(X.shape))
print ('The shape of Y is: ' + str(Y.shape))
print ('I have m = %d training examples!' % (X.shape[1]))

print("sigmoid(-2) = " + str(sigmoid(-2)))
print("sigmoid(0) = " + str(sigmoid(0)))
print("sigmoid(3.5) = " + str(sigmoid(3.5)))
print(sigmoid(np.array([-2, 0, 3.5])))
(n_x, n_y) = layer_sizes(X, Y)
parameters = init_parameters(n_x, n_y)
print("W = " + str(parameters["W"]))
print("b = " + str(parameters["b"]))
A = forward_propagation(X, parameters)
print("Output vector A:", A)
print("cost = " + str(compute_cost(A, Y)))
grads = backward_propagation(A, X, Y)

print("dW = " + str(grads["dW"]))
print("db = " + str(grads["db"]))

# train small dataset
parameters = nn_model(X, Y, num_iterations=50, learning_rate=1.2)
print("W = " + str(parameters["W"]))
print("b = " + str(parameters["b"]))

plot_decision_boundary(X, Y, parameters)

X_pred = np.array([[1, 1, 0, 0],
                   [0, 1, 0, 1]])
Y_pred = predict(X_pred, parameters)
print(f"Coordinates (in the columns):\n{X_pred}")
print(f"Predictions:\n{Y_pred}")