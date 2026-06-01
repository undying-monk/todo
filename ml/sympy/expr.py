from sympy import *
import numpy as np
from sympy.utilities.lambdify import lambdify


# define expression
x, y = symbols('x y')
expr = x**2 + 2*y # x^2 + 2y
pprint(expr)
result = expr.evalf(subs={x:-1,y:2})
print("Expression",result)

# apply numpy array into expressions
f_symb_numpy = lambdify((x,y), expr, 'numpy')
x_array = np.array([1,2,3])
y_array = np.array([1,2,3])
result = f_symb_numpy(x_array,y_array)
print("Apply numpy",result)

# derivative of expression
dfdx = diff(expr, x)  
print("Detivative of x^2 + 2y") # 2x
pprint(dfdx)
result = dfdx.evalf(subs={x:-2, y:1})
print(result)

dfdx_symb_numpy = lambdify((x,y), dfdx, 'numpy')
result = dfdx_symb_numpy(x_array,y_array)
print(result)


# calculate the difference (derivatives)
def f(x):
    return x ** 2

x_array_2 = np.linspace(-5, 5, 100)
dfdx_numerical = np.gradient(f(x_array_2), x_array_2)
print(dfdx_numerical)