Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is being measured by a multimeter: The project mentions the use of a multimeter capable of measuring 0-20 volts DC, which is the output voltage range.
2. There should be an amplification stage and the gain must be provided and justified: The gain stage is described with a non-inverting amplifier configuration, and the gain is calculated to be approximately 22 to achieve the 0-20V output range from a linearization output range of 0.1V to 0.9V.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: The output stage is designed to amplify the signal to the required 0-20V output range.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: A precision metal film resistor of 6.2kΩ is used for linearization at the 50°C midpoint.
5. The NTC is linearized (this is absolutely necessary and fatal if false): The linearization is explicitly mentioned and a resistor value is chosen for this purpose.
6. The architecture should be roughly consisting of the sensor, a linearization stage, amplification, optional filtering, and measurement: The project describes an architecture with a sensor (NTC thermistor), linearization stage (Wheatstone bridge with linearization resistor), amplification (gain stage), and measurement (multimeter).
7. The sensor used is an NTC (this is absolutely necessary and fatal if false): The sensor is specified as an NTC thermistor (Vishay NTCLE100E3).
8. The self-heating effect is taken into account and there the maximum current that passes through the NTC is known: The project does not explicitly state the consideration of the self-heating effect or the maximum current that passes through the NTC, which is a requirement that is not met in the summary provided.

Based on the information provided, the project meets 7 out of 8 requirements. The only requirement not met is the consideration of the self-heating effect and the specification of the maximum current through the NTC.