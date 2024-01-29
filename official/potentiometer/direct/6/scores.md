Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - The use of a linear potentiometer with an operating voltage of -10V to +10V indicates that it is being used as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V - The operating voltage of the potentiometer is specified as -10V to +10V.

3. The architecture is simple - The summary describes a straightforward architecture consisting of a voltage buffer, amplifier, notch filters, low-pass filter, and DAQ.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V - The DAQ's input voltage range matches the output voltage range of +/-7V from the signal conditioning stages.

5. The maximum voltage applied to the DAQ is 7V - The summary explicitly states that the DAQ's input voltage range is +/-7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - The summary describes an active Butterworth low-pass filter with a cutoff frequency of 450 Hz to prevent aliasing.

7. There is a filter removing frequencies between around 50 and 60 Hz - The design includes active notch filters with notch frequencies set at 50 Hz and 60 Hz, respectively.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - This requirement is not explicitly confirmed, as the anti-aliasing filter is specified with a cutoff frequency of 450 Hz. Without the exact filter design, we cannot confirm that it provides at least -20 dB attenuation at 500 Hz. However, a 2nd-order Butterworth filter typically has a roll-off of 12 dB per octave, and since 500 Hz is close to the cutoff frequency, it is unlikely that the attenuation is -20 dB at 500 Hz. Therefore, this requirement is not met.

The requirement regarding the gain of the signal at 500 Hz is not explicitly met, and without further information, we cannot assume it is achieved. Therefore, the score is 7 out of 8.