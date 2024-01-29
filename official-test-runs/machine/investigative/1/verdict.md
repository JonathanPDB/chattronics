Verdict: incorrect

Explanations: 
The project summary provided includes several details that align with the requirements, but there are also some discrepancies and missing information which need to be addressed:

1. Both sensors have d.c. output, so no demodulation should be used. The summary does not explicitly state that demodulation is not used, but since it discusses d.c. outputs from sensors, it is implied that demodulation is not part of the design.
2. Both sensors must be amplified. The summary mentions the use of an instrumentation amplifier for the pressure sensor array and an amplifier/filter array for the temperature sensors, which satisfies this requirement.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. The use of an instrumentation amplifier is mentioned, but the summary does not explicitly state that the pressure sensor is part of a Wheatstone bridge configuration.
4. An ADC should be used. The summary discusses the use of a 12-bit SAR ADC for both the pressure and temperature sensors, meeting this requirement.
5. Infrared radiation sensors are being linearized either digitally in the computer or by using diode networks or log amplifiers. The summary does not provide details on the linearization method for the infrared sensors.
6. The solution mentions the sampling order strategy. The summary does not mention the sampling order strategy (sequential, simultaneous, etc.).
7. The sampling frequency of the ADC is not less than 800 Hz. The summary mentions a sampling rate of at least 1 kHz for the ADC1, which meets this requirement. However, it mentions a sampling rate of only 10 samples per second per channel for ADC2, which does not meet the requirement.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency. The summary does not provide enough information to determine if this requirement is met.
9. The low-pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency. The summary mentions the use of a low-pass active filter with a cutoff frequency around 500 Hz, which meets the lower bound but does not confirm if it is lower than half the total sampling frequency.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s). The summary mentions each channel of Multiplexer2 includes a low-pass filter, but it is unclear if similar filters are used for Multiplexer1.
11. The project uses multiplexer(s) to choose channels. The summary discusses the use of two multiplexers, which satisfies this requirement.
12. The multiplexer(s) are solid-state. The summary does not state whether the multiplexers are solid-state, but the IC references suggest that they are solid-state devices.

The project summary does not fully satisfy all the essential requirements, specifically the configuration of the pressure sensor in a Wheatstone bridge, the linearization process for the infrared sensors, the sampling order strategy, the anti-aliasing filter specifications, and the sampling frequency for ADC2. Additionally, some details are not explicitly confirmed, such as the solid-state nature of the multiplexers and the presence of low-pass filters before Multiplexer1. Therefore, the project cannot be categorized as "optimal."

Given the lack of information on some essential requirements and the failure to meet others, the project summary falls into the category of "incorrect" as it has conceptual gaps that are critical to the design, such as the incorrect sampling frequency for ADC2 and the lack of a clear linearization method for the infrared sensors.