Score: 7
Explanations: 
The provided summary of the project covers the following requirements:

1. The output is being measured by a multimeter: The summary mentions using a high-quality digital multimeter to measure the output voltage.
2. There should be an amplification stage and the gain must be provided and justified: The summary includes a non-inverting operational amplifier configuration with a gain of 4, justified by the desired output range of 0 to 20 Volts given a sensor output of 0 to 5 Volts. However, the summary also mentions a real gain of 19, which contradicts the previous gain of 4 and is not justified or explained in the context of the required output voltage range.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: The gain calculation suggests that the system is designed to amplify the output to this range.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: A parallel resistor (Rp) is chosen to match the resistance of the thermistor at 50 degrees Celsius, indicating that the linearization is targeted for the midrange.
5. The NTC is linearized: The project includes a linearization network with a parallel resistor.
6. The architecture should be roughly consisting of the sensor, a linearization stage, amplification, optional filtering, and measurement: The summary describes an architecture that includes these stages.
7. The sensor used is an NTC: The sensor mentioned is a Vishay NTCLE100E3 thermistor, confirming the use of an NTC.
8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: The summary mentions selecting a series resistor to limit the current and minimize self-heating to less than 1%, although the exact maximum current is not provided.

The requirements not fully covered are:
- The gain of the amplification stage is not consistently justified, as there is a contradiction in the gain values (requirement 2).
- The exact maximum current through the NTC is not provided (part of requirement 8), but the intention to limit self-heating is mentioned.

Given these points, the project summary meets 7 out of the 8 requirements.