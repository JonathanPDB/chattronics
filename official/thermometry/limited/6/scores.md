Score: 8
Explanations: 
The project summary covers the following requirements:

1. The output is being measured by a multimeter: The summary explicitly mentions that the voltage output is readable by a multimeter.
2. There is an amplification stage with gain provided and justified: An Instrumentation Amplifier is used, with a gain set to amplify the bridge's output signal. Additionally, there is an Adjustable Gain Stage with a specified gain range.
3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: A Level Shifter is used to ensure the output voltage starts from 0V, and a Buffer Amplifier is implemented to provide the full 0-20V output range.
4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: A Linearization Resistor is calculated to match the resistance of the thermistor at 50°C.
5. The NTC is linearized: The use of a linearization resistor implies that the NTC is linearized.
6. The architecture follows the required stages: The summary describes a system with a sensor (thermistor), linearization stage, amplification, optional filtering, and measurement.
7. The sensor used is an NTC: The Vishay NTCLE100E3 thermistor is specified as the sensor.
8. The self-heating effect is taken into account: The summary states that the circuit is designed to limit power dissipation below 10mW and keep the biasing current through the thermistor below 1mA.

All the necessary requirements are met according to the project summary provided.