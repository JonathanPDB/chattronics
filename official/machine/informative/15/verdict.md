Verdict: acceptable

Explanations: 
The project design for the Pressure and Temperature Monitoring System for Machinery is largely in line with the requirements, but some discrepancies need to be addressed:

1. Both sensors are assumed to have d.c. output as there is no mention of demodulation, which satisfies requirement 1.
2. Amplification is clearly planned for both sensors, meeting requirement 2.
3. The pressure sensor is conditioned using an instrumentation amplifier, which fulfills requirement 3.
4. An ADC with a specified sampling rate is included in the design, complying with requirement 4.
5. The temperature sensor's signal conditioning includes linearization, but the method is not specified as digital, diode networks, or log amplifiers, which leaves requirement 5 ambiguous.
6. The project does not mention the sampling order strategy (sequential, simultaneous, etc.), leaving requirement 6 unaddressed.
7. The ADC's sampling frequency is specified as "at least 1 kS/s per channel," which satisfies the minimum 800 Hz sampling frequency stated in requirement 7.
8. & 9. The anti-aliasing filter is specified with a cutoff frequency between 100 Hz to 200 Hz, which is below the minimum 400 Hz cutoff frequency required by requirement 9 and may not ensure a -20 dB gain at half the sampling frequency as per requirement 8.
10. The presence of low-pass filters before the multiplexer(s) is implied, but their exact positioning is not clearly stated, leaving requirement 10 partially unclear.
11. The project uses an 8-to-1 analog multiplexer, which satisfies requirement 11.
12. The ADG1606 multiplexer is a solid-state device, in compliance with requirement 12.

Given these points, the verdict for the project is that it is "acceptable" as it can be implemented and doesn't have any fatal issues, but the score isn't perfect due to ambiguities and potential non-compliance with requirements 5, 6, 8, 9, and 10.