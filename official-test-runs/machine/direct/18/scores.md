Score: 11
Explanations: 
1. The sensors have d.c. output as no demodulation process is described. This is implicit for the pressure sensors (strain-gauge) and explicitly stated for the temperature sensors (infrared radiation detectors).
2. The amplification requirement is met for both sensor arrays. The pressure sensors are amplified by a three-op-amp instrumentation amplifier, and the temperature sensors are amplified by non-inverting amplifiers.
3. While the pressure sensors are said to be amplified by an instrumentation amplifier, the use of a Wheatstone bridge is not explicitly mentioned, so this requirement is not verifiably met.
4. An ADC is used in the system, with a 16-bit resolution and a sampling rate of at least 2 kHz per channel.
5. The temperature sensors are being linearized, as indicated by the mention of a microcontroller with a look-up table or polynomial correction, which fulfills the requirement for linearization.
6. The sampling order strategy is not mentioned, so this requirement is not met.
7. The ADC has a sampling rate of at least 2 kHz per channel, which exceeds the minimum requirement of 800 Hz.
8. The anti-aliasing filter is a 4th-order active Butterworth filter with a cutoff frequency of 500 Hz. Given that the sampling rate is at least 2 kHz, the filter would provide at least -24 dB attenuation at half the sampling frequency (1 kHz), exceeding the -20 dB requirement.
9. The cutoff frequency of the anti-aliasing filter is 500 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (at least 1 kHz), satisfying this requirement.
10. The anti-aliasing filter is mentioned, and it is positioned before the multiplexer, which meets the requirement for reducing aliasing.
11. A 16-channel analog multiplexer is used to select among the sensor channels, meeting this requirement.
12. The multiplexer mentioned is an integrated circuit (IC), which implies it is solid-state, thus fulfilling this requirement.

The requirements that are not met are:
- The use of a Wheatstone bridge in conjunction with the pressure sensor (requirement 3).
- The description of the sampling order strategy (requirement 6).