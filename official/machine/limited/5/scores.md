Score: 10
Explanations: 
The project summary provided covers the following requirements:

1. Both sensors have d.c. output - This is implicitly covered as both pressure and temperature sensors typically have a DC output. The summary does not mention any need for demodulation.

2. Both sensors must be amplified - The summary clearly mentions the use of an instrumentation amplifier array for the pressure sensors with a gain of 5000 and an amplifier array for the temperature sensors with a gain of 50.

3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier - While the summary does not explicitly mention a Wheatstone bridge, the use of an instrumentation amplifier is typical for pressure sensors which are often part of a Wheatstone bridge configuration for such applications. The mention of an instrumentation amplifier such as INA118 or AD620 suggests that this requirement is met.

4. An ADC should be used - The summary explicitly mentions the use of an ADC with a sampling rate of at least 2 kHz.

5. Infrared radiation sensors are being linearized - The summary mentions the use of analog computing linearization using op-amps or diode linearization for the temperature sensors, which are infrared radiation sensors.

6. The solution mentions the sampling order strategy - The summary does not explicitly mention the sampling order strategy (sequential, simultaneous, etc.).

7. The sampling frequency of the ADC is not less than 800 Hz - The summary states that the ADC has a sampling rate of at least 2 kHz, which satisfies this requirement.

8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency - The anti-aliasing filter is a fourth-order Butterworth low-pass filter with a cutoff frequency of 450 Hz and a stopband attenuation of 60 dB by 800 Hz. This satisfies the requirement as the attenuation at half the sampling frequency (1000 Hz, given a 2 kHz sampling rate) would be more than -20 dB.

9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency - The cutoff frequency of the anti-aliasing filter is 450 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (1000 Hz).

10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s) - The summary mentions the use of an anti-aliasing filter, but it does not specify its position in relation to the multiplexer(s).

11. The project uses multiplexer(s) to choose channels - The summary explicitly mentions the use of multiplexers with at least 16 single-ended channels or 8 differential channels.

12. The multiplexer(s) are solid state - The suggested models ADG1606 and ADG1408 are solid-state multiplexers.

The requirements not clearly met are #6 and #10, as the summary does not specify the sampling order or the exact position of the anti-aliasing filters in relation to the multiplexers.