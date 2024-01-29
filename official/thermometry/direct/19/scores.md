Score: 7
Explanations: 
The project summary provided covers several of the requirements, but not all. Here's the breakdown:

1. Output measured by a multimeter: The summary specifies the use of a precision digital multimeter (DMM) for output voltage measurement, which meets this requirement.

2. Amplification stage with gain provided and justified: The summary describes a non-inverting operational amplifier configuration with a gain of 10 to achieve a 0-20V range from a 0.5V to 2.5V input. The gain is justified based on the desired output voltage range, fulfilling this requirement.

3. Output voltage range between 0 and 20 Volts: The gain of 10 for the amplification stage is designed to amplify the input voltage range to the specified output voltage range of 0-20V, thus meeting this requirement.

4. NTC linearized by a resistor optimized for the midrange 50ºC: The summary mentions a parallel resistor (Rp) of approximately 3.3k ohms for linearization at the midpoint temperature (50°C), which aligns with this requirement.

5. NTC linearization being absolutely necessary and fatal if false: The summary explicitly states the use of a parallel resistor for linearization and provides details about the expected voltage at 50°C, which confirms that NTC linearization is integral to the system design and is considered.

6. Rough architecture of sensor, linearization stage, amplification, optional filtering, and measurement: The summary outlines a system that includes a sensor (NTC thermistor), a linearization circuit, an amplification stage, and output voltage measurement. Though an explicit filtering stage is not mentioned, the architecture largely follows the required sequence.

7. Use of an NTC sensor: The selected sensor is explicitly stated as the Vishay NTCLE100E3 thermistor, which is an NTC, satisfying this requirement.

8. Self-heating effect taken into account and maximum current through the NTC known: While the summary mentions that self-heating is less than 1%, it does not provide the maximum current that passes through the NTC. Therefore, this requirement is not fully met.

Overall, the project summary meets 7 out of the 8 listed requirements. The only requirement not fully met is the specification of the maximum current through the NTC to account for the self-heating effect (requirement 8).