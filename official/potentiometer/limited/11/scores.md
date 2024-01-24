Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is used as a voltage divider - This is implicitly covered as the potentiometer's operation within its voltage range (-10V to 10V) and its application in angle measurement suggests its use as a voltage divider.

2. The voltage applied to the potentiometer is +/- 10 V - The operating voltage of the potentiometer is specifically stated as -10V to 10V, which meets this requirement.

3. The architecture is simple - The architecture includes a voltage divider, an amplification block (which could act as a voltage buffer), a filtering block (which includes an anti-aliasing filter), and protection before the DAQ, which is a straightforward design for this application.

4. The input voltage of the DAQ is centered at 0, for example, +/- 7V - The amplification block is designed to map a range of -5V to 5V to the DAQ's +/-7V input range, which would center the input voltage at 0.

5. The maximum voltage applied to the DAQ is 7V - The protection block includes TVS diodes with a breakdown voltage slightly above the DAQ's maximum input voltage, and the amplification block's gain ensures that the output does not exceed +/-7V, meeting this critical requirement.

6. There is a low pass filter (or anti-aliasing filter) that avoids aliasing - The filtering block specifies a low-pass filter with a cutoff frequency of 100 Hz or lower, which serves as an anti-aliasing filter.

7. There is a filter removing frequencies between around 50 and 60 Hz - The active Twin-T Notch Filter is designed to attenuate 50 Hz and 60 Hz frequencies, which satisfies this requirement.

8. The filter has a cutoff frequency and order such that the gain of the signal is at least -20 dB at 500 Hz - The summary does not provide enough information on the order of the filter or the exact gain at 500 Hz. The low-pass filter with a cutoff frequency around 250 Hz is mentioned, but without specifying the filter order, it's not possible to confirm that the requirement of at least -20 dB at 500 Hz is met.

Therefore, the project meets 7 out of the 8 requirements listed. The only requirement that is not explicitly met is the filter's performance at 500 Hz, as the information provided is insufficient to make an accurate assessment.