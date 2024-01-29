Score: 10
Explanations: 
The project summary provides a comprehensive overview of the analog instrumentation system for machine monitoring, covering many of the requirements set forth. Let's evaluate them one by one:

1. Both sensors have d.c. output, so no demodulation should be used. The summary details the use of strain-gauge pressure sensors and analog output temperature sensors, which produce a DC output. Therefore, this requirement is met.

2. Both sensors must be amplified. The summary mentions instrumentation amplifiers for pressure sensors and amplifier stages for temperature sensors, indicating that both sensor outputs are indeed amplified. Thus, this requirement is met.

3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. While the summary does not explicitly mention a Wheatstone bridge, it is implied as strain-gauge pressure sensors typically operate within a Wheatstone bridge configuration. The use of instrumentation amplifiers confirms that this requirement is likely met.

4. An ADC should be used. The summary explicitly mentions the use of a multiplexed SAR ADC. This requirement is met.

5. Infrared radiation sensors are being linearized either digitally in the computer or by using diode networks or log amplifiers. The summary mentions linearization mechanisms for temperature sensors but does not specify if they are infrared radiation sensors or the exact linearization method. The TMP006 or TMP007 sensors mentioned are not infrared sensors. Thus, this requirement is not met.

6. The solution mentions the sampling order strategy, if it'll be sequential, simultaneously, etc. The summary does not specify the sampling order strategy. This requirement is not met.

7. The sampling frequency of the ADC is not less than 800 Hz. The summary states the ADC has a sampling rate of at least 1 kHz per channel, which satisfies this requirement.

8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency. The summary mentions second-order low-pass filters for both pressure and temperature sensors but does not provide enough information to calculate the attenuation at half the sampling frequency. However, with a sampling rate of at least 1 kHz and a cutoff frequency of 400-500 Hz, it is reasonable to assume that the filters should meet the -20 dB criterion at 500 Hz (half the sampling rate). This requirement is likely met.

9. The low-pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency. The summary states that the cutoff frequencies for the low-pass filters are 400 Hz and 500 Hz, which falls within the specified range given the sampling rate is at least 1 kHz. This requirement is met.

10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s). The summary mentions the use of low-pass filters, but it does not explicitly state their position relative to the multiplexers. However, it is standard practice to place anti-aliasing filters before the multiplexer(s). This requirement is likely met.

11. The project uses multiplexer(s) to choose channels. The summary explicitly mentions the use of a 16-to-1 analog multiplexer. This requirement is met.

12. The multiplexer(s) are solid state. The summary suggests using the ADG1606 or DG406, which are solid-state multiplexers. This requirement is met.

In summary, the project meets most of the requirements, with the exception of the linearization of infrared radiation sensors (requirement 5) and the specification of the sampling order strategy (requirement 6). All other requirements are either explicitly met or can be reasonably inferred based on standard practices and the information provided.