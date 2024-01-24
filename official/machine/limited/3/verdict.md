Verdict: acceptable

Explanations: 
The project summary presents a comprehensive analog instrumentation system designed for monitoring machine conditions through pressure and temperature sensors. It appears to meet many of the specified requirements, but there are some points that need to be addressed:

1. Both sensors have d.c. output: The pressure sensors are piezoresistive and have a d.c. output. The temperature sensors have a digital output, which is not d.c., but since they are directly interfaced with a microcontroller, no demodulation is used. This meets the requirement.
2. Both sensors must be amplified: Amplification is correctly planned for both the pressure and temperature sensors.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The use of an instrumentation amplifier is mentioned, but the summary does not explicitly state that the pressure sensors are part of a Wheatstone bridge configuration.
4. An ADC should be used: A 12-bit SAR ADC is specified for digitizing the sensor signals.
5. Infrared radiation sensors are being linearized: The summary specifies software-based linearization using a microcontroller, which satisfies the requirement.
6. The solution mentions the sampling order strategy: There is no explicit mention of the sampling order strategy (sequential, simultaneous, etc.).
7. The sampling frequency of the ADC: The ADC's sampling rate is specified as at least 2 kHz per channel, which is above the minimum requirement of 800 Hz.
8. The anti-aliasing filter has a cutoff frequency and order: The filters are mentioned to be second-order Butterworth with a cutoff of 450 Hz, but there is no information about the gain at half the sampling frequency.
9. The low-pass cutoff frequency: The cutoff frequency is correctly set between 400 Hz and less than half the sampling frequency.
10. There are low-pass filters positioned before the multiplexer(s): Low-pass filtering is planned, but their exact positioning relative to the multiplexer is not mentioned.
11. The project uses multiplexers to choose channels: An 8-to-1 analog multiplexer is included in the design.
12. The multiplexers are solid-state: Solid-state multiplexers such as the ADG1608 or MAX4051A are specified.

Based on the provided information, the project meets most of the requirements, but there are some details that are either not addressed or not clearly specified. The absence of explicit mentions of the Wheatstone bridge configuration for the pressure sensor and the exact positioning of the low-pass filters relative to the multiplexer, as well as the sampling order strategy, need clarification. Additionally, the gain of the anti-aliasing filters at half the sampling frequency is not provided.