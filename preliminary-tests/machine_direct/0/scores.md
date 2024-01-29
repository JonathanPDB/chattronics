Score: 13
Explanations: 
The project summary covers the following requirements:

1. Both sensors have d.c. output - Covered implicitly as strain-gauge and infrared sensors typically provide d.c. outputs.
2. Both sensors must be amplified - Covered, the strain-gauge is amplified by an instrumentation amplifier and the infrared sensor by a logarithmic amplifier.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier - Covered, the strain-gauge sensor is a full-bridge (Wheatstone bridge) and is amplified by an instrumentation amplifier.
4. An ADC should be used - Covered, a 12-bit SAR ADC is specified.
5. Infrared sensors are being linearized - Covered, the use of a logarithmic amplifier for the infrared sensor indicates linearization.
6. The solution mentions the sampling order strategy - Covered, it is mentioned that the ADC samples all 16 channels within a 1-second interval, implying sequential sampling.
7. The ADC samples one sensor per time - Covered, as it is implied by the use of a 16-channel multiplexer and the statement about sampling all channels within a 1-second interval.
8. The sampling frequency of the ADC is not less than 800 Hz - Covered, it is stated that the ADC is capable of at least 32 kSPS total throughput, which exceeds the minimum requirement.
9. The sampling frequency of the ADC is higher than 2000 Hz - Covered, with a 32 kSPS and 16 channels, the sampling rate per channel is 2000 Hz, meeting this requirement.
10. The ADC should be a SAR - Covered, a 12-bit SAR ADC is specified.
11. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency - Not covered, the cutoff frequencies for the low-pass filters are specified as 350 Hz and 450 Hz, which do not meet the minimum 400 Hz requirement.
12. There are low-pass filters before the multiplexers - Covered, the inclusion of Sallen-Key Butterworth filters in both subsystems implies they are before the multiplexers.
13. The project uses multiplexers to choose channels - Covered, a 16-channel analog multiplexer is specified.
14. The multiplexers are solid state - Covered implicitly as CD74HC4067 is a solid-state device.

The requirements not met are:
- The low pass cutoff frequency is higher than 400 Hz (requirement 11 is not met because one of the filters has a cutoff frequency of 350 Hz).