Score: 7
Explanations: 
1. Both sensors have d.c. output - The pressure sensor output is mentioned as a maximum of 1 mV, which is a DC value, and the temperature sensor is an IR detector with a digital output (PWM or SMBus), which implies a DC-equivalent signal for the purpose of interfacing with an ADC. Hence, both sensors meet the DC output requirement.
2. Both sensors must be amplified - The pressure sensor is connected to an instrumentation amplifier with a gain of 5000, and the temperature sensor is connected to a non-inverting operational amplifier with a gain of 50, fulfilling the amplification requirement.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier - The use of an instrumentation amplifier suggests a Wheatstone bridge configuration for the strain gauge pressure sensor, although the bridge itself is not explicitly mentioned.
4. An ADC should be used - The system includes an ADS7950 ADC with an integrated multiplexer.
5. Infrared radiation sensors are being linearized - The summary does not explicitly mention linearization of the IR sensor. The use of an operational amplifier with the IR sensor suggests some form of signal conditioning, but without explicit mention of linearization, this requirement is not verifiably met.
6. The solution mentions the sampling order strategy - There is no explicit mention of the sampling order strategy (sequential, simultaneous, etc.).
7. The sampling frequency of the ADC is not less than 800 Hz - The ADC sampling rate is 8 kHz overall, which means 1 kHz per channel, satisfying this requirement.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency - The low-pass filter is a 4th order Butterworth with a cutoff frequency of 500 Hz. Considering a 4th order filter has an attenuation rate of -24 dB/octave, it would have an attenuation of at least -48 dB at 1 kHz (half the sampling frequency of 2 kHz per channel), satisfying this requirement.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency - The cutoff frequency is 500 Hz, which is higher than 400 Hz and lower than 4 kHz (half of the 8 kHz total sampling frequency), satisfying this requirement.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, positioned before the multiplexer(s) - Although low-pass filters are mentioned, their placement in the signal chain relative to the multiplexer is not specified.
11. The project uses multiplexer(s) to choose channels - The system includes an ADG1606 8-to-1 analog multiplexer.
12. The multiplexer(s) are solid state - The ADG1606 is a solid-state device.

Based on the given information, requirements 1, 2, 4, 7, 8, 9, and 11 are met. Requirements 3, 5, 6, 10, and 12 are either not mentioned or not verifiably met based on the provided summary.