Pendulum Angle Measurement System via Potentiometer and DAQ

The project involves designing an analog system to measure the angle of a pendulum using a potentiometer and a Data Acquisition (DAQ) setup. The system architecture comprises a potentiometer, buffer amplifier, scaling amplifier, band-stop filter, anti-aliasing filter, and a DAQ unit for signal sampling and conversion to digital data.

Potentiometer:
- Type: Rotary, Single-Turn
- Resistance Value: 10 kOhms
- Linearity: ±0.5% or better
- Resolution: Limited by the DAQ resolution and noise
- Tolerance: ±5% or better
- Operating Temperature Range: -10°C to +70°C
- Lifecycle: Rated for at least 1 million cycles
- Mounting: Servo mount
- Model Recommendation: Bourns 3590S-2-103L 10K Ohm Rotary Wirewound Precision Potentiometer

Buffer Amplifier:
- Topology: Non-inverting buffer
- Op-Amp Selection: Low offset voltage and low noise, e.g., AD8605 or OPA333
- Gain: 1 (unity gain)
- Power Supply: Must accommodate -10 V to +10 V range
- Input Impedance: High (1 MΩ or higher)
- Bandwidth: At least 10 kHz

Scaling Amplifier:
- Topology: Two-stage with Inverting and Non-Inverting Amplifier
- Inverting Stage: Unity-gain with Rf = Rin = 10 kOhms
- Non-Inverting Stage: Gain of 1.7 (to result in a net scaling gain of -0.7)
- Rin (Non-inverting stage): 10 kOhms
- Rf (Non-inverting stage): 7 kOhms (standard value)
- Op-Amp: General-purpose, low-noise, e.g., LM358 or TL081

Band-Stop Filter (To attenuate 50 and 60 Hz noise):
- Topology: Active Band-stop Filter (Twin-T or Operational Amplifier based)
- Order: Second-order
- Cutoff Frequencies: Notch frequencies at 50 Hz and 60 Hz
- Q Factor: Around 10
- Attenuation: At least 20 dB at notch frequencies
- Components: To be calculated based on op-amp characteristics

Anti-Aliasing Filter:
- Type: Butterworth Low-Pass Filter
- Order: 2nd order
- Cutoff Frequency: 250 Hz (below Nyquist frequency of 500 Hz)
- Attenuation at Cutoff: -3 dB (standard for Butterworth filter)
- Operational Amplifier: Low-noise with good frequency response, e.g., OPAx134 series
- Resistors: Precision metal film resistors, 1% tolerance
- Capacitors: Ceramic or film capacitors, 5% tolerance or better
- Component Values: To be calculated based on op-amp specs

DAQ System:
- Resolution: Minimum 12 bits
- Sampling Rate: 1000 samples per second (minimum)
- ADC Type: Successive Approximation Register (SAR)
- Input Voltage Range: Bipolar, capable of +/- 7 V
- Input Impedance: High (1 MΩ or higher)
- ADC Model: A 12-bit SAR ADC with at least 2 ksps, such as those offered by Texas Instruments or Analog Devices

Overall, the system must ensure that the angular displacement of the pendulum is correctly translated into a voltage range suitable for the DAQ's input, with appropriate filtering to eliminate power line interference and prevent aliasing. The potentiometer will be mechanically adjusted to restrict its electrical range to match the angle of the pendulum from 45 to 135 degrees, with 90 degrees corresponding to 0 V output. This configuration provides accurate angle measurements for analysis.