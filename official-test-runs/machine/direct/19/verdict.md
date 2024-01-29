Verdict: acceptable

Explanations: 
The project design for the Industrial Machine Pressure and Temperature Monitoring System has been reviewed based on the provided requirements. Here are the findings:

1. Both sensors have DC output, and no demodulation is used, satisfying requirement 1.
2. Both sensors are amplified as required; the pressure sensor through a dual-stage instrumentation amplifier and the temperature sensor through a non-inverting amplifier, satisfying requirement 2.
3. The pressure sensor is indeed configured in a Wheatstone bridge and amplified by a dual-stage instrumentation amplifier, satisfying requirement 3.
4. An ADC is used in the design, meeting requirement 4.
5. The infrared temperature sensors are linearized using a microcontroller-based approach, which can include digital linearization such as lookup tables or polynomial approximation algorithms, satisfying requirement 5.
6. The sampling strategy is not explicitly mentioned. Thus, requirement 6 is not confirmed.
7. The ADC's sampling frequency is stated to be a minimum of 1 kSPS per channel, which is above the required 800 Hz, satisfying requirement 7.
8. The anti-aliasing filter is a 4th-order Butterworth with a cutoff frequency of 350 Hz. However, the requirement is that the gain should be at least -20 dB at half the sampling frequency, but without the actual sampling frequency or the filter's response curve, this cannot be confirmed. Therefore, requirement 8 is not fully assessed.
9. The low pass cutoff frequency is stated to be 350 Hz, which is below the 400 Hz minimum required, failing requirement 9.
10. There are low-pass filters positioned before the multiplexer, satisfying requirement 10.
11. The project uses an 8-to-1 analog multiplexer, meeting requirement 11.
12. The multiplexer is implied to be solid-state as most modern multiplexers are, but this is not explicitly stated. Therefore, requirement 12 is not fully assessed.

Given the lack of detail on the sampling strategy, the anti-aliasing filter's performance, and the explicit confirmation of solid-state multiplexers, along with the low pass cutoff frequency being below the required minimum, the project cannot be considered "optimal." However, since the project meets most of the essential requirements and can be implemented with some adjustments, it is categorized as "acceptable."