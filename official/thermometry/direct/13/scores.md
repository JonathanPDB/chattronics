Score: 7
Explanations: 
The project summary covers the following items:

1. The output is measured by a multimeter: The summary specifies the use of a multimeter capable of measuring up to 20V.
2. There is an amplification stage, and the gain is provided and justified: The gain for the signal conditioning amplifier is calculated to be 10 to achieve the desired output range, and another non-inverting amplifier with a gain of 10 is mentioned for the output stage.
3. The output voltage range is 0-20V: The summary states that the system outputs a voltage linearly corresponding to the measured temperature range, with an output range of 0 to 20V.
4. The NTC is linearized by a resistor optimized for the midrange 50ºC: A resistor in parallel with the NTC thermistor is mentioned to linearize its response around 50°C.
5. The NTC is linearized: As noted above, a linearization network is included to ensure the response is linear around the midrange temperature.
6. The architecture consists of the sensor, a linearization stage, amplification, and measurement: The summary outlines a system with a sensor (NTC thermistor), a linearization network, buffer and signal conditioning amplifiers, and an output stage for measurement by a multimeter.
7. The sensor used is an NTC: The Vishay NTCLE100E3 thermistor is specified as the temperature sensor.
8. The self-heating effect is taken into account: The summary mentions that self-heating is managed by keeping the operating current low, although the exact current is not specified.

The aspects not explicitly covered are:
- The exact maximum current that passes through the NTC is not provided, so we cannot confirm if the self-heating effect is within acceptable limits.

Based on the provided information, the project successfully covers 7 out of the 8 requirements.