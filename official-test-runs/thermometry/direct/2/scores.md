Score: 7
Explanations: 
1. The output is being measured by a multimeter: The project mentions the compatibility with a standard multimeter and the design of the output voltage range from 0 to 20 Volts, which implies that the output is indeed measured by a multimeter. (Requirement met)

2. There should be an amplification stage and the gain must be provided and justified: The project describes a Gain Stage with a non-inverting operational amplifier configuration and provides a gain calculation with an adjustable gain between 1 and 20, justified by the need to scale the maximum input voltage to a 20V output. (Requirement met)

3. The output voltage range to be measured by the multimeter is between 0 and 20 Volts: This is explicitly stated in the project summary. (Requirement met)

4. The NTC is linearized by a resistor optimized for the midrange 50 ºC: The project specifies a linearization resistor (R_lin) with a value close to the calculated one for linearity at the 50°C midpoint. (Requirement met)

5. The NTC is linearized: The project mentions the sensitivity of the NTC being variable due to its nature and being linearized at the 50°C midpoint, which shows that linearization is considered and implemented. (Requirement met)

6. The architecture should be roughly consisting of the sensor, a linearization stage, amplification, optional filtering, and measurement: The architecture described includes the sensor (NTC thermistor), a linearization stage (linearization resistor), amplification (buffer amplifier, gain stage, and output amplifier), and measurement (multimeter). Optional filtering is mentioned in the form of a filter capacitor for noise suppression. (Requirement met)

7. The sensor used is an NTC: The project uses an NTCLE100E3103JB0, which is an NTC thermistor. (Requirement met)

8. The self-heating effect is taken into account and the maximum current that passes through the NTC is known: Although the self-heating effect is mentioned by selecting an excitation voltage level between 1V and 5V to minimize it, the exact maximum current passing through the NTC is not provided. Therefore, this requirement is not fully met as the maximum current is not explicitly known.