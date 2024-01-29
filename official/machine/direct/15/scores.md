Score: 8
Explanations: 
The project summary provides a detailed description of the analog sensor acquisition system for pressure and temperature monitoring. The evaluation of the requirements is as follows:

1. Both sensors have d.c. output: The use of strain-gauge pressure sensors and infrared radiation detectors implies d.c. outputs as there is no mention of any frequency components or modulation in their signals. Requirement met.
2. Both sensors must be amplified: The pressure sensor signal is amplified using an instrumentation amplifier with a gain of 4500, and the temperature sensor signal is amplified using a non-inverting operational amplifier circuit with a fixed gain of 50. Requirement met.
3. The pressure sensor must be inserted in a wheatstone bridge and amplified by a instrumentation amplifier: The use of a strain-gauge and an instrumentation amplifier (e.g., AD620) implies that the pressure sensor is indeed part of a bridge configuration and is being amplified correctly. Requirement met.
4. An ADC should be used: A 12-bit multiplexed ADC with at least 2 kHz sampling rate per channel is mentioned. Requirement met.
5. Infrared radiation sensors are being linearized: The summary mentions a non-linear conversion circuit using a logarithmic/exponential method, which implies linearization is taking place. Requirement met.
6. Sampling order strategy: There is no explicit mention of the sampling order strategy (sequential, simultaneous, etc.), so this requirement is not covered in the summary.
7. The sampling frequency of the ADC: An ADC with at least 2 kHz sampling rate per channel is specified, which exceeds the minimum requirement of 800 Hz. Requirement met.
8. Anti-aliasing filter specifications: The low-pass filter with a 400 Hz cutoff frequency is mentioned, but there is no information about the filter order or the gain at half the sampling frequency, so this requirement cannot be confirmed as met.
9. Low pass cutoff frequency: The cutoff frequency of the low-pass filter is specified as 400 Hz, which is higher than the minimum 400 Hz but not explicitly stated as lower than half the total sampling frequency (1 kHz if considering just one channel). This requirement is not fully confirmed.
10. Low-pass filters before the multiplexer(s): Low-pass filters are mentioned, but their position in relation to the multiplexer is not specified. Requirement not confirmed.
11. The project uses multiplexer(s): Dual 8:1 Multiplexer/Demultiplexer IC is mentioned. Requirement met.
12. The multiplexer(s) are solid state: The CD4052BE is a solid-state multiplexer. Requirement met.

The score is based on the 8 requirements that are clearly met in the summary.