Score: 7
Explanations: 
The project summary covers the following requirements:

1. The potentiometer is clearly stated to be used for angular position sensing, which typically involves using it as a voltage divider.

2. The power supply for the buffer is specified to be within the -10V to +10V range, which implies that the potentiometer is also being used within this voltage range.

3. The architecture seems simple and includes a voltage divider (potentiometer), a voltage buffer to prevent loading, an anti-aliasing filter (low pass filter), and a notch filter before interfacing with the DAQ.

4. The DAQ's input voltage range is centered at 0V with a maximum of +/- 7V, as indicated by the design of the attenuator which scales the potentiometer's voltage range to this level.

5. The maximum voltage applied to the DAQ is maintained at 7V due to the inclusion of an attenuator with a calculated attenuation ratio to scale down the voltage.

6. The low pass filter with a cutoff frequency assumed to be 10 Hz would serve as an anti-aliasing filter, which avoids aliasing by attenuating high-frequency components that could otherwise fold back into the frequency band of interest.

7. A notch filter is included to remove frequencies around 50 and 60 Hz, which are common power line noise frequencies.

8. The requirement for the filter to have at least -20 dB gain at 500 Hz is not explicitly stated. However, the notch filter is described to have a notch depth between -20 to -40 dB. This suggests that the notch filter could potentially meet the requirement, but without specific cutoff frequencies and filter order, it is not possible to determine if the -20 dB gain at 500 Hz is achieved.

The score is based on the requirements that are clearly met by the project summary. Requirement 8 is the only item that is not explicitly covered in the summary provided.