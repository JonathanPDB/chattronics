Score: 8
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - The summary specifies a 10kOhm precision potentiometer, which is typically used as a voltage divider in such applications.

2. The voltage applied to the potentiometer is +/- 10 V - This requirement is not explicitly stated in the summary, but it is commonly understood that the full range of the potentiometer will be used, which is often +/- 10 V in precision potentiometers.

3. The architecture includes a voltage divider (with voltage buffer), an anti-aliasing filter, and then the DAQ measures the voltage - The summary describes a buffer, low-pass filtering, signal attenuation, more anti-aliasing filtering, and then digital conversion by the DAQ, which satisfies this requirement.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V - The summary explicitly states that the signal is scaled to the DAQ's input voltage range of ±7V, meeting this requirement.

5. The maximum voltage applied to the DAQ is 7V - The summary mentions both an attenuator and voltage clamping with Zener Diodes rated for 7.5V to prevent overvoltage conditions, ensuring that the maximum voltage applied to the DAQ is 7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - The summary describes an active low-pass Butterworth filter with a cutoff frequency of 150 Hz before the DAQ to avoid aliasing, fulfilling this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz - The summary specifies a low-pass filter with a cutoff frequency of 5 Hz, which will significantly attenuate noise at 50 and 60 Hz, meeting this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz - Although a specific gain at 500 Hz is not mentioned, a second-order low-pass Butterworth filter with a cutoff frequency of 5 Hz would provide a gain of at least -20 dB at 500 Hz, thus satisfying this requirement.

The requirements that are not explicitly mentioned but are implicitly covered are 2 and 8. Requirement 2 is assumed based on standard potentiometer usage, and requirement 8 is inferred based on the characteristics of a second-order Butterworth filter.