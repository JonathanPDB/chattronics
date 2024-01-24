Score: 10
Explanations: 
1. The summary does not explicitly state that no demodulation is used, but since both sensors have d.c. output, we can infer that demodulation is not part of the system design, thus fulfilling this requirement.
2. Amplification is mentioned for both the pressure and temperature sensors, meeting this requirement.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. While the summary mentions the use of instrumentation amplifiers, it does not explicitly state the use of a Wheatstone bridge. Therefore, this requirement is not met.
4. An ADC is mentioned in the summary with specific details regarding its type (SAR), resolution, and sampling rate, thus meeting this requirement.
5. Linearization for the temperature sensors is addressed by mentioning analog piecewise linearization using op-amps or digital linearization via microcontroller. This fulfills the requirement.
6. The summary mentions a 16-to-1 analog multiplexer that sequentially selects signals, which suggests a sequential sampling order strategy, thereby meeting this requirement.
7. The sampling frequency of the ADC is specified as a minimum of 2 kHz per channel, which is well above the 800 Hz minimum requirement.
8. The anti-aliasing filter is described as a fourth-order filter with a cutoff frequency set at 450 Hz. However, the summary does not provide enough information to determine if the gain of the signal is at least -20 dB at half the sampling frequency. This requirement is not met due to insufficient information.
9. The low-pass cutoff frequency is stated to be 450 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (given as 2 kHz per channel, which would make half the sampling frequency 1 kHz). This requirement is met.
10. Low-pass filters are mentioned in both the pressure and temperature signal paths, and are positioned before the multiplexer, which implies that they serve as anti-aliasing filters, fulfilling this requirement.
11. The use of a multiplexer is explicitly mentioned, meeting this requirement.
12. The project uses a CD74HC4067 or equivalent, which are CMOS-based solid-state multiplexers, thus meeting this requirement.