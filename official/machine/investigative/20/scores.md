Score: 7
Explanations: 
1. The sensors are described as having a DC output, which means no demodulation is required, fulfilling requirement 1.
2. Both the pressure and temperature sensors are amplified using instrumentation amplifiers, meeting requirement 2.
3. While a Wheatstone bridge and instrumentation amplifier are mentioned for the pressure sensors, there is no explicit mention of the bridge being implemented, which leaves requirement 3 uncertain.
4. An ADC is used in the system, satisfying requirement 4.
5. There is no mention of linearization methods for the infrared radiation sensors, so requirement 5 is not met.
6. There is no explicit mention of the sampling order strategy, so requirement 6 is not met.
7. The sampling rate is stated to be at least 2 kSPS per channel, which is above the minimum of 800 Hz required by requirement 7.
8. The anti-aliasing filter for the pressure sensor has a cutoff frequency of 800 Hz, which is equal to the Nyquist frequency (half the sampling frequency of 1600 Hz if we consider 2 kSPS per channel), so we cannot confirm that the gain of the signal is at least -20 dB at half the sampling frequency, leaving requirement 8 uncertain.
9. The low pass cutoff frequency for the pressure sensor is 800 Hz, which is higher than 400 Hz but not necessarily lower than half the total sampling frequency, as the exact sampling frequency is not provided (we only know it's at least 2 kSPS per channel), making requirement 9 uncertain.
10. Low-pass filters are mentioned for both pressure and temperature sensors, positioned before the multiplexer, fulfilling requirement 10.
11. Multiplexers are used in the system, satisfying requirement 11.
12. The multiplexer mentioned is a CMOS type, which is a solid-state device, meeting requirement 12.