Pendulum Angle Measurement System Design

This system is designed to calculate the angle of a pendulum using a linear potentiometer and interface the analog output with a Data Acquisition (DAQ) system. The system architecture consists of several key stages: the potentiometer sensor, buffering, low-pass filtering to attenuate unwanted frequencies, signal attenuation to match the DAQ's input voltage range, further anti-aliasing filtering, and finally, digital conversion by the DAQ system.

Potentiometer:
- Model: Bourns 3590S-2-103L 10kOhm precision potentiometer
- Characteristics: 10-turn, linear response, capable of measuring angles from 45 to 135 degrees with good precision and resolution.

Buffer:
- Topology: Unity Gain Buffer (Voltage Follower) using an operational amplifier
- Parameters: High input impedance (> 1MΩ), low output impedance (< 10Ω), unity gain (Gain = 1).

Low-Pass Filtering (To Remove 50 and 60 Hz Noise):
- Filter Type: Active Low-Pass Butterworth Filter
- Order: Second-order (2 poles) for 12 dB/octave roll-off
- Cutoff Frequency: 5 Hz to ensure significant attenuation of 50 and 60 Hz noise
- Components: Calculated based on the cutoff frequency using standard op-amp and passive components (resistors and capacitors).

Attenuator:
- Voltage Divider: R1 = 4.3 kOhms, R2 = 10 kOhms, for an attenuation ratio of 0.7 to scale the -10V to +10V signal to the DAQ's ±7V range.
- Voltage Clamping: Zener Diodes rated for 7.5V to prevent overvoltage conditions.

Anti-Aliasing Filter (Before DAQ):
- Filter Type: Active Low-Pass Butterworth Filter
- Order: Second-order (2 poles)
- Cutoff Frequency: 150 Hz, chosen to avoid aliasing and to accommodate the DAQ's 1000 samples per second sampling rate.
- Components: Designed using standard op-amps, resistors, and capacitors.

Data Acquisition (DAQ):
- ADC Type: Successive Approximation Register (SAR)
- Sampling Rate: At least 2000 samples per second to maintain a margin above the DAQ's rate and to prevent aliasing.
- Resolution: 12-bit, which offers 4096 discrete levels, providing a resolution of less than 0.1 degrees if we assume a full 360-degree swing.
- Input Range: Matched to the attenuator's output range (±7V).

The overall system is designed with standard commercial-grade components, assuming typical laboratory conditions for temperature and humidity. The combination of a high-quality potentiometer, careful signal conditioning, and a suitable ADC in the DAQ system ensures accurate and reliable angle measurements of the pendulum's swing.