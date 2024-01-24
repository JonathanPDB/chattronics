Verdict: acceptable

Explanations: 
The provided summary meets several key requirements but lacks specific information on a few critical aspects. 

1. It is implied that the sensors have a DC output since no mention of demodulation is made.
2. Amplification of both sensors is planned with specified gains.
3. The pressure sensor is intended to be used with an instrumentation amplifier, but there is no explicit mention of it being part of a Wheatstone bridge configuration.
4. An ADC is mentioned, with a recommendation for at least 10-bit resolution.
5. Linearization of the temperature sensors is planned to be handled digitally, which is acceptable.
6. The sampling strategy is not explicitly mentioned (sequential, simultaneous, etc.).
7. The sampling frequency is not explicitly stated, only that the multiplexer supports a sampling rate of "at least 1 kSps per channel," which could imply that the requirement is met if the overall sampling frequency remains above 800 Hz after considering the number of channels.
8. The anti-aliasing filter is suggested to have a cutoff frequency at or slightly below 400 Hz, but there is no information regarding the attenuation at half the sampling frequency.
9. The cutoff frequency of the low-pass filter is at the lower bound of the acceptable range (400 Hz) but lacks details on whether it is less than half the total sampling frequency.
10. Low-pass filters are planned, but their positioning before the multiplexer is not confirmed.
11. An 8:1 multiplexer is included in the design, satisfying this requirement.
12. The multiplexer (ADG1608) is a solid-state device.

While the design outline is generally in line with the requirements, the lack of explicit confirmation on the Wheatstone bridge configuration for the pressure sensor, the sampling strategy, the exact sampling frequency, the anti-aliasing filter's attenuation at half the sampling frequency, and the positioning of the low-pass filters before the multiplexer prevents the project from being categorized as "optimal." The project can be considered "acceptable" if it is assumed that the missing details are addressed in the final implementation.