Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is being measured by a multimeter - This requirement is implicitly met as the output voltage range is specified to be 0-20V which is a standard range for multimeter measurements, and it mentions that multimeter requirements include high input impedance and resolution capable of reading small voltage changes.

2. There should be an amplification stage and the gain must be provided and justified - The gain of the instrumentation amplifier is stated to be 25 V/V, which is assumed from estimated sensitivity across the thermistor network. This implies that an amplification stage is present and the gain is provided and justified.

3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts - This is explicitly stated in the "Voltage Output" section.

4. The NTC is linearized by a resistor optimized for the midrange 50 ºC - The linearization method is described using a parallel resistor with the thermistor to linearize response at the midpoint of the temperature range (50°C).

5. The NTC is linearized - As per the previous point, the linearization method is specified, meaning this requirement is met.

6. The architecture should be roughly consisting of the sensor, a linearization stage, amplification, optional filtering, and measurement - The project summary describes an NTC thermistor as the sensor, a resistor network for linearization, buffer and instrumentation amplifiers for amplification, a low-pass filter, and voltage output for measurement, which aligns with the required architecture.

7. The sensor used is an NTC - The sensor specified is the Vishay NTCLE100E3 NTC thermistor, fulfilling this requirement.

8. The self-heating effect is taken into account and there the maximum current that passes through the NTC is known - While the self-heating effect is mentioned by stating the dissipation constant, the actual maximum current through the NTC is not provided. Therefore, this requirement is not met.

Requirements not met: 8

Requirements met: 1, 2, 3, 4, 5, 6, 7.