Score: 7
Explanations: 
1. The project summary indicates that strain-gauge-based pressure sensors and infrared radiation detectors are used, which typically have a d.c. output. However, there is no explicit mention of avoiding demodulation, so this requirement is implicitly met.
2. Amplification is clearly mentioned for both sensors, with a gain of 5000 for pressure sensors and a unity gain for temperature sensors.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. While amplification is mentioned, there is no specific mention of a Wheatstone bridge or the use of an instrumentation amplifier.
4. An ADC is specified with a resolution of 16-bit or higher and a sampling rate of at least 800 samples per second per channel.
5. Linearization for infrared radiation sensors is mentioned to be done digitally, which satisfies the requirement.
6. The sampling order strategy is not explicitly mentioned, such as whether it will be sequential or simultaneous.
7. The sampling frequency of the ADC is stated to be at least 800 Hz per channel, which meets the requirement.
8. The anti-aliasing filter is mentioned to be a 4th order Butterworth filter with a cutoff frequency of 400 Hz, but there is no information given about the gain at half the sampling frequency to evaluate if it meets the -20 dB requirement.
9. The cutoff frequency for the low pass filter is given as 400 Hz, which is higher than 400 Hz as required, but since the total sampling frequency is not provided, it cannot be confirmed if it's lower than half the total sampling frequency.
10. Low-pass filters are mentioned and are positioned before the multiplexer, which implies their use as anti-aliasing filters.
11. The use of an 8-to-1 multiplexer is clearly mentioned, satisfying this requirement.
12. The multiplexer is specified as a CMOS-based analog multiplexer, which is a solid-state device.

Based on the above, the project summary successfully covers requirements 1 (implicitly), 2, 4, 5, 7, 10, and 11. Requirements 3, 6, 8, 9, and 12 are not clearly met based on the information provided.