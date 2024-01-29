Score: 8
Explanations: 
1. Both sensors have d.c. output: The pressure sensor is a strain gauge, and the temperature sensor is an infrared radiation detector, both typically provide d.c. output, so this requirement is met.
2. Both sensors must be amplified: The design includes instrumentation amplifiers with specific gains for both sensors, meeting this requirement.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The pressure sensor is a strain gauge, which is typically used in Wheatstone bridge configurations, and the design includes the use of an instrumentation amplifier, fulfilling this requirement.
4. An ADC should be used: The design specifies the use of a SAR ADC, meeting this requirement.
5. Infrared radiation sensors are being linearized: The design does not explicitly mention linearization of the infrared sensor output. This requirement is not met.
6. The solution mentions the sampling order strategy: There is no explicit mention of the sampling order strategy (sequential, simultaneous, etc.); therefore, this requirement is not met.
7. The sampling frequency of the ADC is not less than 800 Hz: The design does not specify the ADC sampling rate, so this requirement is not met.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency: The filter is a 2nd order Sallen-Key with a cutoff frequency of 600 Hz, which implies that at 800 Hz (half of 1600 Hz, the assumed minimum sampling frequency due to requirement 7), the attenuation would be more than -20 dB, fulfilling this requirement.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency is 600 Hz, which is higher than 400 Hz. However, since the sampling frequency is not specified, it cannot be determined if it is less than half the total sampling frequency. This requirement is not met.
10. There are low-pass filters positioned before the multiplexer(s): The design includes an anti-aliasing filter, which is a type of low-pass filter typically placed before the ADC, which implies it is before the multiplexer in the signal chain. This requirement is met.
11. The project uses multiplexer(s) to choose channels: The design specifies the use of a 16:1 analog CMOS multiplexer. This requirement is met.
12. The multiplexer(s) are solid state: The design specifies a CMOS multiplexer, which is a solid-state device. This requirement is met.

The requirements that are not met are 5, 6, 7, and 9. The total score is 8 out of 12.