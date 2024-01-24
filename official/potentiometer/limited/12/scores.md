Score: 6
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider: The description mentions the use of a potentiometer for measuring the angle of a pendulum. It is implied that the potentiometer is part of the input conditioning and thus would act as a voltage divider in this context.

2. The voltage applied to the potentiometer is +/- 10 V: This is not explicitly mentioned, but given the power supply for the buffer amplifier is ±15V, it is reasonable to assume that the potentiometer could be supplied with +/- 10 V. However, without explicit confirmation, this requirement is not met.

3. The architecture includes a voltage divider, an anti-aliasing filter, and then the DAQ measures the voltage: The system does have a voltage divider (potentiometer), an anti-aliasing filter (4th-order Butterworth active filter), and a DAQ block.

4. The input voltage of the DAQ is centered at 0, +/- 7V: The DAQ input range is stated as ±7V, which implies it is centered at 0.

5. The maximum voltage applied to the DAQ is 7V: The DAQ input range is specified as ±7V, which meets this requirement.

6. There is a low pass filter (anti-aliasing filter) that avoids aliasing: A 4th-order Butterworth active filter with a cutoff frequency of 400 Hz is mentioned, which serves as the anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz: A notch filter with an attenuation depth of at least 20 dB at the notch frequencies (50 & 60 Hz) is included in the design.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: The anti-aliasing filter is a 4th-order Butterworth filter with a cutoff frequency of 400 Hz. A 4th-order Butterworth filter has a slope of 24 dB/octave beyond the cutoff frequency. Since 500 Hz is close to the cutoff frequency, the attenuation at 500 Hz would not be -20 dB. Therefore, this requirement is not met.

In summary, the requirements 1, 3, 4, 5, 6, and 7 are met based on the provided project summary. Requirements 2 and 8 are not met. Therefore, the score is 6.