Verdict: incorrect

Explanations: 
The project design for the Analog Data Acquisition System meets many of the requirements, but there are some discrepancies that need to be addressed:

1. The project does not explicitly mention if demodulation is avoided, but since it does mention the use of sensors with a DC output and no demodulation circuitry is described, it can be inferred that this requirement is met.
2. Amplification for both sensors is clearly described, fulfilling this requirement.
3. The pressure sensor is not explicitly stated to be part of a Wheatstone bridge configuration, which is a requirement. However, strain-gauge pressure sensors typically imply a bridge configuration, so this may be an oversight in the description rather than an omission in the design.
4. An ADC is included in the design, meeting this requirement.
5. The linearization of the infrared radiation sensors is not described. This is a requirement and is critical for accurate temperature measurement, so this is a significant omission.
6. The sampling order strategy is not mentioned. It is unclear whether the ADC will sample the sensors sequentially, simultaneously, or using another strategy.
7. The ADC sampling rate meets the minimum requirement of 800 Hz, as it is stated to be at least 2 kSPS per channel.
8. The anti-aliasing filter for the pressure sensor has a cutoff frequency of 800 Hz, which is exactly half of the sampling frequency. This could potentially allow aliasing, as the requirement is for the gain of the signal to be at least -20 dB at half the sampling frequency. Without further specifications on the filter's response curve, this might not be sufficient.
9. The low-pass filter for the temperature sensor has a cutoff frequency of 10 Hz, which is lower than the required 400 Hz minimum.
10. It is unclear if there are low-pass filters positioned before the multiplexer, as this is not explicitly stated.
11. The project uses multiplexers to choose channels, meeting this requirement.
12. The multiplexers are solid state, which satisfies this requirement.

Overall, the project design has several strengths but fails to explicitly meet some of the necessary requirements, particularly regarding the linearization of infrared sensors, the Wheatstone bridge configuration for the pressure sensors, the anti-aliasing filter's effectiveness, and the low pass filter's cutoff frequency for the temperature sensor.