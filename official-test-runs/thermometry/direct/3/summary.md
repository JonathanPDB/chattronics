Analog Water Temperature Measurement System

This system measures the temperature of water using an NTC thermistor (Vishay NTCLE100E3) and outputs a voltage that can be measured by an external multimeter.

Sensor Block:
- Sensor: Vishay NTCLE100E3 NTC Thermistor
- Estimated Resistance at 25°C: 10kΩ (based on typical values)
- Estimated Resistance at Midpoint (50°C): 5kΩ (based on typical curve for NTCLE100E3)
- Self-Heating: Controlled by biasing with a low current to keep the effect less than 1%

Parallel Linearization Resistor:
- Linearization Resistor Value: 5kΩ precision resistor with a low temperature coefficient (<50 ppm/°C)
- The resistor value is calculated using the formula Rlin = R(50°C) * sqrt(R(10°C)/R(90°C))

Constant Current Source:
- Target Bias Current: 100 µA to limit self-heating
- Power Supply: Assumed to be ±15V for op-amp based current source
- Reference Voltage: 2.5V precision voltage reference
- Set Resistor (Rset): 24.9 kΩ (closest standard E96 series value to the calculated 25 kΩ)

Instrumentation Amplifier:
- Op-amp Choice: AD620 or similar with a high input impedance (>1MΩ)
- Power Supply: Assumed to be +24V single supply
- Gain Calculation: Maximum gain of 20 to amplify a 1V input to 20V output
- Gain Resistor (Rg) for AD620: 2.6kΩ for a gain of 20

Low-Pass Filter:
- Filter Type: Second-order active low-pass Butterworth filter
- Cutoff Frequency: 5 Hz to attenuate noise while preserving signal integrity
- Resistor Value: 316 kΩ (closest standard E96 series value)
- Capacitor Value: 0.1 µF
- Op-amp Choice for Filter: TL072 or similar with a bandwidth greater than 50 Hz

Voltage Output Stage:
- Op-amp Choice: LM324 or similar
- Power Supply: Assumed to be ±15V (if single supply, virtual ground created with voltage divider)
- Gain Setting Resistors: R1 = 20kΩ, R2 = 1kΩ for a gain of 21 (a little above 20 to account for component tolerances)
- Level Shifting: Implemented if signal from filter stage does not start from 0V (additional components as needed)
- Output Resistor: 100Ω for short-circuit protection

Please note that the values and assumptions made here are based on typical specifications and common practice. They should be validated and, if necessary, adjusted for the actual characteristics of the components used and the specific conditions of the application. Calibration of the system using known temperature references is crucial for accurate temperature measurements.