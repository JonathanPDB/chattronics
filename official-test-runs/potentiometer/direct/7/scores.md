Score: 7
Explanations: 
The project summary addresses several of the required items for the electronics project:

1. The potentiometer is used as a voltage divider: The summary does not explicitly state that the potentiometer is used as a voltage divider, but it implies it by stating that the voltage varies linearly with the angle, which is typical for a voltage divider application.
2. The voltage applied to the potentiometer is +/- 10V: The power supply for the operational amplifiers is +/- 10V, which suggests that this is also the voltage applied to the potentiometer.
3. The architecture includes a voltage divider, an anti-aliasing filter, and DAQ measurement: The summary describes a voltage follower (buffer), a gain stage, a notch filter, a low-pass filter, and a DAQ, which meets the requirement for a simple architecture.
4. The input voltage of the DAQ is centered at 0, +/- 7V: The gain stage is designed to output a voltage range of +/- 7V, which matches the DAQ's maximum input voltage specification.
5. The maximum voltage applied to the DAQ is 7V: This is explicitly stated in the DAQ section as the maximum input voltage.
6. There is a low-pass filter that avoids aliasing: The anti-aliasing filter is a Sallen-Key low-pass filter with a cutoff frequency of approximately 210 Hz, which is suitable for the given sampling rate of 1000 samples per second.
7. There is a filter removing frequencies between around 50 and 60 Hz: The band-pass filter section describes an active twin-T notch filter topology specifically for attenuating 50 Hz and 60 Hz frequencies.
8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The low-pass filter has a cutoff frequency of 200 Hz, which would result in a gain of less than -20 dB at 500 Hz due to the roll-off rate of -12 dB/octave.

All the requirements seem to be met based on the provided summary, except for the explicit mention of the potentiometer being used as a voltage divider, which is inferred rather than directly stated.