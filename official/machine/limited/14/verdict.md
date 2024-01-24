Verdict: incorrect

Explanations: 
The project review for the Pressure and Temperature Data Acquisition System for Machine Monitoring indicates that the project has met most of the necessary requirements, but there are a few critical points that need to be addressed:

1. The pressure sensors are indeed being used in conjunction with instrumentation amplifiers, as required. However, there is no explicit mention of a Wheatstone bridge configuration for the pressure sensors, which is necessary for requirement 3.

2. Amplification of both sensor signals is confirmed, satisfying requirement 2.

3. The ADC is present and meets the minimum sampling frequency of 800 Hz per channel, satisfying requirement 7.

4. The anti-aliasing filter is specified with a cutoff frequency of 400 Hz, which is the minimum acceptable value as per requirement 9. However, the requirement 8 for the gain of the signal to be at least -20 dB at half the sampling frequency is not explicitly confirmed.

5. The temperature sensors are not being linearized because they provide a digital output. This configuration does not meet requirement 5, which specifies linearization of infrared radiation sensors.

6. The sampling strategy (sequential or simultaneous) is not described, which leaves requirement 6 unaddressed.

7. The multiplexers are indeed solid-state, as required by requirement 12, and are used to select channels, satisfying requirement 11.

8. Low-pass filters are placed before the multiplexer inputs, which aligns with requirement 10.

Because the project did not meet some of the essential requirements (specifically requirement 3, 5, 6, and possibly 8), the verdict is that the project is "incorrect".