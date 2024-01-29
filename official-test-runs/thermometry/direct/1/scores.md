Score: 6
Explanations: 
The project summary covers many of the requirements, but does not explicitly mention all of them:

1. The output is being measured by a multimeter: This is indicated as the voltage output is meant to be read by a high-impedance digital multimeter.
2. There should be an amplification stage and the gain must be provided and justified: The summary mentions a non-inverting amplifier configuration to scale the output to the 0-20V range, but does not provide the exact gain value or the justification.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: The project is designed to scale the bridge output to the 0-20V range, which meets this requirement.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: The summary states that linearization is achieved with a fixed resistor in parallel at the midpoint (50°C), based on the thermistor's characteristics.
5. The NTC is linearized: As mentioned in point 4, linearization is explicitly addressed in the project.
6. The architecture consists of the sensor, a linearization stage, amplification, optional filtering, and measurement: The project includes all these stages, meeting the requirement.
7. The sensor used is an NTC: The project uses a Vishay NTCLE100E3 thermistor, which is an NTC.
8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: The summary mentions that the self-heating effect is kept under 1% by minimizing power dissipation, but it does not provide the maximum current value.

The project summary does not provide specific values or calculations for the gain of the amplification stage or the maximum current through the NTC, which are necessary to fully meet requirements 2 and 8. The summary does cover requirements 1, 3, 4, 5, 6, and 7. Therefore, the score is 6.