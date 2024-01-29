Score: 7
Explanations: 
1. The project summary does not explicitly mention the use of a multimeter to measure the output, but it is implied that the output voltage is within a range that a multimeter could measure (0-20V).
2. There is a gain stage mentioned with a gain of 20, designed to scale an input of 0-1V to an output of 0-20V. However, the justification for the gain is based on an assumption of the input from the buffer stage, which is not explicitly provided.
3. The output voltage range of 0-20V is explicitly mentioned, which fits the requirement.
4. Linearization is achieved with a resistor in parallel with the NTC, optimized for the 50°C midrange, which meets the requirement.
5. The summary explicitly states that linearization is achieved, which is a critical requirement.
6. The architecture is roughly as required: NTC sensor, linearization stage, buffer (to avoid loading effects), amplification, and optional filtering before the measurement stage.
7. The sensor used is explicitly stated to be a Vishay NTCLE100E3 NTC thermistor.
8. The self-heating effect is acknowledged by the use of a buffer amplifier to avoid loading the thermistor network, but the maximum current through the NTC is not specified.

The given project summary successfully covers 7 out of the 8 listed requirements. The only requirement not explicitly met is the specification of the maximum current through the NTC thermistor (requirement 8).