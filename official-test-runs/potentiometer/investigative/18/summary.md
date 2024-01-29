Pendulum Angle Measurement System Design

This design outlines an analog instrumentation system for measuring the angle of a pendulum using a linear potentiometer. The system consists of the following main blocks: Position Sensor (Potentiometer), Signal Conditioning Circuit, Buffer Amplifier, Anti-Aliasing Filter, and Analog-to-Digital Converter (ADC).

Position Sensor (Potentiometer):
- Type: Spectra Symbol SoftPot Membrane Potentiometer
- Operating Voltage: Up to 10V
- Resistance: 10kΩ
- Range of Motion: 45 to 135 degrees
- Linearity: ±0.5%
- Operating Temperature Range: -35°C to +80°C

Signal Conditioning Circuit:
- The circuit is designed to create a voltage divider with the potentiometer, outputting a voltage signal that swings linearly with the pendulum's angle within the range of +/- 7V to match the DAQ's input.
- Voltage Divider: A fixed resistor of 10kΩ is used in series with the potentiometer to achieve a midpoint voltage at the 90-degree position.

Buffer Amplifier:
- Topology: Voltage follower using an operational amplifier with rail-to-rail output capability.
- Op-Amp: Selected for rail-to-rail output to accommodate the +/- 7V output range, with a suggested supply voltage of +/- 15V for sufficient headroom.
- Gain: Unity (1) since the goal is to buffer the signal, not to amplify it.
- Feedback Resistor: A value between 10kΩ to 100kΩ for stability, if necessary.

Anti-Aliasing Filter:
- Topology: 2nd Order Butterworth Low-Pass Filter or 3rd Order Butterworth Low-Pass Filter
- Cutoff Frequency: Recommended between 200 Hz and 250 Hz, below the Nyquist frequency of 500 Hz (half of the DAQ's 1000 samples per second).
- Roll-off Rate: -12 dB/octave for 2nd order, -18 dB/octave for 3rd order.

Analog-to-Digital Converter (ADC):
- ADC Type: 12-bit Successive Approximation Register (SAR) ADC, suitable for the 1000 samples/sec DAQ system.
- Resolution: 12-bit (~3.4 mV per bit over +/- 7V range)
- Sampling Rate: At least 1000 samples/sec
- Input Voltage Range: +/- 7V
- Input Impedance: >10kΩ
- Communication Interface: Parallel or SPI

This project design assumes standard temperature, power supply, and mounting conditions. If additional specific requirements are later provided, components and calculations will need to be reviewed to meet those specific conditions.