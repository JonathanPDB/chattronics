Score: 8
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider: The summary states a linear taper potentiometer is used, which is commonly implemented as a voltage divider.
2. The voltage applied to the potentiometer is +/- 10 V: The power supply for the Op-Amp in the buffer stage is ±10V, which implies the potentiometer is also working within this range.
3. The architecture includes a voltage divider, an anti-aliasing filter, and then the DAQ measures the voltage: The project includes all these elements.
4. The input voltage of the DAQ is centered at 0, +/- 7V: The summary specifies the DAQ's input voltage range as +/- 7V.
5. The maximum voltage applied to the DAQ is 7V: This is explicitly stated in the DAQ specifications.
6. There is a low-pass filter (anti-aliasing filter) that avoids aliasing: A second-order active low-pass Butterworth filter with a cutoff frequency of approximately 150 Hz is described, which would serve as an anti-aliasing filter.
7. There is a filter removing frequencies between around 50 and 60 Hz: The summary describes a Twin-T Notch Filter specifically designed to attenuate 50 Hz and 60 Hz frequencies.
8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The anti-aliasing filter is described to have an attenuation of -40 dB or more at the Nyquist frequency, which if the Nyquist frequency is twice the cutoff (300 Hz), then at 500 Hz the attenuation would certainly be more than -20 dB.

All 8 requirements are covered by the project summary.