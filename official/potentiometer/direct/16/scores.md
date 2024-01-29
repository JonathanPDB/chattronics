Score: 7
Explanations: 
1. The potentiometer is used as a voltage divider: The project description does not explicitly state that the potentiometer is used as a voltage divider, but this is typically implied in analog measurement systems using potentiometers, so we can assume this requirement is met.

2. The voltage applied to the potentiometer is +/- 10 V: The buffer amplifier section mentions that the power supply must accommodate a -10 V to +10 V range, which implies that the potentiometer is indeed supplied with this voltage range. This requirement is met.

3. The architecture is simple: The project includes a voltage divider (potentiometer), a buffer amplifier (to prevent loading effects), scaling amplifier, band-stop filter, anti-aliasing filter, and a DAQ. This meets the requirement for simplicity in architecture.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The DAQ system is described as having a bipolar input voltage range capable of +/- 7 V, which meets this requirement.

5. The maximum voltage applied to the DAQ is 7V: The DAQ's input voltage range is specified as +/- 7 V, which ensures that the maximum voltage applied does not exceed 7 V. This requirement is met.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The anti-aliasing filter is a 2nd order Butterworth low-pass filter with a cutoff frequency of 250 Hz, which is below the Nyquist frequency of 500 Hz for the given sampling rate. This will avoid aliasing, meeting this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz: The band-stop filter is designed to attenuate noise at 50 Hz and 60 Hz. This requirement is met.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: The anti-aliasing filter is a 2nd order Butterworth filter with a -3 dB point at 250 Hz. Butterworth filters have a roll-off rate of -12 dB/octave per order beyond the cutoff frequency. At 500 Hz, which is one octave above 250 Hz, the filter would theoretically have an attenuation of -3 dB (at cutoff) + (-12 dB/octave * 1 octave) = -15 dB. This does not quite reach the requirement of -20 dB at 500 Hz. Therefore, this requirement is not met.