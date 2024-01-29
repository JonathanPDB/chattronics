Score: 6
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider: The summary explicitly mentions a linear rotary potentiometer, which is typically used as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: There is no explicit mention of the voltage applied to the potentiometer, so this requirement is not verifiably met.

3. The architecture is simple: The summary indicates a simple architecture with a voltage divider (potentiometer), buffer amplifier, scaling amplifier (which could act as a voltage buffer), and an anti-aliasing filter before DAQ measurement.

4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The DAQ's input voltage range is specified as at least +/- 7 V, which satisfies this requirement.

5. The maximum voltage applied to the DAQ is 7V: The DAQ's input voltage range is specified as at least +/- 7 V, which implies that the maximum voltage applied to the DAQ is 7V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The summary specifies a recommended first-order cutoff frequency of 50 Hz and a second-order cutoff frequency, indicating the presence of an anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz: A notch filter with a cutoff frequency of 50 Hz and 60 Hz is specified, which meets this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The summary does not provide enough information about the order and cutoff frequency to ensure that the gain of the signal is at least -20 dB at 500 Hz. Without specific values for the filter order and cutoff frequency, it is unclear whether this requirement is met.

Based on the information provided, the project summary meets 6 out of the 8 requirements listed. Requirements 2 and 8 are not verifiably met based on the information given.