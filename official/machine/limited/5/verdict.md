Verdict: unfeasible

Explanations: 
The project summary provides a comprehensive design for an analog instrumentation system for pressure and temperature monitoring. However, it does not explicitly address all of the requirements listed. Here is a breakdown of how the project aligns with the requirements:

1. Both sensors have d.c. output - The project summary does not explicitly state that both sensors have a d.c. output, which is a fatal flaw if false.
2. Both sensors must be amplified - The project includes amplifier arrays for both pressure and temperature sensors, fulfilling this requirement.
3. Wheatstone bridge and instrumentation amplifier for the pressure sensor - The project mentions using an instrumentation amplifier but does not specify the use of a Wheatstone bridge.
4. Use of an ADC - The project includes an ADC with a sampling rate of at least 2 kHz, meeting this requirement.
5. Infrared radiation sensors linearization - The project mentions analog computing linearization and diode linearization for temperature sensors, which satisfies this requirement.
6. Sampling order strategy - The project summary does not mention the sampling order strategy, which is necessary for a complete design.
7. ADC sampling frequency - The ADC has a sampling rate of at least 2 kHz, which is above the required 800 Hz.
8. Anti-aliasing filter's gain at half the sampling frequency - The project specifies a stopband attenuation of 60 dB by 800 Hz, which satisfies the requirement of at least -20 dB at half the sampling frequency.
9. Low-pass filter cutoff frequency - The anti-aliasing filter has a cutoff frequency of 450 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (given that the sampling frequency is at least 2 kHz).
10. Low-pass filters before the multiplexer - The project includes anti-aliasing filters, but it is not clear whether they are positioned before the multiplexer, as required.
11. Use of multiplexers - The project includes multiplexers with specifications that suggest they are used to choose channels, fulfilling this requirement.
12. Multiplexers are solid state - The suggested models, such as the ADG1606 and ADG1408, are solid-state multiplexers.

Overall, the project fails to explicitly confirm the use of d.c. output for both sensors and does not detail the sampling order strategy or the positioning of the anti-aliasing filters in relation to the multiplexer. Additionally, the use of a Wheatstone bridge for the pressure sensor is not specified. These shortcomings are critical to the project's alignment with the stated requirements.