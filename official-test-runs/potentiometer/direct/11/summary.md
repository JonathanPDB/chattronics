Pendulum Angle Measurement System Design

Potentiometer-Based Angle Sensor:
- Type: Linear rotary potentiometer.
- Model: Bourns 3590S-2-103L 10K Ohm Rotary Wirewound Precision Potentiometer.
- Operating Voltage: -10V to 10V.
- Operating Current: 1 mA at 10V (I = V/R).
- Temperature Range: Assuming 0°C to 40°C for room temperature operation.
- Sensitivity and Resolution: As fine as the ADC in the DAQ system can detect. Practical sensitivity will depend on mechanical dither and electrical noise.
- Angle Measurement Range: 45 to 135 degrees, with 90 degrees being the steady position.

Amplification Block:
- Topology: Inverting operational amplifier (op-amp) configuration.
- Gain: -0.7 (For the range from -5V to 5V mapped to DAQ's +/-7V input range).
- Input Resistor (Rin): 10 kOhms.
- Feedback Resistor (Rf): 7 kOhms (Gain = -Rf/Rin = -7/10 = -0.7).
- Op-amp: Precision op-amp with low offset voltage (e.g., OPA2277).
- Power Supply: +/- 15 V (common for most precision op-amps).
- Protection: 1 kOhm series resistor and clamping diodes (e.g., 1N4148) at the output.

Filtering Block:
- Topology: Active Twin-T Notch Filter for attenuation of 50 and 60 Hz frequencies.
- Notch Frequencies: 50 Hz and 60 Hz with at least 20 dB attenuation.
- Quality Factor (Q): Approximately 10 to 15 for sharp notches without excessive peaking.
- Bandwidth: 5 Hz around each notch frequency.
- Passband: 0 to 20 Hz to include the pendulum's motion frequencies.
- Low-Pass Filter Cutoff Frequency: 100 Hz or lower for anti-aliasing purposes.

Protection Block:
- Components: Transient Voltage Suppression (TVS) diodes, series resistors, and a low-pass filter.
- TVS Diodes: Bidirectional with a breakdown voltage slightly above the DAQ's maximum input voltage (e.g., P6KE8.2CA).
- Series Resistors: Between 100 to 500 ohms (e.g., 330 ohms, 1/4 Watt, 1% tolerance metal film resistor).
- Low-Pass Filter: Cutoff frequency around 250 Hz with a 330-ohm resistor and a 1.5 uF capacitor (X7R dielectric ceramic capacitor, 1.5 uF, 25V).

Analog to Digital Converter (ADC) Block:
- Type: Successive Approximation Register (SAR) ADC.
- Resolution: At least 12 bits for a 4096-level differentiation across the input voltage range.
- Sampling Rate: 1000 samples per second or higher to match the DAQ's rate.
- Input Voltage Range: +/- 7 V to match the DAQ's input voltage range.
- Interface: SPI or I2C for communication with data processing systems.

Note: Component values and specifications provided are based on standard practices and assumptions made due to the lack of specific user requirements. These are subject to change based on empirical testing and further details of the application environment.