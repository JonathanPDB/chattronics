Verdict: unfeasible

Explanations: 
The project design provides a comprehensive approach to monitoring pressure and surface temperature variations on a machine. However, several critical requirements have not been explicitly met or clarified:

1. The design assumes DC output from both sensors without specifying this, which is a fatal requirement. Lack of explicit confirmation means we cannot be certain that demodulation is not necessary.
2. Amplification of both sensors is mentioned, with a gain of 5000 for the pressure sensor and unity gain for the temperature sensor, fulfilling this requirement.
3. The pressure sensor is to be amplified, but there's no explicit mention of a Wheatstone bridge or an instrumentation amplifier, which is a fatal requirement.
4. The use of an ADC is mentioned, fulfilling this requirement.
5. Linearization of infrared sensors is planned to be done digitally, meeting this requirement.
6. The sampling order strategy is not mentioned, which is not a fatal requirement, but it makes the design less comprehensive.
7. The minimum sampling frequency of the ADC is stated as at least 800 samples per second per channel, satisfying this requirement.
8. There is no mention of the anti-aliasing filter's attenuation at half the sampling frequency, so it's unclear if it meets the -20 dB requirement.
9. The low-pass filter's cutoff frequency is stated to be 400 Hz, which meets the requirement of being higher than 400 Hz but does not confirm if it is less than half the total sampling frequency.
10. Low-pass filters are mentioned, but their exact positioning in relation to the multiplexer(s) is not specified, which is crucial for anti-aliasing purposes.
11. The use of an 8-to-1 multiplexer is included, satisfying this requirement.
12. The multiplexer is specified as CMOS-based, which implies solid-state construction, meeting this requirement.

Due to the lack of explicit confirmation of DC sensor outputs, absence of a Wheatstone bridge and instrumentation amplifier for the pressure sensor, and incomplete information about the anti-aliasing filters and their placement, the project cannot be considered optimal. Some of these issues are fatal requirements, which significantly impacts the feasibility of the project as designed.