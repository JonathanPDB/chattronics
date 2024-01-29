Score: 9
Explanations: 
1. The project summary does not explicitly state that the sensors have d.c. output without the need for demodulation. However, it is implied by the type of sensors mentioned (strain gauge pressure transducers and infrared thermometer ICs), which typically provide a d.c. output. Assumption is made based on typical characteristics of these sensors.
2. Both types of sensors are mentioned to be connected to amplifier arrays, indicating they are amplified.
3. There is no explicit mention of the pressure sensor being inserted in a Wheatstone bridge, but it is a typical application for strain gauge pressure transducers. Since this is a common practice and not explicitly contradicted, it is assumed that the requirement is met.
4. The use of ADCs is described for both sensor arrays, indicating this requirement is met.
5. The summary mentions digital low-pass filters for temperature sensors but does not explicitly mention linearization methods. Although temperature sensors typically require linearization, the requirement specifies that the linearization should be mentioned as being performed either digitally or by using diode networks or log amplifiers. This requirement is not clearly met in the summary.
6. The summary does not explicitly mention the sampling order strategy, such as whether it will be sequential or simultaneous.
7. The sampling frequency of the ADC is mentioned to be at least 1 kHz, which is above the 800 Hz minimum requirement.
8. The anti-aliasing filter for the temperature sensor array is mentioned to have a cutoff frequency around 500 Hz. However, there is no explicit information on the gain of the signal at half the sampling frequency, nor the order of the filter to determine the attenuation rate.
9. The low-pass cutoff frequency for the temperature sensor array is mentioned to be around 500 Hz, which is higher than 400 Hz and, assuming the sampling rate is at least 1 kHz as mentioned, it would be lower than half the total sampling frequency.
10. Low-pass filters are mentioned to be positioned before the multiplexers, indicating the presence of anti-aliasing filters to reduce aliasing.
11. The project uses multiplexers to choose channels, as described in the summary for both Multiplexer1 and Multiplexer2.
12. The multiplexers mentioned (e.g., ADG708, ADG732, 74HC4051, CD4051, MAX14752, MAX14750, CD74HC4067) are solid-state devices.

The project summary covers the following items: 1 (assumed), 2, 3 (assumed), 4, 7, 9, 10, 11, and 12. Items 5, 6, and 8 are not clearly met based on the information provided.