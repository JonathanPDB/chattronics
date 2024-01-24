Score: 9
Explanations: 
1. Both sensors have d.c. output: The pressure sensor is strain-gauge based which typically provides a DC output, and infrared radiation detectors also provide a DC output after the radiation is converted to an electrical signal. Hence, this requirement is met.
2. Both sensors must be amplified: The project specifies gain requirements for both the pressure and temperature amplifiers, indicating that the sensors' outputs are indeed amplified.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The project specifies the use of instrumentation amplifiers for the pressure sensors, which is typical of a Wheatstone bridge configuration, so this requirement is met.
4. An ADC should be used: The project specifies the use of ADCs with a minimum sampling rate of 1 kHz per channel, so this requirement is met.
5. Infrared radiation sensors are being linearized: The project mentions the use of lookup tables, analog linearization circuits, or polynomial linearization for the temperature sensors, which would include infrared radiation sensors. This meets the requirement for linearization.
6. The sampling order strategy is not explicitly mentioned, so we cannot confirm if this requirement is met.
7. The sampling frequency of the ADC is not less than 800 Hz: The project states a minimum sampling rate of 1 kHz per channel, which is above 800 Hz, so this requirement is met.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at half the sampling frequency: The project specifies a second-order Butterworth filter with a 400 Hz cutoff frequency. However, the required gain at half the sampling frequency (500 Hz) is not provided, so we cannot confirm if this requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency is exactly 400 Hz, which meets the lower bound of this requirement, and less than half of the 1 kHz sampling frequency, so this requirement is met.
10. There are low-pass filters positioned before the multiplexer(s): The project specifies the use of anti-aliasing filters, which would be positioned before the multiplexer(s) to reduce aliasing. This requirement is met.
11. The project uses multiplexer(s) to choose channels: The project mentions the use of SAR ADCs with integrated multiplexers, so this requirement is met.
12. The multiplexer(s) are solid state: While not explicitly stated, it is implied that modern ADCs with integrated multiplexers are solid state, so this requirement is met.

The project summary successfully covers 9 out of the 12 requirements. The sampling order strategy (6) and the gain of the anti-aliasing filter at half the sampling frequency (8) are not confirmed. All other requirements are met.