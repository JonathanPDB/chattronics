Score: 6
Explanations: 
The project covers the following requirements:

1. Both sensors have DC output: This is implicit in the description of the sensors, as there's no mention of AC or varying signals requiring demodulation.
2. Both sensors must be amplified: This is met as the pressure signal is amplified using a dual-stage instrumentation amplifier with a total gain of 4900, and the temperature signal is amplified with a gain of 50.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: Although the use of a Wheatstone bridge is not explicitly mentioned, the use of instrumentation amplifiers suggests a bridge configuration might be used. However, due to the lack of explicit confirmation, this requirement cannot be considered as met.
4. An ADC should be used: This requirement is met as the project includes a 16-bit SAR ADC.
5. Infrared sensors are being linearized: The project meets this requirement by using a microcontroller-based approach for linearization.
6. The solution mentions the sampling order strategy: No specific sampling order strategy is mentioned, so this requirement is not met.
7. The sampling frequency of the ADC is not less than 800 Hz: The sampling rate is stated to be a minimum of 1 kSPS per channel, which exceeds the 800 Hz requirement.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency: The filter is 4th-order Butterworth with a cutoff frequency of 350 Hz. However, the project does not indicate that the gain at half the sampling frequency is at least -20 dB. This requirement is not explicitly met as we cannot determine if the -20 dB attenuation is achieved without additional information.
9. The low-pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency of the anti-aliasing filter is 350 Hz, which does not meet the requirement of being higher than 400 Hz.
10. There are low-pass filters before the multiplexer(s): The project includes anti-aliasing filters, but it's not clear whether they are positioned before the multiplexer(s), as required.
11. The project uses multiplexer(s) to choose channels: This requirement is met with the inclusion of an 8-to-1 analog multiplexer.
12. The multiplexer(s) are solid state: As the type of multiplexer is not specified, we cannot confirm this requirement is met based solely on the given information.

Based on the information provided, the project summary successfully covers 6 out of the 12 requirements listed.
