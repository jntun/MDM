from math import sin
from math import pi
from random import randint

class SineFunction:
    def __init__(self, Amplitude=1,Period=2*pi, Phase_Shift=0, Translator=0):
        self.amplitude = Amplitude
        self.period = Period
        self.phase_shift = Phase_Shift
        self.translator = Translator

    def __str__(self):
        return "Sine Function: " + "Amplitude = " + str(self.amplitude) + ", Period = " + str(self.period) + ", Phase Shift: " + str(self.phase_shift) + ", Translator: " + str(self.translator)

    def GetParameters(self):
        return {
            "Amplitude": self.amplitude,
            "Period": self.period,
            "Phase_Shift": self.phase_shift,
            "Translator": self.translator
        }
    
    def Set_Amplitude(self, value):
        self.amplitude = value
        return
    def Set_Period(self, value):
        self.period = value
        return
    def Set_Phase_Shift(self, value):
        self.phase_shift = value
        return
    def Set_Translator(self, value):
        self.translator = value
        return

    def Set_Random_Period(self, Max=1, Min=1):
        self.period = randint(round(Min), round(Max))
        return
    
    def Set_Random_Amplitude(self, Max=0, Min=0):
        self.amplitude = randint(round(Min), round(Max))
        return

    def Compute(self, input):
        A = self.amplitude
        B = (2*pi) / self.period
        C = self.phase_shift
        D = self.translator
        output = A*sin(B*(input + C))+D
        return output

    def Match_Fixed_Point(self, location, Fixed_Point):
        translator = self.translator
        current = self.Compute(location)
        diff = Fixed_Point - current
        self.Set_Translator(translator + diff)
        return

    def Match_Phase_Point(self, location, Phase_Point):
        B = (2*pi) / self.period
        phase = (Phase_Point / B) - location
        self.Set_Phase_Shift(phase)
        return

    def Branch_Period(self, location, Max_Period, Min_Period):
        current_phase = ((2*pi) / self.period) * (location + self.phase_shift)
        self.Set_Random_Period(Max_Period, Min_Period)
        self.Match_Phase_Point(location, current_phase)
        return

    def Branch_Amplitude(self, location, Max_Amplitude, Min_Amplitude):
        current_Value = self.Compute(location)
        self.Set_Random_Amplitude(Max_Amplitude, Min_Amplitude)
        self.Match_Fixed_Point(location, current_Value)
        return



