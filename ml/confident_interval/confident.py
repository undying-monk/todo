import numpy as np
import pandas as pd
import scipy
import matplotlib.pyplot as plt
import statsmodels.formula.api as smf

df = pd.read_csv('data/rideshare_2022_cleaned.csv', parse_dates=['trip_start_timestamp', 'date'])
print(df.head())

daily_rides = df.groupby('date').size().reset_index(name='daily_rides')
print(daily_rides)
mean_rides_per_day = daily_rides['daily_rides'].mean()
std_rides_per_day = daily_rides['daily_rides'].std()

print(f'Mean number of rides per day: {mean_rides_per_day:.2f}')
print(f'Standard deviation: {std_rides_per_day:.2f}')

# plt.figure(figsize=(18,6))
# # Plot the histogram of the daily rides
# plt.bar(daily_rides['date'], daily_rides['daily_rides'], label='Rides per Day')

# # # Plot the mean value as a horizontal line
# plt.axhline(y=mean_rides_per_day, c='r', label=f'Mean Rides per Day')

# # plt.ylabel('Rides per Day', fontsize=16)
# # plt.xlabel('Date', fontsize=16)
# # plt.xlim(min(daily_rides['date']), max(daily_rides['date']))
# plt.legend(fontsize=14)
# plt.show()

confidence = 0.95
alpha = 1-confidence
critical_value = scipy.stats.t.ppf(1-alpha/2, df=len(daily_rides)-1)
print(f"For the confidence interval of {confidence}, the critical value is {critical_value}")

total_days = daily_rides['date'].count()
confidence_interval = critical_value * std_rides_per_day / np.sqrt(total_days)
print(f"With a {100 * confidence}% confidence you can say that your error will be no more than {confidence_interval:.4f} rides per day.")

# plt.figure(figsize=(18,6))
# # Plot the histogram of the daily rides
# plt.bar(daily_rides['date'], daily_rides['daily_rides'], label='Rides per Day')

# Plot the mean value as a horizontal line
# plt.axhline(y=mean_rides_per_day, c='r', label=f'Mean Rides per Day +/- {confidence}% Confidence Interval')
# # Plot the confidence interval around the line
# plt.fill_between(daily_rides['date'], mean_rides_per_day-confidence_interval,
#                  mean_rides_per_day+confidence_interval, color='r', alpha=0.2)

# plt.ylabel('Rides per Day', fontsize=16)
# plt.xlabel('Date', fontsize=16)
# plt.xlim(min(daily_rides['date']), max(daily_rides['date']))
# plt.legend(fontsize=14)
# plt.show()

# Select the data only for holidays
daily_rides_holidays = daily_rides[daily_rides['date'] > '2022-12-17']
# Compute sample mean and standard deviation for holidays
mean_rides_per_day_holidays = daily_rides_holidays['daily_rides'].mean()
std_rides_per_day_holidays = daily_rides_holidays['daily_rides'].std()

print(f'Mean number of rides per day: {mean_rides_per_day_holidays:.2f} +/- {std_rides_per_day_holidays:.2f}')
# Get the confidence interval for the population mean for the holidays.
critical_value_holidays =  scipy.stats.t.ppf(1 - alpha/2, df=len(daily_rides_holidays)-1)
total_days_holidays = daily_rides_holidays['date'].count()
confidence_interval_holidays = critical_value_holidays * std_rides_per_day_holidays / np.sqrt(total_days_holidays)
print(f"With a {100 * confidence}% confidence you can say that your error will be no more than {confidence_interval_holidays} rides per day.")

# plt.figure(figsize=(18,6))

# Plot the histogram of the daily rides, the mean and the confidence interval
# plt.bar(daily_rides['date'], daily_rides['daily_rides'], label='Rides per Day')
# plt.axhline(y=mean_rides_per_day, color='C0', label=f'Mean Rides per Day +/- {confidence}% Confidence Interval')
# plt.fill_between(daily_rides['date'], mean_rides_per_day-confidence_interval,
#                  mean_rides_per_day+confidence_interval, color='C0', alpha=0.3)

# # Plot the histogram of the daily rides, the mean and the confidence interval for the holiday season
# plt.bar(daily_rides_holidays['date'], daily_rides_holidays['daily_rides'], label='Rides per Day (Holidays)')
# plt.axhline(y=mean_rides_per_day_holidays, color='C1', label='Mean Rides per Day (Holidays) +/- {confidence}% Confidence Interval')

# plt.fill_between(daily_rides_holidays['date'], mean_rides_per_day_holidays-confidence_interval_holidays,
#                  mean_rides_per_day_holidays+confidence_interval_holidays, color='C1', alpha=0.5)

# plt.ylabel('Rides per Day', fontsize=16)
# plt.xlabel('Date', fontsize=16)
# plt.xlim(min(daily_rides['date']), max(daily_rides_holidays['date']))
# plt.legend(fontsize=14)
# plt.show()

# Two Sample t-test
WEEKDAYS = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']

daily_rides['weekday'] = daily_rides['date'].apply(lambda x: WEEKDAYS[x.weekday()])

# group by on weekday and daily rides, and return numerical summaries
weekday_summary = daily_rides.groupby('weekday')['daily_rides'].describe()
# Reorder the series given weekdays
weekday_summary = weekday_summary.reindex(WEEKDAYS)
print(weekday_summary)

# Create two series, one for the numbers of rides on every friday and saturday and one for the other days
fridays_and_saturdays = daily_rides[daily_rides['weekday'].isin(['Friday', 'Saturday'])]['daily_rides']
other_days = daily_rides[daily_rides["weekday"].isin(["Monday", "Tuesday", "Wednesday", "Thursday", "Sunday"])]["daily_rides"]
# Note that these series contain all days, not the summary from the dataframe in the previous cell
print(f"Number of datapoints for Fridays and Saturdays: {len(fridays_and_saturdays)}")
print(f"Number of datapoints for other days: {len(other_days)}")

#Calculate the t
t = scipy.stats.ttest_ind(a=fridays_and_saturdays,b=other_days, alternative='greater')
print(t)

# since we got p-value is 2.6591073725083493e-67 < 0.05, we rejects the Null Hypothesis


# fig, ax = plt.subplots(1,2, figsize=(12,4))
# df.plot.scatter('fare','trip_seconds', ax=ax[0])
# df.plot.scatter('fare','trip_miles', ax=ax[1])

# plt.show()

# Create the model
model = smf.ols(formula='fare ~ trip_seconds + trip_miles', data=df)

# Fit the model
result = model.fit()

# Display the results
print(result.summary())
print(result.params)

starting_fare = result.params["Intercept"]
price_per_second = result.params["trip_seconds"]
price_per_mile = result.params["trip_miles"]

print(f"The starting fare is {starting_fare:.3} USD. In addition the ride costs {price_per_second*60:.3} USD per minute and {price_per_mile:.3} USD per mile.")

def fare_calculator(trip_time, trip_distance): # b + W1X1 + W2X2
    return starting_fare + price_per_second * trip_time + price_per_mile * trip_distance

sample_trip_duration = 10 * 60 # 10 minutes
sample_trip_distance = 10 # 10 miles

sample_fare = fare_calculator(sample_trip_duration, sample_trip_distance)

print(f"For a {sample_trip_distance} mile trip that takes {sample_trip_duration/60} minutes, you would pay around {sample_fare:.3} USD.")

# Get the x and y data that you used to fit the model and drop nan values
x_y = df[["trip_miles", "trip_seconds", "fare"]].dropna()

# Change this row if you want to choose another x variable (trip_miles or trip_seconds) to plot
x_variable = "trip_seconds"

# Get the plotting data
x_plot =  x_y[x_variable]
y_plot =  x_y["fare"]
y_result = result.predict()

# Plot the data
plt.scatter(x_plot, y_plot, label="Original Data")
plt.scatter(x_plot, y_result, label="Prediction")
plt.xlabel(" ".join(x_variable.split("_")).title(), fontsize=14)
plt.ylabel("Fare", fontsize=14)
plt.legend(fontsize=14)

plt.show()