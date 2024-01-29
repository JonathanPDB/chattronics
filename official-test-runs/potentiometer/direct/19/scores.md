Score: 7
Explanations: 
The project summary covers the following requirements:

1. The use of a potentiometer as a voltage divider is implicit in the description of the sensitivity of the potentiometer, which suggests that it is being used to provide a voltage output proportional to the angle of the pendulum.

2. The voltage applied to the potentiometer is not explicitly stated to be +/- 10V. However, the sensitivity is approximately 0.222 V/degree based on a voltage range of 20V, which implies that the potentiometer may be operating over a range of +/- 10V. Without explicit confirmation, this requirement is not clearly met.

3. The architecture includes a voltage divider, an offset adjustment, an amplifier block, and filters, which is relatively simple and meets the requirement.

4. The input voltage of the DAQ is centered at 0V for the 90 degrees position, which meets the requirement.

5. The maximum voltage applied to the DAQ is 7V, as the amplifier gain is set to 0.7 to scale the signal to fit within the DAQ's input range.

6. There is an anti-aliasing filter with a cutoff frequency of 450 Hz, which serves to avoid aliasing as required.

7. The band-stop filter is designed to attenuate frequencies around 50 and 60 Hz, meeting this requirement.

8. The cutoff frequency of the anti-aliasing filter is 450 Hz, which is close to the requirement of at least -20 dB at 500 Hz. Assuming the filter has at least a second-order response, it is likely to meet the -20 dB attenuation at 500 Hz due to the roll-off rate of 12 dB/octave or better. However, without explicit information on the order and actual attenuation at 500 Hz, this requirement is not clearly met.

Overall, the project summary successfully covers 7 out of the 8 requirements, with requirement 2 being ambiguous due to lack of explicit information and requirement 8 not being clearly met without additional data on the filter's performance at 500 Hz.