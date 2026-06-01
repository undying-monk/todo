import matplotlib.pyplot as plt
import numpy as np


x = np.array([-4,0,8,12])
y = (12 - x) / 2 

plt.plot(x,y)
plt.scatter(x,y, color="red", alpha=0.5, s=50)
plt.title("x+2y=12")
plt.xlabel("x")
plt.ylabel("y")
plt.legend()

for i, txt in enumerate(x):
    plt.annotate(f"""({x[i]}, {y[i]})""", (x[i], y[i]), textcoords="offset points", xytext=(0,10), ha='center')

# plt.axline((0, 0), slope=0.5, color='red', label='Slope = 2')
plt.show()
# print(x,y)