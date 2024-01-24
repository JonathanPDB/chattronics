Score: 9
Explanations: 
1. Both sensors have d.c. output: The pressure sensor provides a d.c. output as it's a piezoresistive strain-gauge type. The temperature sensor's output is digital, which also implies a d.c. characteristic. Therefore, no demodulation is used. (Requirement met)
2. Both sensors must be amplified: Both the pressure and temperature sensors have amplification blocks with specified gains. (Requirement met)
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The amplification block for the pressure sensor uses an instrumentation amplifier, which is typical for Wheatstone bridge applications. (Requirement met)
4. An ADC should be used: A high-resolution ADC (12-bit or higher) is mentioned for use with the temperature sensor, and it's implied that the pressure sensor would also be connected to an ADC. (Requirement met)
5. Infrared radiation sensors are being linearized: The linearization of the infrared temperature sensor is done digitally using a microcontroller. (Requirement met)
6. The solution mentions the sampling order strategy: The summary does not explicitly mention the sampling order strategy (sequential, simultaneous, etc.). (Requirement not met)
7. The sampling frequency of the ADC is not less than 800 Hz: The ADC sampling rate is at least 2 kHz per channel, which is above the 800 Hz minimum requirement. (Requirement met)
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency: The cutoff frequency for the filters is 450 Hz, but there is no explicit mention of the gain at half the sampling frequency (1 kHz). Without more information, we cannot confirm this requirement is met. (Requirement not met)
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency is set at 450 Hz which is higher than 400 Hz and lower than half the total sampling frequency (1 kHz). (Requirement met)
10. There are low-pass filters positioned before the multiplexer(s): Both sensor blocks include low-pass filters, but it is not explicitly stated that these are positioned before the multiplexers. (Requirement not met)
11. The project uses multiplexer(s): An 8-to-1 analog multiplexer is mentioned. (Requirement met)
12. The multiplexer(s) are solid state: The multiplexers mentioned (ADG1608 or MAX4051A) are solid state. (Requirement met)