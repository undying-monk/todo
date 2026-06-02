import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from matplotlib import colors

# Set a seed so that the results are consistent.
np.random.seed(3)


def sigmoid(z):
    return 1/(1+np.exp(-z))

def layer_sizes(X,Y):
    return (X.shape[0], Y.shape[0])

def init_parameters(n_x,n_y):
    W = np.random.randn(n_y,n_x) * 0.01
    b = np.zeros((n_y, 1))
    return {
        "W": W,
        "b": b
    }

def forward_propagation(X, parameters):
    W = parameters["W"] 
    b = parameters["b"]
    Z = np.matmul(W, X) + b
    Y_hat = sigmoid(Z)
    return Y_hat

def compute_cost(Y_hat, Y):
    m = Y_hat.shape[1]
    logloss = -Y*np.log(Y_hat) - (1-Y)*np.log(1-Y_hat)
    return np.sum(logloss) / m

def backward_propagation(Y_hat,X,Y):
    m = X.shape[1]

    dZ = Y_hat - Y
    dW = (1/m) * np.dot(dZ, X.T)
    db = (1/m) * np.sum(dZ, axis=1, keepdims = True)

    return {
        "dW": dW,
        "db": db
    }

def update_parameters(parameters, learning_rate, grads):
    dW = grads["dW"]
    db = grads["db"]
    W = parameters["W"] 
    b = parameters["b"]

    W = W - learning_rate * dW
    b = b - learning_rate * db
    return {
        "W": W,
        "b": b
    }

def nn_model(X, Y, num_iterations = 100, learning_rate=1.2):
    # define layer of neural networks
    n_x, n_y = layer_sizes(X,Y)
    print("Define layer", n_x, n_y)
    # init parameters
    parameters = init_parameters(n_x,n_y)
    print("Init parameters", parameters)

    for i in range(num_iterations):
        Y_hat = forward_propagation(X, parameters)
        cost = compute_cost(Y_hat, Y)
        grads = backward_propagation(Y_hat, X, Y)
        parameters = update_parameters(parameters, learning_rate, grads)
        print("Cost at iteration %i %f" % (i,cost))
    return parameters

def predict(X, parameters):
    result = forward_propagation(X, parameters)
    return result > 0.5

def plot_decision_boundary(X, Y, parameters):
    W = parameters["W"]
    b = parameters["b"]

    fig, ax = plt.subplots()
    plt.scatter(X[0, :], X[1, :], c=Y, cmap=colors.ListedColormap(['blue', 'red']));
    
    x_line = np.arange(np.min(X[0,:]),np.max(X[0,:])*1.1, 0.1)
    ax.plot(x_line, - W[0,0] / W[0,1] * x_line + -b[0,0] / W[0,1] , color="black")
    plt.plot()
    plt.show()


# Visualization
# fig, ax = plt.subplots()
# xmin, xmax = -0.2, 1.4
# x_line = np.arange(xmin, xmax, 0.1)
# # Data points (observations) from two classes.
# ax.scatter(0, 0, color="b")
# ax.scatter(0, 1, color="r")
# ax.scatter(1, 0, color="b")
# ax.scatter(1, 1, color="b")
# ax.set_xlim([xmin, xmax])
# ax.set_ylim([-0.1, 1.1])
# ax.set_xlabel('$x_1$')
# ax.set_ylabel('$x_2$')
# # One of the lines which can be used as a decision boundary to separate two classes.
# ax.plot(x_line, x_line + 0.5, color="black")
# plt.plot()
# plt.show()