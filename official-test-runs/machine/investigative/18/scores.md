Score: 10
Explanations: 
The project summary covers the following requirements:

1. Both sensors have DC output, as indicated by the output signal specifications for the pressure sensor (1 mV max) and the temperature sensor (100 mV max). No mention of demodulation is made, satisfying the DC output requirement.
2. Both sensors are amplified, with recommended gains of 1000x for the pressure sensor and 10x for the temperature sensor.
3. The pressure sensor is amplified by an instrumentation amplifier, with models such as AD620, INA118, or INA128/INA129 series suggested, and a specified gain of 1000x. Although not explicitly stated, the use of a strain-gauge pressure sensor typically implies the use of a Wheatstone bridge.
4. An ADC is used with a specified minimum sampling rate of 12.8 ksamples/s aggregate across 16 channels, which implies the use of an ADC.
5. The temperature sensors are non-contact infrared radiation detectors, and while the summary does not explicitly mention linearization methods, the digital data processing capability of the computer could be used for digital linearization.
6. The sampling order strategy is implicitly covered by mentioning a 16-to-1 analog multiplexer and specifying an aggregate sampling rate, suggesting sequential sampling.
7. The specified minimum sampling rate of 12.8 ksamples/s aggregate across 16 channels results in 800 samples per second per channel, satisfying the minimum sampling frequency requirement.
8. The anti-aliasing filter is a 2nd or 4th order Butterworth filter with a cutoff frequency of 500 Hz. However, the requirement for a gain of at least -20 dB at half the sampling frequency is not explicitly confirmed. If we assume the sampling frequency to be 12.8 kHz (from the aggregate rate), half of this is 6.4 kHz. We need to ensure a -20 dB point at 6.4 kHz, but the provided cutoff frequency is 500 Hz, which would typically ensure a -20 dB point well before 6.4 kHz. Therefore, this requirement may not be met without further information on the filter's characteristics at 6.4 kHz.
9. The anti-aliasing filter's cutoff frequency is 500 Hz, which is higher than 400 Hz and lower than half the total sampling frequency (6.4 kHz), satisfying this requirement.
10. The anti-aliasing filter is mentioned, which would be positioned before the multiplexer to reduce aliasing, satisfying this requirement.
11. The project uses a 16-to-1 analog multiplexer (ADG1606 or CD74HC4067), satisfying the requirement for channel selection.
12. The multiplexers mentioned (ADG1606 or CD74HC4067) are solid-state devices.

The requirements that are not explicitly confirmed or met are:

8. The anti-aliasing filter's characteristics at half the sampling frequency (6.4 kHz) are not provided to confirm the -20 dB gain requirement.
5. The exact linearization method for the infrared sensors is not explicitly stated.

Overall, the project summary successfully covers 10 out of the 12 requirements, with requirements number 5 and 8 not being explicitly confirmed.