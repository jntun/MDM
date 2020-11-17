from random import randint

class LinearFunction:
    def __init__(self, Slope=0, Translator=0):
        self.slope = Slope
        self.translator = Translator
    
    def __str__(self):
        return "Slope: " + str(self.slope) + ", Intercept: " + str(self.translator)

    def Get_Parameters(self):
        return {
            "Slope": self.slope,
            "Intercept": self.translator
        }
    def Set_Slope(self, value):
        self.slope = value
        return
    def Set_Translator(self, value):
        self.translator = value
        return

    def Set_Random_Slope(self, Max=0, Min=0):
        self.slope = randint(round(Min), round(Max))
        return
    
    def Compute(self, input):
        M = self.slope
        B = self.translator
        output = M*input + B
        return output
    
    def Match_Fixed_Point(self, location, Fixed_Point):
        translator = self.translator
        current = self.Compute(location)
        diff = Fixed_Point - current
        self.Set_Translator(translator + diff)
        return
    
    def Branch_Function(self, location, Max_Slope=0, Min_Slope=0):
        target = self.Compute(location)
        self.Set_Random_Slope(Max_Slope, Min_Slope)
        self.Match_Fixed_Point(location, target)
        return

