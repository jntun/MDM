from LinearFunction import LinearFunction
from SineFunction import SineFunction

class StockFunction:
    def __init__(self):
        self.Wave = SineFunction()
        self.Volatility = SineFunction()
        self.Trend = LinearFunction()
    
    def __str__(self):
        return "Wave: [" + str(self.Wave) + "], Volatility: [" + str(self.Volatility) + "], Trend: [" + str(self.Trend) + "]"
    
    def Get_Members(self):
        return {
            "Wave": self.Wave.GetParameters(),
            "Volatility": self.Volatility.GetParameters(),
            "Trend": self.Trend.Get_Parameters()
        }
    
    def Compute(self, value):
        result = self.Wave.Compute(value) + self.Volatility.Compute(value) + self.Trend.Compute(value)
        return result
