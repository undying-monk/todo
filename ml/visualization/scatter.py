import matplotlib.pyplot as plt
import numpy as np

# Use scatter to show relationship between 2 variables, help to indentify correlation sample (+,-,none)

# Number of hours
x1 = np.array([2,4,6,7, 7.5, 8, 8.5, 9, 10])
x2 = np.array([1,3,5,6, 7.5, 7.5, 8, 9, 9.5])
# Test score
y1 = np.array([1,4, 7, 8, 8.5, 9, 9.5, 9.8, 10])
y2 = np.array([2,3.5,6.5, 7.5, 8, 9, 9.2, 9.6, 9.8])

plt.xlabel("Study hours")
plt.ylabel("Score")
plt.scatter(x1,y1, color="red", alpha=0.5, s=50, label="CLASS A")
plt.scatter(x2,y2, color="blue", alpha=0.5, s=50, label="CLASS B")
plt.legend()
plt.show()