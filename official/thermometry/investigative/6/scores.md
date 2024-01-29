Score: 4
Explanations: 
1. The output is measured by a multimeter: The project summary describes a voltage range of 0 to 20V output suitable for measurement by a multimeter, but it does not explicitly state that the output is being measured by a multimeter.

2. There is an amplification stage with a gain provided and justified: The project summary includes an Instrumentation Amplifier Block with a gain calculation, where Av = 20 is justified based on an assumed Vin_max of approximately 1V from the thermistor.

3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: This is explicitly covered in the Instrumentation Amplifier Block, which specifies a voltage range of 0 to 20V output.

4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: The project summary includes a Linearization Resistor Block, but it does not explicitly state that the resistor is optimized for the midrange of 50 ºC.

5. The NTC is linearized: The presence of a Linearization Resistor Block implies that linearization is part of the design, but there is no explicit statement or evidence that the linearization is effective (i.e., that the NTC is indeed linearized).

6. The architecture consists of the sensor, linearization stage, amplification, optional filtering, and measurement: The project summary describes a Thermistor Sensor Block, Linearization Resistor Block, Buffer Op-Amp Block, Instrumentation Amplifier Block, Anti-Aliasing Filter Block, and Output Stage Block, which align with the required architecture.

7. The sensor used is an NTC: The project summary specifies the use of a Vishay NTCLE100E3 thermistor, which is an NTC thermistor.

8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: The project summary does not mention the self-heating effect or provide information on the maximum current passing through the NTC. Therefore, this requirement is not met.

The project summary covers requirements 2, 3, 6, and 7. Requirements 1, 4, 5, and 8 are not explicitly met based on the information provided.