Verdict: acceptable

Explanations: 
The project summary provides a comprehensive design for an analog instrumentation system to monitor pressure and surface temperature. 

1. Both sensors are presumed to have a d.c. output, as the pressure sensor is stated to give a millivolt output per psi, and the temperature sensor (Melexis MLX90614ESF-BAA-000-TU-ND) has a digital output, which is not the same as a d.c. output but is acceptable since it doesn't require demodulation.
2. Amplification is addressed for both the pressure sensor, with the use of an instrumentation amplifier, and the temperature sensor signals, which are linearized using precision resistors and op-amps.
3. The pressure sensor's inclusion in a Wheatstone bridge is not explicitly mentioned, but the use of an instrumentation amplifier suggests that it could be part of such a configuration.
4. An ADC is included in the design, with specifications that suggest it is appropriate for the application.
5. The temperature sensor's nonlinearity is addressed through piecewise-linear approximation, meeting the requirement for linearization.
6. The summary does not explicitly mention the sampling order strategy (sequential, simultaneous, etc.).
7. The ADC's sampling rate is specified to be at least 2 kHz per channel, which exceeds the minimum requirement of 800 Hz.
8. The anti-aliasing filter has a cutoff frequency of 400 Hz and provides 60 dB of attenuation starting at 800 Hz, which is sufficient to meet the requirement for the gain of the signal to be at least -20 dB at half the sampling frequency.
9. The low-pass cutoff frequency meets the requirement of being higher than 400 Hz and lower than half the total sampling frequency.
10. Low-pass filters are included in the design and are specified to be positioned before the multiplexer.
11. The use of a multiplexer is included in the design, meeting the requirement.
12. The multiplexer suggested (ADG1606) is a solid-state device.

However, despite meeting most requirements, the temperature sensor choice has a digital output, which is not aligned with the requirement for both sensors to have d.c. outputs. Additionally, the lack of explicit mention of the sampling order strategy is a minor omission that could be clarified with additional information. 

Given that the temperature sensor output is not d.c. and the sampling order strategy is not specified, the project cannot be considered optimal. However, these issues are not fatal to the project's implementation, as the digital output of the temperature sensor does not require demodulation and can be easily integrated into the system. The lack of a specified sampling order strategy is a detail that can be resolved in the design phase.