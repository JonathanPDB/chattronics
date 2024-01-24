Verdict: acceptable

Explanations: 
The project description addresses many of the requirements, but there are some critical aspects that are either not met or not clearly specified:

1. The pressure sensors are strain-gauge types, which typically provide a DC output, so this meets the requirement of no demodulation needed.
2. Amplification is mentioned for the pressure sensors using instrumentation amplifiers, which satisfies the amplification requirement.
3. The pressure sensor is indeed being used in conjunction with an instrumentation amplifier, but it is not explicitly stated that it is in a Wheatstone bridge configuration. This is a critical detail for strain-gauge pressure sensors.
4. An ADC is used, meeting this requirement.
5. The infrared temperature sensors have a built-in 17-bit ADC and use a second-order polynomial correction for linearization. However, it is not clear if this digital linearization is done in the computer or if diode networks or log amplifiers are used for analog linearization. This requirement may not be fully met.
6. The sampling order strategy is not mentioned, which is important for understanding how the system handles multiple sensor inputs.
7. The ADC sampling rate is specified as 1 kSps per channel, which is above the minimum requirement of 800 Hz.
8. The low-pass filter has a stop-band attenuation of -40 dB at 800 Hz, which is well beyond the -20 dB requirement at half the sampling frequency, assuming a 1 kSps rate.
9. The low-pass filter's cutoff frequency is specified as 500 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (given as 1 kSps per channel), meeting this requirement.
10. The presence of low-pass filters is mentioned, but it is not explicitly stated that they are positioned before the multiplexer(s).
11. Multiplexers are used, satisfying this requirement.
12. The multiplexers recommended are solid-state devices.

Given the information provided, the project is close to meeting all requirements but has some gaps in the details provided, particularly in the temperature sensor linearization method and the sampling order strategy. Assuming a Wheatstone bridge is used for the pressure sensors would be speculative without explicit confirmation.