Score: 8
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider: The project specifies a precision potentiometer with a 10 kOhm resistance and a linear response, used to create a voltage range of -5 V to 5 V, indicating it is used as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V: The project mentions a supply voltage of -10 V to 10 V for the potentiometer.

3. The architecture includes a voltage divider, an anti-aliasing filter, and a DAQ system: The project describes a system that includes a potentiometer (voltage divider), a unity gain buffer (voltage follower), a low-pass filter with a notch filter, an anti-aliasing filter, and a DAQ system.

4. The input voltage of the DAQ is centered at 0, i.e., +/- 7V: The project specifies that the DAQ system's input range is configured to match the ±7 V input range.

5. The maximum voltage applied to the DAQ is 7V: The gain stage calculation suggests a gain of 1.4 to match the DAQ input range from the potentiometer output range, which means the maximum voltage applied to the DAQ would be 7 V as required.

6. There is a low-pass filter (or anti-aliasing filter) to avoid aliasing: The project includes a 2nd order Butterworth low-pass filter with a cutoff frequency of 50 Hz, which is one-tenth of the Nyquist frequency, suggesting it serves as an anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz: The project describes a notch filter designed to filter out noise at 50 Hz and 60 Hz.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz: The project specifies a 2nd order active Butterworth filter with a cutoff frequency of 100 Hz, which would provide at least -20 dB attenuation at 500 Hz due to the nature of Butterworth filters' frequency response.

All eight requirements have been successfully covered by the project summary.