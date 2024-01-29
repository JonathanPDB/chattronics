Score: 7
Explanations: 
1. The summary does not explicitly state that the sensors have d.c. output, but the pressure sensor's sensitivity is given in mV/V/psi, which typically implies a DC output. It's not clear for the temperature sensor, but since a digital output sensor (Melexis MLX90614ESF-BAA) is recommended, it doesn't fit the DC output requirement.
2. Amplification is mentioned for both sensor types: a gain of 5000 is needed for the pressure sensor, and the temperature sensor signals are linearized using op-amps, which implies amplification.
3. The pressure sensor is to be amplified by an instrumentation amplifier, but there is no mention of it being inserted in a Wheatstone bridge configuration.
4. An ADC is mentioned, specifically a SAR ADC with a sampling rate of at least 2 kHz per channel.
5. Nonlinear correction for the temperature sensors is addressed using piecewise-linear approximation with precision resistors and op-amps, which satisfies the requirement for linearization.
6. The summary does not mention the sampling order strategy (sequential, simultaneous, etc.).
7. The ADC sampling frequency is stated to be at least 2 kHz per channel, which is above the minimum requirement of 800 Hz.
8. The anti-aliasing filter is described as a 4th order Butterworth low-pass filter with a cutoff frequency of 400 Hz and 60 dB of attenuation starting at 800 Hz, which implies at least -20 dB at half the sampling frequency of 1 kHz (given the sampling frequency is at least 2 kHz).
9. The cutoff frequency for the low-pass filter is exactly 400 Hz, which is higher than the required 400 Hz but not explicitly stated to be lower than half the total sampling frequency.
10. There is mention of anti-aliasing filters, which are positioned before the ADC, but it is not explicitly stated that they are positioned before the multiplexer(s).
11. The project uses a multiplexer, specifically a 16-channel ADG1606.
12. The multiplexer model ADG1606 is a solid-state device.

Overall, the project summary covers most of the requirements, but there are some areas where the information is not explicitly stated or does not meet the requirements as stated.