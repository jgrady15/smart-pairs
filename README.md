# Stock analysis software that recommends trading strategies using ChatGPT Model 0613 
Calculates the covariance, correlation, and cointegration between stock pairs using the Engle-Granger two step method using ADF-test to test if the pair is stationary and an OLS regression model to determine the linear relationship between the pairs.

# Requirements:
Must have Docker software installed on your computer. If you are using macOS, then you must have the brew docker casket installed on your computer. 

## To run on macOS:
```
./run.sh
```
# Or Windows:
```
./run.ps1
```


<!-- Use Python for:
Loading and preprocessing the stock price data (handling dates, null values, etc)
Calculating covariance, and correlation between the stock returns
Handling any API calls to obtain data
Plotting charts and visualizations

Use C++ for:
Analyzing the spread time series (volatility, mean reversion)
Performing the Engle-Granger two-step method for cointegration, which uses
OLS Regression and the ADF test for regression residuals
Efficiently calculating the ratio between stock prices
Computing the rolling means for the spread time series
Performing linear regression to quantify mean reversion speed
Any core math-heavy computations that require speed

Use ChatGPT for:
Evaluating criteria and determining if stocks qualify as a tradable pair -->
