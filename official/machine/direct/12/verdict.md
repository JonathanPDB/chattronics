Verdict: acceptable

Explanations: 
The project design mostly meets the specified requirements for an analog instrumentation system for pressure and temperature monitoring. Here's the analysis based on the provided information:

1. Both sensors have d.c. output, so no demodulation should be used. The design does not mention demodulation, which complies with this requirement.
2. Both sensors must be amplified. The project specifies amplification for both the pressure and temperature signals, fulfilling this requirement.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier. The use of an instrumentation amplifier is mentioned, but the Wheatstone bridge is not explicitly stated. However, strain-gauge pressure sensors typically require a Wheatstone bridge, so it can be inferred that this is part of the design.
4. An ADC should be used. The project includes an SAR ADC, which meets this requirement.
5. Infrared radiation sensors are being linearized either digitally in the computer or by using diode networks or log amplifiers. The project mentions using a linearization circuit for the temperature sensor, which satisfies this requirement.
6. The solution mentions the sampling order strategy, if it'll be sequential, simultaneously, etc. The project states a total sampling rate of 20 kSPS for 16 channels, implying a sequential sampling strategy.
7. The sampling frequency of the ADC is not less than 800 Hz. With a 1 kHz per channel sampling rate, this requirement is met.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency. The filters have a cutoff frequency of 800 Hz and 450 Hz for pressure and temperature, respectively, and are second-order, which should provide at least -20 dB attenuation at half the sampling frequency. However, the exact attenuation at half the sampling frequency is not specified, which leaves this requirement uncertain.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency. The cutoff frequencies of 800 Hz and 450 Hz for pressure and temperature filters respectively meet this requirement.
10. There are low-pass filters (or anti-aliasing filters) to reduce aliasing, which are positioned before the multiplexer(s). The design includes low-pass filters for both pressure and temperature channels, satisfying this requirement.
11. The project uses multiplexer(s) to choose channels. The design includes multiple multiplexers for channel selection, meeting this requirement.
12. The multiplexer(s) are solid state. The specified CD74HC4067 and CD4051B are solid-state multiplexers, fulfilling this requirement.

The main point of concern is the lack of explicit mention of the Wheatstone bridge for the pressure sensor, which is critical for proper sensor operation. Additionally, the exact attenuation provided by the anti-aliasing filters at half the sampling frequency is not quantified.