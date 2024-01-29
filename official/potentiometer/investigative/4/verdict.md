Verdict: incorrect

Explanations: 
The project description provides a detailed approach to measuring the angle of a pendulum with a linear potentiometer and processing the signal for a DAQ system. Here are the evaluations based on the requirements:

1. The potentiometer is used as a voltage divider, which meets the first requirement.

2. The voltage applied to the potentiometer is +/- 10 V, aligning with the second requirement.

3. The architecture includes a voltage buffer, anti-aliasing filter, and a DAQ measures the voltage, which is consistent with the third requirement.

4. The input voltage of the DAQ is centered at 0, with a range of +/- 7V, meeting the fourth requirement.

5. The maximum voltage applied to the DAQ is 7V, which is stated as an absolute necessity, and the design respects this limit, thus fulfilling the fifth requirement.

6. An anti-aliasing filter is mentioned to avoid aliasing, but the cutoff frequency is stated to be just below 500 Hz. This is problematic because the sampling rate is 1000 samples per second, and according to the Nyquist theorem, the cutoff frequency should be below half the sampling rate to avoid aliasing. This could potentially be a non-fatal issue if the actual cutoff frequency is below 500 Hz, but the description is unclear.

7. The design incorporates a low-pass filter to remove frequencies around 50 and 60 Hz, satisfying the seventh requirement.

8. The filter has a cutoff frequency of 40 Hz, which is well below 500 Hz, and it should provide at least -20 dB attenuation at 500 Hz, meeting the eighth requirement.

However, the cutoff frequency for the anti-aliasing filter is not explicitly stated to be below half of the sampling rate (500 Hz) and this might not provide sufficient attenuation to avoid aliasing. This is a critical aspect that needs clarification. If the cutoff frequency is indeed below 500 Hz, then it would meet the requirements.

Additionally, the voltage divider ratio chosen does not seem to be correct. To scale a +/- 10V signal to +/- 7V, the voltage divider must reduce the voltage by 30%, not scale it with the given ratio of R1 = 30kΩ and R2 = 70kΩ, which would actually increase the voltage. This is a fatal issue because it violates the requirement that the maximum voltage applied to the DAQ must not exceed 7V.