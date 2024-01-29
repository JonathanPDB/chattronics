Verdict: unfeasible

Explanations: 
The project's description indicates a well-thought-out system for measuring the angle of a pendulum using a potentiometer for voltage division and various signal conditioning stages before analog-to-digital conversion. However, there are a few points that need to be addressed to determine if the project meets the requirements:

1. The potentiometer is used as a voltage divider, which aligns with the first requirement.
2. The voltage applied to the potentiometer is +/- 10 V, satisfying the second requirement.
3. The architecture includes a voltage divider, voltage buffer, low-pass filters, and DAQ, which is simple and meets the third requirement.
4. The input voltage of the DAQ is not explicitly stated to be centered at 0V, nor is the range mentioned. The requirement specifies an input voltage centered at 0V with a range of +/- 7V, which is essential for a correct verdict.
5. The maximum voltage applied to the DAQ is not specified. It is crucial that the DAQ does not receive more than +/- 7V to avoid damage and ensure accurate measurements.
6. There is an anti-aliasing filter with a cutoff frequency of 100 Hz, which should help avoid aliasing and thus meets the sixth requirement.
7. A low-pass filter is mentioned for attenuating 50 and 60 Hz noise, which fulfills the seventh requirement.
8. The stopband attenuation is stated to be > 60 dB beyond 500 Hz, which implies that the gain of the signal at 500 Hz would be at least -20 dB, meeting the eighth requirement.

The main concerns are the lack of specific mention of the DAQ's input voltage range and maximum voltage, which are critical for assessing the project's feasibility. Without this information, we cannot confirm that the DAQ's input requirements are met, which is a potential fatal issue. Therefore, the project is currently categorized as "unfeasible" until these details are clarified.