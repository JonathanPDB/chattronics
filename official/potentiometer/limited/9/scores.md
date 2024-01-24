Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - The summary implies this as it mentions the potentiometer is used to measure the pendulum's angle, which is a common application of potentiometers as voltage dividers.
2. The voltage applied to the potentiometer is +/- 10 V - This is explicitly stated in the project description.
3. The architecture is simple with a voltage divider, an anti-aliasing filter, and then the DAQ measures the voltage - The overall architecture follows this simplicity.
4. The input voltage of the DAQ is centered at 0, +/- 7V - The offset adjustment and amplifier blocks ensure the signal fits the DAQ's input voltage range.
5. The maximum voltage applied to the DAQ is 7V - The gain setting of the differential amplifier (7/10) and the non-inverting amplifier gain (slightly less than 1.4) ensure that the output does not exceed +/- 7V.
6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - An active low-pass Sallen-Key filter with a cut-off frequency of 150 Hz is used for anti-aliasing.
7. There is a filter removing frequencies between around 50 and 60 Hz - Active band-stop (notch) filters at both 50 Hz and 60 Hz are used to remove power line noise.
8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - This is not explicitly stated. While there is mention of attenuation at the center frequencies of the notch filters, there is no information on the gain at 500 Hz. Without further details, we cannot confirm that the gain at 500 Hz meets the -20 dB requirement.

The requirements that are not explicitly covered are:
8. The gain at 500 Hz has not been specified, so we cannot confirm if it is at least -20 dB.

Given the information provided, we can confirm 7 out of 8 requirements are met.