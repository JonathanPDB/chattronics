Score: 8
Explanations: 
1. The potentiometer is used as a voltage divider: The potentiometer is mentioned as the primary sensor with a voltage range of -10V to +10V, indicating it is used as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: This is explicitly stated in the Potentiometer Sensor Block description.

3. The architecture includes a voltage divider, an anti-aliasing filter, and then the DAQ measures the voltage: The provided summary includes a potentiometer (voltage divider), scaling amplifier (which can act as a voltage buffer), band-stop filter, anti-aliasing filter, and a DAQ, satisfying this requirement.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The scaling amplifier block is designed to scale the voltage from -10V/+10V to -7V/+7V, which matches this requirement.

5. The maximum voltage applied to the DAQ is 7V: The scaling amplifier block ensures that the voltage does not exceed -7V/+7V, fulfilling this requirement.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: An anti-aliasing filter with a 4th order Butterworth response and a cutoff frequency of 100 Hz is specified, which meets this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz: A band-stop filter block is mentioned with suggestions for both passive and active topologies to attenuate 50 Hz and 60 Hz, which satisfies this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: Although a 4th order Butterworth low pass filter at 100 Hz is specified, there is no explicit information on the gain at 500 Hz. A 4th order Butterworth filter has a slope of -24 dB/octave beyond the cutoff frequency. Since 500 Hz is approximately two octaves above 100 Hz, the gain will be -48 dB at 500 Hz, which is well beyond the -20 dB requirement. This requirement is thus met.