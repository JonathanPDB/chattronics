Score: 8
Explanations: 
The project summary provides details on the electronic system designed to measure the angle of a pendulum. Based on the requirements provided, the following items are covered:

1. The potentiometer is used as a voltage divider: The summary explicitly states the use of a precision linear potentiometer for angle measurement, which implies its use as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: The potentiometer's voltage range is clearly stated as -10V to +10V.

3. The architecture is simple: The summary describes a simple architecture consisting of a voltage divider, offset adjustment, gain stage, low pass filter, notch filter, anti-aliasing filter, and DAQ system.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The summary confirms the DAQ input voltage range is ±7V.

5. The maximum voltage applied to the DAQ is 7V: The gain stage is designed to scale the -5V to 5V signal to -7V to 7V, ensuring the maximum voltage applied to the DAQ is 7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The summary describes two low-pass filters, one with a cutoff frequency of 100 Hz and one with a cutoff frequency of 250 Hz, which serve as anti-aliasing filters.

7. There is a filter removing frequencies between around 50 and 60 Hz: The summary includes an active multiple feedback notch filter designed to target and remove frequencies at 50 Hz and 60 Hz.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: Although the summary mentions a second-order Butterworth filter with a cutoff frequency of 100 Hz, the actual attenuation at 500 Hz is not stated. However, a second-order Butterworth filter typically has a roll-off rate of -12 dB/octave. Since 500 Hz is more than one octave above the cutoff frequency, the attenuation will be more than -24 dB, thus satisfying the requirement of at least -20 dB at 500 Hz.

All eight requirements are successfully covered by the given project summary.