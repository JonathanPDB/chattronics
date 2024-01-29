Score: 7
Explanations: 
The project summary covers the following requirements:

1. The output being measured by a multimeter is not explicitly stated, but it is implied that the system's output is in a voltage range that is compatible with multimeter measurement (0-20V after amplification and level shifting).

2. There is an amplification stage with a non-inverting amplifier configuration, and the gain is provided as 20, justified by the need to scale a voltage change of 1V to 20V.

3. The output voltage range after the amplification and level shifting is specified to be 0-20 Volts, which is within the required measurement range.

4. The NTC is linearized by a resistor optimized for the midrange 50ºC, as indicated by the selection of a parallel resistor close to the thermistor's resistance at that temperature.

5. Linearization of the NTC is mentioned and is a critical part of the design, as the resistor is chosen to linearize the response at the midpoint temperature of 50ºC.

6. The architecture consists of the sensor (NTC thermistor), a linearization stage (parallel resistor), amplification (non-inverting amplifier), optional filtering (anti-aliasing filter), and measurement, which aligns with the requirement.

7. The sensor used is explicitly named as an NTC thermistor (Vishay NTCLE100E3).

8. The self-heating effect is acknowledged, as the summary suggests the need for short-circuit protection and the use of a series output resistor (Rs), but it does not provide the maximum current that passes through the NTC, which is a requirement for considering the self-heating effect.

The summary does not provide detailed information about the actual maximum current through the NTC to evaluate the self-heating effect fully. Therefore, requirement 8 is not met.

The project summary does not cover the following requirement:

- The maximum current that passes through the NTC is not known, which is necessary to consider the self-heating effect.

Therefore, the score is 7 out of 8.