Verdict: acceptable

Explanations: 
The project summary provides a detailed outline of an analog instrumentation system for monitoring pressure and temperature. However, to assess its compliance with the given requirements, the following points must be addressed:

1. Both sensors have d.c. output, so no demodulation should be used. - The summary does not explicitly mention demodulation, which is good since it is not needed for d.c. output sensors.
2. Both sensors must be amplified - The summary describes amplification for both pressure and temperature sensors, meeting this requirement.
3. The pressure sensor must be inserted in a wheatstone bridge and amplified by an instrumentation amplifier - The summary mentions the use of an instrumentation amplifier for the pressure sensor but does not explicitly state the use of a Wheatstone bridge. This is a potential issue if not implemented.
4. An ADC should be used - The summary includes the use of an ADC with a minimum 16-bit resolution and a sampling rate greater than 10 ksps, satisfying this requirement.
5. Infrared radiation sensors are linearized either digitally in the computer or by using diode networks or log amplifiers - The summary does not detail the linearization method for the infrared sensors, which is a critical requirement.
6. The solution mentions the sampling order strategy, if it'll be sequential, simultaneously, etc. - The summary does not specify the sampling order strategy.
7. The sampling frequency of the ADC is not less than 800 Hz - The summary states a total rate of over 10 ksps to handle eight channels at 800 Hz each, which fulfills this requirement.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency - The summary does not provide enough detail to evaluate if this requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency - The summary mentions a low-pass filter with a cutoff frequency significantly above 400 Hz and below half the ADC sampling rate, which seems to satisfy this requirement.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s) - The summary describes low-pass filters in the signal conditioning stage, but it is not clear if they are positioned before the multiplexers.
11. The project uses multiplexer(s) to choose channels - The summary includes the use of an 8:1 analog multiplexer, complying with this requirement.
12. The multiplexer(s) are solid state - The suggested models for the multiplexer are solid state, fulfilling this requirement.

In conclusion, while the project summary addresses many of the requirements, it lacks explicit details on the Wheatstone bridge for the pressure sensor, the linearization method for the infrared sensors, the sampling order strategy, and specific information about the anti-aliasing filters' placement and specifications. Therefore, the project cannot be considered optimal as it stands.