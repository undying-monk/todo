import matplotlib.pyplot as plt
import numpy as np

# Present distribution of quantitative data, group by bins and counts how many fall in each bin (range)

scores = np.random.normal(loc=80,scale=10, size=100)
scores = np.clip(scores, 0, 100)

plt.hist(scores, bins=10, color="lightgreen",edgecolor="black")
plt.xlabel("Scores")
plt.ylabel("Number of students")

plt.show()
