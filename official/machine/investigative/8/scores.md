Score: 8
Explanations: 
1. Both sensors have d.c. output: The pressure sensor is a strain-gauge type which typically provides a DC output, and the infrared radiation detectors also provide a DC output, so this requirement is met.
2. Both sensors must be amplified: The design specifies the use of instrumentation amplifiers for both pressure and temperature sensors, so this requirement is met.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The project discusses using an instrumentation amplifier with a strain gauge pressure sensor, which implies a Wheatstone bridge configuration, so this requirement is met.
4. An ADC should be used: The project specifies the use of a SAR ADC, so this requirement is met.
5. Infrared radiation sensors are being linearized: The project does not specifically mention linearization methods for the infrared sensors, so this requirement is not clearly met.
6. The solution mentions the sampling order strategy: The project mentions using a multiplexer and a recommended sampling rate but does not specify the sampling order strategy (sequential, simultaneous, etc.), so this requirement is not met.
7. The sampling frequency of the ADC is not less than 800 Hz: The project specifies a sampling rate of at least 1 kHz per channel, which exceeds 800 Hz, so this requirement is met.
8. The anti-aliasing filter has a cutoff frequency and order: The project specifies a third-order Butterworth filter with a cutoff frequency of 500 Hz, but it does not mention the gain of the signal at half the sampling frequency, so we cannot confirm this requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: With a cutoff frequency of 500 Hz and a sampling rate of at least 1 kHz, this requirement is met.
10. There are low-pass filters positioned before the multiplexer(s): The project mentions the use of a low pass filter but does not specify its position relative to the multiplexer, so this requirement is not clearly met.
11. The project uses multiplexer(s) to choose channels: The project specifies the use of an analog multiplexer IC, so this requirement is met.
12. The multiplexer(s) are solid state: The project specifies using an analog multiplexer IC such as ADG508F or CD74HC4067, which are solid-state devices, so this requirement is met.