Score: 5
Explanations: 
The project summary covers the following requirements:

1. The output voltage range to be measured by the multimeter is mentioned to be between 0 and 20 Volts, fulfilling requirement 3.
2. There is mention of an amplification stage with a required gain (Av) of 20, which is provided and justified by assuming a maximum sensor voltage (Vsensor_max) of 1V, fulfilling requirement 2.
3. Linearization is addressed by using a resistor in parallel with the NTC thermistor, optimized for 50°C, which matches requirement 4. However, the explicit value of the resistor or the actual resistance of the NTC at 50°C is not provided, so the exact linearization cannot be verified.
4. The use of an NTC thermistor sensor is explicitly stated, fulfilling requirement 7.
5. Self-heating is considered, with a power dissipation limit set to below 0.1W, which suggests that the maximum current through the NTC is known, fulfilling requirement 8.
6. The architecture described in the summary follows the sensor, linearization stage, amplification, optional filtering, and measurement, which aligns well with requirement 6.

The requirements that are not explicitly fulfilled in the summary are:
- Requirement 5 states that the NTC must be linearized, and while the summary mentions linearization, it does not provide evidence that the linearization is absolutely correct, so this cannot be confirmed.
- Requirement 1 about the output being measured by a multimeter is not directly stated, but it is implied that a multimeter is used for reading the voltage range.

Overall, the project summary explicitly covers 5 out of the 8 listed requirements. Since the linearization is mentioned but not fully justified with values to confirm its accuracy, we cannot count requirement 5 as fulfilled. The output being measured by a multimeter is implied but not directly stated, so requirement 1 is not counted.