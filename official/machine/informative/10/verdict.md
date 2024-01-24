Verdict: incorrect

Explanations: 
The project summary outlines a monitoring system for pressure and temperature using analog sensors, with dedicated signal conditioning, amplification, and analog-to-digital conversion stages. However, the project falls short of meeting several of the specified requirements:

1. The summary does not mention the need to avoid demodulation for d.c. output sensors. However, since both sensors (pressure and temperature) are assumed to have d.c. outputs, this is implicitly met.
2. Both sensor types are being amplified, which satisfies this requirement.
3. The pressure sensor is amplified by an instrumentation amplifier, but there is no explicit mention of its insertion in a Wheatstone bridge, which is a fatal issue.
4. An ADC is used for both sensor types, meeting this requirement.
5. The summary does not specify the linearization technique for the infrared radiation sensors. The requirement for linearization is essential, and its omission is a fatal flaw.
6. The summary does not mention the sampling order strategy (sequential, simultaneous, etc.).
7. The sampling frequency for the pressure ADC is at least 2 kHz, satisfying the requirement. However, the temperature ADC has a range of 1 kHz to 10 kHz, which could potentially be below the required 800 Hz if set to the lower end, but if set to 1 kHz or higher, it would be compliant.
8. The anti-aliasing filter for pressure has a cutoff frequency of 500 Hz, which does not explicitly meet the -20 dB gain at half the sampling frequency requirement. This is a critical issue, as the exact attenuation at the Nyquist frequency is not provided.
9. The low-pass filter for temperature has a cutoff frequency of 400 Hz, which is just at the lower limit of the requirement but does not guarantee that it is less than half the total sampling frequency.
10. The summary does not specify if the low-pass filters are positioned before the multiplexer(s), which is essential to reduce aliasing.
11. The use of multiplexers is not mentioned in the summary, which is a fatal flaw since it is a necessary requirement.
12. Even if multiplexers were included, there is no mention that they are solid-state, which is also a requirement.

Due to the project not meeting several essential requirements, particularly the absence of a Wheatstone bridge for the pressure sensor, unspecified linearization techniques for infrared sensors, lack of detail on the sampling strategy and positioning of filters, and no mention of multiplexers, the verdict for this project is "incorrect."