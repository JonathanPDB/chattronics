Score: 7
Explanations: 
The project summary covers the following requirements:

1. Both sensors have d.c. output: The piezoresistive strain gauge pressure transducer and the MLX90614 IR sensor both provide d.c. outputs.
2. Both sensors must be amplified: The pressure sensor uses an instrumentation amplifier with a gain of 4500 to 5000x, and the temperature sensor is amplified with an op-amp.
3. The pressure sensor is inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The use of a piezoresistive strain gauge implies a Wheatstone bridge configuration, and it is mentioned that an instrumentation amplifier is used.
4. An ADC should be used: Both the pressure and temperature sensor arrays have dedicated ADCs.
5. Infrared radiation sensors are being linearized: Although the method of linearization is not explicitly mentioned, the use of a high-precision op-amp for the IR sensor suggests that linearization is being considered.
6. The solution does not mention the sampling order strategy explicitly.
7. The sampling frequency of the ADC for the pressure sensor is at least 2 kHz, which satisfies the requirement of not being less than 800 Hz. The temperature ADC sampling rate is 1 kHz to 10 kHz, so it also satisfies this requirement.
8. The anti-aliasing filter requirement is not explicitly stated in terms of gain at half the sampling frequency.
9. The low-pass filter for the pressure sensor has a cutoff frequency of 500 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (1 kHz), but for the temperature sensor, the cutoff frequency is 400 Hz, which is not explicitly stated as lower than half the sampling frequency.
10. There are low-pass filters mentioned for both sensors, but their positioning relative to the multiplexer is not stated.
11. The project uses multiplexers: This is not explicitly mentioned in the summary.
12. The multiplexer(s) are solid state: This is not explicitly mentioned in the summary.

From the summary provided, the project seems to implicitly satisfy several requirements but does not explicitly address some key details such as the sampling order strategy, the specifics of the anti-aliasing filters, and the use and type of multiplexers. Hence, the score reflects the requirements that are either explicitly met or can be reasonably inferred from the information given.