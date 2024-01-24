Score: 7
Explanations: 
1. **The potentiometer is used as a voltage divider**: The summary indicates that a linear taper potentiometer is used, and it is common for such a potentiometer to be used as a voltage divider.

2. **The voltage applied to the potentiometer is +/- 10 V**: The summary mentions a 10V range over 270 degrees for the potentiometer, which implies a +/-10V supply.

3. **The architecture is simple**: The architecture described includes a voltage divider, a signal buffer, a low pass filter, a voltage scaler, and an anti-aliasing filter before the DAQ, which is in accordance with the requirement.

4. **The input voltage of the DAQ is centered in 0, for example, +/- 7V**: The summary specifies a voltage scaler that converts the -10 to +10 volts to within the DAQ's +/- 7V range.

5. **The maximum voltage applied to the DAQ is 7V**: The voltage scaler is specifically designed to ensure that the maximum voltage applied to the DAQ does not exceed +/- 7V.

6. **There is a low pass filter (or anti-aliasing filter) that avoids aliasing**: The summary describes an anti-aliasing filter with a second-order active low-pass Butterworth filter with a cutoff frequency of 100 Hz, which is suitable for avoiding aliasing in this context.

7. **There is a filter removing frequencies between around 50 and 60 Hz**: The low pass filter described as having "more than 40 dB at 50 Hz and 60 Hz" would effectively remove frequencies in this range.

8. **The filter has a cutoff frequency and order such that the gain of the signal is of at least -20 dB at 500 Hz**: The summary does not provide enough information to determine the gain at 500 Hz. A 4th order Butterworth filter with a cutoff frequency of 5 Hz would have a steep roll-off, but without the filter's transfer function or additional data, it is not possible to confirm the gain at 500 Hz. Therefore, we cannot positively confirm this requirement is met.