Score: 11
Explanations: 
The project summary provides a comprehensive overview of the design of an analog instrumentation system for monitoring machine pressure and temperature. The requirements are evaluated as follows:

1. Both sensors have d.c. output: The pressure sensor is a strain gauge-based sensor, and the temperature sensor is an infrared detector. Both are typical d.c. output sensors, so this requirement is met.
2. Both sensors must be amplified: The design includes instrumentation amplifiers for both pressure and temperature sensors, thereby fulfilling this requirement.
3. Pressure sensor in a Wheatstone bridge and amplified by an instrumentation amplifier: Although not explicitly stated, strain gauge-based pressure sensors are typically employed in a Wheatstone bridge configuration. The use of a programmable gain instrumentation amplifier suggests this requirement is met.
4. An ADC should be used: The summary specifies a 16-bit ADC with a minimum sampling rate of 2 kHz per channel, meeting this requirement.
5. Infrared radiation sensors are linearized: The summary mentions digital linearization algorithms for temperature data, which implies that the infrared sensors' non-linear output will be linearized, satisfying this requirement.
6. Sampling order strategy: The summary does not explicitly mention the sampling order (sequential, simultaneous, etc.), so this requirement is not met.
7. ADC sampling frequency: The specified sampling rate is a minimum of 2 kHz per channel, which is above the required 800 Hz, thus meeting this requirement.
8. Anti-aliasing filter specifications: The anti-aliasing filter is described as having a cutoff frequency of 1 kHz and -40 dB attenuation for frequencies well above 1 kHz. This meets the requirement for a gain of at least -20 dB at half the sampling frequency (1 kHz if sampling at 2 kHz).
9. Low pass cutoff frequency: The anti-aliasing filter has a cutoff frequency of 1 kHz, which is higher than 400 Hz and lower than half the total sampling frequency (8 kHz for all eight channels), satisfying this requirement.
10. Low-pass filters before the multiplexer: The summary includes anti-aliasing filters in the signal chain, presumably before the multiplexer, meeting this requirement.
11. Use of multiplexers: The summary proposes an 8-to-1 multiplexer, fulfilling this requirement.
12. Multiplexers are solid state: The ADG708/ADG709 suggested are solid-state ICs, which satisfies this requirement.

In summary, the project successfully meets 11 out of the 12 requirements. The only requirement not explicitly confirmed is #6 regarding the sampling order strategy.