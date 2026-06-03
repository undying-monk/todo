import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.datasets import make_blobs 

# Set a seed so that the results are consistent.
np.random.seed(3)

def sigmoid(z):
    return  1 / (1+np.exp(-z))

def layer_sizes(X, Y):
    n_x = X.shape[0]
    n_h = 2
    n_y = Y.shape[0]

    return (n_x, n_h, n_y)

# W1: hidden layers have 2 perceptron, so each perceptron has 2 weights
# W2: output layer has 1 perceptron, with 2 weights
def init_parameters(n_x, n_h, n_y):
    W1 = np.random.randn(n_h,n_x) * 0.01
    b1 = np.zeros((n_h,1))
    W2 = np.random.randn(n_y,n_h) * 0.01
    b2 = np.zeros((n_y,1))
    return {
        "W1": W1,
        "b1": b1,
        "W2": W2,
        "b2": b2,
    }

def forward_propagation(X, parameters):
    W1 = parameters["W1"]
    b1 = parameters["b1"]
    W2 = parameters["W2"]
    b2 = parameters["b2"]

    Z1 = np.matmul(W1, X) + b1
    A1 = sigmoid(Z1)
    Z2 = np.matmul(W2, A1) + b2
    A2 = sigmoid(Z2)

    cache = {
        "Z1": Z1,
        "Z2": Z2,
        "A1": A1,
        "A2": A2,
    }
    return A2, cache

def compute_cost(Y, A2):
    m = A2.shape[1]
    logloss = - np.multiply(np.log(A2),Y) - np.multiply(np.log(1 - A2),1 - Y)
    cost = 1/m * np.sum(logloss) 
    print(logloss.shape)
    print(np.isinf(logloss).any())

    return cost

def backward_propagation(X,Y, parameters, cache):
  # First, retrieve W from the dictionary "parameters".
    W1 = parameters["W1"]
    W2 = parameters["W2"]
    
    # Retrieve also A1 and A2 from dictionary "cache".
    A1 = cache["A1"]
    A2 = cache["A2"]

    dZ1 = A1 - Y
    dZ2 = A2 - Y

    dW2 = 1/m * np.dot(dZ2 , X.T)
    db2 = (1/m) * np.sum(dZ2, axis=1, keepdims = True)

    # (2,1) * (1,2000)
    dZ1 = np.dot(W2.T, dZ2) * A1 * (1 - A1)
    dW1 = np.dot(dZ1, X.T)
    db1 = np.sum(dZ1, axis=1, keepdims = True)

    grads = {"dW1": dW1,
             "db1": db1,
             "dW2": dW2,
             "db2": db2}
    return grads

def update_parameters(parameters, learning_rate, grads):
    W1 = parameters["W1"]
    W2 = parameters["W2"]
    b1 = parameters["b1"]
    b2 = parameters["b2"]

    dW1 = grads["dW1"]
    db1 = grads["db1"]
    dW2 = grads["dW2"]
    db2 = grads["db2"]

    W1 = W1 - learning_rate * dW1
    b1 = b1 - learning_rate * db1

    W2 = W2 - learning_rate * dW2
    b2 = b2 - learning_rate * db2
    return {
        "W1": W1,
        "b1": b1,
        "W2": W2,
        "b2": b2,
    }

def nn_model(X , Y, n_h, num_iterations = 100, learning_rate = 1.2):
    # define neuron network layers
    (n_x, _, n_y) = layer_sizes(X,Y)

    # init parameters
    parameters = init_parameters(n_x, n_h, n_y)

    # loop: forward propagation + backward propagation
    for i in range(num_iterations):
        A2, cache = forward_propagation(X, parameters)
        cost = compute_cost(Y, A2)
        # print("Cost at i %i %f" % (i, cost))

        grads = backward_propagation(X,Y, parameters, cache)
        parameters = update_parameters(parameters, learning_rate, grads)
    return parameters

def predict(X, parameters):
    A2, cache = forward_propagation(X, parameters)
    predictions = A2 > 0.5
    return predictions


# Dataset
m = 2000
samples, labels = make_blobs(n_samples=m, 
                             centers=([2.5, 3], [6.7, 7.9], [2.1, 7.9], [7.4, 2.8]), 
                             cluster_std=1.1,
                             random_state=0)
labels[(labels == 0) | (labels == 1)] = 1
labels[(labels == 2) | (labels == 3)] = 0
X = np.transpose(samples)
Y = labels.reshape((1, m))

print ('The shape of X is: ' + str(X.shape))
print ('The shape of Y is: ' + str(Y.shape))
print ('I have m = %d training examples!' % (m))

(n_x, n_h, n_y) = layer_sizes(X, Y)
print("The size of the input layer is: n_x = " + str(n_x))
print("The size of the hidden layer is: n_h = " + str(n_h))
print("The size of the output layer is: n_y = " + str(n_y))

parameters = init_parameters(n_x, n_h, n_y)
print("W1 = " + str(parameters["W1"]))
print("b1 = " + str(parameters["b1"]))
print("W2 = " + str(parameters["W2"]))
print("b2 = " + str(parameters["b2"]))

A2, cache = forward_propagation(X, parameters)
print("forward_propagation",A2)

learning_rate = 1.2
cost = compute_cost(A2, Y)
print("cost = ", cost)

grads = backward_propagation(X, Y, parameters, cache)
print("dW1 = " + str(grads["dW1"]))
print("db1 = " + str(grads["db1"]))
print("dW2 = " + str(grads["dW2"]))
print("db2 = " + str(grads["db2"]))

parameters_updated = update_parameters(parameters, learning_rate, grads)
print("W1 updated = " + str(parameters_updated["W1"]))
print("b1 updated = " + str(parameters_updated["b1"]))
print("W2 updated = " + str(parameters_updated["W2"]))
print("b2 updated = " + str(parameters_updated["b2"]))

parameters = nn_model(X, Y, n_h=2, num_iterations=3000, learning_rate=1.2)
print("W1 = " + str(parameters["W1"]))
print("b1 = " + str(parameters["b1"]))
print("W2 = " + str(parameters["W2"]))
print("b2 = " + str(parameters["b2"]))

W1 = parameters["W1"]
b1 = parameters["b1"]
W2 = parameters["W2"]
b2 = parameters["b2"]

X_pred = np.array([[2, 8, 2, 8], [2, 8, 8, 2]])
Y_pred = predict(X_pred, parameters)

print(f"Coordinates (in the columns):\n{X_pred}")
print(f"Predictions:\n{Y_pred}")