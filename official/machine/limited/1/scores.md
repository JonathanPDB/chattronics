Score: 8
Explanations: 
1. The project summary does not specify whether the sensors produce a direct current (DC) output, so this requirement cannot be confirmed.
2. Both sensors are being amplified, as indicated by "Amplifier_Pressure" with a gain of 1000 and "Amplifier_Temp" with a gain of 40.
3. The pressure sensor is mentioned to be a strain-gauge type, which is typically used in a Wheatstone bridge configuration, and an instrumentation amplifier (INA114 or AD620) is specified.
4. An ADC with a sampling rate of at least 6400 Hz is included in the system design.
5. While it mentions that the infrared radiation sensors' nonlinear scale will be linearized, the summary does not specify the exact method (digitally, diode networks, or log amplifiers), so this requirement is not clearly met.
6. The summary indicates a sampling rate of 6400 Hz for 8 channels, implying sequential sampling, but does not explicitly mention the sampling order strategy.
7. The sampling frequency of the ADC is specified to be at least 6400 Hz, which exceeds the minimum requirement of 800 Hz.
8. The anti-aliasing filter is mentioned to be included before the ADC, but the summary does not provide the cutoff frequency or order necessary to determine if the gain of the signal is at least -20 dB at half the sampling frequency.
9. Both low-pass filters for pressure and temperature signals have a cutoff frequency of 400 Hz, which is above the required 400 Hz but does not specify if it is lower than half the total sampling frequency (3200 Hz in this case).
10. It mentions that there are low-pass filters before the ADC, serving as anti-aliasing filters.
11. The project uses multiplexers to choose channels, as indicated by the presence of MUX_Pressure, MUX_Temp, and MUX_Combined.
12. The multiplexers mentioned (DG408, CD4051B, ADG719) are all solid-state devices.