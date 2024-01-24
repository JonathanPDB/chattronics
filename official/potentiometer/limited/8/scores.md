Score: 7
Explanations: 
1. The potentiometer is used as a voltage divider: This is implicitly met as the potentiometer converts the pendulum's angular position into a corresponding voltage signal.

2. The voltage applied to the potentiometer is +/- 10 V: This is not explicitly stated, but the use of a level shifter and a gain stage suggests that the input voltage range is indeed +/- 10 V.

3. The architecture is simple: The system includes a voltage divider (with a voltage buffer), a level shifter, notch filters, an anti-aliasing filter, a gain stage, and a DAQ. This meets the requirement for a simple architecture.

4. The input voltage of the DAQ is centered at 0, for example, +/- 7V: The level shifter and gain stage are designed to ensure that the voltage range is within the DAQ's +/- 7 V input range.

5. The maximum voltage applied to the DAQ is 7V: The design explicitly ensures that the voltage does not exceed +/- 7 V.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing: The design includes an anti-aliasing filter with a second-order Butterworth low-pass filter, which satisfies this requirement.

7. There is a filter removing frequencies between around 50 and 60 Hz: The design includes active twin-T notch filters for both 50 Hz and 60 Hz, meeting this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The anti-aliasing filter is a second-order Butterworth low-pass filter with a cutoff frequency just below 500 Hz, which should provide at least -20 dB of attenuation at 500 Hz. However, without the exact cutoff frequency and the order of the filter, it is not possible to confirm this requirement with certainty.