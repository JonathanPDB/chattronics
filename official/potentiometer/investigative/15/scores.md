Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - The description includes a potentiometer with specific characteristics, implying its use as a voltage divider in the context of the pendulum angle measurement.

2. The voltage applied to the potentiometer is +/- 10 V - While the exact voltage applied to the potentiometer is not specified, the design choices are made with consideration of a power supply between -10V and +10V, suggesting that the potentiometer is indeed operating within this range.

3. The architecture includes a voltage divider (with or without a voltage buffer), an anti-aliasing filter, and then the DAQ measures the voltage - The design incorporates a potentiometer (voltage divider), a buffer stage, and an anti-aliasing filter before the DAQ.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V - The buffer stage is configured with a gain of 0.7 to scale the potentiometer's ±10V output to the DAQ's ±7V input range.

5. The maximum voltage applied to the DAQ is 7V - The buffer stage ensures that the scaled voltage does not exceed ±7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - The anti-aliasing filter is present with a cutoff frequency just above 500 Hz, which meets the requirement to avoid aliasing.

7. There is a filter removing frequencies between around 50 and 60 Hz - The low-pass filter with a cutoff frequency of 5 Hz and an attenuation of ≥ 60 dB at 50 Hz effectively removes frequencies in the 50-60 Hz range.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - This requirement is not explicitly confirmed. The anti-aliasing filter has a cutoff frequency just above 500 Hz, but there is no information on the attenuation at 500 Hz, so it's not possible to confirm if it meets the -20 dB gain requirement at that frequency.

Overall, the project meets 7 out of the 8 listed requirements. The only requirement that is not explicitly confirmed is the filter's attenuation at 500 Hz.