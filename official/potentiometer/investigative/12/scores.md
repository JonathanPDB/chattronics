Score: 4
Explanations: 
The project summary provides a detailed overview of the pendulum angle measurement system design, including the selection of components and topology of circuits. Here's how it aligns with the given requirements:

1. The potentiometer is used as a voltage divider: This requirement is implicitly met as the potentiometer is typically used as a voltage divider in such applications, but the specific configuration is not detailed.

2. The voltage applied to the potentiometer is +/- 10 V: The project summary mentions a maximum voltage of 7 volts for the potentiometer, which does not meet the requirement of +/- 10 V.

3. The architecture is simple with a voltage divider, anti-aliasing filter, and then the DAQ: The summary describes a voltage follower, low-pass filter with notch, gain stage, and anti-aliasing filter before the DAQ, which aligns with the architecture requirement.

4. The input voltage of the DAQ is centered in 0, +/- 7V: There is no explicit mention of the input voltage of the DAQ being centered at 0V or its range being +/- 7V, so this requirement is not verifiably met.

5. The maximum voltage applied to the DAQ is 7V: The project does not explicitly state the maximum voltage applied to the DAQ, but the use of a gain stage adjustable up to 1.4 suggests that the output voltage could potentially exceed 7V if the input exceeds 5V. Therefore, this requirement is not clearly met.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The summary mentions an anti-aliasing filter with a 2nd order active low-pass design and a cutoff frequency of 250 Hz, which meets this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz: The summary specifies notch filters for 50 Hz and 60 Hz, satisfying this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The low-pass filter is mentioned to have a cutoff at 400 Hz, and it is either 2nd or 3rd order Butterworth. Without the specific order and transfer function, it cannot be confirmed if the gain at 500 Hz is at least -20 dB. However, considering a 2nd order Butterworth filter has a rolloff of 12 dB/octave, and a 3rd order has 18 dB/octave, and the cutoff frequency is near 500 Hz, it is likely that this requirement is met.

In summary, items 1, 3, 6, and 7 are met. Items 2, 4, 5, and 8 are not clearly met based on the information given.