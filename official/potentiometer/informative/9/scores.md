Score: 6
Explanations: 
The project summary addresses the following requirements:

1. The potentiometer is used as a voltage divider - The summary mentions the use of a potentiometer in a voltage divider configuration.
2. The voltage applied to the potentiometer is +/- 10 V - The buffer amplifier has a power supply range of -10V to +10V, which implies the voltage applied could be within this range.
3. The architecture is simple - The summary outlines a simple architecture consisting of a voltage divider, buffer amplifier, anti-aliasing filter, and DAQ system.
4. The input voltage of the DAQ is centered in 0 - The voltage divider is adjustable to set the midpoint at 0 V for the 90-degree steady position, which meets the requirement.
5. The maximum voltage applied to the DAQ is 7V - The summary does not explicitly state that the voltage applied to the DAQ will be limited to a maximum of 7V. There is a risk that the potentiometer's output, when scaled down, could exceed this value, especially since the operating range of the potentiometer is expected to be from -3.33V to 3.33V, which does not cover the full +/- 10V range. Therefore, this requirement is not verifiably met based on the provided information.
6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - The summary describes a 2nd-order active low-pass filter with a cutoff frequency of 100 Hz, which would serve as an anti-aliasing filter.
7. There is a filter removing frequencies between around 50 and 60 Hz - The summary states that the anti-aliasing filter provides attenuation at 50/60 Hz of >20 dB, which meets this requirement.
8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - The summary does not provide sufficient information to evaluate the attenuation at 500 Hz. The roll-off rate is 12 dB/octave, but without the exact filter design, it's unclear if the -20 dB requirement at 500 Hz is met.

Requirements 5 and 8 are not clearly met based on the information provided. The remaining requirements appear to be addressed adequately by the project summary.