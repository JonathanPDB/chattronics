Verdict: acceptable

Explanations: 
The project description provided meets several of the requirements but does fail to address a few critical ones:

1. Both sensors have d.c. output: The description does not specify if the sensors' outputs are indeed d.c., which is a fatal requirement if false. The assumption here is that they are d.c. since no demodulation is mentioned.
2. Both sensors must be amplified: The pressure sensor is amplified by an instrumentation amplifier with a gain of 5000, and the temperature sensor is amplified with a gain of 50. This meets the requirement.
3. The pressure sensor must be inserted in a Wheatstone bridge and amplified by an instrumentation amplifier: The use of an AD620 (an instrumentation amplifier) suggests compliance with this requirement. However, the Wheatstone bridge is not explicitly mentioned, which is a critical part of the pressure sensor signal conditioning.
4. An ADC should be used: The project specifies the use of a 16-bit SAR ADC with a sampling rate of at least 2 kSPS, fulfilling this requirement.
5. Infrared radiation sensors are being linearized: The project mentions that nonlinearity correction for the temperature sensor is addressed digitally post-ADC or with a DAC for analog correction. This meets the requirement.
6. The solution mentions the sampling order strategy: The project does not explicitly discuss the sampling order strategy (sequential, simultaneous, etc.).
7. The sampling frequency of the ADC is not less than 800 Hz: The ADC sampling rate is specified as at least 2 kSPS, which exceeds the minimum requirement.
8. The anti-aliasing filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at half the sampling frequency: With a fourth-order Butterworth filter and a cutoff frequency of 500 Hz, this requirement is met.
9. The low pass cutoff frequency is higher than 400 Hz and lower than half the total sampling frequency: The cutoff frequency is 500 Hz, which meets this requirement given the 2 kSPS sampling rate.
10. There are low-pass filters positioned before the multiplexer(s): The project specifies RC low-pass filters with a cutoff frequency of 500 Hz for both sensor types, positioned in the signal conditioning block, presumably before the multiplexer.
11. The project uses multiplexers to choose channels: The project describes the use of a 16-to-1 CMOS analog multiplexer, which satisfies this requirement.
12. The multiplexers are solid state: The CD74HC4067 is indeed a solid-state device, meeting this requirement.

The project description does not explicitly mention the use of a Wheatstone bridge for the pressure sensors, nor does it clarify the sampling strategy. These unaddressed points could lead to potential issues in implementation. Additionally, the d.c. output of the sensors is assumed but not confirmed in the description, which is a critical requirement.