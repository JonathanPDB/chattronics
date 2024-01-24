Verdict: acceptable

Explanations: 
The project description provides a comprehensive overview of an analog instrumentation system designed to monitor pressure and surface temperature. The design incorporates a variety of components, including pressure and temperature sensors, amplifiers, filters, a multiplexer, and an ADC, all of which are necessary for the proper functioning of the system. Here's how the design aligns with the provided requirements:

1. The pressure sensor (Honeywell Sensotec Model TJE) and the temperature sensor (MLX90614ESF-BAA) both have DC outputs, satisfying the first requirement.
2. Both sensors are amplified; the pressure sensor via an instrumentation amplifier with a gain of 5000, and the temperature sensor through a pre-amplifier with a gain of 50.
3. The pressure sensor is indeed configured with a Wheatstone bridge and amplified by an instrumentation amplifier.
4. An ADC with a sampling rate of at least 2 kSPS is included in the system, fulfilling the ADC requirement.
5. The temperature sensor output is linearized using a piecewise linearization technique, meeting the requirement for linearization.
6. The sampling order strategy is not explicitly mentioned, but the inclusion of a multiplexer implies sequential sampling.
7. The ADC's sampling frequency is at least 2 kSPS, well above the minimum requirement of 800 Hz.
8. The anti-aliasing filter for the pressure sensor is a 4th-order Butterworth filter with a cutoff frequency of 450 Hz, which should ensure at least -20 dB gain at half the sampling frequency (1 kSPS), satisfying this requirement.
9. The low-pass cutoff frequency for both pressure and temperature sensors is within the specified range (higher than 400 Hz and lower than half the total sampling frequency).
10. Low-pass filters are positioned before the multiplexer, reducing the risk of aliasing.
11. An 8-to-1 analog multiplexer is used to select sensor channels, complying with the requirement for multiplexers.
12. The multiplexer models suggested (CD4051B or DG408) are solid-state devices.

All the essential requirements have been met. However, the project description lacks detail on the anti-aliasing filter's order and the exact attenuation at half the sampling frequency, which are necessary to fully evaluate requirement 8. Given that the project is theoretically sound and implementable, but with some omissions and assumptions, the verdict is "acceptable."