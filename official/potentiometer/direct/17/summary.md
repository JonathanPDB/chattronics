Pendulum Angle Measurement System Design

The proposed system is designed for calculating the angle of a pendulum using a potentiometer. The system comprises several analog components for signal conditioning before the signal is digitized by a Data Acquisition System (DAQ). Here is a detailed design for each block of the system:

Potentiometer:
- Type: Linear taper
- Resistance: 10 kOhms
- Power Rating: Standard 0.5 W
- Linearity: 1% or better
- Temperature Range: Industrial grade, -40°C to +85°C
- Life Cycle: Rated for 1 million cycles or more
- Mounting: Panel mount, robust mechanical design
- Suggested Model: Bourns 3590 Precision Potentiometer

Buffer Amplifier:
- Topology: Voltage Follower (Unity Gain Buffer)
- Op-Amp Choice: General-Purpose Precision Op-Amp (e.g., LM324 or MCP6024)
- Gain: 1 (Unity Gain)
- Power Supply: Powered by the same -10 to 10 Volts source as the potentiometer
- Bandwidth: A few hundred Hz to exceed the maximum frequency of interest
- Output Current: Capable of supplying at least 10 mA
- Noise: Low noise figure, less than 10 nV/√Hz

Scaling Amplifier:
- Topology: Non-inverting amplifier configuration using an operational amplifier
- Gain Calculation: Approximate gain of 1.05, determined by the voltage swing of the potentiometer and the DAQ's input range
- Resistor Values for Gain: R1 (feedback) ≈ 2.35kΩ and R2 (grounded) = 47kΩ
- Op-Amp Selection: Precision op-amp with rail-to-rail output (e.g., TLV2371)

Notch Filter (50 Hz and 60 Hz):
- Topology: Active Twin-T Notch Filter
- Notch Frequency: Centered at 50 Hz and 60 Hz, respectively
- Attenuation: -40 dB at the notch frequencies
- Q Factor: 10 for both filters
- Bandwidth: Approximately 5 Hz around each notch frequency
- Op-Amps: Low-noise, precision op-amps (e.g., LM358 or TL072)
- Resistors and Capacitors: 1% metal film resistors; NP0/C0G ceramic or film capacitors

Low-Pass Filter:
- Type: Active Low-Pass Butterworth Filter
- Order: Second-order (Two-pole)
- Cutoff Frequency: 200 Hz
- Topology: Sallen-Key or Multiple Feedback
- Resistor Values: R1 = R2 = 10 kOhm
- Capacitor Values: C1 and C2 = 82 nF (nearest standard value to 79.58 nF)
- Attenuation Rate: -12 dB per octave

Data Acquisition System (DAQ):
- ADC Type: Successive Approximation Register (SAR)
- Resolution: 12 bits
- Sampling Rate: Minimum 1000 samples per second
- Input Voltage Range: At least +/- 7 V to match the conditioned signal
- Interface: SPI or I2C

All components should be selected with consideration for availability and environmental conditions they will operate under. Precision resistors and capacitors are recommended for all filter designs to ensure accuracy and stability. The op-amps should provide low noise and low distortion to maintain signal integrity through the system. The design prioritizes simplicity and robustness while satisfying the project's requirements.