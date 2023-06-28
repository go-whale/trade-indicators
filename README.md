[![Coverage Status](https://coveralls.io/repos/github/go-whale/trade-indicators/badge.svg)](https://coveralls.io/github/go-whale/trade-indicators)

# Trading indicators

Note: this is beta version 
Top trading indicators for go projects

## EMA (Exponential Moving Average):

EMA = Price(t) * k + EMA(y) * (1 − k)

where:
t=today
y=yesterday
N=number of days in EMA
k = 2 / (N+1)

and for the first EMA(y) we need calculate SMA:
SMA = (Price(1) + Price(2) + ... + Price(n)) / N

EMAs are commonly used in conjunction with other indicators to confirm significant market moves and to gauge their validity. For traders who trade intraday and fast-moving markets, the EMA is more applicable. Quite often, traders use EMAs to determine a trading bias. If an EMA on a daily chart shows a strong upward trend, an intraday trader’s strategy may be to trade only on the long side.

## Volatility:

The term “volatility” refers to the statistical measure of the dispersion of returns during a certain period of time for stocks, security, or market index. The volatility can be calculated either using the standard deviation or the variance of the security or stock.

The formula for daily volatility is computed by finding out the square root of the variance of a daily stock price.

Daily Volatility Formula is represented as Volatility = sqrt(Variance)


## RSI calculation (Relative Strength index):

The relative strength index (RSI) is a technical indicator used in the analysis of financial markets. It is intended to chart the current and historical strength or weakness of a stock or market based on the closing prices of a recent trading period.
Let’s understand how to calculate and graph the RSI indicator now. While you can easily calculate the RSI indicator value with the python code, for explanation purposes we will do it manually.

|  Date | Close (1) | Change (2) | Gain (3) | Loss (4) | Avg Gain (5) | Avg Loss (6) | RS (7) | 14-day RSI (8) |
|:-----:|:---------:|:----------:|:--------:|:--------:|:------------:|:------------:|:------:|:--------------:|
| 24-04 |   283.46  |            |          |          |              |              |        |                |
| 25-04 |   280.69  |    -2.77   |   0.00   |   2.77   |              |              |        |                |
| 26-04 |   285.48  |    4.79    |   4.79   |   0.00   |              |              |        |                |
| 27-04 |   294.08  |    8.60    |   8.60   |   0.00   |              |              |        |                |
| 30-04 |   293.90  |    -0.18   |   0.00   |   0.18   |              |              |        |                |
| 01-05 |   299.92  |    6.02    |   6.02   |   0.00   |              |              |        |                |
| 02-05 |   301.15  |    1.23    |   1.23   |   0.00   |              |              |        |                |
| 03-05 |   284.45  |   -16.70   |   0.00   |   16.70  |              |              |        |                |
| 04-05 |   294.09  |    9.64    |   9.64   |   0.00   |              |              |        |                |
| 07-05 |   302.77  |    8.68    |   8.68   |   0.00   |              |              |        |                |
| 08-05 |   301.97  |    -0.80   |   0.00   |   0.80   |              |              |        |                |
| 09-05 |   306.85  |    4.88    |   4.88   |   0.00   |              |              |        |                |
| 10-05 |   305.02  |    -1.83   |   0.00   |   1.83   |              |              |        |                |
| 11-05 |   301.06  |    -3.96   |   0.00   |   3.96   |              |              |        |                |
| 14-05 |   291.97  |    -9.09   |   0.00   |   9.09   |     3.13     |     2.52     |  1.24  |      55.37     |
| 15-05 |   284.18  |    -7.79   |   0.00   |   7.79   |     2.91     |     2.90     |  1.00  |      50.07     |
| 16-05 |   286.48  |    2.30    |   2.30   |   0.00   |     2.86     |     2.69     |  1.06  |      51.55     |
| 17-05 |   284.54  |    -1.94   |   0.00   |   1.94   |     2.66     |     2.64     |  1.01  |      50.20     |
| 18-05 |   276.82  |    -7.72   |   0.00   |   7.72   |     2.47     |     3.00     |  0.82  |      45.14     |
| 21-05 |   284.49  |    7.67    |   7.67   |   0.00   |     2.84     |     2.79     |  1.02  |      50.48     |
| 22-05 |   275.01  |    -9.48   |   0.00   |   9.48   |     2.64     |     3.27     |  0.81  |      44.69     |
| 23-05 |   279.07  |    4.06    |   4.06   |   0.00   |     2.74     |     3.03     |  0.90  |      47.47     |
| 24-05 |   277.85  |    -1.22   |   0.00   |   1.22   |     2.54     |     2.90     |  0.88  |      46.71     |
| 25-05 |   278.85  |    1.00    |   1.00   |   0.00   |     2.43     |     2.70     |  0.90  |      47.45     |
| 29-05 |   283.76  |    4.91    |   4.91   |   0.00   |     2.61     |     2.50     |  1.04  |      51.05     |
| 30-05 |   291.72  |    7.96    |   7.96   |   0.00   |     2.99     |     2.32     |  1.29  |      56.29     |
| 31-05 |   284.73  |    -6.99   |   0.00   |   6.99   |     2.78     |     2.66     |  1.05  |      51.12     |
| 01-06 |   291.82  |    7.09    |   7.09   |   0.00   |     3.09     |     2.47     |  1.25  |      55.58     |
| 04-06 |   296.74  |    4.92    |   4.92   |   0.00   |     3.22     |     2.29     |  1.40  |      58.41     |
| 05-06 |   291.13  |    -5.61   |   0.00   |   5.61   |     2.99     |     2.53     |  1.18  |      54.17     |

Step 1: Closing Price
We will take the closing price of the stock for 30 days. The closing price is mentioned in column (1).

Step 2: Changes in Closing Price
We then compare the closing price of the current day with the previous day’s closing price and note them down. Thus, from the table, for 25-04, we get the change in price as (280.69 - 283.46) = -2.77.

Similarly, for 26-04,
Change in price = (Current closing price - Previous closing price) = (285.48 - 280.6) = 4.79.

We will then tabulate the results in the column mentioned as “Change (2)”. In this manner, we calculated the change in price.

Step 3: Gain and Loss
We will now create two sections depending on the fact the price increased or decreased, with respect to the previous day’s closing price.

If the price has increased, we note down the difference in the “Gain” column and if it’s a loss, then we note it down in the “Loss” column.

For example, on 26-04, the price had increased by 4.79. Thus, this value would be noted in the “Gain” column.

If you look at the data for 25-04, there was a decrease in the price by 2.77. Now, while the value is written as negative in the “change” column, we do not mention the negative sign in the “Loss” column. And only write it as 2.77. In this manner, the table for the columns “Gain (3)” and “Loss (4)” is updated.

Step 4: Average Gain and Loss
In the RSI indicator, to smoothen the price movement, we take an average of the gains (and losses) for a certain period.

While we call it an average, a little explanation would be needed. For the first 14 periods, it is a simple average of the values.

To explain it, we will look at the average gain column.

Thus, in the table, the first 14 values would be from (25-04) to (14-05) which is, (0.00 + 4.79 + 8.60 + 0.00 + 6.02 + 1.23 + 0.00 + 9.64 + 8.68 + 0.00 + 4.88 + 0.00 + 0.00 + 0.00)/14 = 3.13.

Now, since we are placing more emphasis on the recent values, for the next set of values, we use the following formula,

[(Previous avg. gain)*13)+ current gain)]/14
Thus, for (15-05), we will calculate the average gain as [(3.13*13)+0.00]/14 = 2.91.

Similarly, we will calculate the average Loss too.

Based on these formulae, the table is updated for the columns “Avg Gain (5)” and “Avg Loss (6)”.

Step 5: Calculate RS
Now, to make matters simple, we add a column called “RS” which is simply, (Avg Gain)/(Avg Loss). Thus, for 14-05,

RS = (Avg Gain)/(Avg Loss) = 3.13/2.52 = 1.24.

In this manner, the table for the column “RS (7)” is updated. In the next step, we finally work out the RSI values.

Step 6: Calculation of RSI
RSI = [100 - (100/{1+ RS})]
For example, for (14-05),
RSI = [100 - (100/{1+ RS})] = [100 - (100/{1+ 1.24})] = 55.37.