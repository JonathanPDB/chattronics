Score: 8
Explanations: 
1. The potentiometer is used as a voltage divider: The summary specifies the use of a Bourns 3590 Precision Potentiometer, which would typically be used as a voltage divider in this context, so this requirement is met.

2. The voltage applied to the potentiometer is +/- 10 V: The power supply requirements mention a stable -10V to +10V supply for the potentiometer, fulfilling this requirement.

3. The architecture is simple: The architecture includes a voltage divider, offset adjustment, amplification blocks, low-pass filters, and an ADC within the DAQ. Although it includes additional blocks for offset adjustment and amplification, the basic requirement of having a voltage divider, anti-aliasing filter, and DAQ is met.

4. The input voltage of the DAQ is centered at 0, for example, +/- 7V: The ADC block specifies an input voltage range of +/-7V, which meets this requirement.

5. The maximum voltage applied to the DAQ is 7V: The input voltage range of the ADC is given as +/-7V, ensuring the maximum voltage applied to the DAQ is 7V, thus meeting this requirement.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The system includes two low-pass filters (LowPassFilter1 and LowPassFilter2) and an AntiAliasingFilter block, all designed to avoid aliasing, satisfying this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz: The LowPassFilter1 is designed with a cutoff frequency of 40 Hz to attenuate 50 Hz interference, and the LowPassFilter2 with a cutoff frequency of 100 Hz to attenuate 60 Hz interference, meeting this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The filters mentioned are second-order, which have a roll-off rate of 12 dB/octave. However, since the cutoff frequency is 100 Hz for the highest filter, the attenuation at 500 Hz would be approximately 12 dB + 20*log10(500/100) = 12 dB + 20*log10(5) ≈ 12 dB + 14 dB = 26 dB, which exceeds the requirement of -20 dB at 500 Hz.

All eight requirements have been met based on the project summary provided.