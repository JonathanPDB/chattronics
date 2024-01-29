Score: 6
Explanations: 
The project summary covers the following requirements:

1. The output is being measured by a multimeter: This requirement is explicitly mentioned in the summary, stating that the system produces an analog voltage output that can be measured by a multimeter.
2. There should be an amplification stage and the gain must be provided and justified: The summary describes the use of an instrumentation amplifier with adjustable or fixed gain. Although the specific gain value is not provided, the requirement for amplification is met.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: The summary clearly states that the output voltage range is 0 to 20 Volts.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: This requirement is not explicitly mentioned in the summary. There is no discussion of a linearization resistor or method.
5. The NTC is linearized (this is absolutely necessary and fatal if false): As in item 4, the summary does not address linearization, which is a critical requirement for the project.
6. The architecture should consist of the sensor, a linearization stage, amplification, optional filtering, and measurement: The summary describes an architecture with a sensor, amplification, filtering, and measurement but fails to mention a linearization stage.
7. The sensor used is an NTC: The summary specifies the use of a Vishay NTCLE100E3 NTC thermistor.
8. The self-heating effect is taken into account, and the maximum current that passes through the NTC is known: The summary mentions a current range of 100 µA to 1 mA to minimize self-heating, which satisfies this requirement.

The requirements that were not met are:
- Requirement 4: There is no mention of a resistor optimized for linearization at the midrange temperature of 50 ºC.
- Requirement 5: The summary does not confirm that the NTC is linearized, which is a critical omission.

Therefore, the project meets 6 out of the 8 requirements.