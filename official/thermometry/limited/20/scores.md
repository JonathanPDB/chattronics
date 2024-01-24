Score: 5
Explanations: 
The project summary covers the following requirements:

1. The output voltage range to be measured by the multimeter is specified as between 0 and 20 Volts, which aligns with the requirement.
2. The gain is provided and justified through the use of an instrumentation amplifier and a variable gain stage with a potentiometer for calibration. However, the exact gain value is not provided, so it's not clear if the amplification is appropriate.
3. The NTC thermistor is linearized using a fixed resistor based on the resistance at the midpoint temperature of 50°C, fulfilling the requirement for linearization.
4. The sensor used is an NTC thermistor, which is a requirement.
5. The self-heating effect is acknowledged, and measures to minimize it below 1% are mentioned, indicating that the maximum current through the NTC is known and controlled.
6. The architecture is described as consisting of the sensor, linearization stage, amplification, and optional filtering, which is in line with the requirement.
7. The output is intended to be measured by a multimeter, which is part of the project's design.

The project summary does not explicitly mention the following:

- The exact gain value of the amplification stage and the justification for it (Requirement 2).
- It is assumed that the linearization is successful, but there is no explicit statement or evidence provided that the NTC is indeed linearized (Requirement 5). This could be inferable from the use of the linearization resistor, but without data or a graph showing the linearized response, we cannot be conclusive.

Therefore, the project meets requirements 1, 3, 4, 5, 6, and 7. Requirement 2 is partially met, but without an explicit gain value, it is not fully satisfied. Requirement 5 is also not explicitly confirmed as successful linearization, but it is implied by the approach taken. Since the requirements state that linearization is "absolutely necessary and fatal if false," and because we can't definitively confirm it, we must err on the side of caution and not count it fully satisfied. Consequently, we count 5 out of the 7 requirements as fully met.