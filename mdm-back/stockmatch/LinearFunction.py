
class LinearFunction:
    def __init__(self, Slope=0, Intercept=0):
        self.slope = Slope
        self.intercept = Intercept
    
    def __str__(self):
        return "Slope: " + str(self.slope) + ", Intercept: " + str(self.intercept)

    def Get_Parameters(self):
        return {
            "Slope": self.slope,
            "Intercept": self.intercept
        }
    def Set_Slope(self, value):
        self.slope = value
    def Set_Intercept(self, value):
        self.intercept = value

    def Compute(self, input):
        M = self.slope
        B = self.intercept
        output = M*input + B
        return output