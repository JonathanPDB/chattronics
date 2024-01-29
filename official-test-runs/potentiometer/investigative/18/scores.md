Score: 6
Explanations: 
1. The potentiometer is used as a voltage divider: The signal conditioning circuit description indicates that a fixed resistor is used in series with the potentiometer to create a voltage divider, fulfilling this requirement.

2. The voltage applied to the potentiometer is +/- 10 V: The operating voltage of the potentiometer is up to 10V, which satisfies the requirement for a +/- 10V range.

3. The architecture includes a voltage divider, anti-aliasing filter, and DAQ: The system consists of the main blocks including a potentiometer (used as a voltage divider), buffer amplifier, anti-aliasing filter, and ADC (DAQ), meeting the requirement for a simple architecture.

4. The input voltage of the DAQ is centered at 0, +/- 7V: The signal conditioning circuit is designed to output a voltage signal that swings within the range of +/- 7V to match the DAQ's input, which meets this requirement.

5. The maximum voltage applied to the DAQ is 7V: The ADC input voltage range is specified as +/- 7V, ensuring that the maximum voltage applied to the DAQ does not exceed 7V.

6. There is a low pass filter (anti-aliasing filter) that avoids aliasing: The design includes a 2nd or 3rd order Butterworth low-pass filter with a cutoff frequency below the Nyquist frequency of 500 Hz, which will help avoid aliasing.

7. There is a filter removing frequencies between around 50 and 60 Hz: The project summary does not explicitly mention a filter designed to remove frequencies specifically between 50 and 60 Hz, so this requirement is not verifiably met.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The roll-off rate for a 2nd order filter is -12 dB/octave and for a 3rd order filter is -18 dB/octave. However, without the exact cutoff frequency value, it is not possible to confirm if the -20 dB at 500 Hz requirement is met. Typically, a 2nd order filter would need to have a cutoff frequency well below 500 Hz to achieve this attenuation. Since the recommended cutoff frequency is between 200 Hz and 250 Hz, it is likely that this requirement is met, but it is not explicitly confirmed.

Since the project summary does not explicitly confirm requirements 7 and 8, and there is uncertainty regarding the exact performance of the filter at 500 Hz, these requirements are not counted as met.