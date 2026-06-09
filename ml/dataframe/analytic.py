import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv("data/rideshare_2022.csv")
# print(df.head())
# print(df.info())

columns_of_interest = ['Trip Start Timestamp', 'Trip Seconds',
       'Trip Miles', 'Fare', 'Tip', 'Additional Charges', 'Trip Total', 'Shared Trip Authorized',
       'Trips Pooled', 'Pickup Centroid Latitude', 'Pickup Centroid Longitude', 'Dropoff Centroid Latitude',
       'Dropoff Centroid Longitude']

df = df[columns_of_interest]

# Rename all the columns to not include whitespace
df = df.rename(columns={i: "_".join(i.split(" ")).lower() for i in df.columns})

df['date'] = pd.to_datetime(df['trip_start_timestamp'])
print(df.head())

# Select the column which you want to plot.
column_to_plot = 'date'

# Plot the histogram of the desired column
# df.hist(column_to_plot, density=True)

# Select the column which you want to plot.
column_to_plot = 'tip'

# Plot the histogram of the desired column
# df.hist(column_to_plot, density=True, bins = 100)

tippers = df['tip'] > 0

# Count the number of tippers
number_of_tippers = tippers.sum()
print(number_of_tippers)

# Count the total number of rides
total_rides = len(df)

# Calculate the fraction of people who tip
fraction_of_tippers = number_of_tippers / total_rides
print(f'The percentage of riders who tip is {fraction_of_tippers*100:.0f}%.')

# Create a dataframe That only consists of tippers (conditioned on the boolean series)
df_tippers = df[tippers]
print(df_tippers['tip'])

# Now re-plot the above histogram, but only for tippers
# df_tippers.hist('tip', density=True, bins = 100)

# Extracting the day of the week is simple when you have it in datetime format.
df['weekday'] = df["date"].dt.day_name()

# Count the number of rides each day
daily_ride_counts = df['weekday'].value_counts()
print(daily_ride_counts)

WEEKDAYS = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']

# Reorder the series given weekdays
daily_ride_counts = daily_ride_counts.reindex(WEEKDAYS)

df_tippers = df[df['tip'] > 0]
# Count the number of tips given each day
daily_tippers_counts = df_tippers['weekday'].value_counts()

# Reorder the series given weekdays
daily_tippers_counts = daily_tippers_counts.reindex(WEEKDAYS)
print(daily_tippers_counts)


df_daily_aggregation = pd.concat([daily_ride_counts, daily_tippers_counts], axis=1, keys=['ride_count', 'tippers_count'])
df_daily_aggregation["tips_percentage"] = df_daily_aggregation['tippers_count'] / df_daily_aggregation['ride_count'] * 100
print(df_daily_aggregation)

plt.show()

