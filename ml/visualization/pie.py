import matplotlib.pyplot as plt
import numpy as np

# Circular chart divides into slices that show percentage of total, good for visualization among categories

y1 = np.array([35,15,5,45])
myExplode = [0,0,0,0.1]
colors = ["red","green","blue","cyan"]
plt.pie(y1, labels=["Sugars", "Sault", "Ginger", "carrot"], autopct="%1.1f%%",colors=colors, explode=myExplode, startangle=90)
plt.show()