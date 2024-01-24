Score: 7
Explanations: 
1. Both sensors have d.c. output: The piezoresistive strain gauge pressure sensor and the MLX90614 IR temperature sensor both provide DC output signals. (Requirement met)
2. Both sensors must be amplified: Both the pressure sensor and the IR temperature sensor are connected to instrumentation amplifiers, which implies they are being amplified. (Requirement met)
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: While the use of a Wheatstone bridge is not explicitly mentioned, the use of piezoresistive strain gauge pressure sensors typically implies a Wheatstone bridge configuration. The presence of an instrumentation amplifier also satisfies the amplification part of this requirement. (Requirement met)
4. An ADC should be used: The project summary specifies the use of SAR ADCs for both the pressure and temperature sensors. (Requirement met)
5. Infrared radiation sensors are being linearized: The summary does not provide explicit information on the linearization of the IR sensors. Therefore, this requirement is not met.
6. The solution mentions the sampling order strategy: The project summary does not specify the sampling order strategy (sequential, simultaneous, etc.). (Requirement not met)
7. The sampling frequency of the ADC is not less than 800 Hz: The project summary mentions a sampling rate of ≥ 800 Hz per channel for signals up to 400 Hz, which satisfies this requirement. (Requirement met)
8. The anti-aliasing filter has the required gain at half the sampling frequency: The low-pass filter for the pressure has a cutoff frequency of 200 Hz. With a sampling rate of 800 Hz, half the sampling frequency is 400 Hz. However, there's no explicit information on the gain at 400 Hz. Therefore, it's not possible to evaluate if this requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency for the pressure sensor's filter is 200 Hz, which does not meet the requirement of being higher than 400 Hz. (Requirement not met)
10. There are low-pass filters positioned before the multiplexer(s): The project includes low-pass filters for both pressure and temperature signals, but it's not explicitly stated that they are positioned before the multiplexers. (Requirement not met)
11. The project uses multiplexer(s) to choose channels: An 8:1 multiplexer is used in the project, satisfying this requirement. (Requirement met)
12. The multiplexer(s) are solid state: The CD74HC4051 is a CMOS logic multiplexer, which is a solid-state device. (Requirement met)

The requirements that are not met or cannot be evaluated due to lack of explicit information are 5, 6, 8, 9, and 10.