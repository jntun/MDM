from matplotlib import pyplot as plt
from math import sin
from math import pi
from math import ceil
from math import floor
from numpy import linspace
from SineFunction import SineFunction
from LinearFunction import LinearFunction
from StockFunction import StockFunction
from random import randint
from random import seed as randseed

randseed()

'''
def append_interval(domain, codomain, StockFunctionInstance, start_int, stop_int):
    i = start_int
    while i < stop_int:
        result = StockFunctionInstance.Compute(domain[i])
        codomain.append(result)
        i = i + 1
    return 

def algorithm(Max, Piece, domain, codomain, Stocks):
    Upper_Bound = 10
    Lower_Bound = 0
    i = 0
    max_slope = 0
    min_slope = 0
    while i < Max:
        Current_Value = Stocks.Compute(domain[i])
        Flip_Probability = (Current_Value - Lower_Bound)/ (Upper_Bound - Lower_Bound) * 100
        probability_value = randint(0, 100)
        if probability_value <= Flip_Probability:
            max_slope = 0
            min_slope = floor((Lower_Bound - Current_Value) / (domain[Piece] - domain[0]))
        else:
            max_slope = ceil((Upper_Bound - Current_Value) / (domain[Piece] - domain[0]))
            min_slope = 0
        stock.Trend.Branch_Function(domain[i], max_slope, min_slope)
        stock.Wave.Branch_Period(domain[i], 3, 1)
        stock.Wave.Branch_Amplitude(domain[i], 3, 1)
        append_interval(domain, codomain, Stocks, i, i + Piece)
        i = i + Piece
    return

x = linspace(0,10,1000)
y = []

stock = StockFunction()
stock.Trend.Set_Slope(0)
stock.Volatility.Set_Amplitude(0)

algorithm(1000, 200, x, y, stock)


plt.plot(x,y)
plt.show()
'''
