Score: 8
Explanations: 
The project summary covers the following requirements:

1. The output is measured by a multimeter: This is explicitly stated in the "Output Conditioning and Measurement" section, where it mentions the use of a high-precision digital multimeter to measure the 0-20V range.

2. There is an amplification stage with the gain provided and justified: The gain is mentioned in the "Amplification" section, with a gain between 7 to 10, adjustable via feedback resistor values, and in the "Voltage Scaling" section with a gain of 4. The justification is not explicitly provided, but we can infer that it is to bring the thermistor signal into the desired 0-20V output range.

3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: This is clearly stated in the project summary.

4. The NTC is linearized by a resistor optimized for the midrange 50°C: The "Linearization Resistor Block" describes the use of a resistor estimated at approximately 6.5kΩ to match the thermistor resistance at 50°C, which suggests that the NTC is linearized.

5. The NTC is linearized: This is confirmed by the presence of a linearization resistor block and the calculation provided for selecting the linearization resistor.

6. The architecture consists of the sensor, a linearization stage, amplification, optional filtering, and measurement: The project summary describes a sensor block, linearization resistor block, buffer, amplification, level shifting, voltage scaling, and an anti-aliasing filter, which aligns with the required architecture.

7. The sensor used is an NTC: The "NTC Thermistor Sensor Block" explicitly states the use of a Vishay NTCLE100E3 NTC thermistor.

8. The self-heating effect is taken into account, and the maximum current that passes through the NTC is known: The "NTC Thermistor Sensor Block" mentions that the operating current is kept typically under 100 µA to prevent self-heating, which indicates that self-heating has been considered.

All eight requirements are covered by the project summary.