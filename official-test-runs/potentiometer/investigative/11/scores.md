Score: 7
Explanations: 
The project summary covers the following items:

1. The potentiometer is used as a voltage divider: The summary specifies a 10 kOhm linear potentiometer used for angle measurement, which implicitly indicates it is used as a voltage divider.
2. The voltage applied to the potentiometer is +/- 10 V: This is explicitly mentioned as the potentiometer outputs a voltage range from -10V at 45 degrees to +10V at 135 degrees.
3. The architecture is simple: The project comprises four main blocks: Potentiometer, Buffer, Low-Pass Filter, and Notch Filter, which interface with the DAQ system. This implies a simple architecture following the requirements.
4. The input voltage of the DAQ is centered in 0, for example, +/- 7V: The buffer stage is designed to attenuate the potentiometer's output to fit the DAQ's input range of +/-7V.
5. The maximum voltage applied to the DAQ is 7V: The project summary explicitly states that the buffered potentiometer signal range is accommodated to +/-7V, ensuring the maximum voltage applied to the DAQ is 7V.
6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: An active low-pass Sallen-Key filter with a cutoff frequency of 50 Hz is proposed, which serves as an anti-aliasing filter.
7. There is a filter removing frequencies between around 50 and 60 Hz: A notch filter is suggested for creating notches at 50 Hz and 60 Hz, which would remove these frequencies.
8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The summary does not provide sufficient information to verify this requirement. The cutoff frequency for the low-pass filter is 50 Hz with a roll-off of -12 dB/octave, but without knowing the exact configuration of the filter, the gain at 500 Hz cannot be determined. Additionally, the notch filter's depth is specified to be at least 20 dB, but this is related to the notches at 50 and 60 Hz, not the gain at 500 Hz.

Based on the provided summary, the project successfully covers 7 out of the 8 requirements. Requirement 8 cannot be verified due to insufficient information regarding the filter's gain at 500 Hz.