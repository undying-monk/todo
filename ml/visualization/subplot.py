import matplotlib.pyplot as plt
import numpy as np

# Axe is a single plot
# Figure is the entire canvas

figure, axes = plt.subplots(2,2)
x = np.array([-2,-1,0,1,2])
axes[0,0].plot(x, x*2)
axes[0,0].set_title("X*2")

axes[0,1].plot(x, x**2)
axes[0,1].set_title("X^2")

axes[1,0].plot(x, x**3)
axes[1,0].set_title("Cubic")

axes[1,1].plot(x, np.sqrt(x))
axes[1,1].set_title("Square root")

plt.show()