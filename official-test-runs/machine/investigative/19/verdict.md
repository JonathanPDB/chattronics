Verdict: acceptable

Explanations: 
The provided project summary addresses many of the required elements, but there are some critical issues that need to be addressed:

1. Both sensors are specified to have d.c. output, and no demodulation is mentioned, fulfilling this requirement.
2. Amplification is mentioned for both the pressure sensors (instrumentation amplifiers) and temperature sensors (transimpedance amplifiers), satisfying this requirement.
3. The pressure sensor is mentioned to be used with instrumentation amplifiers, which is typical in Wheatstone bridge applications, but the direct inclusion of a Wheatstone bridge is not explicitly stated.
4. ADCs are used for both pressure and temperature sensors, fulfilling this requirement.
5. The temperature sensors are infrared radiation detectors, but the summary does not explicitly mention linearization using diode networks, log amplifiers, or digital methods.
6. The sampling order strategy is not mentioned, which is necessary for understanding how the sensors are read by the ADC.
7. The sampling frequency of the ADC for pressure is mentioned to be ≥ 2 kHz, which is above the minimum required 800 Hz.
8. & 9. The anti-aliasing filter for the pressure sensor has a cutoff frequency of 500 Hz, which is higher than 400 Hz and lower than half the sampling frequency, but the gain at half the sampling frequency is not specified.
10. Low pass filters are mentioned but are not explicitly positioned before the multiplexers in the signal chain.
11. Multiplexers are included in the design, satisfying this requirement.
12. The multiplexers mentioned are solid state, fulfilling this requirement.

Given the above points, the project summary does not fully meet all the specified requirements. Key information is missing, such as the explicit mention of a Wheatstone bridge for the pressure sensor, the linearization method for the infrared sensors, the sampling order strategy, the positioning of the anti-aliasing filters, and the gain specification at half the sampling frequency for the anti-aliasing filters. Therefore, the project cannot be categorized as "optimal."

However, since most of the essential requirements are met or only partially unfulfilled, and the project does not contain any fatal issues, the project falls into the "acceptable" category. It can be implemented with some modifications to meet all the requirements.