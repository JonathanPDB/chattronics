Score: 7
Explanations: 
The project summary successfully covers the following requirements:

1. The use of a multimeter for measuring the output voltage is implied as the output voltage range is specified to be 0 to 20 Volts, which is a standard range for multimeters.
2. An amplification stage is present, with the selection of either an AD623 or INA128/INA129 instrumentation amplifier, and the gain is set post-empirical testing, which implies that the gain will be provided and justified based on the testing results.
3. The output voltage range is clearly stated to be between 0 and 20 Volts, which meets the specified requirement.
4. The NTC is linearized with a parallel resistor optimized for the midrange temperature of 50ºC, as detailed in the linearization stage section with calculations.
5. Linearization of the NTC is explicitly mentioned and calculated, ensuring that this requirement is absolutely necessary and met.
6. The architecture consists of the sensor (NTC thermistor), a linearization stage (parallel resistor), amplification (instrumentation amplifier), and optional filtering (RC low-pass filter), which aligns with the required architecture.
7. The sensor used is explicitly stated as a Vishay NTCLE100E3 NTC thermistor, meeting the requirement for an NTC sensor.
8. The self-heating effect is not explicitly addressed in the summary, and the maximum current passing through the NTC is not known, therefore this requirement is not met.

The project summary does not explicitly mention the self-heating effect or the maximum current through the NTC, which is required to be known. Therefore, this requirement is not fulfilled. All other requirements are covered by the project summary.