Score: 6
Explanations: 
The project summary covers the following requirements:

1. The output is measured by a multimeter: The project description indicates an output voltage range of 0 to 20 Volts, suitable for reading by a multimeter. (Requirement 3)

2. Amplification stage: The project includes a buffer amplifier and an output scaling amplifier, indicating an amplification stage. However, the gain for the buffer amplifier is mentioned to be unity (1), and the output scaling amplifier has a hypothetical gain of 20, which is not justified in the text. (Requirement 2)

3. Output voltage range: The summary specifies a 0-20V output range, which meets the requirement for the multimeter measurement range. (Requirement 3)

4. NTC linearization: The project uses a parallel resistor (Rp) to linearize the response of the NTC thermistor around the midpoint of the temperature range, optimized for 50ºC. (Requirement 4)

5. Linearization necessity: The linearization of the NTC is explicitly mentioned, and it is clear that this is a critical component of the design. (Requirement 5)

6. Architecture: The project follows the required architecture of sensor, linearization stage, amplification, optional filtering, and measurement. (Requirement 6)

7. Sensor used is an NTC: The project specifically mentions the use of a Vishay NTCLE100E3 NTC thermistor. (Requirement 7)

8. Self-heating effect: The self-heating effect is taken into account by limiting the current through the NTC using a high-value series resistor or constant-current source, although the maximum current is not specified. (Requirement 8)

The requirements not covered in the summary are:

- Justification of the gain of the amplification stage: The gain for the output scaling amplifier is mentioned as a hypothetical value of 20, but there is no justification provided for this value. (Requirement 2)

- Maximum current through NTC: While the self-heating effect is addressed, the exact maximum current is not specified in the summary. (Requirement 8)

As the project summary does not explicitly state the maximum current through the NTC, and the gain of the output scaling amplifier is not justified, requirements 2 and 8 are not fully satisfied. All other requirements are met according to the information provided.