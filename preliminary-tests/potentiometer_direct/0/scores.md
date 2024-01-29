Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is inserted into a voltage divider: Although the summary does not explicitly mention a voltage divider, the potentiometer's use suggests that it is part of a voltage divider configuration to provide a variable output voltage.

2. The voltage applied to the potentiometer is +/- 10 V: The power supply for the buffer stage, which is presumably connected to the potentiometer, is +/- 10 V, indicating that the requirement is met.

3. The architecture includes a voltage divider, an anti-aliasing filter, and then the DAQ measures the voltage: The summary specifies that there is a buffer stage, a filter design (notch filter), and an ADC with an anti-aliasing filter. The anti-aliasing filter is mentioned as part of the ADC, which implies it is in place before DAQ measurement.

4. The input voltage of the DAQ is centered in 0, +/- 7V: The ADC is said to be configurable to accept +/- 7V, which meets this requirement.

5. The maximum voltage applied to the DAQ is 7V: The non-inverting amplifier has a gain of 0.7 to map the potentiometer's maximum voltage of 10V down to the DAQ's maximum of 7V, satisfying this requirement.

6. There is a low-pass filter (or anti-aliasing filter) that avoids aliasing: The anti-aliasing filter is mentioned as part of the ADC specification, indicating this requirement is met.

7. There is a filter removing frequencies between around 50 and 60 Hz: The Active Twin-T Notch Filter is designed to attenuate 50 Hz and 60 Hz by at least 40 dB, which satisfies this requirement.

8. The order of the filter is or adds up to be at least 3: The summary does not provide explicit information about the order of the filters (low-pass and notch). The notch filter's complexity is not specified, and the anti-aliasing filter's order is not mentioned. Therefore, we cannot confirm that the combined order of the filters meets or exceeds 3.

The requirements that are not explicitly satisfied due to lack of information are:
- The order of the filters adding up to at least 3 (requirement 8).