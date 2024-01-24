Score: 8
Explanations: 
The project summary covers the following requirements:

1. Both sensors have d.c. output: The strain-gauge pressure sensor has a maximum output of 1mV, which is a d.c. value. The infrared radiation detector (Melexis MLX90614) typically provides a digital output, which implies that the output is already in a form that does not require demodulation. Therefore, this requirement is met.

2. Both sensors must be amplified: The strain-gauge pressure sensor is amplified by instrumentation amplifiers with a gain of approximately 4000, and the transimpedance amplifiers are used for the temperature sensors to convert current to voltage, thus amplifying the signal. This requirement is met.

3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The use of instrumentation amplifiers is mentioned, but there is no explicit mention of a Wheatstone bridge configuration. However, since strain-gauge sensors are commonly used in Wheatstone bridge configurations, this can be implicitly assumed. This requirement is met.

4. An ADC should be used: ADCs are mentioned for both pressure and temperature with a sampling rate of ≥ 2 kHz for pressure. This requirement is met.

5. Infrared radiation sensors are being linearized: There is no explicit mention of linearization for the infrared radiation sensors. Without a clear statement on whether linearization is done digitally or through hardware, this requirement is not met.

6. The solution mentions the sampling order strategy: There is no explicit mention of the sampling order strategy (sequential, simultaneous, etc.). This requirement is not met.

7. The sampling frequency of the ADC is not less than 800 Hz: The sampling rate for the ADC is given as ≥ 2 kHz for pressure, which is above 800 Hz. This requirement is met.

8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency: The cutoff frequency of the low pass filter is specified as 500 Hz, but there is no information about the gain at half the sampling frequency. Without this information, this requirement cannot be confirmed as met.

9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The low pass filter cutoff frequency is 500 Hz, which is higher than 400 Hz. Assuming the sampling frequency is the minimum stated 2 kHz, 500 Hz is less than half the total sampling frequency (1 kHz), thus this requirement is met.

10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s): The summary mentions low pass filters for both pressure and temperature, but it does not specify their placement in relation to the multiplexers. Without explicit placement information, this requirement cannot be confirmed as met.

11. The project uses multiplexer(s) to choose channels: Multiplexers are mentioned for both pressure and temperature sensors. This requirement is met.

12. The multiplexer(s) are solid state: The types of multiplexers mentioned (e.g., CD4051B CMOS or ADG704) are solid-state devices. This requirement is met.

The project summary successfully covers 8 out of the 12 requirements.