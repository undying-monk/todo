import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from nn_model import nn_model,predict

# Set a seed so that the results are consistent.
np.random.seed(3)

adv = pd.read_csv("./data/tvmarketing.csv")
# adv.plot(x='TV', y='Sales', kind='scatter', c='black')
adv_norm = (adv - np.mean(adv,axis=0)) / np.std(adv)
# adv_norm.plot(x='TV', y='Sales', kind='scatter', c='black')
X_norm = adv_norm['TV']
Y_norm = adv_norm['Sales']

X_norm = np.array(X_norm).reshape((1,len(X_norm)))
Y_norm = np.array(Y_norm).reshape((1,len(Y_norm)))

parameters_simple = nn_model(X_norm, Y_norm, num_iterations=30, learning_rate=1.2)
print("W = " + str(parameters_simple["W"]))
print("b = " + str(parameters_simple["b"]))

X_pred = np.array([50, 120, 280])
Y_pred = predict(adv["TV"], adv["Sales"], parameters_simple, X_pred)
print(f"TV marketing expenses:\n{X_pred}")
print(f"Predictions of sales:\n{Y_pred}")

# visualize scatters
# plt.scatter(X_norm, Y_norm, color='black')
# plt.show()

# fig, ax = plt.subplots()
# plt.scatter(adv["TV"], adv["Sales"], color="black")

# plt.xlabel("$x$")
# plt.ylabel("$y$")
# X_line = np.arange(np.min(adv["TV"]),np.max(adv["TV"])*1.1, 0.1)
# Y_line = predict(adv["TV"], adv["Sales"], parameters_simple, X_line)
# ax.plot(X_line, Y_line, "r") # line
# ax.plot(X_pred, Y_pred, "bo") # 1 single marker
# plt.plot()
# plt.show()