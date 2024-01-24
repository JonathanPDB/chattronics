Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - The summary mentions a precision potentiometer which is typically used as a voltage divider in such applications.
2. The voltage applied to the potentiometer is +/- 10 V - The power supply for the potentiometer is specified to be between -10V and +10V.
3. The architecture is simple: there is a voltage divider (with or without a voltage buffer), an anti-aliasing filter, and then the DAQ measures the voltage - The summary describes a voltage follower (buffer), low-pass filters, and an ADC interfaced with a DAQ.
4. The input voltage of the DAQ is centered in 0, for example, +/- 7V - No explicit information is given about the centering of the DAQ input voltage.
5. The maximum voltage applied to the DAQ is 7V - The ADC resolution is specified for a +/- 7V input range, which implies that the maximum voltage applied to the DAQ is indeed 7V.
6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - An active second-order Butterworth low-pass filter with a cutoff frequency of 100 Hz is described, which would serve as an anti-aliasing filter.
7. There is a filter removing frequencies between around 50 and 60 Hz - The summary specifies a low-pass filter with a cutoff frequency of approximately 100 Hz, which would attenuate 50 and 60 Hz noise.
8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz - The anti-aliasing filter is described to have a stopband attenuation of > 60 dB beyond 500 Hz, which meets the requirement.

The requirements that are not explicitly covered in the summary are:

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V - Since this is not mentioned, we cannot assume it is met.

Given the information provided and the assumptions that can be made based on standard practices in analog electronics design, the project covers 7 out of 8 requirements.