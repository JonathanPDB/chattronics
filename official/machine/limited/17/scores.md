Score: 11
Explanations: 
1. Requirement met: The pressure sensor has a d.c. output, and the temperature sensor, though digital, has an output that can be considered d.c. equivalent for this context (no demodulation needed for digital signals).
2. Requirement met: Both pressure and temperature sensors are connected to amplifier arrays to amplify their signals.
3. Requirement met implicitly: The pressure sensor's output is suitable for a Wheatstone bridge, and an instrumentation amplifier is suggested for the first amplification stage.
4. Requirement met: An ADC is used in the system.
5. Requirement met: Non-linearity correction for the temperature sensor is done using a microcontroller-based lookup table, which is a digital approach to linearization.
6. Requirement not met: There is no explicit mention of the sampling order strategy, whether it is sequential, simultaneous, etc.
7. Requirement met: The sampling rate of the ADC is 2 kHz per channel, which is above the minimum requirement of 800 Hz.
8. Requirement met: The anti-aliasing filter has a cutoff frequency of 500 Hz for pressure and 450 Hz for temperature signals. With a sampling rate of 2 kHz, the Nyquist frequency is 1 kHz. The filters, being second-order, provide at least -12 dB/octave roll-off. By 1 kHz (half the sampling frequency), the attenuation should be more than -20 dB, meeting the requirement.
9. Requirement met: The low-pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency (1 kHz), which is in accordance with the given filter cutoff frequencies of 500 Hz for pressure and 450 Hz for temperature signals.
10. Requirement met: Low-pass filters are used to reduce aliasing and are placed before the multiplexer.
11. Requirement met: The project uses a CD74HC4067 16-channel CMOS analog multiplexer.
12. Requirement met: The CD74HC4067 is a solid-state multiplexer.

The project summary covers 11 out of 12 requirements, with the only missing information being the sampling order strategy.