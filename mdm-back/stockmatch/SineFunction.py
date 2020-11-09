from math import sin
from math import pi

class SineFunction:
    def __init__(self, Amplitude=1,Period=2*pi, Phase_Shift=0, Vertical_Shift=0):
        self.amplitude = Amplitude
        self.period = Period
        self.phase_shift = Phase_Shift
        self.vertical_shift = Vertical_Shift

    def __str__(self):
        return "Sine Function: " + "Amplitude = " + str(self.amplitude) + ", Period = " + str(self.period) + ", Phase Shift: " + str(self.phase_shift) + ", Vertical Shift: " + str(self.vertical_shift)

    def GetParameters(self):
        return {
            "Amplitude": self.amplitude,
            "Period": self.period,
            "Phase_Shift": self.phase_shift,
            "Vertical_Shift": self.vertical_shift
        }
    
    def Set_Amplitude(self, value):
        self.amplitude = value
    def Set_Period(self, value):
        self.period = value
    def Set_Phase_Shift(self, value):
        self.phase_shift = value
    def Set_Vertical_Shift(self, value):
        self.vertical_shift = value

    def Compute(self, input):
        A = self.amplitude
        B = (2*pi) / self.period
        C = self.phase_shift
        D = self.vertical_shift
        output = A*sin(B*(input + C)+D)
        return output

