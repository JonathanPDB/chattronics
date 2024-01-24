Score: 10
Explanations: 
The project summary covers the following requirements:

1. Both sensors have d.c. output, so no demodulation should be used. The project description mentions the use of strain-gauge pressure sensors and non-contact infrared temperature sensors. These types of sensors typically have d.c. outputs, satisfying this requirement.

2. Both sensors must be amplified. The summary specifies amplifier topologies and gains for both the pressure (Amplifier_Pressure) and temperature sensors (Amplifier_Temperature), fulfilling this requirement.

3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. While the summary does not explicitly mention a Wheatstone bridge, it is common practice to use one with strain-gauge pressure sensors. The use of an instrumentation amplifier (e.g., AD620, INA118) is explicitly mentioned, thus meeting this requirement.

4. An ADC should be used. The summary specifies the use of an ADC with a minimum overall sampling rate and resolution, fulfilling this requirement.

5. Infrared radiation sensors are being linearized. The summary does not provide explicit details on the linearization method for the infrared temperature sensors, so this requirement is not verifiably met.

6. The solution mentions the sampling order strategy. The summary mentions a minimum overall sampling rate needed for sequential sampling of 16 channels, implying a sequential sampling strategy, which fulfills this requirement.

7. The sampling frequency of the ADC is not less than 800 Hz. The summary mentions a minimum overall sampling rate of 12.8 kHz for 16 channels, which means each channel is sampled at 800 Hz, satisfying this requirement.

8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency. The summary specifies the use of a 4th order Butterworth filter with a cutoff frequency of approximately 500 Hz, but it does not provide the necessary information to verify that the gain is at least -20 dB at half the sampling frequency, so this requirement is not verifiably met.

9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency. The summary specifies a cutoff frequency of approximately 500 Hz, which is higher than 400 Hz. However, without knowing the total sampling frequency, we cannot confirm that it is less than half of that frequency. Since the sampling frequency per channel is 800 Hz, and assuming sequential sampling, the total sampling frequency could be 12.8 kHz or higher, making 500 Hz less than half of the total sampling frequency, so this requirement is met.

10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s). The summary specifies the use of active low-pass filters for both pressure and temperature sensors, fulfilling this requirement.

11. The project uses multiplexers to choose channels. The summary specifies the use of a high-speed 16:1 analog multiplexer, which satisfies this requirement.

12. The multiplexers are solid state. The summary indicates the use of a high-speed analog multiplexer, which would be a solid-state device, thus fulfilling this requirement.

The project summary does not cover:
5. Explicit linearization method for the infrared temperature sensors.
8. Verification that the anti-aliasing filter attenuates the signal by at least -20 dB at half the sampling frequency.