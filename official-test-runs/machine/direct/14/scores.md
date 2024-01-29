Score: 10
Explanations: 
The project summary covers the following requirements:

1. The pressure sensors have a d.c. output, as indicated by the maximum signal of 1 mV corresponding to maximum pressure. (Requirement 1)
2. Both the pressure and temperature sensors are mentioned to be amplified. The pressure sensor's signal is explicitly amplified using a three-op-amp instrumentation amplifier configuration. However, the temperature sensor's output is digital and does not require amplification in the analog domain. (Requirement 2)
3. The pressure sensors are inserted in a Wheatstone bridge configuration and amplified by an instrumentation amplifier. This is evident from the use of precision resistors and op-amps for amplifying the pressure sensor's signal. (Requirement 3)
4. An ADC is used, with a sampling rate of ≥1 kHz per channel. (Requirement 4)
5. The temperature sensors' output is digital, hence linearization can be handled digitally. (Requirement 5)
6. The project summary does not explicitly mention the sampling order strategy. (Requirement 6)
7. The sampling frequency of the ADC is not less than 800 Hz, as it is stated to be ≥1 kHz per channel. (Requirement 7)
8. The anti-aliasing filter has a cutoff frequency of 400 Hz, which is less than half the sampling rate of 1 kHz, but the project summary does not provide information on the gain of the signal at half the sampling frequency, so this requirement cannot be confirmed. (Requirement 8)
9. The low pass cutoff frequency is stated to be 400 Hz, which complies with being higher than 400 Hz and lower than half the total sampling frequency (500 Hz). (Requirement 9)
10. Low-pass filters are positioned before the multiplexer inputs, as stated for Multiplexer2. (Requirement 10)
11. Multiplexers are used to choose channels, as indicated by the description of Multiplexer1 and Multiplexer2. (Requirement 11)
12. The suggested CD4051B CMOS multiplexers are solid-state devices. (Requirement 12)

The requirements that are not covered in the summary or cannot be confirmed are:
- Requirement 6: The sampling order strategy is not mentioned.
- Requirement 8: The gain of the signal at half the sampling frequency is not specified.

Therefore, the project meets 10 out of the 12 requirements.