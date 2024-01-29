Pendulum Angle Measurement System Design

This system design details the electronics components and architecture required to measure the angular position of a pendulum using a potentiometer and to interface the analog signal with a Data Acquisition (DAQ) system.

1. Potentiometer:
- Type: Precision conductive plastic potentiometer
- Total Resistance: 10 kΩ
- Electrical Rotation: At least 270 degrees to cover the pendulum's 45 to 135 degrees range
- Linearity: ±0.1% or better
- Power Rating: 0.5 to 1 W
- Operating Temperature Range: -55°C to +125°C
- Example Component: Bourns 3549 series

2. Buffer Stage:
- Configuration: Voltage follower with input attenuation
- Input Resistance (Rin): 1 MΩ
- Output Resistance (Rout): < 100 Ω
- Gain (A): 0.7 (to scale the potentiometer's ±10V output to the DAQ's ±7V input range)
- Bandwidth: ≥ 10 times the highest frequency of pendulum movement
- Slew Rate: ≥ 0.5 V/µs
- Suggested Op-amp: General-purpose operational amplifier with a suitable supply voltage (e.g., TL072 or LM358)

3. Calibration Circuit:
- Configuration: Non-inverting op-amp with variable gain and offset
- Op-amp: Rail-to-rail precision op-amp (e.g., OPAx177 or AD8476)
- Resistors: R1 = 10kΩ, R2 = 10kΩ, R3 = 10kΩ, Rs (series protection resistor) = 1kΩ
- Potentiometers: P1 = 0-10kΩ (gain adjustment), P2 = 0-10kΩ (offset adjustment)
- Zener Diodes: D1, D2 = 6.8V (output voltage protection)

4. Low-Pass Filter:
- Topology: Active Butterworth low-pass filter, 2nd order
- Cutoff Frequency (f_c): 5 Hz
- Op-amp: TL072 or LM358
- Resistors: R = 10 kΩ
- Capacitors: C = 3.1831 μF
- Attenuation at 50 Hz: ≥ 60 dB
- Quality Factor (Q): 0.707 for maximally flat passband

5. Anti-Aliasing Filter:
- Topology: Active Butterworth low-pass filter, 2nd order
- Cutoff Frequency: Approximately 150 Hz
- Op-amp: TL072 or LM358
- Components chosen to meet required cutoff frequency and attenuation characteristics

6. Data Acquisition (DAQ) Input - ADC Specifications:
- Type: Successive Approximation Register (SAR) ADC
- Resolution: 12 bits or higher
- Sampling Rate: ≥ 2000 samples per second
- Input Voltage Range: Bipolar, at least -10V to +10V
- Integral Nonlinearity (INL) and Differential Nonlinearity (DNL): < 1 LSB
- Anti-Aliasing Filter Cutoff: Just above 500 Hz (Nyquist frequency)
- Interface: SPI or I2C (to be compatible with a wide range of microcontrollers and DAQ systems)

Calculations for the Low-Pass Filter components:
f_c = 1/(2π√(R1*R2*C1*C2))
Assuming R1 = R2 and C1 = C2,
R = 10 kΩ and C = 3.1831 μF, the cutoff frequency (f_c) is approximately 5 Hz.

The design components, parameters, and calculations provided ensure the accurate measurement of a pendulum's angle and its proper interfacing with the DAQ system without overdriving or damaging the components. A calibration routine is required to ensure the voltage output of the potentiometer accurately corresponds to the angular range of the pendulum. All electronic design choices are made with consideration of a power supply between -10V and +10V, the DAQ's voltage and sampling rate limitations, and the need to filter out 50/60 Hz noise.