Score: 7
Explanations: 
The project summary provides detailed information on the design of a pendulum angle measurement system. Based on the provided information, the following requirements are met:

1. The potentiometer is used as a voltage divider: The summary specifies a single-turn linear potentiometer with a sensitivity of 0.2222V/degree, which indicates that it is indeed used as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: The summary states that the potentiometer's response is from -10V at 45 degrees to +10V at 135 degrees, confirming this requirement.

3. The architecture is simple: The provided architecture includes a voltage divider (potentiometer), a voltage follower (buffer), a band-stop filter, an anti-aliasing filter, and a DAQ system, which is a straightforward and typical design for such applications.

4. The input voltage of the DAQ is centered in 0: The DAQ's input voltage range is configurable to +/-7 V, which implies that it can be centered around 0V.

5. The maximum voltage applied to the DAQ is 7V: The summary explicitly states that the DAQ's input voltage range is configurable to +/-7 V, ensuring that the maximum voltage does not exceed 7V.

6. There is a low pass filter (anti-aliasing filter) that avoids aliasing: The summary describes a second-order Butterworth low-pass filter with a cutoff frequency of 50 Hz, which serves as an anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz: The project includes a band-stop (twin-T notch) filter specifically designed to attenuate frequencies at 50 Hz and 60 Hz.

8. The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz: The anti-aliasing filter's cutoff frequency is 50 Hz, and it is a second-order Butterworth filter. However, there is not enough information to confirm that the gain at 500 Hz is at least -20 dB without the actual transfer function or further calculations. The -20 dB requirement typically implies that the signal is attenuated to 1/10 of its original strength at the specified frequency. With a second-order filter, the attenuation rate is -12 dB/octave beyond the cutoff frequency. Since 500 Hz is roughly two octaves above 50 Hz, the attenuation would be approximately -24 dB, which would satisfy the requirement. However, without explicit confirmation, this requirement is considered not verifiable from the information provided.

Based on the information given, the project meets 7 out of the 8 requirements. The only requirement that is not explicitly confirmed is the attenuation of the signal at 500 Hz, which could be inferred but is not directly stated.