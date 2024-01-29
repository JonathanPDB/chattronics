Score: 7
Explanations: 
1. The potentiometer is used as a voltage divider and is directly connected to the DAQ after signal conditioning, implying that it is serving this role.

2. The voltage applied to the potentiometer is not explicitly mentioned to be +/- 10 V, but since the op-amps for the buffer and scaling amplifier are powered by the same -10 V to +10 V source, it is reasonable to infer that the potentiometer is also operating within this voltage range.

3. The architecture includes a voltage divider (potentiometer), a voltage scaling stage (scaling amplifier), a buffer amp, a band-stop filter, and an anti-aliasing filter before the DAQ measures the voltage, meeting this requirement.

4. The input voltage of the DAQ is centered at 0V, as stated that the potentiometer output at 90 degrees (steady position) is 0V, and the DAQ's input range is +/- 7V.

5. The maximum voltage applied to the DAQ is 7V, as the scaling amplifier has been configured to adjust the potentiometer's output range to fit the DAQ's input range of -7V to +7V.

6. There is an active Butterworth low-pass filter with a cutoff frequency of 250 Hz, serving as an anti-aliasing filter to avoid aliasing.

7. The project includes a Twin-T Notch Filter for 50 Hz and 60 Hz, which would remove frequencies around these values, satisfying this requirement.

8. The filter requirement is not explicitly met as there are no details about the gain of the signal at 500 Hz. The -20 dB gain requirement at 500 Hz cannot be confirmed without additional information about the filter's response at that frequency. However, given that the cutoff frequency is 250 Hz for a second-order filter, there is a possibility that the gain at 500 Hz could be close to or beyond -20 dB. Since the exact response is not provided, this requirement cannot be fully verified.