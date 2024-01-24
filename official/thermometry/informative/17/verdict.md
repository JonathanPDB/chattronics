Verdict: unfeasible

Explanations: 
The project outline presents a comprehensive approach to designing a temperature measurement system using an NTC thermistor. The requirements are generally addressed, but there are elements of the design that need clarification or adjustment to ensure they align with the requirements:

1. The use of the Vishay NTCLE100E3 thermistor satisfies the requirement of using an NTC sensor.
2. The constant current source is a good design choice to minimize self-heating effects in the NTC thermistor, but the exact maximum current value is not provided, which is necessary to confirm that self-heating is accounted for.
3. The inclusion of a buffer amplifier and instrumentation amplifier provides the necessary amplification stage. However, the specific gain value is not mentioned. The justification for the gain chosen is critical to ensure the output voltage range matches the required 0 to 20 Volts.
4. The project mentions a linear voltage change from 0V at 10°C to 20V at 90°C, which implies a linear relationship between temperature and voltage; however, it does not explicitly state that the NTC is linearized, nor does it provide details on the resistance value used for linearization at the midrange 50°C. Linearization is a crucial requirement and must be clearly addressed.
5. The level shifter is included to map the voltage range to 0-20V, but again, specific details or calculations are not provided to ensure that the output indeed falls within this range.
6. A microcontroller and ADC are mentioned, but they are not necessary for the analog output to be measured by a multimeter; this adds complexity and is not required by the given requirements.
7. Filtering is suggested, which is good practice, but it is optional according to the requirements.
8. The protection mechanisms mentioned are appropriate for safeguarding the system.

The project outline is well thought out but lacks the specificity and explicit confirmation of some of the critical requirements, particularly the linearization of the NTC thermistor and the precise gain of the amplification stage. Moreover, the inclusion of a microcontroller and ADC contradicts the requirement for an analog output measurable by a multimeter, adding unnecessary complexity.