Verdict: acceptable

Explanations: 
The project overview describes a pendulum angle measurement system that uses a linear potentiometer for sensing, followed by signal conditioning circuits, and a DAQ system for data acquisition. The requirements specified in the question must be compared against the project details provided:

1. The potentiometer is used as a voltage divider: The project uses a linear potentiometer, which is consistent with the requirement for a voltage divider.
2. The voltage applied to the potentiometer is +/- 10 V: The project description does not specify the voltage applied to the potentiometer. This information is critical to ensure the output voltage does not exceed the DAQ's limits.
3. Simple architecture with a voltage divider, anti-aliasing filter, and DAQ measurement: The project includes a voltage divider, buffer amplifier, scaling amplifier, filtering, and a DAQ system. This is somewhat more complex than the simple architecture described in the requirement but still meets the basic intent.
4. The input voltage of the DAQ is centered at 0, for example, +/- 7V: The scaling amplifier is set to achieve a gain of 0.7 to scale the potentiometer's voltage output to +/- 7V, which meets this requirement.
5. The maximum voltage applied to the DAQ is 7V: The gain of the scaling amplifier suggests that the maximum voltage to the DAQ will be 7V, provided that the input voltage to the potentiometer does not exceed +/- 10V, which is not verified in the description.
6. Low pass filter (anti-aliasing filter) to avoid aliasing: The project includes an active second-order Butterworth low-pass filter with a cutoff frequency of around 100 Hz, which should help to avoid aliasing.
7. Filter removing frequencies between around 50 and 60 Hz: The project includes notch filters for 50 and 60 Hz, which meet this requirement.
8. Filter gain of at least -20 dB at 500 Hz: The project specifies >20 dB attenuation at the notch frequencies, but it does not provide information on the gain at 500 Hz. The cutoff frequency for the anti-aliasing filter is around 100 Hz, which implies that by 500 Hz, the gain should be significantly below -20 dB, but this is not explicitly confirmed.

The project description is mostly aligned with the requirements. However, it lacks specific confirmation of the applied voltage to the potentiometer and the filter performance at 500 Hz. These are critical points to ensure compliance with the maximum voltage to the DAQ and the filter performance at higher frequencies.