Score: 7
Explanations: 
The project summary covers the following requirements:
1. The output is being measured by a multimeter: The summary explicitly mentions that the voltage output is to be read by a standard multimeter.
2. There should be an amplification stage and the gain must be provided and justified: There is a gain stage mentioned, and a general justification for the gain is given based on the voltage range required (0 to 20V), but the exact gain value is not provided as it depends on the actual voltage change across the thermistor.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: This is clearly stated in the project summary in the output stage section.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: The summary specifies the use of a 6.2k ohm precision resistor for linearization at the midpoint temperature of 50 degrees Celsius.
5. The NTC is linearized: The linearization process is explicitly mentioned with the use of a parallel resistor.
6. The architecture consists of sensor, linearization stage, amplification, optional filtering, and measurement: The project includes all these stages, with the sensor being the NTC thermistor, the linearization stage using a parallel resistor, the amplification stage including a buffer and gain stage, optional filtering with a low-pass filter, and the measurement is the voltage output for the multimeter.
7. The sensor used is an NTC: The sensor model Vishay NTCLE100E3 is an NTC thermistor, fulfilling this requirement.
8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: The summary does not provide information on the self-heating effect or the maximum current passing through the NTC, so this requirement is not met.

The project does not meet the requirement for providing the maximum current through the NTC to account for self-heating effects.