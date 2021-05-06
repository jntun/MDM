package market

import (
    "fmt"
    "math"

    "github.com/Nastyyy/mdm-back/config"
)

type sine struct {
        amplitude  float64
        period     float64
        phaseShift float64
        translator float64
        x          float64
}

type linear struct {
        slope      float64
        translator float64
        x          float64
}

type PriceFunction struct {
        wave *sine
        volatility *sine
        trend *linear
}

func (sin sine) b() float64 {
        return (2*math.Pi) / sin.period
}

func (sin sine) s() float64 {
       return math.Sin(sin.b()*(sin.x + sin.phaseShift))
}

// SetAmp() takes a target amplitude, a and calculates the new translator value, d
func (sin sine) SetAmp(a float64) {
        d := (sin.amplitude-a) * sin.s() + sin.translator

        sin.amplitude = a
        sin.translator = d
}

func (sin *sine) value(i float64) float64 {
        sin.x = i
        return sin.amplitude * sin.s() + sin.translator
}

func (lin *linear) value(i float64) float64 {
        lin.x = i
        return lin.slope*lin.x+lin.translator
}

func (pf sine) String() string {
        return fmt.Sprintf("Sine | Amp: %f | Perd: %f | PhShf: %f | Trans: %f | X: %f", pf.amplitude, pf.period, pf.phaseShift, pf.translator, pf.x)
}

func (lin linear) String() string {
        return fmt.Sprintf("Linr | Slope: %f | Intercept: %f | X: %f", lin.slope, lin.translator, lin.x)
}

func (pf PriceFunction) value(i float64) float64 {
        config.StockLog(fmt.Sprintf("[pf.value] | pf.wave.value: %f + pf.volatility.value: %f + pf.trend.value: %f |", pf.wave.value(i), pf.volatility.value(i), pf.trend.value(i)))
        return pf.wave.value(i) + pf.volatility.value(i) + pf.trend.value(i)
}

func (pf *PriceFunction) NextPrice(i float64) float32 {
        price := pf.value(i)
        if price < 0 {
            price = 0.0
        }
        return float32(price)
}

func (pf PriceFunction) String() string {
        return fmt.Sprintf("\tWave > [%s]\n\tVola > [%s]\n\tTrnd > [%s]", pf.wave, pf.volatility, pf.trend)
}

func GeneratePriceFunc() *PriceFunction {
        //wave := &sine{1.0, 2*math.Pi, 0, 0, 0}
        wave := &sine{0.01, 0.1, 0, 0, 0}
        vola := &sine{0.01, 10, 0, 0, 0}
        trnd := &linear{0.0, 1.0, 0}
        return &PriceFunction{wave, vola, trnd}
}
