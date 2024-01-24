Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output is measured by a multimeter: This is implicitly covered as the final stage includes circuit protection and filtering suitable for connection to a multimeter.
2. Amplification stage with gain provided and justified: The non-inverting operational amplifier configuration is used with a calculated gain to map the temperature range to a 0-20V output; the gain is 0.25V/°C for an 80°C span which is justified.
3. Output voltage range between 0 and 20 Volts: The gain calculation ensures the output voltage is within this range for the temperature span considered.
4. NTC linearized at midrange 50°C: A parallel resistor of approximately 10kΩ is used for linearizing the NTC at 50°C.
5. NTC linearization is absolutely necessary and is covered: The use of a linearization resistor confirms this requirement is met.
6. Architecture: The project consists of a sensor (NTC), linearization stage (resistor), amplification (operational amplifier), and optional buffer/output stage which could be considered as filtering and measurement. This matches the required architecture.
7. The sensor used is an NTC: The Vishay NTCLE100E3103JB0 is specified as the sensor, meeting this requirement.
8. Self-heating effect is taken into account with maximum current known: Self-heating is limited to 1%, and a low power dissipation thermistor is selected, although the exact maximum current is not specified.

The requirements that are not explicitly met include:
- While self-heating is mentioned, the exact maximum current through the NTC is not provided. Without this value, we cannot confirm that requirement 8 is fully met.

Therefore, the project summary successfully covers 7 out of the 8 requirements.