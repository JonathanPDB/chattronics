Score: 4
Explanations: 
1. The potentiometer is used as a voltage divider: The summary describes the potentiometer outputting a linear voltage change from -5 V to +5 V as the pendulum moves, which implies the use of a voltage divider configuration.

2. The voltage applied to the potentiometer is +/- 10 V: This requirement is not explicitly met. The potentiometer output is described as ranging from -5 V to +5 V, but there's no information on the voltage applied to the potentiometer itself. This could potentially mean that the voltage supplied is indeed +/- 10 V, but this is not confirmed in the summary.

3. The architecture is simple with a voltage divider, anti-aliasing filter, and DAQ: The summary outlines a simple architecture including a potentiometer (voltage divider), a unity-gain buffer (voltage buffer), low-pass filter, anti-aliasing filter, and a DAQ system.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The summary does not specify the input voltage range of the DAQ. However, it does mention that the potentiometer outputs -5 V to +5 V, which is centered on 0 but does not meet the +/- 7V range.

5. The maximum voltage applied to the DAQ is 7V: Since the potentiometer output voltage range is -5 V to +5 V, the maximum voltage applied to the DAQ does not exceed 7 V, which meets this requirement.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The system includes an anti-aliasing filter block with a specified cutoff frequency of 250 Hz, which should help prevent aliasing given the DAQ's sampling rate of 1000 samples per second.

7. There is a filter removing frequencies between around 50 and 60 Hz: The summary does not explicitly mention a filter that specifically targets the 50 to 60 Hz range, although a 4th-order Butterworth filter is mentioned which could potentially attenuate these frequencies. However, without specific details on the filter design, we cannot confirm this requirement is met.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: A second-order Butterworth low-pass filter with a cutoff frequency of 200 Hz is mentioned, which would typically provide at least -20 dB attenuation at 500 Hz. However, without the actual filter design or transfer function, this cannot be confirmed with certainty.

In summary, the project meets requirements 1, 3, 5, and 6, but does not explicitly meet requirements 2, 4, 7, and 8 due to a lack of specific information or confirmation.