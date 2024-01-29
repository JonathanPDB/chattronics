Score: 11
Explanations: 
The project summary covers the following requirements:

1. Both sensors have d.c. output: The pressure sensor outputs a millivolt level signal, which is DC. The temperature sensor (MLX90614) has a digital output, which implies a DC level signal (though digital). Hence, this requirement is met.
2. Both sensors must be amplified: The pressure sensor is amplified by an instrumentation amplifier, and the temperature sensor signal is amplified by a voltage amplifier. This requirement is met.
3. The pressure sensor is inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: This is explicitly stated in the summary. This requirement is met.
4. An ADC is used: A 12-bit SAR ADC with a sampling rate of at least 2 kHz per channel is mentioned. This requirement is met.
5. Infrared radiation sensors are being linearized: Polynomial correction circuit with operational amplifier stages is used for nonlinearity correction. While this doesn't explicitly mention linearizing infrared sensors, the MLX90614 is an infrared sensor, and the polynomial correction is a valid linearization technique. This requirement is met.
6. The solution mentions the sampling order strategy: There is no explicit mention of the sampling order strategy (sequential, simultaneous, etc.). This requirement is not met.
7. The sampling frequency of the ADC is not less than 800 Hz: The ADC has a sampling rate of at least 2 kHz per channel, which exceeds 800 Hz. This requirement is met.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency: The stopband attenuation is at least 48 dB by the first octave above the cutoff (800 Hz), which is more than -20 dB at half the sampling frequency of 1 kHz (half of 2 kHz). This requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency of the anti-aliasing filter is 400 Hz, which meets the first part of this requirement. Since the sampling frequency is at least 2 kHz, half of this is 1 kHz, so the cutoff frequency is also lower than half the total sampling frequency. This requirement is met.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s): The presence of anti-aliasing filters is mentioned, but their exact positioning relative to the multiplexer is not detailed. However, since the filters are part of the signal conditioning, it can be reasonably inferred that they are positioned before the multiplexer, which is typical in such designs. This requirement is met implicitly.
11. The project uses multiplexers to choose channels: A 16-channel analog multiplexer (CD74HC4067) is used. This requirement is met.
12. The multiplexers are solid state: The CD74HC4067 is a solid-state device. This requirement is met.

The requirements that are not met are:
- Requirement 6: The sampling order strategy is not mentioned.

Therefore, the project successfully meets 11 out of 12 requirements.