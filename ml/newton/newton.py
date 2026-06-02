import numpy as np
import matplotlib.pyplot as plt
from matplotlib.gridspec import GridSpec

# e^x - log(x)
def f_example_1(x):
    return np.exp(x) - np.log(x)

def dfdx(x):
    return np.exp(x) - 1/x

def d2fdx2(x):
    return np.exp(x) + 1/(x**2)


x_0 = 1.6
print(f"f({x_0}) = {f_example_1(x_0)}")
print(f"f'({x_0}) = {dfdx(x_0)}")
print(f"f''({x_0}) = {d2fdx2(x_0)}")


def plot_f(x_range, y_range, f, ox_position):
    x = np.linspace(*x_range, 100)
    fig, ax = plt.subplots(1,1,figsize=(8,4))

    ax.set_ylim(*y_range)
    ax.set_xlim(*x_range)
    ax.set_ylabel('$f\,(x)$')
    ax.set_xlabel('$x$')
    ax.spines['left'].set_position('zero')
    ax.spines['bottom'].set_position(('data', ox_position))
    ax.spines['right'].set_color('none')
    ax.spines['top'].set_color('none')
    ax.xaxis.set_ticks_position('bottom')
    ax.yaxis.set_ticks_position('left')
    ax.autoscale(enable=False)
    
    pf = ax.plot(x, f(x), 'k')
    plt.plot()
    plt.show()
    return fig, ax

plot_f([0.001, 2.5], [-0.3, 13], f_example_1, 0.0)

def newtons_method(x, dfdx, d2fdx2, num_iterations=100):
    for i in range(num_iterations):
        x = x - dfdx(x) / d2fdx2(x)
        print("Newtons at %i %f" %(i,x))
    return x

def gradient_descent(x, dfdx, learning_rate=1.2, num_iterations = 100):
    for i in range(num_iterations):
        x = x - learning_rate *dfdx(x) 
        print("Gradient at %i %f" %(i,x))
    return x


x_initial = 1.6
newtons_example_1 = newtons_method(x_initial, dfdx, d2fdx2, num_iterations=25)
print("Newton's method result: x_min =", newtons_example_1)

gd_example_1 = gradient_descent(x_initial, dfdx, learning_rate=0.1, num_iterations=25)
print("Gradient descent result: x_min =", gd_example_1) 