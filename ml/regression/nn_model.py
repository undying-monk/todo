import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

# Set a seed so that the results are consistent.
np.random.seed(3)

def init_layers(X,Y):
    return (X.shape[0], Y.shape[0])

def init_parameters(n_x,n_y):
    W = np.random.randn(n_y, n_x) * 0.01
    b = np.zeros((n_y, 1))

    parameters = {
        "W": W,
        "b": b
    }
    return parameters

def forward_propagation(X, parameters):
    W = parameters["W"]
    b = parameters["b"]

    Y_hat = np.matmul(W,X) + b
    return Y_hat

def backward_propagation(Y_hat,X, Y):
    m = X.shape[1]

    dZ = Y_hat - Y
    dW = 1/m * np.dot(dZ, X.T)
    dB = 1/m * np.sum(dZ, axis=1, keepdims = True)

    grads = {
        "dW": dW,
        "dB": dB,
    }
    return grads

def update_parameters(parameters, grads, learning_rate):
    W = parameters["W"]
    b = parameters["b"]
    W = W - learning_rate * grads["dW"]
    b = b - learning_rate * grads["dB"]
    parameters = {
        "W": W,
        "b": b
    }
    return parameters

def compute_cost(Y_hat, Y):
    m = Y_hat.shape[1]
    cost = np.sum((Y_hat - Y)**2) / (2*m)
    return cost

def nn_model(X, Y, num_iterations=100,learning_rate=0.01):
    # define layer of neural networks
    
    n_x, n_y = init_layers(X,Y)
    print("Define layer", n_x, n_y)
    # init parameters
    parameters = init_parameters(n_x,n_y)
    print("Init parameters", parameters)

    # loop: forward propagation + backward propagation, update parameters
    for i in range(0, num_iterations):
        Y_hat = forward_propagation(X_norm, parameters)
        cost = compute_cost(Y_hat, Y)

        grads = backward_propagation(Y_hat, X_norm, Y_norm)
        parameters = update_parameters(parameters, grads, learning_rate)

        print("Cost after iteration %i: %f" %(i, cost))
    return parameters

def predict(X,Y, parameters,X_pred):
    W = parameters["W"]
    b = parameters["b"]
    if isinstance(X, pd.Series):
        X_pred_norm = ((X_pred - np.mean(X)) / np.std(X)).reshape(1, len(X_pred))
    else:
        X_mean = np.array(np.mean(X)).reshape((len(X),1))
        X_std = np.array(np.std(X)).reshape((len(X),1))
        X_pred_norm = ((X_pred - X_mean)/X_std)
    Y_pred_norm = np.matmul(W,X_pred_norm) + b
    # Use the same mean and standard deviation of the original training array Y.
    Y_pred = Y_pred_norm * np.std(Y) + np.mean(Y)
    return Y_pred[0]


adv = pd.read_csv("./data/tvmarketing.csv")
# adv.plot(x='TV', y='Sales', kind='scatter', c='black')
adv_norm = (adv - np.mean(adv,axis=0)) / np.std(adv)
adv_norm.plot(x='TV', y='Sales', kind='scatter', c='black')
X_norm = adv_norm['TV']
Y_norm = adv_norm['Sales']

X_norm = np.array(X_norm).reshape((1,len(X_norm)))
Y_norm = np.array(Y_norm).reshape((1,len(Y_norm)))

parameters_simple = nn_model(X_norm, Y_norm, num_iterations=30, learning_rate=1.2)
print("W = " + str(parameters_simple["W"]))
print("b = " + str(parameters_simple["b"]))

# visualize scatters
# plt.scatter(X_norm, Y_norm, color='black')
# plt.show()

X_pred = np.array([50, 120, 280])
Y_pred = predict(adv["TV"], adv["Sales"], parameters_simple, X_pred)
print(f"TV marketing expenses:\n{X_pred}")
print(f"Predictions of sales:\n{Y_pred}")