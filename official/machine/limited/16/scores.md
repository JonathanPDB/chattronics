Score: 10
Explanations: 
1. The pressure sensor output is 1 mV/V, which is typically a DC output, and the IR temperature sensor has a linear analog output, which implies a DC output as well. Therefore, no demodulation is needed, fulfilling this requirement.
2. Both sensor arrays have amplification stages mentioned: the pressure sensor array uses an instrumentation amplifier with a gain of 5000, and the temperature sensor array has a gain stage for scaling the output.
3. The pressure sensor is likely inserted in a Wheatstone bridge as it is a strain-gauge sensor, and it is mentioned to be amplified by an instrumentation amplifier, which is typical for a Wheatstone bridge output.
4. An ADC with a sampling rate of at least 2 kSPS per channel is specified, satisfying the requirement for an ADC.
5. The temperature sensor array includes an analog linearization stage using operational amplifiers, fulfilling the requirement for linearization.
6. The summary does not specify the sampling order strategy (sequential, simultaneous, etc.).
7. The specified sampling frequency of the ADC is at least 2 kSPS per channel, which is above the minimum requirement of 800 Hz.
8. The anti-aliasing filter for the pressure sensor array is specified with a cutoff frequency of 400 Hz. However, there's no mention of the gain of the signal at half the sampling frequency, thus this requirement is not verifiably met.
9. The low-pass cutoff frequency for the pressure sensor is mentioned as 400 Hz, which meets the requirement of being higher than 400 Hz. The temperature sensor array's cutoff frequency is 500 Hz, which is also higher than 400 Hz and lower than half the total sampling frequency (assuming the total sampling frequency is at least 2 kSPS).
10. Low-pass filters are mentioned for both sensor arrays, positioned before the multiplexer(s), to reduce aliasing.
11. Multiplexers (CD4051B and ADG708 or similar) are included to select channels for each sensor array, satisfying the requirement for multiplexers.
12. The multiplexers mentioned (CD4051B and ADG708) are solid-state devices.

The score is 10 because items 6 and 8 are not verifiably met based on the provided information.