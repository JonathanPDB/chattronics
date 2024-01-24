Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - The summary mentions the use of a potentiometer to vary linearly with the angle of the pendulum, which implies its use as a voltage divider.
2. The voltage applied to the potentiometer is +/- 10 V - The summary does not explicitly state the voltage applied to the potentiometer, but it mentions a -10V to +10V output range, which suggests the correct voltage is applied.
3. The architecture is simple - The summary describes a simple architecture including a voltage divider, a buffer, a low-pass filter, and the DAQ.
4. The input voltage of the DAQ is centered in 0, +/- 7V - The voltage divider is designed to scale the potentiometer's output to the DAQ's -7V to +7V range.
5. The maximum voltage applied to the DAQ is 7V - The voltage divider ensures that the maximum voltage applied to the DAQ is 7V.
6. There is a low-pass filter (or anti-aliasing filter) - A second-order active low-pass Butterworth filter with a cutoff frequency of 40 Hz is implemented as an anti-aliasing filter.
7. There is a filter removing frequencies between around 50 and 60 Hz - The low-pass filter is designed to attenuate 50 Hz and 60 Hz noise specifically, which meets this requirement.
8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz - The requirement is for a filter that provides at least -20 dB attenuation at 500 Hz, but the summary only mentions a cutoff frequency of 40 Hz without specifying the attenuation at 500 Hz. Therefore, it is not possible to confirm if this requirement is met without additional information about the filter design.

The requirements that are not explicitly met or cannot be confirmed with the provided information are:

- The filter's attenuation at 500 Hz is not specified, so we cannot confirm if it meets the -20 dB attenuation requirement.
