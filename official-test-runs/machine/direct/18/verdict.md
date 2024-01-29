Verdict: acceptable

Explanations: 
The project summary provides a comprehensive overview of a pressure and temperature monitoring system design, including details on the sensor arrays, signal conditioning, multiplexing, data conversion, and processing. The following points address the requirements given:

1. Both sensors have d.c. output: The pressure sensors are strain-gauge based and the temperature sensors are infrared radiation detectors, both of which typically provide d.c. outputs. This requirement is met.
2. Both sensors must be amplified: The design includes an instrumentation amplifier array for the pressure sensors and a non-inverting amplifier array for the temperature sensors. This requirement is met.
3. The pressure sensor must be inserted in a wheatstone bridge and amplified by an instrumentation amplifier: The summary specifies the use of a three-op-amp instrumentation amplifier for each pressure sensor, which implies the use of a Wheatstone bridge configuration. This requirement is met.
4. An ADC should be used: A 16-bit SAR ADC with a sampling rate of at least 2 kHz per channel is specified. This requirement is met.
5. Infrared radiation sensors are being linearized: The design mentions using a microcontroller with a look-up table or polynomial correction for non-linear correction, which addresses linearization. However, it does not mention the use of diode networks or log amplifiers, but since digital linearization is allowed, this requirement is met.
6. The solution mentions the sampling order strategy: The summary does not explicitly mention the sampling order strategy (sequential, simultaneous, etc.). This requirement is not met.
7. The sampling frequency of the ADC is not less than 800 Hz: The ADC has a specified sampling rate of at least 2 kHz per channel, exceeding the minimum requirement. This requirement is met.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency: The summary specifies a 4th-order Butterworth filter with a cutoff frequency of 500 Hz, which would provide at least -24 dB attenuation at 1 kHz (half of 2 kHz sampling frequency). This requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency of the anti-aliasing filter is 500 Hz, which is higher than 400 Hz and less than 1 kHz (half the total sampling frequency of 2 kHz). This requirement is met.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s): The design includes a 4th-order active Butterworth filter, but it is not explicitly stated that this filter is positioned before the multiplexer. This requirement is not clearly met.
11. The project uses multiplexer(s) to choose channels: A 16-channel analog multiplexer/demultiplexer IC is specified. This requirement is met.
12. The multiplexer(s) are solid state: The CD74HC4067 is a solid-state multiplexer. This requirement is met.

Due to the lack of explicit mention of the sampling order strategy and the unclear positioning of the anti-aliasing filters with respect to the multiplexers, the project falls short of the "optimal" category. However, since these issues are not fatal and can potentially be addressed with minor clarifications or design adjustments, the project is categorized as "acceptable".