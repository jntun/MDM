from LinearFunction import LinearFunction
from SineFunction import SineFunction

class StockFunction:
    def __init__(self):
        self.wave = SineFunction()
        self.volatility = SineFunction()
        self.trend = LinearFunction()
    
    def __str__(self):
        return "Wave: [" + str(self.wave) + "], Volatility: [" + str(self.volatility) + "], Trend: [" + str(self.trend) + "]"
    
    def Get_Members(self):
        return {
            "Wave": self.wave.GetParameters(),
            "Volatility": self.volatility.GetParameters(),
            "Trend": self.trend.Get_Parameters()
        }
    
    def Compute(self, value):
        result = self.wave.Compute(value) + self.volatility.Compute(value) + self.trend.Compute(value)
        return result


test = StockFunction()
test.wave.Set_Period(1)
print(test.Get_Members())