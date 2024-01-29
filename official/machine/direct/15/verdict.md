Verdict: acceptable

Explanations: 
The project design for an Analog Sensor Acquisition System for Pressure and Temperature Monitoring is detailed and covers many requirements for a robust system. However, there are several critical points that need to be addressed to determine the feasibility of the project:

1. Both sensors have d.c. output, so no demodulation should be used. The summary does not mention demodulation, which is compliant with this requirement.
2. Both sensors must be amplified. The pressure sensor is amplified by an instrumentation amplifier, and the temperature sensor is amplified by a non-inverting operational amplifier circuit, fulfilling this requirement.
3. The pressure sensor is correctly inserted in a Wheatstone bridge and amplified by an instrumentation amplifier.
4. An ADC is used with at least 2 kHz sampling rate per channel, which meets the requirement.
5. Infrared radiation sensors are being linearized using a non-linear conversion circuit which appears to handle the linearization requirement, although it is not explicitly stated if it is digital or analog linearization.
6. The solution does not mention the sampling order strategy (sequential, simultaneous, etc.).
7. The ADC's sampling frequency of at least 2 kHz per channel satisfies the minimum requirement of 800 Hz.
8. The anti-aliasing filter's cutoff frequency is 400 Hz, but there is no information on the order of the filter or the gain at half the sampling frequency to ensure at least -20 dB.
9. The low pass cutoff frequency at 400 Hz is higher than the minimum 400 Hz, but without knowing the total sampling frequency, it is not clear if it is lower than half the total sampling frequency.
10. Low-pass filters are mentioned, but their placement before the multiplexer is not specified.
11. Multiplexers are used to choose channels, fulfilling this requirement.
12. The multiplexers are solid-state, as per the given example (CD4052BE).

Given the information provided, the project seems to have missed some details that are critical to fully assess it against the requirements. Specifically, the sampling order strategy and the explicit placement of low-pass filters before the multiplexers are not addressed. Additionally, the exact nature of the linearization for the infrared sensors and the specifics of the anti-aliasing filter's performance are not adequately detailed.