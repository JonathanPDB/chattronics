Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider: The summary specifies a Bourns 3590S 10 kOhm linear potentiometer, which is typically used in a voltage divider configuration.

2. The voltage applied to the potentiometer is +/- 10 V: The buffer stage mentions a dual power supply of +/-10V, which implies that this is the voltage applied to the potentiometer.

3. The architecture is simple: The summary describes a simple architecture including a buffer stage (voltage follower), a band-stop filter, an anti-aliasing filter, and a differential amplifier, which meets the requirement of having a voltage divider, anti-aliasing filter, and DAQ measurement setup.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The differential amplifier is designed with a gain of 1.4 to map the -5V to +5V potentiometer output to -7V to +7V for the DAQ input.

5. The maximum voltage applied to the DAQ is 7V: As mentioned above, the differential amplifier ensures that the output does not exceed +/- 7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: A second-order Sallen-Key Low-Pass Filter with a cutoff frequency of 200 Hz is included, which serves as an anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz: The active Twin-T Notch Filter is designed to attenuate 50 Hz and 60 Hz by -40 dB, which fulfills this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: This requirement is not explicitly met. The anti-aliasing filter has a cutoff frequency of 200 Hz, but the order and the actual gain at 500 Hz are not provided. A second-order filter typically has a roll-off rate of -12 dB/octave past the cutoff frequency. To meet the requirement of at least -20 dB at 500 Hz, we would need to calculate the actual gain based on the filter design, which is not provided in the summary. Therefore, we cannot confirm if this requirement is met.

The summary does not provide explicit information on the order of the anti-aliasing filter and the gain at 500 Hz to meet the -20 dB requirement. Therefore, this requirement is not confirmed to be met.