from matplotlib import pyplot as plt
from math import sin
from math import pi
from numpy import linspace
from SineFunction import SineFunction
from LinearFunction import LinearFunction
from StockFunction import StockFunction

x = linspace(0,3,1000)
y = []

stock = StockFunction()

stock.wave.Set_Period(1)
stock.volatility.Set_Period(.1)
stock.volatility.Set_Amplitude(.2)
stock.trend.Set_Slope(1)

for num in x:
    y.append(stock.Compute(num))





plt.plot(x,y)
plt.show()
