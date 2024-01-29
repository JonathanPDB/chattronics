Score: 7
Explanations: 
The project summary addresses the following items:

1. The potentiometer is used as a voltage divider. This is explicitly stated in the description of the potentiometer section.

2. The voltage applied to the potentiometer is +/- 10 V. This is implied as the standard range for a potentiometer, and the attenuator section mentions scaling from -10V to +10V down to -7V to +7V, implying the initial range is indeed +/- 10 V.

3. The architecture is simple with a voltage divider, anti-aliasing filter, and DAQ. The system description outlines a voltage divider (potentiometer and attenuator), a buffer, a notch filter, an anti-aliasing filter, and a DAQ, which fits the requirement.

4. The input voltage of the DAQ is centered at 0, with a range of +/- 7V. This is explicitly stated in the DAQ section.

5. The maximum voltage applied to the DAQ is 7V. This is confirmed by the attenuator design meant to scale the voltage down to +/- 7V and the DAQ's input voltage range.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing. The anti-alias filter description specifies an active Sallen-Key Low-Pass Filter with a cutoff frequency of 400 Hz, which is appropriate for preventing aliasing.

7. There is a filter removing frequencies around 50 and 60 Hz. The notch filter section describes a Twin-T Notch Filter designed to attenuate both 50 Hz and 60 Hz.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz. The anti-aliasing filter is a second-order filter with a cutoff frequency of 400 Hz. A second-order filter has a roll-off rate of 12 dB/octave, which means that at 500 Hz (1 octave above 400 Hz), the attenuation would be 12 dB. This does not meet the requirement of at least -20 dB at 500 Hz.

The project meets 7 out of the 8 requirements. The only requirement not met is number 8, as the anti-aliasing filter does not provide enough attenuation at 500 Hz.