Score: 7
Explanations: 
The project summary provides a detailed high-level architecture of a pendulum angle measurement system that incorporates a potentiometer, buffer, voltage divider, low-pass filter, notch filters, attenuator/level shifter, and a DAQ. Here is how it stacks up against the given requirements:

1. The potentiometer is used as a voltage divider - The potentiometer is indeed used to alter the voltage based on the pendulum's angle, implying its use as a voltage divider.
2. The voltage applied to the potentiometer is +/- 10 V - The summary specifies that the voltage range of the potentiometer is -10V to +10V.
3. The architecture includes a voltage divider (with buffer), an anti-aliasing filter, and then the DAQ measures the voltage - This is fulfilled as the summary describes a buffer, voltage divider, and a low-pass filter before the DAQ.
4. The input voltage of the DAQ is centered at 0, for example, +/- 7V - The voltage divider is specifically designed to scale the voltage to within the DAQ's range of +/-7V.
5. The maximum voltage applied to the DAQ is 7V - The use of a voltage divider to scale the voltage to +/-7V satisfies this requirement.
6. There is a low-pass filter (or anti-aliasing filter) that avoids aliasing - A second-order Sallen-Key low-pass filter with a cutoff frequency of 10 Hz is present, which would function as an anti-aliasing filter.
7. There is a filter removing frequencies between around 50 and 60 Hz - The project includes notch filters at 50 Hz and 60 Hz to attenuate these frequencies.
8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - The requirement is not explicitly met in the summary. The low-pass filter has a cutoff frequency of 10 Hz, but the order and damping ratio do not guarantee a -20 dB gain at 500 Hz. The notch filters attenuate at 50 Hz and 60 Hz, but there is no mention of the gain at 500 Hz.

Requirement 8 is not explicitly covered, and without additional information on the filter design, we cannot conclude that it ensures the specified gain at 500 Hz. All other requirements are met based on the information provided.