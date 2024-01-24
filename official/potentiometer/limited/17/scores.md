Score: 7
Explanations: 
1. The potentiometer is used as a voltage divider: The summary does not explicitly state that the potentiometer is used as a voltage divider, but given the context of the design, it is implied that the potentiometer's function is to vary the voltage based on the pendulum's angle, which is characteristic of a voltage divider configuration.

2. The voltage applied to the potentiometer is +/- 10 V: This requirement is met as the buffer amplifier is powered by the same -10 to +10 volts source as the potentiometer, implying that the potentiometer is also supplied with +/- 10V.

3. The architecture is simple: The design includes a voltage divider (potentiometer), a buffer, scaling amplifier, notch filter, low-pass filter, and DAQ. This is a simple and typical signal conditioning architecture for such applications, thus meeting this requirement.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The DAQ's input voltage range is specified to be at least +/- 7V, which meets this requirement.

5. The maximum voltage applied to the DAQ is 7V: This requirement is met as the DAQ's input voltage range is specified to be at least +/- 7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The design includes a second-order active low-pass Butterworth filter with a cutoff frequency of 200 Hz, which would act as an anti-aliasing filter, thus meeting this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz: A specific notch filter is included in the design with notch frequencies centered at 50 Hz and 60 Hz, with an attenuation of -40 dB, meeting this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: The cutoff frequency of the low-pass filter is 200 Hz, and being a second-order filter, it has an attenuation rate of -12 dB per octave. To calculate the attenuation at 500 Hz, we must note that 500 Hz is a little more than one octave above 200 Hz (one octave is a doubling in frequency). The attenuation would be slightly more than -12 dB (one octave above the cutoff) but well below -24 dB (two octaves above). Therefore, it does not meet the -20 dB requirement at 500 Hz.