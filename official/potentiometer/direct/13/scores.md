Score: 7
Explanations: 
1. The potentiometer is used as a voltage divider - This is implied as it is being used to measure the angle of a pendulum and its output is used in further signal conditioning stages.
2. The voltage applied to the potentiometer is +/- 10 V - This is explicitly mentioned as the potentiometer should handle a power supply range from -10V to +10V.
3. The architecture is simple - The described system includes a potentiometer, level shifter (op-amp), low-pass filter, notch filters, and a DAQ, which corresponds to a simple architecture for the task.
4. The input voltage of the DAQ is centered in 0, for example, +/- 7V - The non-inverting op-amp level shifter circuit is used to center the voltage swing within the DAQ's input voltage range, which is +/- 7V.
5. The maximum voltage applied to the DAQ is 7V - The non-inverting amplifier configuration is proposed with a gain of 1.4 to scale the potentiometer's output voltage to within the DAQ's acceptable input range of +/- 7V.
6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - A 2nd order active low-pass Butterworth filter is used with a cutoff frequency just below 500 Hz to prevent aliasing.
7. There is a filter removing frequencies between around 50 and 60 Hz - Twin-T Notch Filter designs are chosen for 50 Hz and 60 Hz to attenuate power line interferences.
8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - This requirement is not explicitly covered. The anti-aliasing filter is mentioned to have a cutoff frequency just below 500 Hz, but there is no information on the gain of the signal at 500 Hz. Without specific values or a clear indication that it meets the -20 dB at 500 Hz requirement, we cannot assume it is met.

All requirements except for the 8th one seem to be met based on the information provided.