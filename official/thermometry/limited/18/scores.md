Score: 7
Explanations: 
The project summary covers the following requirements:

1. The use of a multimeter capable of measuring voltages in the 0-20V range is clearly stated, as precision multimeters like Fluke 287 or Fluke 289 are mentioned.

2. An amplification stage is included, with a gain of 10, which is justified by the need to convert a 2V span into a 20V span.

3. The output voltage range is specifically set to match 0-20V, as indicated by the inclusion of an offset adjustment and the design of the output stage.

4. The linearization of the NTC thermistor is addressed with the use of a linearization resistor estimated at 5kΩ, optimized for the midrange 50°C.

5. The NTC is explicitly mentioned as being linearized, which is a critical requirement.

6. The architecture is described as consisting of the sensor (NTC thermistor), a linearization stage, amplification, and an optional filtering stage before measurement, which matches the required structure.

7. The sensor used is identified as an NTC thermistor (Vishay NTCLE100E3 series), satisfying this requirement.

8. While the self-heating effect is not explicitly mentioned, the summary does indicate the use of a buffer with a high input impedance operational amplifier, suggesting that the current through the NTC thermistor is considered to minimize self-heating. However, without clear evidence of the maximum current, this requirement is not fully met.

Requirements not fully met:

8. The maximum current through the NTC is not specified, so it is not possible to confirm that the self-heating effect has been adequately taken into account.

Based on the information provided, 7 out of the 8 requirements have been met. The project summary does not provide a specific value for the maximum current through the NTC, which is necessary to confirm that self-heating has been considered.