Score: 11
Explanations: 
1. The project summary explicitly states that the pressure sensor provides an output of 2 mV/V with a recommended excitation of 10 VDC, which implies a DC output. The requirement is met.

2. Both the pressure and temperature sensors are amplified. The pressure sensor is amplified by an instrumentation amplifier with a gain of 5000, and the temperature sensor by a non-inverting amplifier with a gain of 50. The requirement is met.

3. The pressure sensor is part of an instrumentation amplifier setup, which is typically used in conjunction with a Wheatstone bridge, though the bridge itself is not explicitly mentioned. Given that strain-gauge pressure sensors are commonly used in a Wheatstone bridge configuration and the use of an instrumentation amplifier is indicated, this requirement is implicitly met.

4. An ADC is mentioned with a resolution of 12-bit and a sampling rate of at least 2 kSPS. The requirement is met.

5. The temperature sensor array uses a piecewise linearization technique with operational amplifiers to linearize the output of the infrared radiation sensors. The requirement is met.

6. The project summary does not explicitly mention the sampling order strategy (sequential, simultaneous, etc.). The requirement is not met.

7. The ADC sampling rate is stated to be at least 2 kSPS, which is higher than the minimum requirement of 800 Hz. The requirement is met.

8. The pressure sensor's anti-aliasing filter is a 4th-order low-pass Butterworth filter with a cutoff frequency of 450 Hz. A 4th-order filter has an attenuation rate of -24 dB/octave, which would provide at least -20 dB at half the sampling frequency (1000 Hz for a 2 kSPS rate). The requirement is met.

9. The low-pass cutoff frequencies for both pressure and temperature sensors are 450 Hz and 482 Hz, respectively, which are higher than 400 Hz and lower than half the total sampling frequency (1000 Hz for a 2 kSPS rate). The requirement is met.

10. The project summary mentions low-pass filters for both the pressure and temperature sensors, which are presumed to be positioned before the multiplexer. The requirement is met.

11. An 8-to-1 analog multiplexer is used to select one of the eight sensor signals. The requirement is met.

12. The project summary specifies the use of model CD4051B or DG408 for the multiplexer, which are solid-state devices. The requirement is met.