Verdict: incorrect

Explanations: 
The project summary provided outlines an analog instrumentation system for monitoring pressure and temperature. However, there are several issues with the provided information:

1. Both sensors have d.c. output: This requirement is assumed to be met since the sensors mentioned have a d.c. output.
2. Both sensors must be amplified: The pressure sensor is amplified by an instrumentation amplifier, and the temperature sensor output is scaled, which implies amplification.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The summary does not specifically mention the use of a Wheatstone bridge, which is a fatal flaw if omitted.
4. An ADC should be used: The summary mentions a 12-bit ADC with a sampling rate of at least 2 kSPS per channel, which meets this requirement.
5. Infrared radiation sensors are being linearized: The summary mentions analog linearization for the temperature sensor, which meets this requirement.
6. The sampling order strategy is not mentioned, which could be an important detail for sequential or simultaneous sampling considerations.
7. The sampling frequency of the ADC is not less than 800 Hz: The summary states a sampling rate of at least 2 kSPS per channel, which satisfies this requirement.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency: There is no mention of the gain at half the sampling frequency, which is a critical detail for anti-aliasing filter performance.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequencies for the filters are mentioned as 400 Hz and 500 Hz, which are higher than 400 Hz but without knowing the total sampling frequency, it's unclear if they are lower than half the total sampling frequency.
10. There are low-pass filters before the multiplexer(s): The summary does mention low-pass filters, but it does not specify their position relative to the multiplexer(s), which is crucial to reduce aliasing.
11. The project uses multiplexers to choose channels: The summary mentions the use of multiplexers for both the pressure and temperature arrays, which meets this requirement.
12. The multiplexers are solid-state: Both mentioned multiplexers (CD4051B and ADG708) are solid-state devices, which satisfies this requirement.

Given the missing details and the critical omission of the Wheatstone bridge, the project cannot be considered fully compliant with the requirements provided. The absence of the Wheatstone bridge for the pressure sensor is a fatal flaw, and the lack of information on the anti-aliasing filter's performance at half the sampling frequency is another significant oversight.