Score: 6
Explanations: 
The project summary covers the following items based on the provided information:

1. The potentiometer is used as a voltage divider (explicitly mentioned).
2. The voltage applied to the potentiometer is +/- 10 V (explicitly mentioned).
3. The architecture includes a voltage divider (the potentiometer), an op-amp (voltage buffer), an anti-aliasing filter, and a DAQ system (implicitly fulfilled by the summary of the components and their functions).
4. The input voltage of the DAQ is centered in 0, for example, +/- 7V. The amplifier has a gain of 1.4, which implies that for a 5V input from the potentiometer, the output would be 7V. However, the maximum input voltage range for the DAQ is not explicitly mentioned, so we cannot confirm this requirement is fully met.
5. The maximum voltage applied to the DAQ is 7V (cannot be confirmed as the DAQ input voltage range is not stated).
6. There is a low pass filter (Sallen-Key Butterworth Low-Pass Filter) that avoids aliasing (explicitly mentioned).
7. There is a filter removing frequencies between around 50 and 60 Hz (explicitly mentioned with separate notch filters for 50 Hz and 60 Hz).
8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz. Given that the anti-aliasing filter has a cutoff frequency of 150 Hz and is a 2nd order Butterworth, it would have a slope of -40 dB/decade. At 500 Hz, which is roughly one decade above the cutoff, the attenuation should be about -40 dB, fulfilling this requirement.

Requirements not fully met:
4. The input voltage of the DAQ is centered in 0, for example, +/- 7V (not explicitly confirmed).
5. The maximum voltage applied to the DAQ is 7V (not explicitly confirmed).