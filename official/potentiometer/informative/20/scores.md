Score: 7
Explanations: 
The project overview provides information about the design and components involved in the pendulum angle measurement system. The following points assess how the project aligns with the given requirements:

1. The potentiometer is used as a voltage divider. This is implicit in the description of the linear potentiometer used to measure the angle of the pendulum.
2. The voltage applied to the potentiometer is +/- 10 V. This is not explicitly stated but is implied by the scaling amplifier's gain, which scales the voltage down to +/- 7 V for the DAQ.
3. The architecture includes a voltage divider (potentiometer), a voltage buffer (buffer amplifier), an anti-aliasing filter, and a notch filter before the DAQ measures the voltage.
4. The input voltage of the DAQ is centered at 0, with a range of +/- 7V. This is achieved by the scaling amplifier's gain of 0.7.
5. The maximum voltage applied to the DAQ is 7V. This is ensured by the scaling amplifier, which scales the potentiometer's +/- 10 V down to +/- 7 V.
6. There is an anti-aliasing filter that avoids aliasing, described as a second-order Butterworth low-pass filter with a cutoff frequency of around 100 Hz.
7. There is a filter removing frequencies around 50 and 60 Hz, which is achieved by the notch filters described.
8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz. The anti-aliasing filter is a second-order Butterworth low-pass with a cutoff frequency of around 100 Hz, which may not satisfy the -20 dB requirement at 500 Hz without specific values for the components.

The requirements regarding the cutoff frequency and order of the filter to achieve at least -20 dB at 500 Hz are not fully met, as there are no specific values for the filter components provided to calculate the actual attenuation at 500 Hz. Therefore, this requirement is not verifiable based on the information provided, and it is assumed not met.

All other requirements are either explicitly met or can be reasonably inferred from the information given.

Requirement 8 is the only one that is not clearly met, as the exact attenuation at 500 Hz cannot be determined without further information.