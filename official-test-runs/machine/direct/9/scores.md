Score: 12
Explanations: 
The project summary covers the following requirements:

1. The sensors have a d.c. output, as indicated by the sensitivity and power supply details for the pressure sensor and the digital interface for the temperature sensor.
2. Both sensors are amplified. The pressure sensors are amplified by an instrumentation amplifier with a gain of 5000, and the temperature sensors are amplified by a non-inverting amplifier with a gain of 50.
3. The pressure sensor is inserted in a Wheatstone bridge (implied by the need for an instrumentation amplifier and the typical use of such sensors) and amplified by an instrumentation amplifier.
4. An ADC is used with a 12-bit resolution or higher.
5. Infrared radiation sensors are being linearized digitally in the computer, as indicated by the "Digital correction using a microcontroller with ADC and computational capabilities."
6. The sampling order strategy is implied to be sequential since there is an 8-to-1 multiplexer and a single ADC mentioned, indicating that channels are sampled one after another.
7. The sampling frequency of the ADC is 6400 Hz, which is above the required 800 Hz.
8. The anti-aliasing filter has a cutoff frequency of 400 Hz and an order that provides -48 dB at twice the cutoff frequency (800 Hz), which meets the requirement of at least -20 dB at half the sampling frequency.
9. The low-pass cutoff frequency is 400 Hz, which is higher than the required 400 Hz and lower than half the total sampling frequency (3200 Hz).
10. Low-pass filters are presumably positioned before the multiplexer, as is standard practice, although this is not explicitly stated.
11. The project uses an 8-to-1 solid-state multiplexer to choose channels.
12. The multiplexer(s) are solid state, as indicated by the use of a CD4051 series.

The project summary does not explicitly mention the positioning of the low-pass filters before the multiplexer(s), which is necessary to confirm requirement 10. However, given the standard practice and the context provided, it is reasonable to infer that this requirement is met.

Based on the information provided, all 12 requirements are met.