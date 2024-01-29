Score: 8
Explanations: 
The project summary addresses several of the requirements but does not explicitly cover all of them:

1. Both sensors have DC output: The pressure sensor typically has a DC output, and the temperature sensor's output interface is digital (I2C), which implies DC characteristics. Thus, no demodulation is needed. (Requirement met)
2. Both sensors must be amplified: Gain calculations are provided for both the pressure and temperature sensors, indicating amplification is designed into the system. (Requirement met)
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The use of an instrumentation amplifier for pressure sensors is mentioned, which suggests a Wheatstone bridge configuration, but the bridge itself is not explicitly stated. (Requirement implicitly met)
4. An ADC should be used: The use of an ADC is implied in the context of gain calculations and the mention of ADC resolution. (Requirement implicitly met)
5. Infrared radiation sensors are being linearized: Linearization is mentioned to be achieved using Look-Up Table or Polynomial Approximation in software within the microcontroller. No mention of diode networks or log amplifiers is made. (Requirement met)
6. Sampling order strategy: No specific mention of sampling order strategy (sequential, simultaneous, etc.) is made. (Requirement not met)
7. Sampling frequency of the ADC: A switching speed of the multiplexer supports a sampling rate of at least 1 kSps per channel, which implies a sampling frequency of at least 1000 Hz, meeting the 800 Hz minimum requirement. (Requirement met)
8. Anti-aliasing filter: A second-order Butterworth filter is suggested with a cutoff frequency set at or slightly below 400 Hz. This does not meet the requirement of at least -20 dB at half the sampling frequency, as the sampling frequency is at least 1000 Hz, and half of that would be 500 Hz. (Requirement not met)
9. Low pass cutoff frequency: The cutoff frequency is set at or slightly below 400 Hz, which meets the requirement of being higher than 400 Hz but does not meet the requirement of being lower than half the total sampling frequency (which is at least 1000 Hz, so the cutoff should be below 500 Hz). (Requirement not met)
10. Low-pass filters before the multiplexer: The presence of low-pass filters is mentioned, but their position relative to the multiplexer is not stated. (Requirement not explicitly met)
11. The project uses multiplexer(s): An 8:1 multiplexer is suggested, meeting the requirement. (Requirement met)
12. The multiplexer(s) are solid state: The ADG1608 is a solid-state multiplexer. (Requirement met)

Overall, the project summary sufficiently covers 8 out of the 12 requirements. Specific details regarding the sampling order strategy, the exact configuration of the anti-aliasing filter relative to the sampling frequency, the position of the low-pass filters, and the explicit mention of a Wheatstone bridge for the pressure sensor are either not met or not clearly stated.