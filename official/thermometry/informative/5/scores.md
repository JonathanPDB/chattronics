Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is being measured by a multimeter - The summary indicates that a precision multimeter is used for direct measurement of the output voltage.
2. There should be an amplification stage and the gain must be provided and justified - The summary explains a non-inverting operational amplifier configuration with a gain of 20 to achieve an output voltage range of 0 to 20 Volts, which is justified by the signal conditioning stage outputting a 0 to 1 Volt signal.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts - The project summary explicitly states that the output voltage range that correlates with the temperature range is 0 to 20 Volts.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC - The summary mentions using a parallel resistor to linearize the thermistor response at the midpoint (50°C), although the resistance value needs to be calculated or obtained from the datasheet.
5. The NTC is linearized - This is implied by the inclusion of a linearization resistor in the design.
6. The architecture should roughly consist of the sensor, a linearization stage, amplification, optional filtering, and measurement - The summary describes the use of a sensor (NTC thermistor), a linearization stage (parallel resistor), amplification (operational amplifier), filtering (Butterworth filter), and measurement (multimeter).
7. The sensor used is an NTC - The Vishay NTCLE100E3 series NTC thermistor is proposed for temperature sensing.
8. The self-heating effect is taken into account, and the maximum current that passes through the NTC is known - This requirement has not been explicitly mentioned or addressed in the project summary.

The project summary does not provide information about the self-heating effect or the maximum current through the NTC, which is a requirement for this review. Therefore, 7 out of 8 requirements are met.