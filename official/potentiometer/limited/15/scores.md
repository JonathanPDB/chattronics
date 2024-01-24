Score: 8
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider: The potentiometer is a 10 kOhms single-turn rotary potentiometer, indicating its use as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: The power supply range for the op-amps is a minimum of +/-10V, which implies that the potentiometer is also supplied with this voltage range for proper operation.

3. The architecture includes a voltage divider, an anti-aliasing filter, and DAQ measurements: The summary mentions a potentiometer (voltage divider), a buffer (voltage follower), a level shifter, notch filter, amplifier, and an anti-aliasing filter before the DAQ measurements, satisfying this requirement.

4. The input voltage of the DAQ is centered at 0, +/- 7V: The level shifter adjusts the output range from -5V/+5V to -7V/+7V, centering the DAQ's input voltage at 0.

5. The maximum voltage applied to the DAQ is 7V: The level shifter ensures that the output voltage range is within +/-7V, which is the maximum voltage for the DAQ.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The system includes an active Butterworth low-pass filter with a cutoff frequency of 100 Hz to prevent aliasing.

7. There is a filter removing frequencies between around 50 and 60 Hz: The Twin-T Notch Filter is designed to attenuate noise at 50 and 60 Hz with a center frequency of 55 Hz.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The anti-aliasing filter has an attenuation of at least -40 dB at the Nyquist Frequency (500 Hz), which more than satisfies the requirement of -20 dB at 500 Hz.

All 8 requirements have been successfully covered by the project summary.