Score: 11
Explanations: 
The project summary covers the following requirements:

1. The pressure sensors are strain-gauge sensors with a DC output, as required.
2. Amplification is mentioned for both the pressure sensors, using instrumentation amplifiers, and for the temperature sensors, even though the latter are digital, a buffering stage is considered which can be seen as a form of amplification for the digital signals to ensure proper levels for the multiplexer and ADC.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier, which is implied since strain-gauge pressure sensors typically are used in a Wheatstone bridge configuration, and the use of instrumentation amplifiers is explicitly mentioned.
4. An ADC is used, with a type and sampling rate specified (SAR ADC with 1 kSps per channel).
5. Infrared radiation sensors are linearized using a second-order polynomial correction, which satisfies the requirement of linearization.
6. The sampling order strategy is not explicitly mentioned, so this requirement is not verifiable from the provided summary.
7. The sampling frequency of the ADC is stated as 1 kSps (1000 samples per second), which is above the 800 Hz minimum requirement.
8. The anti-aliasing filter is specified to have a stop-band attenuation of -40 dB at 800 Hz, which satisfies the requirement of at least -20 dB at half the sampling frequency (400 Hz in this case).
9. The low-pass filter has a cutoff frequency of 500 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (500 Hz is less than 500 Hz, which would be half of 1000 Hz sampling frequency), satisfying this requirement.
10. The presence of low-pass (anti-aliasing) filters is mentioned, and they are positioned before the multiplexer, as required.
11. The project uses multiplexers to choose channels, with specific models suggested.
12. The multiplexers mentioned (CD74HC4051 and ADG726) are solid-state devices.

The following requirements are not covered or not clearly addressed in the summary:

6. The sampling order strategy is not mentioned.