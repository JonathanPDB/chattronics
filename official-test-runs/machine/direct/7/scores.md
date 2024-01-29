Score: 9
Explanations: 
1. Both sensors are described as having a d.c. output. The pressure sensor outputs a maximum of 1 mV, and the temperature sensor outputs an I²C-compatible digital interface which can be assumed to be d.c. for the purpose of this system. (Requirement 1 reached)

2. Both sensors are amplified. The pressure sensor is amplified using a three-op-amp instrumentation amplifier with a gain of 5000, and the temperature sensor is amplified with a gain of 50. (Requirement 2 reached)

3. The pressure sensor is described as industrial-grade strain-gauge which is typically used in a Wheatstone bridge configuration, and it is amplified by an instrumentation amplifier. The use of a three-op-amp instrumentation amplifier is consistent with this requirement. (Requirement 3 reached)

4. An ADC with a sampling rate of ≥2 kHz per channel is specified, which meets the requirement for an ADC. (Requirement 4 reached)

5. There is no explicit mention of how the infrared radiation sensors are being linearized. The temperature sensor array output is described as I²C-compatible, which suggests digital processing, but there's no specific mention of digital linearization, diode networks, or log amplifiers. (Requirement 5 not reached)

6. The solution does not mention the sampling order strategy (sequential, simultaneously, etc.). (Requirement 6 not reached)

7. The ADC sampling frequency is specified as ≥2 kHz per channel, which is well above the minimum of 800 Hz. (Requirement 7 reached)

8. The anti-aliasing filter for the pressure sensor is a 2nd-order lowpass Butterworth filter with a cutoff frequency of 400 Hz. The temperature sensor uses a 4th-order filter with the same cutoff frequency. However, there is no information provided to confirm that the gain of the signal is at least -20 dB at half the sampling frequency. (Requirement 8 not reached)

9. The lowpass filter cutoff frequency is specified as 400 Hz, which is higher than the minimum of 400 Hz but without knowing the ADC's exact sampling frequency, it's not possible to confirm if it's lower than half the total sampling frequency. Assuming the ADC sampling frequency is exactly 2 kHz, then the cutoff frequency meets this requirement. (Requirement 9 reached)

10. Low-pass filters are mentioned for both pressure and temperature signals, and given the design, it can be inferred that they are positioned before the multiplexer(s). (Requirement 10 reached)

11. A 16:1 differential multiplexer is specified, which meets the requirement to use multiplexer(s) to choose channels. (Requirement 11 reached)

12. The multiplexer described, such as the Analog Devices ADG1606, is indeed a solid-state device. (Requirement 12 reached)