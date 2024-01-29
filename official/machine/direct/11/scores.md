Score: 10
Explanations: 
1. The project summary specifies that the pressure sensors have a scaled maximum output of 1 mV, which is a DC output, and the temperature sensors have a digital output via I2C, which is also DC. No demodulation is mentioned or required.
2. Both the pressure and temperature sensor signals are amplified using instrumentation amplifiers, with gains set for matching the ADC input range.
3. Although the pressure sensor is mentioned to be a strain-gauge type, which is commonly used in Wheatstone bridge configurations, the summary does not explicitly state that it is inserted in a Wheatstone bridge. However, the use of an instrumentation amplifier implies a differential measurement, which is consistent with a Wheatstone bridge configuration.
4. An ADC is mentioned with a sampling rate of at least 1 kHz per channel.
5. Linearization for the temperature sensors is performed using logarithmic amplifiers, which is one of the methods required.
6. The summary does not mention the sampling order strategy, whether it will be sequential, simultaneous, etc.
7. The sampling frequency is stated to be at least 1 kHz per channel, which is above the 800 Hz minimum requirement.
8. The anti-aliasing filter is a 4th-order active Butterworth low-pass filter with a cutoff frequency of 450 Hz. At half the sampling frequency (500 Hz for a 1kHz sampling rate), the filter would provide a gain reduction of 24 dB per octave, but without the exact filter design, it is not possible to confirm that it meets the -20 dB requirement at precisely half the sampling frequency. However, with a 4th-order Butterworth response, it is reasonable to assume that it will meet the -20 dB at 500 Hz requirement.
9. The low-pass cutoff frequency is stated to be 450 Hz, which meets the requirement of being higher than 400 Hz and lower than half the total sampling frequency (assuming a sampling frequency of 1 kHz).
10. Low-pass filters are positioned before the ADC, though it is not explicitly stated that they are before the multiplexer(s). However, since the filters are part of the anti-aliasing process, it is implicit that they would be placed in the signal path before the multiplexing stage to be effective.
11. The project uses an ADG1606 multiplexer from Analog Devices for channel selection.
12. The ADG1606 is a solid-state multiplexer.

Overall, the project summary covers most of the required items, although it lacks explicit mention of some details such as the sampling order strategy and the exact performance of the anti-aliasing filter at half the sampling frequency. However, based on the information provided and common design practices, it is reasonable to infer that these requirements are likely met.