Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is measured by a multimeter: The summary explicitly mentions the use of a multimeter with a high input impedance and a resolution suitable for the 0-20V range.
2. There is an amplification stage with gain provided and justified: The gain stage is described with an assumed gain of approximately 10 to map the input voltage to the desired output range, with component values provided (R1, R2, Rtrim).
3. The output voltage range is 0-20 Volts: The system is designed to output a voltage in this range, with level shifting and output stages mentioned to ensure that the output starts at 0V.
4. The NTC is linearized for the midrange 50ºC: A parallel resistor is used to linearize the response at the midpoint, with an approximate value given (3.3 kΩ).
5. The NTC is linearized: This is explicitly stated and is a critical aspect of the design.
7. The sensor used is an NTC: The Vishay NTCLE100E3 NTC thermistor is specified.
8. Self-heating effect is considered: The summary states that a low excitation current is used to keep the self-heating effect below 1%.

The project summary does not explicitly cover these requirements:

6. The architecture is not explicitly described in the order mentioned (sensor, linearization, amplification, optional filtering, measurement), but all elements are present.
8. The maximum current passing through the NTC is not given, only that the self-heating effect is below 1%. Without specific current values, we cannot confirm if the maximum current is known.

Based on the summary, the system is well-designed and meets most of the requirements. However, there is a lack of specific details on the architecture's order and the maximum current through the NTC.