Verdict: acceptable

Explanations: 
The project summary provided outlines a comprehensive analog instrumentation system for monitoring pressure and temperature. Here's how it addresses the requirements:

1. Both sensors have d.c. output, as indicated by the strain-gauge pressure sensors and infrared radiation detectors.
2. Both sensor types are amplified, with gains of 5000 for pressure and 50 for temperature.
3. The pressure sensor is suggested to be paired with an instrumentation amplifier; however, the use of a Wheatstone bridge for the pressure sensor is not explicitly mentioned.
4. An ADC with a minimum sampling rate of 2 kHz per channel is included, exceeding the 800 Hz requirement.
5. Linearization of the infrared sensor output is mentioned, with options for both analog and digital methods.
6. A sequential sampling order strategy is implied by the use of a 16-to-1 analog multiplexer.
7. The sampling frequency exceeds the minimum requirement of 800 Hz, with a minimum of 2 kHz per channel specified.
8. The anti-aliasing filter is fourth-order with a cutoff frequency of 450 Hz, but the gain at half the sampling frequency is not explicitly stated.
9. The low-pass cutoff frequency of 450 Hz is within the specified range.
10. Low-pass filters are positioned before the multiplexer to reduce aliasing.
11. The project uses a multiplexer to choose channels.
12. The multiplexer is solid-state, as indicated by the suggested CD74HC4067 model.

However, the requirement for the Wheatstone bridge configuration for the pressure sensor is not explicitly confirmed, and the gain at half the sampling frequency for the anti-aliasing filter is not provided. Therefore, while the project meets most of the requirements, it falls short on fully confirming some essential details.