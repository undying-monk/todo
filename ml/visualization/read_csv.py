import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

df = pd.read_csv("./data.csv")
category = df["Type1"].value_counts(ascending=True)

plt.barh(category.index, category.values)
plt.show()