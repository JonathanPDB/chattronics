Score: 9
Explanations: 
The project summary covers the following requirements:

1. Since both sensors (pressure and temperature) are mentioned as having a DC output and no demodulation is discussed, it is implied that requirement 1 is met.
2. Both the pressure and temperature sensors are mentioned to be amplified, which satisfies requirement 2.
3. The pressure sensor is specified to be amplified by an instrumentation amplifier, which meets requirement 3.
4. An ADC is mentioned to be used in the system, fulfilling requirement 4.
5. Linearization for the infrared sensors is mentioned, though it is not explicitly stated how it is done. However, the mention of a "custom designed" linearization network suggests that one of the methods (diode networks or log amplifiers) could be used. This is a bit ambiguous, but given the benefit of the doubt, requirement 5 is met.
6. The sampling order strategy is not specifically mentioned, so requirement 6 is not met.
7. The ADC sampling rate is stated to be "at least 1 kS/s per channel," which is more than the 800 Hz minimum, satisfying requirement 7.
8. The anti-aliasing filter is described as a 4th order Butterworth filter with a cutoff frequency between 100 Hz to 200 Hz, but without the ADC sampling rate, the -20 dB gain at half the sampling frequency cannot be confirmed. However, if we assume the minimum sampling rate of 1 kS/s per channel as mentioned, the cutoff frequency would not meet the -20 dB at 500 Hz (half of 1 kS/s). Thus, requirement 8 is not met.
9. The low-pass cutoff frequency is mentioned to be 500 Hz, which is higher than 400 Hz but without the exact sampling frequency, it's unclear if it's lower than half the total sampling frequency. If we assume the sampling frequency is at least 1 kS/s per channel, the cutoff frequency does not meet this requirement, so requirement 9 is not met.
10. The anti-aliasing filters are mentioned to be positioned before the multiplexer, fulfilling requirement 10.
11. The project uses an 8-to-1 analog multiplexer, which satisfies requirement 11.
12. The multiplexer suggested (ADG1606) is a solid-state device, meeting requirement 12.

The requirements that are not met are 6, 8, and 9. Therefore, the project successfully covers 9 out of 12 requirements.