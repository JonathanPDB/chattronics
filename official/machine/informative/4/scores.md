Score: 10
Explanations: 
The project summary provides a detailed outline of an analog instrumentation system for pressure and temperature monitoring. Based on the provided information, the following requirements are met:

1. Both sensors have d.c. output - The pressure sensor is a strain-gauge sensor, and the temperature sensor is an infrared sensor with a 0-5V output. No mention of demodulation suggests they have d.c. outputs. (Requirement met)
2. Both sensors must be amplified - The pressure signal conditioning uses an instrumentation amplifier, and the temperature signal conditioning employs a non-inverting amplifier configuration, indicating amplification for both. (Requirement met)
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier - The use of an instrumentation amplifier suggests this requirement is met, although a Wheatstone bridge is not explicitly mentioned. (Requirement implicitly met)
4. An ADC should be used - An ADC with a resolution of 16-bit and a sampling rate greater than 10 ksps is mentioned. (Requirement met)
5. Infrared radiation sensors are being linearized - Linearization is not explicitly mentioned, but the use of a non-inverting amplifier configuration for signal conditioning could implicitly address this. (Requirement not explicitly met)
6. The solution mentions the sampling order strategy - The ADC section specifies a total rate greater than 10 ksps to handle eight channels at 800 Hz each, suggesting a sequential sampling strategy. (Requirement met)
7. The sampling frequency of the ADC is not less than 800 Hz - The total rate is specified to be greater than 10 ksps for eight channels, which implies each channel is sampled at least at 800 Hz. (Requirement met)
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency - Specific values for the filter cutoff frequency and order are not provided, only that the low-pass filter for the pressure sensor has a "fc significantly above 400 Hz and below half the ADC sampling rate." (Requirement not explicitly met)
9. The low-pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency - The pressure sensor's low-pass filter meets this requirement, but the temperature sensor's low-pass filter cutoff frequency is stated as 100 Hz, which does not meet the requirement. (Requirement partially met)
10. There are low-pass filters positioned before the multiplexer(s) - Low-pass filters are mentioned for both the pressure and temperature signal conditioning stages. (Requirement met)
11. The project uses multiplexers - An 8:1 analog multiplexer is specified. (Requirement met)
12. The multiplexers are solid state - CMOS analog multiplexers are mentioned, which are solid state. (Requirement met)

In summary, the project meets most of the requirements, with some not explicitly addressed or partially met. The requirements that are not met or uncertain due to lack of explicit information are 5 and 8, while requirement 9 is only partially met due to the temperature sensor's filter cutoff frequency.