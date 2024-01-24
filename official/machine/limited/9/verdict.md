Verdict: unfeasible

Explanations: 
The project summary addresses the design of an analog instrumentation system for pressure and temperature monitoring. Let's evaluate it against the provided requirements:

1. Both sensors have d.c. output, so no demodulation should be used. The provided sensors (pressure and temperature) have a d.c. output, satisfying this requirement.

2. Both sensors must be amplified. The pressure sensor is amplified by an instrumentation amplifier with a gain of 5000, and the temperature sensor is amplified post non-linear correction with a gain of 50, fulfilling this requirement.

3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. The summary does not explicitly mention the use of a Wheatstone bridge, which is a fatal issue.

4. An ADC should be used. A 12-bit ADC with a sampling rate of 6400 Hz is used, meeting this requirement.

5. Infrared radiation sensors are being linearized either digitally in the computer or by using diode networks or log amplifiers. Digital correction using a microcontroller is mentioned for the temperature sensors, which meets this requirement.

6. The solution mentions the sampling order strategy. The ADC is mentioned to sample at a rate of 6400 Hz for eight channels, which implies a sequential strategy; however, the summary is not explicit about the sampling order.

7. The sampling frequency of the ADC is not less than 800 Hz. The ADC sampling rate is 6400 Hz, which is above the required 800 Hz.

8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency. The low-pass filter has a -48 dB stopband rejection at 800 Hz, which is more than sufficient.

9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency. The cutoff frequency is exactly 400 Hz, which is higher than 400 Hz but not lower than half the total sampling frequency (3200 Hz), which is a fatal issue.

10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s). The summary mentions a low-pass filter but does not specify its position in the signal chain relative to the multiplexer(s).

11. The project uses multiplexer(s) to choose channels. An 8-to-1 solid-state analog multiplexer is included, meeting this requirement.

12. The multiplexer(s) are solid state. The CD4051 series solid-state analog multiplexer is used, which satisfies this requirement.

The project summary fails to explicitly mention the use of a Wheatstone bridge for the pressure sensor and the exact position of the anti-aliasing filters relative to the multiplexer(s). Additionally, the low-pass filter does not meet the requirement of having a cutoff frequency lower than half the total sampling frequency. These are critical issues that would need to be addressed for the project to be feasible.