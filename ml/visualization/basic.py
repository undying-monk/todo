# ------- file: myplot.py ------
import matplotlib.pyplot as plt
import numpy as np

x = np.array([10,20,30,40])
y1 = np.array([2023,2024,2025,2024])
y2 = np.array([2025,2026,2020,2021])

line_style = dict(marker="o", markersize=10, markerfacecolor="red", linestyle="solid", linewidth="4")
plt.title("Example", fontsize=20, family="Arial", fontweight="bold", color="red")

plt.xlabel("Numbers", fontsize=16)
plt.ylabel("Years", fontsize=16)
plt.tick_params(axis="both", colors="red")
plt.grid(axis="both", linestyle="--")
plt.plot(x, y1, color="blue", **line_style)
plt.plot(x, y2, color="cyan", **line_style)
plt.show()