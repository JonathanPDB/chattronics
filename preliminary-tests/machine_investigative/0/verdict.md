Verdict: incorrect

Explanations: 
The project summary describes an analog instrumentation system to monitor pressure and temperature in an industrial setting. The design includes pressure sensors, temperature sensors, signal conditioning circuits, multiplexers, and an ADC, as well as a data interface to a computer. Here's how the project aligns with the requirements:

1. Both sensors have d.c. output: The pressure sensors likely have a d.c. output as they are strain-gauge based. However, the temperature sensors are infrared detectors with a digital signal output (PWM or I2C), which does not meet the requirement of having a d.c. output. This is a fatal flaw.
2. Both sensors must be amplified: The pressure sensors are amplified, but the temperature sensor signal conditioning does not mention an amplification stage, only a buffer amplifier for impedance matching.
3. Pressure sensor in a Wheatstone bridge and amplified by an instrumentation amplifier: The design suggests using a stable excitation voltage and an instrumentation amplifier, which is compatible with a Wheatstone bridge configuration, though the bridge itself isn't explicitly mentioned.
4. An ADC is used: The design includes an ADC, the ADS1115.
5. Infrared radiation sensors are linearized: The summary does not mention any linearization method for the infrared sensors, which is a requirement.
6. Sampling order strategy is not mentioned: There is no information regarding the sampling order strategy (sequential, simultaneous, etc.).
7. ADC samples one sensor per time: The project uses an 8-to-1 multiplexer, which implies that the ADC samples one sensor at a time.
8. Sampling frequency of the ADC is not less than 800 Hz: The ADC sampling rate is 1 kHz, which meets this requirement.
9. Sampling frequency of the ADC is higher than 2000 Hz: The ADC sampling rate is 1 kHz, which does not meet this requirement.
10. The ADC should be a SAR: The summary does not specify if the ADS1115 is a SAR ADC, so this information is missing.
11. Low pass cutoff frequency: The pressure signal conditioning mentions a low-pass filter above 400 Hz, but the exact cutoff is not specified, and it's unclear if it meets the requirement of being lower than half the total sampling frequency.
12. Low-pass filters before the multiplexers: The summary mentions low-pass filters but does not specify their placement relative to the multiplexers.
13. The project uses multiplexers: The design includes an 8-to-1 multiplexer.
14. Multiplexers are solid state: The suggested models (DG408, ADG1606) are solid state.

Based on the information provided, the critical issues are the lack of d.c. output for temperature sensors, no explicit mention of a low-pass filter before the multiplexers, and no linearization method for the infrared sensors. Additionally, the sampling frequency is not higher than 2000 Hz, and the ADC type (SAR) is not confirmed. These issues make the project not fully compliant with the specified requirements.