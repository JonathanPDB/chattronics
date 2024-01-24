Score: 11
Explanations: 
1. Both sensors have d.c. output: The project summary does not explicitly mention that the sensors have d.c. output, but given the nature of a strain-gauge pressure sensor and an IR thermopile sensor, it is generally understood that they provide d.c. output signals proportional to the measured quantities.

2. Both sensors must be amplified: Amplification is clearly mentioned for both the pressure sensor ("Amplifier_Pressure" with a gain of 1000) and the temperature sensor ("Amplifier_Temp" with a gain of approximately 50).

3. The pressure sensor must be inserted in a wheatstone bridge and amplified by an instrumentation amplifier: The use of a strain-gauge pressure sensor and an instrumentation amplifier (such as the AD620 or INA333) implies the use of a Wheatstone bridge, which is typical for strain-gauge sensors.

4. An ADC should be used: The summary explicitly mentions the use of a Successive Approximation Register (SAR) ADC with a sampling rate of 20 kSPS.

5. Infrared radiation sensors are being linearized: There is mention of a linearization circuit in the "Amplifier_Temp" section, indicating that the non-linear response of the IR sensor is being addressed, though the exact method is not detailed.

6. The solution mentions the sampling order strategy: The summary indicates that there are 16 channels being sampled at 1 kHz per channel, which implies a sequential sampling strategy, as simultaneous sampling of all channels would not require a 20 kSPS ADC for 16 channels.

7. The sampling frequency of the ADC is not less than 800 Hz: The ADC's sampling rate is mentioned to be 20 kSPS total, which is well above the 800 Hz minimum requirement.

8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency: The cutoff frequencies for the low-pass filters are given as 800 Hz and 450 Hz, and both are second-order Butterworth filters. However, without knowing the sampling frequency for each individual channel, we cannot verify the -20 dB gain at half the sampling frequency.

9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequencies for both filters (800 Hz and 450 Hz) meet the requirement of being higher than 400 Hz. Since the sampling frequency per channel is 1 kHz, the cutoff frequencies are also lower than half the total sampling frequency (10 kHz) for each channel.

10. There are low-pass filters positioned before the multiplexers: The description includes active low-pass Butterworth filters for both the pressure and temperature signals, which would be positioned before the multiplexers to reduce aliasing.

11. The project uses multiplexers to choose channels: The summary clearly states the use of multiplexers for both pressure and temperature channels.

12. The multiplexers are solid state: The models mentioned (CD74HC4067, CD4051B, DG202) are all solid-state multiplexers.