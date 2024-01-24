Score: 4
Explanations: 
The project summary provides a reasonable level of detail for a review. Here is the evaluation based on the provided requirements:

1. The output is being measured by a multimeter: The summary indicates that a high-resolution digital multimeter (DMM) is used for measuring the output voltage, which fulfills this requirement.

2. There should be an amplification stage, and the gain must be provided and justified: The summary describes a differential amplifier with a gain of approximately 200 to achieve a 0 to 20V output from a millivolt-level input, and a non-inverting amplifier configuration to map the bridge voltage (0V to 4V) to 0 to 20V output. The gain values and their purposes are provided and justified.

3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: The gain stage is specifically calculated to ensure the output voltage range is 0 to 20V, which meets this requirement.

4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: There is no explicit mention of linearization or the optimization of a resistor for this purpose, so this requirement is not met.

5. The NTC is linearized (this is absolutely necessary and fatal if false): As above, there is no explicit mention of linearization, so this requirement is not met.

6. The architecture should roughly consist of the sensor, a linearization stage, amplification, optional filtering, and measurement: The architecture described includes the sensor, amplification, filtering, and measurement stages, but lacks an explicit linearization stage, so this requirement is partially met.

7. The sensor used is an NTC (this is absolutely necessary and fatal if false): The sensor mentioned is indeed an NTC thermistor (Vishay NTCLE100E3), which satisfies this requirement.

8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: The summary does not mention the self-heating effect or the maximum current through the NTC, therefore this requirement is not met.

In summary, the project meets requirements 1, 2, 3, and 7. Requirements 4, 5, 6, and 8 are not met based on the information provided.