from regression import linear
import matplotlib.pyplot as plt
import numpy as np

x_train = np.array([1.0, 2.0])   #features
y_train = np.array([300.0, 500.0])   #target value
# linear.plot_gradient(x_train, y_train)

alpha = 0.01
num_inters = 10000
w,b,j_history, param_history = linear.f_gradient_descent(x_train, y_train, 0, 0, alpha, num_inters)

# linear.plot_cost_vs_iters(j_history)
# linear.plot_cost_contour(x_train, y_train, param_history)
linear.plot_cost_contour(x_train, y_train, param_history, w_range=[180, 220, 0.5], b_range=[80, 120, 0.5],
            contours=[1,5,10,20],resolution=0.5)
