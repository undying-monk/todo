import pandas as pd
import statsmodels.formula.api as smf

# 1. Prepare sample data
data = {
    'Salary': [50000, 60000, 65000, 70000, 85000],
    'Experience': [1, 3, 4, 5, 8]
}
df = pd.DataFrame(data)

# 2. Define the model using R-style formula (y ~ x)
# This automatically adds a coefficient/intercept to your model WX+b
model = smf.ols(formula='Salary ~ Experience', data=df)

# 3. Fit the model
results = model.fit()

# 4. View the statistical summary
print(results.summary())
