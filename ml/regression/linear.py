# ------- file: myplot.py ------
import matplotlib.pyplot as plt
import numpy as np
import math, copy
from utils import myplot

plt.style.use('./regression/mystyle.mplstyle')

# 1/2m * (y^ - y)^2
# m is number of examples
# w is weight
# b is bias
def f_cost(x,y,w,b):
    m = x.shape[0]
    total_cost = 0
    for i in range (m):
        f_wb = w*x[i] + b
        total_cost += (f_wb - y[i])**2
    return total_cost / 2 * m

# w = w - alpha * gradient
# b = b - alpha * gradient
# gradient = d/dw J(w,b) = 1 / m * Sum[(f_wb(x^i) - y^i) * x^i ]
def f_gradient(x,y,w,b): # return dj_dw, dj_db
    m = x.shape[0]
    dj_dw = 0
    dj_db = 0
    for i in range (m):
        f_wb = w*x[i] + b
        dj_dw += (f_wb - y[i])*x[i] 
        dj_db += (f_wb - y[i])
    dj_dw /= m
    dj_db /= m
    return dj_dw, dj_db


# w = w - alpha * gradient
# b = b - alpha * gradient
def f_gradient_descent(x,y,w_in,b_in, alpha, num_inters): # return w,b
    w = w_in
    b = b_in
    j_history = []
    param_history = []
    for i in range (num_inters):
        if i == 100000:
            return w,b,j_history,param_history
        dj_dw, dj_db = f_gradient(x,y,w,b)
        w = w - alpha * dj_dw
        b = b - alpha * dj_db
        
        j_history.append(f_cost(x,y,w,b))
        param_history.append([w,b])

    return w,b,j_history, param_history


def plot_gradient(x,y):
    myplot.plt_gradients(x,y,f_cost,f_gradient)
    plt.show()

def plot_cost_vs_iters(j_history):
    myplot.plt_cost_vs_iters(j_history)
    plt.show()

def plot_cost_contour(x_train, y_train, p_hist, w_range, b_range, contours, resolution):
    fig, ax = plt.subplots(1,1, figsize=(12, 4))
    myplot.plt_cost_contour(x_train, y_train, p_hist, ax, f_cost, w_range, b_range)
    plt.show()