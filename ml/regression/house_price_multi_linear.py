import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from nn_model import nn_model,predict

# Set a seed so that the results are consistent.
np.random.seed(3)

df = pd.read_csv("./data/house_prices_train.csv")
X_multi = df[['GrLivArea', 'OverallQual']]
Y_multi = df['SalePrice']

X_multi_norm = (X_multi - np.mean(X_multi,axis=0))/np.std(X_multi)
Y_multi_norm = (Y_multi - np.mean(Y_multi,axis=0))/np.std(Y_multi)
X_multi_norm = np.array(X_multi_norm).T
Y_multi_norm = np.array(Y_multi_norm).reshape((1, len(Y_multi_norm)))

print ('The shape of X: ' + str(X_multi_norm.shape))
print ('The shape of Y: ' + str(Y_multi_norm.shape))
print ('I have m = %d training examples!' % (X_multi_norm.shape[1]))
parameters_multi = nn_model(X_multi_norm, Y_multi_norm, num_iterations=100)

print("W = " + str(parameters_multi["W"]))
print("b = " + str(parameters_multi["b"]))

W_multi = parameters_multi["W"]
b_multi = parameters_multi["b"]