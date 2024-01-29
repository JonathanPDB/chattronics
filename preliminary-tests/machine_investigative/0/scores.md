Score: 7
Explanations: 
The project summary covers the following requirements:

1. Both sensors have d.c. output, so no demodulation should be used. The pressure sensor output is specified as 1 mV/psi, which is a DC signal. However, the temperature sensor output is specified as a digital signal (PWM or I2C), which does not inherently indicate a DC signal. Since the requirement specifies "no demodulation," and PWM could be considered a form of modulation, this requirement is not met.

2. Both sensors must be amplified. The pressure sensor is amplified, as indicated by the gain requirement of 5000. The temperature sensor uses a buffer amplifier, which is a form of amplification (unity-gain). Therefore, this requirement is met.

3. The pressure sensor must be inserted in a wheatstone bridge and amplified by an instrumentation amplifier. While the use of an instrumentation amplifier is indicated, there is no explicit mention of a Wheatstone bridge configuration. This requirement is not met.

4. An ADC should be used. An ADC (ADS1115 or equivalent) is specified, so this requirement is met.

5. Infrared radiation sensors are being linearized either digitally in the computer or by using diode networks or log amplifiers. There is no explicit mention of linearization methods for the infrared sensors, so this requirement is not met.

6. The solution mentions the sampling order strategy, if it'll be sequential, simultaneously, etc. There is no mention of the sampling order strategy, so this requirement is not met.

7. The ADC samples one sensor per time and does not do it all in the same channel. The use of an 8-to-1 multiplexer implies that the ADC samples one sensor at a time, as the multiplexer selects one channel at a time for the ADC. This requirement is met.

8. The sampling frequency of the ADC is not less than 800 Hz. The ADC sampling rate is specified as 1 kHz, which satisfies this requirement.

9. The sampling frequency of the ADC is higher than 2000 Hz. The specified sampling rate is 1 kHz, which does not meet this requirement.

10. The ADC should be a SAR. There is no indication that the ADC (ADS1115 or equivalent) is a Successive Approximation Register (SAR) ADC, so this requirement is not met.

11. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency. The low-pass filter for the pressure signal conditioning is specified to be "slightly above 400 Hz," which meets the first part of the requirement. However, since the sampling frequency is 1 kHz, half of that would be 500 Hz, making the cutoff frequency requirement not met.

12. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexers. Low-pass filters are mentioned for both signal conditioning circuits, but their position relative to the multiplexers is not specified. This requirement is not met.

13. The project uses multiplexers to choose channels. An 8-to-1 multiplexer is specified, so this requirement is met.

14. The multiplexers are solid state. While specific multiplexer models (e.g., DG408, ADG1606) are suggested and these are typically solid-state devices, there is no explicit mention that they are solid state. This requirement is implicitly met due to the nature of the suggested models.

In summary, out of the 14 requirements, 7 are met (2, 4, 7, 8, 13, and implicitly 14), while the others are not.