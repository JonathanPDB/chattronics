Score: 11
Explanations: 
1. Both sensors (pressure and temperature) have DC outputs as the pressure sensor output is mV level from the strain gauge and the temperature sensor is an infrared radiation type, commonly with a DC output. No mention of demodulation indicates that the requirement is met.
2. Amplification is required and is covered by the instrumentation amplifier array with a gain of 5000, meeting the requirement for the pressure sensor, and by the implied need for signal conditioning for the temperature sensor to match the ADC’s voltage range.
3. The pressure sensor is a strain-gauge based transducer, which is typically used with a Wheatstone bridge, and amplified by an instrumentation amplifier, meeting the requirement.
4. An ADC with a sampling rate greater than 1 kHz per channel is used in the design, fulfilling this requirement.
5. Non-linear signal conditioning for the infrared radiation temperature sensors is mentioned, with options for both analog circuitry (using op-amps for linearization) and digital processing (via microcontroller), satisfying this requirement.
6. There is no explicit mention of the sampling order strategy (sequential or simultaneous), hence this requirement is not met.
7. The sampling frequency of the ADC is greater than 1 kHz per channel, which is above the minimum requirement of 800 Hz, so this requirement is met.
8. The anti-aliasing filter is a second-order low pass Butterworth filter with a cutoff frequency of 500 Hz. Since the sampling frequency is greater than 1 kHz, the gain at half the sampling frequency (500 Hz) would be in the stop band of the filter, providing at least -20 dB attenuation, which satisfies this requirement.
9. The low pass cutoff frequency of 500 Hz is higher than 400 Hz and lower than half the total sampling frequency (which is greater than 1 kHz), meeting this requirement.
10. Low-pass filters are mentioned to be part of the design, and they would be positioned before the multiplexer(s) as part of the signal conditioning, meeting this requirement.
11. An 8-to-1 analog multiplexer is used in the design, fulfilling this requirement.
12. The multiplexer mentioned is a solid-state device, such as the ADG408/ADG409 or CD74HC4051, satisfying this requirement.