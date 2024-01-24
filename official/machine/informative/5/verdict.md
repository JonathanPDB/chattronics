Verdict: unfeasible

Explanations: 
The project has several commendable aspects that align with the requirements, but there are critical areas where it falls short or lacks specificity.

1. Both sensors have d.c. output, which complies with requirement 1.
2. The use of instrumentation amplifiers with specified gains for both sensors satisfies requirement 2.
3. The pressure sensor is mentioned to be used alongside an instrumentation amplifier, but there is no explicit mention of a Wheatstone bridge, which is a fatal flaw in meeting requirement 3.
4. An ADC is included in the design, fulfilling requirement 4.
5. Linearization for infrared radiation sensors is not clearly addressed in the design, which is a fatal flaw against requirement 5.
6. The sampling order strategy is not discussed, leaving requirement 6 unmet.
7. The ADC sampling rate is not specified, failing to meet requirement 7.
8. & 9. The anti-aliasing filter design is specified with a cutoff frequency of 600 Hz, which is higher than 400 Hz and less than half the unspecified ADC sampling frequency, but without the sampling frequency, it is not possible to confirm if it satisfies requirements 8 and 9.
10. The anti-aliasing filter is mentioned, but its position relative to the multiplexer(s) is not specified, which is crucial in meeting requirement 10.
11. The project uses a multiplexer, which meets requirement 11.
12. The multiplexer is solid-state, fulfilling requirement 12.

The project fails to meet several critical requirements and lacks specificity in some areas, rendering the design incomplete and not fully practical in its current form.