Water Temperature Measurement System Using an NTC Thermistor

Overview:
The system is designed to measure the temperature of water in a beaker using a Vishay NTCLE100E3 NTC thermistor and output a voltage between 0 and 20 Volts, corresponding to the temperature range of 10°C to 90°C. The output voltage is measured using a multimeter. The system is fully analog.

NTC Thermistor:
We will use a standard 10k ohm NTC thermistor at 25°C, which typically has a beta value of around 3950K. This will provide the resistance change necessary for temperature measurement in the specified range.

Linearization Network:
To linearize the response of the thermistor, a parallel resistor is used. The value of this resistor is chosen to match the resistance of the NTC at the midpoint temperature (50°C), which we've estimated at 3.3k ohms. Linearization is essential for a more accurate voltage output relative to the temperature.

- R_parallel (for linearization at 50°C): 3.3k ohms

Buffer Amplifier:
A voltage follower (unity gain buffer) topology is suggested. The gain is set to unity (Gain = 1) to isolate the linearization network from the scaling amplifier. The amplifier should have high input impedance to minimize loading effects and low output impedance to drive the scaling amplifier.

Scaling Amplifier:
A non-inverting amplifier topology with a gain sufficient to scale the maximum expected voltage from the linearization network (assumed to be less than 1V) to the full range of the multimeter (0-20 Volts) is suggested. With an assumed input voltage of 1V, the gain required is 20.

- Amplifier Gain (Av): 20
- Rin (Input Resistor): 10 kΩ
- Rf (Feedback Resistor): 190 kΩ

Power Supply for the Scaling Amplifier:
- Assuming a standard single-supply voltage of 24V, which provides enough headroom for the amplifier.

Output Voltage Measurement:
A high-precision digital multimeter with at least 0.1% measurement accuracy and an input impedance of 10 MΩ or greater is suggested to ensure accurate voltage measurements. Use shielded cables to connect the multimeter to avoid noise and interference.

Environmental and Protective Considerations:
- Implement transient voltage suppressor (TVS) diodes or varistors for transient protection at the output.
- Use a robust enclosure and conformal coating on the PCB if operating in harsh environments.

Please note that the given component values and suggestions are based on standard practices and typical thermistor characteristics. Actual design and final component selection should be validated with the specific NTC thermistor datasheet values and calibrated for precise temperature measurement.