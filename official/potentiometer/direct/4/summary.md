Pendulum Angle Measurement System Design

This document summarizes the design of an analog instrumentation system for calculating the angle of a pendulum using a potentiometer, which outputs its measurements to a Data Acquisition System (DAQ).

PENDULUM_POTENTIOMETER:
- Potentiometer Type: Linear Rotary Potentiometer
- Model Suggestion: Bourns 3590S-2-103L 10K Ohm Rotary Wirewound Precision Potentiometer
- Key Characteristics: 10 kOhms, ±0.5% linearity, capable of 0-270 degree rotation (mechanical range will be limited to desired 45 to 135-degree range).

SIGNAL_CONDITIONING:
- Offset Adjustment: Not required, the potentiometer output at 90 degrees (steady position) is 0V, which is the center of the DAQ's input range.
- Voltage Scaling: A scaling amplifier is required to adjust the potentiometer's output range (-5V to +5V) to the DAQ's input range (-7V to +7V).
  - Scaling Amplifier Configuration: Non-inverting amplifier with R1 = 10kΩ and R2 = 4kΩ to achieve a gain of 1.4.

BAND_STOP_FILTER:
- Twin-T Notch Filter for 50 Hz and 60 Hz combined with an Active Band-Reject Filter for sharper attenuation.
  - 50 Hz Twin-T Components: Capacitor value at 10 nF; Resistor value calculated at 318.3 kΩ, using standard value of 330 kΩ.
  - 60 Hz Twin-T Components: Capacitor value at 10 nF; Resistor value calculated at 265.3 kΩ, using standard value of 270 kΩ.
  - Operational Amplifier for Active Stage: GBWP > 1 MHz, with a gain of 1 at non-notch frequencies.

BUFFER_AMP:
- Configuration: Unity-gain buffer using an operational amplifier.
- Op-Amp Selection: LM358 or TL071, powered by the same -10 V to +10 V source as the potentiometer.

SCALING_AMP:
- Configuration: Non-inverting amplifier with gain set to 1.4 to scale the voltage range of -5V to +5V from the potentiometer to -7V to +7V for the DAQ.
- Component Values: R1 = 10kΩ, R2 = 4kΩ.

ANTI_ALIASING_FILTER:
- Filter Type: Active Butterworth Low-Pass Filter
- Filter Order: Second-order (2 poles)
- Cutoff Frequency (f_c): 250 Hz
- Configuration: Sallen-Key topology.
- Operational Amplifier: TL072 or similar low-noise op-amp.
- Precision components: 1% tolerance for resistors and capacitors.

DAQ:
- ADC Type: Successive Approximation Register (SAR) ADC
- Required Resolution: 10 bits, as calculated from the desired voltage change per degree of pendulum's motion and the system's full-scale voltage range.
- Sampling Rate: 1000 samples per second or higher (2 kHz recommended for better performance).
- Input Range: Compatible with +/- 7 V from the signal conditioning stage.
- Linearity: INL and DNL less than 1 LSB for accurate measurement.

Overall, the system is designed to measure the angle of the pendulum ranging from 45 to 135 degrees, where 90 degrees corresponds to the steady position. The output from the potentiometer is conditioned to fit the DAQ's input range of +/- 7 V, with careful consideration given to noise attenuation and signal integrity.