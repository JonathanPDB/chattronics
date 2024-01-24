Verdict: acceptable

Explanations: 
The project summary describes an analog instrumentation system for monitoring pressure and surface temperature, including the use of a pressure sensor array and temperature sensor array, instrumentation amplifier array, non-linear signal conditioning, low pass filter array, multiplexer, and ADC. The summary details various components and strategies for implementing the system. Here's the analysis of the requirements:

1. The pressure and temperature sensors output DC signals, satisfying the first requirement.
2. An instrumentation amplifier array with a gain of 5000 is suggested, indicating that the sensor outputs are to be amplified, meeting the second requirement.
3. The pressure sensor is not explicitly mentioned to be inserted in a Wheatstone bridge, but since it is a strain-gauge-based transducer, it is typically used with a Wheatstone bridge, and an instrumentation amplifier is suggested. This is generally compliant, but the lack of explicit mention of the Wheatstone bridge configuration is a minor concern.
4. An ADC with a greater than 1 kHz sampling rate per channel is mentioned, fulfilling the fourth requirement.
5. Non-linear signal conditioning is discussed, with options for analog circuitry using operational amplifiers and digital processing, which meets the fifth requirement.
6. The summary does not mention the sampling order strategy, leaving out important information on how the system will handle multiple sensors.
7. The ADC sampling rate is greater than 1 kHz per channel, which exceeds the minimum requirement of 800 Hz, satisfying the seventh requirement.
8. & 9. The low pass filter has a cutoff frequency of 500 Hz, which is higher than 400 Hz but does not explicitly state that the gain of the signal is at least -20 dB at half the sampling frequency, nor does it confirm that the cutoff frequency is less than half the total sampling frequency.
10. Low-pass filters are mentioned, and their positioning before the multiplexer(s) suggests a correct approach to reducing aliasing.
11. An 8-to-1 analog multiplexer is described, meeting the requirement of using multiplexers to choose channels.
12. The multiplexers described are solid state.

While the project mostly aligns with the requirements, the lack of explicit mention of the Wheatstone bridge configuration for the pressure sensor (requirement 3), the absence of details on the sampling order strategy (requirement 6), and the incomplete information on the low pass filter specifications in relation to the ADC sampling rate (requirements 8 and 9) prevent the project from being categorized as "optimal". However, since these issues do not constitute fatal flaws and the system can still be implemented effectively, the verdict is "acceptable".