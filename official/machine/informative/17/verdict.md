Verdict: acceptable

Explanations: 
The project description provides an extensive overview of the proposed design for an analog instrumentation system to monitor pressure and temperature. However, the summary lacks explicit confirmation of several critical requirements. Here is the assessment based on the requirements provided:

1. Both sensors have d.c. output - The summary does not explicitly confirm that both sensors have a d.c. output, which is a fatal requirement.
2. Both sensors must be amplified - The design includes amplification for both the pressure sensor (via instrumentation amplifier) and the temperature sensor.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier - The use of an instrumentation amplifier is mentioned, but there is no explicit mention of a Wheatstone bridge for the pressure sensor.
4. An ADC should be used - The design incorporates an ADC with a sampling rate and resolution that meet the requirements.
5. Infrared radiation sensors are being linearized - The summary mentions linearization algorithms for temperature data but does not specify the method, which could be digital, diode networks, or log amplifiers.
6. The solution mentions the sampling order strategy - The summary does not discuss the sampling order strategy (sequential, simultaneous, etc.).
7. The sampling frequency of the ADC is not less than 800 Hz - The ADC sampling rate is specified as a minimum of 2 kHz per channel, satisfying this requirement.
8. The anti-aliasing filter has a cutoff frequency and order - The anti-aliasing filter is described as a second-order filter with a cutoff frequency of 1 kHz, which should provide the required attenuation at half the sampling frequency.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency - The cutoff frequency of 1 kHz meets this requirement.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s) - The summary specifies low-pass filters, but it is not clear if they are positioned before the multiplexers.
11. The project uses multiplexer(s) to choose channels - An 8-to-1 multiplexer is included in the design.
12. The multiplexer(s) are solid state - The suggested multiplexers (ADG708/ADG709) are solid state.

The summary provides a comprehensive design that appears to meet most requirements. However, it lacks explicit confirmation of d.c. output for the sensors, the use of a Wheatstone bridge for the pressure sensor, the linearization method for the infrared sensors, the sampling order strategy, and the positioning of the anti-aliasing filters relative to the multiplexers. These omissions and uncertainties prevent the design from being classified as 'optimal.'