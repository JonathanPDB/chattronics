Score: 6
Explanations: 
The project summary provided covers several of the requirements:

1. The potentiometer is used as a voltage divider: This is implicitly covered because potentiometers are typically used as voltage dividers, and the sensitivity calculation (1.11 mV/degree) suggests that it's being used in this way.

2. The voltage applied to the potentiometer is +/- 10 V: The voltage rating of the suggested potentiometer is greater than 10 V (25 V), so it can handle a +/- 10 V supply.

3. The architecture includes a voltage divider, an anti-aliasing filter, and then the DAQ measures the voltage: The summary mentions a potentiometer (voltage divider), a buffer, gain blocks (attenuation and level shifting), and an anti-aliasing filter before the DAQ measures the voltage.

4. The input voltage of the DAQ is centered at 0, for example, +/- 7V: The summary does not explicitly mention the final output voltage range being centered at 0V or being +/- 7V. The level shifter suggests that there is an intention to shift the voltage, but the exact range is not specified.

5. The maximum voltage applied to the DAQ is 7V: The summary does not specify the maximum voltage applied to the DAQ. The gain stages and level shifter suggest that the voltage is being adjusted, but without specific values, we cannot confirm compliance with this requirement.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The summary mentions an active Sallen-Key Low-Pass Filter, which serves as an anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz: The summary mentions the possibility of adding a Notch Filter at 50/60 Hz if required, which would meet this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The cutoff frequency of the anti-aliasing filter is 250 Hz, which is a 2nd order filter. Without the exact component values or the transfer function, we cannot calculate the gain at 500 Hz to confirm if it meets the -20 dB requirement. However, a 2nd order filter typically has a roll-off of -12 dB/octave, so it is plausible that the filter could meet this requirement.

Given the information provided, it is not possible to confirm whether requirements 4, 5, and 8 are fully met due to a lack of specific details. Requirement 4 is partially met with the presence of a level shifter, but the exact centered voltage range is not specified. Requirement 5 is not confirmed without knowing the final output voltage range. Requirement 8 is plausible but not confirmed without the filter's transfer function or component values.