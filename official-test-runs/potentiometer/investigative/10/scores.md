Score: 7
Explanations: 
The project summary describes a pendulum angle measurement system with several components, including a potentiometer, buffer stage, scaling amplifier, low pass filter, notch filters, and a DAQ system. Let's evaluate the project based on the given requirements:

1. The potentiometer is used as a voltage divider: The summary does not explicitly state that the potentiometer is used as a voltage divider, but given that it's a common application for potentiometers in angle measurement systems, we can assume this requirement is met.

2. The voltage applied to the potentiometer is +/- 10 V: The potentiometer is rated for 10 V across terminals, but typically operated at 5 V. The power supply for the op-amp in the buffer stage is typically +/- 15 V. This does not directly confirm the voltage applied to the potentiometer; however, it seems the design can support the +/- 10 V requirement.

3. The architecture includes a voltage divider, an anti-aliasing filter, and a DAQ: The summary mentions a potentiometer (implied voltage divider), a buffer stage (which can act as a voltage follower), and a DAQ. There is also mention of a low pass filter which can serve as an anti-aliasing filter. Thus, this requirement is met.

4. The input voltage of the DAQ is centered at 0, with a range of +/- 7V: The summary states that the input range of the DAQ is configurable to ±7 V, which meets this requirement.

5. The maximum voltage applied to the DAQ is 7V: As the input range of the DAQ is configurable to ±7 V, it is implied that the maximum voltage applied to the DAQ will not exceed 7V. This requirement is met.

6. There is a low pass filter to avoid aliasing: A low pass Butterworth filter with a cutoff frequency of 100 Hz is mentioned, which would serve as an anti-aliasing filter. Therefore, this requirement is met.

7. There is a filter removing frequencies between around 50 and 60 Hz: The design includes notch filters for 50 Hz and 60 Hz with the specified attenuation. This requirement is met.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The summary does not provide enough information to determine if the gain of the signal is at least -20 dB at 500 Hz. The cutoff frequency for the low pass filter is given as 100 Hz, but without the transfer function or additional information about the filter design (e.g., component values, specific topology beyond "Butterworth"), we cannot confirm whether this requirement is met.

In conclusion, the project meets most of the requirements, but there is insufficient information to confirm requirement 8. Assuming the potentiometer is indeed used as a voltage divider and the design supports the +/- 10 V requirement, the project meets 7 out of 8 requirements.