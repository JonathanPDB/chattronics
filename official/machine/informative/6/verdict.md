Verdict: acceptable

Explanations: 
The project description provided does not explicitly state that demodulation is not used, which is a critical requirement. However, since both sensors are described as having a d.c. output, it can be inferred that demodulation is not part of the design, fulfilling requirement 1.

Amplification is mentioned for both the pressure and temperature sensors, satisfying requirement 2.

The pressure sensor is described to be amplified by an instrumentation amplifier, which is in line with requirement 3.

An ADC with a sampling rate of 12.8 kHz is specified, which exceeds the minimum 800 Hz per channel, satisfying requirement 4 and 7.

Linearization of infrared radiation sensors is not explicitly mentioned, so requirement 5 cannot be confirmed to be met.

The summary does not detail the sampling order strategy (sequential or simultaneous), which does not fulfill requirement 6.

The anti-aliasing filter is described with a cutoff frequency of approximately 500 Hz for the pressure sensor and around 1 Hz for the temperature sensor, but it does not specifically address the requirement of having at least -20 dB gain at half the sampling frequency, leaving requirement 8 unverified.

The low-pass cutoff frequency is stated to be set at 500 Hz for the pressure sensor, which is higher than 400 Hz but not explicitly less than half the total sampling frequency, which is required to determine if requirement 9 is met.

Requirement 10 is addressed with the inclusion of low-pass filters, but their position relative to the multiplexer(s) is not explicitly stated.

Multiplexers are mentioned in the design, satisfying requirement 11, and they are described as solid state, which meets requirement 12.