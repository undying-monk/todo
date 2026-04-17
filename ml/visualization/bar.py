import matplotlib.pyplot as plt
import numpy as np

# compare categories by present each category as bar chart

x = np.array(["A","B","C","D"])
y1 = np.array([50,200,125,300])

plt.bar(x, y1)
plt.show()