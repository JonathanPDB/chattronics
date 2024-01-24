Verdict: acceptable

Explanations: 
The project summary provided details a comprehensive analog instrumentation system for pressure and temperature monitoring. However, it does not meet all the specified requirements necessary for a favorable verdict.

1. Both sensors are implied to have d.c. output as no demodulation is mentioned.
2. Both sensors are amplified - the pressure sensor through an instrumentation amplifier with a gain of 20,000 V/V, and the temperature sensor with a gain of approximately unity.
3. The pressure sensor is included in a full bridge configuration and amplified by an instrumentation amplifier, which is consistent with the wheatstone bridge requirement.
4. AD converters are used for both pressure (ADC_Pressure) and temperature (ADC_Temp) sensors.
5. The method for linearizing the infrared radiation sensors is not explicitly mentioned, which is a requirement.
6. The sampling order strategy is not defined in the summary.
7. The sampling frequency for the pressure ADC is specified as a minimum of 1 kSPS per channel, which is above the required 800 Hz. The temperature ADC sampling rate is not clearly defined as per the channel but suggests a total rate that is likely above 800 Hz.
8. There is no mention of the anti-aliasing filter's attenuation at half the sampling frequency, thus this requirement is not clearly met.
9. The low pass filter cutoff frequencies for pressure (350 Hz) and temperature (450 Hz) are within the required range.
10. Low-pass filters are mentioned, but their positioning regarding the multiplexer is not specified.
11. Multiplexers are used to select channels.
12. The multiplexers mentioned are solid state.

The project is missing critical information regarding the linearization of the infrared sensors, the anti-aliasing filter specifications, the exact sampling frequency for the temperature sensor, the sampling order strategy, and the positioning of the low-pass filters relative to the multiplexers. These omissions and ambiguities lead to a conclusion that the project cannot be categorized as "optimal."