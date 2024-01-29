Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is measurable by a multimeter: The summary explicitly states that the output voltage range is 0 to 20V, which is within the capabilities of a standard multimeter.
2. Amplification stage with gain provided and justified: The gain stage is described with a calculated gain of 20 to achieve a 0-20V output from a 0.1V to 1V range. Additionally, there's mention of an adjustable DC bias in the level shifter stage, which might also contribute to the overall gain.
3. Output voltage range is between 0 and 20 Volts: The summary specifies an output voltage range of 0 to 20V.
4. NTC is linearized by a resistor optimized for the midrange 50ºC: The linearization resistor is chosen to match the resistance of the NTC at 50ºC, which implies linearization around the midrange.
5. NTC is linearized: As mentioned above, the project summary describes the use of a linearization resistor to achieve this requirement.
6. Architecture outline: The project summary lists the sensor (NTC thermistor), a linearization stage (linearization resistor), amplification (gain stage and level shifter), optional filtering (low-pass filter), and measurement, which aligns with the requested architecture.
7. Sensor used is an NTC: The sensor is explicitly stated to be a Vishay NTCLE100E3 NTC thermistor.
8. Self-heating effect is considered: The summary does not mention the self-heating effect or the maximum current through the NTC, which is a required element to evaluate whether this effect is accounted for.

The requirements not explicitly covered in the summary are the maximum current through the NTC to consider the self-heating effect. Due to the absence of this information, the score is 7.