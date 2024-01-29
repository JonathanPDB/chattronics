Score: 10
Explanations: 
1. The summary does not explicitly mention that the sensors have DC output, but since strain-gauge pressure sensors and infrared temperature sensors typically have DC output, this requirement is met implicitly.
2. Both sensors are amplified, as indicated by the presence of amplifier blocks for both pressure and temperature.
3. The pressure sensor is mentioned to be amplified by an instrumentation amplifier, which is typically part of a Wheatstone bridge configuration, thus fulfilling this requirement.
4. An ADC is used for both pressure and temperature sensor arrays, as mentioned in their respective sections.
5. There is no explicit mention of linearizing the infrared radiation sensors. This requirement is not met.
6. The sampling rate for the ADC is described as ≥1 kHz per channel, suggesting sequential sampling. However, there is no explicit mention of the sampling order strategy, so this requirement is not met.
7. The sampling frequency of the ADC for the pressure sensors is ≥8 kHz and ≥1 kHz for the temperature sensor array, which is above the required 800 Hz.
8. The anti-aliasing filter for the pressure sensors has a cutoff frequency of 400 Hz and is a 4th-order filter, which results in a roll-off rate of -24 dB/octave. This should ensure a gain of at least -20 dB at half the sampling frequency, fulfilling this requirement.
9. The low-pass filters for both pressure and temperature have cutoff frequencies of 400 Hz and 450 Hz, respectively, which are higher than 400 Hz and lower than half the total sampling frequency, assuming the sampling frequency is at least 1 kHz per channel.
10. Low-pass filters are included before the multiplexer(s) to reduce aliasing, as mentioned in the filter block sections for pressure and temperature.
11. The project uses multiplexers to choose channels for both the pressure and temperature sensor arrays.
12. The multiplexers are described as analog multiplexers, which are inherently solid-state devices, thus fulfilling this requirement.