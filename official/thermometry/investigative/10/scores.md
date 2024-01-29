Score: 6
Explanations: 
The project summary covers several of the given requirements, but not all are explicitly stated:

1. The output is being measured by a multimeter. This is mentioned in the overview, "output a voltage easily measurable by a multimeter."
2. There should be an amplification stage and the gain must be provided and justified. The gain is provided as a two-stage approach with initial gain of 100 followed by a gain of 2. It is justified by assuming a 5V max input from the Wheatstone bridge and the need to fit the 0-20V output range.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts. This is explicitly stated in the section "Output Voltage (0-20V)."
4. The NTC is linearized by a resistor optimized for the midrange 50ºC. This is stated in the "Linearizing Resistor" section, although the exact value determination is deferred to datasheet consultation.
5. The NTC is linearized. This is the same as point 4 and is covered.
6. The architecture is roughly consisting of the sensor, a linearization stage, amplification, optional filtering, and measurement. This is explicitly stated in the "Architecture Overview."
7. The sensor used is an NTC. This is mentioned in the "Thermistor" section with the specific part number NTCLE100E3103GB0.
8. The self-heating effect is taken into account and there the maximum current that passes through the NTC is known. This requirement is not explicitly covered in the summary provided. There is no mention of the self-heating effect or the maximum current passing through the NTC.

The score is 6, as requirements 1, 2, 3, 4, 5, 6, and 7 are covered, but requirement 8 is not addressed in the project summary.