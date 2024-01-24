Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is being measured by a multimeter: The use of a standard multimeter for reading the output voltage is mentioned.

2. There should be an amplification stage and the gain must be provided and justified: An amplification stage with a gain of 20 is described to scale a 0-1V signal to a 0-20V range, which is justified based on the expected voltage range from the sensor.

3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: The gain stage is designed to amplify the signal to a 0-20V range, suitable for multimeter measurement.

4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: The linearization resistor block indicates matching the thermistor resistance at 50°C with a parallel resistor to linearize the response at the midpoint of the desired temperature range.

5. The NTC is linearized: The linearization process is explicitly mentioned as part of the design.

6. The architecture should be roughly consisting of the sensor, a linearization stage, amplification, optional filtering, and measurement: The architecture is described following these stages, including a buffer amplifier, gain stage, level shifter, low-pass filter, and output amplifier.

7. The sensor used is an NTC: The NTC thermistor Vishay NTCLE100E3 is specified as the temperature sensor.

8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: This requirement is not explicitly addressed in the summary. There is no mention of the maximum current through the NTC or considerations for the self-heating effect.

Based on the provided summary, the project fulfills 7 out of the 8 requirements. No information is given regarding the self-heating effect and the maximum current through the NTC, which is a critical requirement for the project.