import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
# A function to create a dataset.
from sklearn.datasets import make_blobs 
from matplotlib import colors
from nn_model import *

# Dataset
n_samples = 1000
samples, labels = make_blobs(n_samples=n_samples, 
                             centers=([2.5, 3], [6.7, 7.9]), 
                             cluster_std=1.4,
                             random_state=0)

X_larger = np.transpose(samples)
Y_larger = labels.reshape((1,n_samples))

# plt.scatter(X_larger[0, :], X_larger[1, :], c=Y_larger, cmap=colors.ListedColormap(['blue', 'red']))
# plt.show()

# train model
parameters_larger = nn_model(X_larger,Y_larger, num_iterations=100, learning_rate=1.2)
print("W = " + str(parameters_larger["W"]))
print("b = " + str(parameters_larger["b"]))

plot_decision_boundary(X_larger,Y_larger, parameters_larger)