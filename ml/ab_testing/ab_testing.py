import math
import numpy as np
import pandas as pd
from scipy import stats
import w4_unittest

# Load the data from the test using pd.read_csv
data = pd.read_csv("data/background_color_experiment.csv")

# Print the first 10 rows
print(data.head(10))

print(f"The dataset size is: {len(data)}")

control_sd_data = data[data['user_type'] == 'control']['session_duration']
variation_sd_data = data[data['user_type'] == 'variation']['session_duration']

print(f"{len(control_sd_data)} users saw the original website with an average duration of {control_sd_data.mean():.2f} minutes\n")
print(f"{len(variation_sd_data)} users saw the new website with an average duration of {variation_sd_data.mean():.2f} minutes")

# X_c stores the session tome for the control group and X_v, for the variation group. 
X_c = control_sd_data.to_numpy()
X_v = variation_sd_data.to_numpy()
print(f"The first 10 entries for X_c are:\n{X_c[:20]}\n")
print(f"The first 10 entries for X_v are:\n{X_v[:20]}\n")

def get_stats(X):
    """
    Calculate basic statistics of a given data set.

    Parameters:
    X (numpy.array): Input data.

    Returns:
    tuple: A tuple containing:
        - n (int): Number of elements in the data set.
        - x (float): Mean of the data set.
        - s (float): Sample standard deviation of the data set.
    """
    n = len(X)
    x = X.mean()
    s = X.std(ddof=1)
    return (n,x,s)

w4_unittest.test_get_stats(get_stats)

n_c, x_c, s_c = get_stats(X_c)
n_v, x_v, s_v = get_stats(X_v)

print(f"For X_c:\n\tn_c = {n_c}, x_c = {x_c:.2f}, s_c = {s_c:.2f} ")
print(f"For X_v:\n\tn_v = {n_v}, x_v = {x_v:.2f}, s_v = {s_v:.2f} ")


def degrees_of_freedom(n_v, s_v, n_c, s_c):
    """Computes the degrees of freedom for two samples.

    Args:
        control_metrics (estimation_metrics_cont): The metrics for the control sample.
        variation_metrics (estimation_metrics_cont): The metrics for the variation sample.

    Returns:
        numpy.float: The degrees of freedom.
    """
    s_c_n_c = s_c**2 / n_c
    s_v_n_v = s_v**2 / n_v

    numerator = (s_c_n_c + s_v_n_v)**2
    denominator = (s_c_n_c**2 / (n_c-1)) +  (s_v_n_v**2 / (n_v-1))
    return numerator/denominator


w4_unittest.test_degrees_of_freedom(degrees_of_freedom)

d = degrees_of_freedom(n_v, s_v, n_c, s_c)
print(f"The degrees of freedom for the t-student in this scenario is: {d:.2f}")

def t_value(n_v, x_v, s_v, n_c, x_c, s_c):
    # As you did before, let's split the numerator and denominator to make the code cleaner.
    # Also, let's compute again separately s_c^2/n_c and s_v^2/n_v.

    s_v_n_v = s_v**2/n_v
    s_c_n_c = s_c**2/n_c

    # Compute the numerator for the t-value as given in the formula above
    numerator = x_v - x_c

    # Compute the denominator for the t-value as given in the formula above. You may use np.sqrt to compute the square root.
    denominator = np.sqrt(s_c_n_c + s_v_n_v)
    
    ### END CODE HERE ###

    t = numerator/denominator

    return t

w4_unittest.test_t_value(t_value)

t = t_value(n_v, x_v, s_v, n_c, x_c, s_c)
print(f"The t-value for this experiment is: {t:.2f}")

t_10 = stats.t(df=10)
cdf = t_10.cdf(1.21)
print(f"The CDF for the t-student distribution with 10 degrees of freedom and t-value = 1.21, or equivalently P(t_10 < 1.21) is equal to: {cdf:.2f}")


# 𝑝=𝑃(𝑡𝑑>𝑡)= 1−CDF 𝑡𝑑(𝑡)
def p_value(d, t_value):
    # Load the t-student distribution with $d$ degrees of freedom. Remember that the parameter in the stats.t is given by df.
    t_d = stats.t(df=d)

    # Compute the p-value, P(t_d > t). Remember to use the t_d.cdf with the proper adjustments as discussed above.
    p = 1 - t_d.cdf(t_value)
    ### END CODE HERE ###

    return p

w4_unittest.test_p_value(p_value)
print(f"The p-value for t_15 with t-value = 1.10 is: {p_value(15, 1.10):.4f}")
print(f"The p-value for t_30 with t-value = 1.10 is: {p_value(30, 1.10):.4f}")

def make_decision(X_v, X_c, alpha = 0.05):
    n_v, x_v, s_v = get_stats(X_v)
    n_c, x_c, s_c = get_stats(X_c)

    # Also, remember that x_c and x_v are not used in this computation
    d = degrees_of_freedom(n_v,s_v,n_c,s_c)
    
    # Compute the t-value
    t = t_value(n_v, x_v, s_v,n_c, x_c, s_c)

    # Compute the p-value for the t-student distribution with d degrees of freedom
    p = p_value(d, t)


    if p < alpha:
        return 'Reject H_0'
    else:
        return 'Do not reject H_0'

w4_unittest.test_make_decision(make_decision)
alphas = [0.06, 0.05, 0.04, 0.01]
for alpha in alphas:
    print(f"For an alpha of {alpha} the decision is to: {make_decision(X_v, X_c, alpha = alpha)}")