Verdict: incorrect

Explanations: 
The project description provided outlines a comprehensive analog data acquisition system for pressure and temperature monitoring. The project is reviewed based on the provided requirements:

1. Both sensors have a d.c. output. The pressure sensor output is specified to be a maximum of 1 mV, which is a d.c. output. The temperature sensor output is digital (PWM or SMBus), which does not require demodulation but does not meet the requirement of having a d.c. output for direct amplification.

2. Both sensors must be amplified. The pressure sensor is connected to an instrumentation amplifier with a specified gain, fulfilling this requirement. However, the temperature sensor is a digital output sensor, and its signal does not go through an analog amplification stage as required.

3. The pressure sensor is inserted in a Wheatstone bridge configuration and amplified by an instrumentation amplifier, meeting this requirement.

4. An ADC is used with a specified sampling rate and resolution, meeting this requirement.

5. The project description does not mention linearization for the infrared radiation sensors, which is a necessary requirement. It is unclear whether the digital output is already linearized internally or if additional linearization is needed.

6. The sampling order strategy is not explicitly mentioned, which is information that would be useful to determine the overall system performance.

7. The sampling frequency of the ADC is 1 kHz per channel, which is above the 800 Hz minimum requirement.

8. The anti-aliasing filter has a 4th order Butterworth filter with a cutoff frequency of 500 Hz, which should ensure a gain of at least -20 dB at half the sampling frequency (400 Hz), meeting this requirement.

9. The low-pass cutoff frequency is 500 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (4 kHz), fulfilling this requirement.

10. There are low-pass filters positioned before the multiplexer, meeting this requirement.

11. The project uses an 8-to-1 analog multiplexer, meeting this requirement.

12. The multiplexer is solid-state, meeting this requirement.

Overall, the project meets most of the requirements but fails to meet the essential requirement of having both sensors with a d.c. output for direct amplification and does not specify linearization for the infrared sensors. The sampling order strategy is also not mentioned.