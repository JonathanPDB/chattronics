Verdict: acceptable

Explanations: 
The project summary provided describes a comprehensive analog instrumentation system for monitoring pressure and surface temperature variations. However, to determine the appropriateness of the project, we must evaluate it against the specified requirements:

1. No indication of demodulation being used is given, which is in line with the requirement for DC output sensors.
2. Amplification of both sensors is planned, with detailed gain calculations provided.
3. The pressure sensor is planned to be used in conjunction with an instrumentation amplifier, but there is no explicit mention of a Wheatstone bridge configuration.
4. An ADC is included in the design, fulfilling this requirement.
5. The temperature sensors are infrared radiation detectors, but there is no explicit mention of linearization techniques such as digital linearization, diode networks, or log amplifiers.
6. The sampling order strategy is not mentioned, only the sampling rate.
7. The suggested sampling rate of at least 2 kHz per channel exceeds the minimum requirement of 800 Hz.
8. & 9. A low pass Butterworth filter with a cutoff frequency of 500 Hz is included, which satisfies the requirement of being higher than 400 Hz and less than half the total sampling frequency (assuming a minimum of 2 kHz per channel). The summary does not provide details on the gain of the signal at half the sampling frequency to confirm a -20 dB gain.
10. The low pass filter is mentioned, but its position relative to the multiplexer is not specified.
11. The project includes the use of a multiplexer.
12. The type of multiplexer is specified as solid-state.

Given the information provided, the project does not explicitly fulfill all the necessary requirements, particularly the Wheatstone bridge for the pressure sensor (requirement 3), the linearization of infrared sensors (requirement 5), the sampling order strategy (requirement 6), and the specific positioning of the low pass filter before the multiplexer (requirement 10). Moreover, the specification of the anti-aliasing filter's performance at half the sampling frequency (requirement 8) is missing. Therefore, the project can be considered as having potential but cannot be ranked as optimal due to these gaps.