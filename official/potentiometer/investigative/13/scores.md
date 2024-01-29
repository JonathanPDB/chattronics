Score: 7
Explanations: 
1. The potentiometer is used as a voltage divider: The summary indicates the use of a linear rotary potentiometer powered by a -10 to +10 V source, which implies it is being used as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: The power source for the potentiometer is specifically mentioned to be -10 to +10 V.

3. The architecture is simple: The summary describes a straightforward architecture that includes a voltage divider, voltage buffer, anti-aliasing filter, and ADC, which is in line with the requirement.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The ADC is specified to have an input range that accommodates at least +/-7 V, which suggests that the output of the voltage divider (potentiometer) and buffer is centered around 0 V and within the required range.

5. The maximum voltage applied to the DAQ is 7V: This is explicitly stated in the ADC requirements, ensuring that the maximum voltage does not exceed 7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: An anti-alias filter with a recommended cutoff frequency of 150 Hz is mentioned, which would help avoid aliasing.

7. There is a filter removing frequencies between around 50 and 60 Hz: The summary calls for special attention to attenuate 50 and 60 Hz interference, suggesting the inclusion of a filter to remove these frequencies.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: The cutoff frequency is recommended to be 150 Hz, which is below 500 Hz. However, without knowing the filter's order or the exact type of filter (Butterworth or Chebyshev), we cannot precisely confirm that the gain at 500 Hz will be at least -20 dB. Generally, a second-order filter provides a roll-off of -12 dB/octave, which may not be enough to guarantee a -20 dB gain at 500 Hz. More details or calculations are needed to confirm this requirement.