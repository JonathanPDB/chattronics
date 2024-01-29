Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is measured by a multimeter: This requirement is implicit as the design includes an output stage that provides a 0-20V range, which is a standard range for multimeter measurements.
2. There is an amplification stage with a gain provided: The gain is explicitly stated as 200 in the differential amplifier section and 10 in the final output stage. The justification for the gain is given by mapping the Wheatstone bridge output to the 0-20V output range.
3. The output voltage range is between 0 and 20 Volts: The final output stage is designed to provide this range, as explicitly stated in the summary.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: The linearization resistor is mentioned, which implies that it is chosen to linearize the NTC around 50°C.
5. The NTC is linearized: The summary mentions the use of a linearization resistor to create a linear response, satisfying this requirement.
6. The architecture consists of the sensor, a linearization stage, amplification, optional filtering, and measurement: The summary describes each of these stages, fulfilling this requirement.
7. The sensor used is an NTC: The NTCLE100E3 thermistor is explicitly mentioned as the temperature sensor.
8. The self-heating effect is taken into account, and the maximum current passing through the NTC is known: The self-heating effects are mentioned to be kept below 1% by using high-value resistors in the bridge circuit, although the exact maximum current is not stated.

The project summary does not explicitly provide the exact maximum current through the NTC (requirement 8). However, it does mention that self-heating effects are kept below 1%, which implies that the maximum current is considered but not quantified. Since an exact numeric value is not provided, it is not possible to evaluate the accuracy of this statement. Therefore, requirement 8 is not fully satisfied.