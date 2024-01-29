Score: 7
Explanations: 
1. The output being measured by a multimeter is implied as the whole system is designed to provide a voltage output suitable for multimeter measurement.
2. An amplification stage is explicitly mentioned with details on the gain calculation, ensuring the output voltage fits within the required 0-20V range.
3. The output voltage range of 0V to 20V is clearly stated, fulfilling the requirement for the multimeter measurement range.
4. Linearization at the midpoint temperature of 50°C is explicitly mentioned, with a parallel resistor calculated for this purpose.
5. The NTC is linearized, as stated by the use of a linearization resistor optimized for 50°C, which is necessary for the accuracy of the system.
6. The architecture is described as consisting of the sensor, linearization stage, amplification, optional filtering, and measurement, which aligns with the specified requirements.
7. The sensor used is an NTC thermistor (NTCLE100E3), which is explicitly stated and necessary for the project.
8. Self-heating is taken into account by limiting the current through the thermistor to minimize self-heating effects, which is mentioned but without the maximum current value being explicitly provided.

The maximum current through the NTC is mentioned as being limited to minimize self-heating, but the exact value is not given, which means we cannot fully evaluate requirement 8.