Design of a Portable Low-Frequency Vibration Measurement Device

The design of the portable low-frequency vibration measurement device is based on the following components and calculations:

Piezoelectric Accelerometer:
- Sensor type: PCB Piezotronics 352A21 (substitute for the specified 100 pC/g sensitivity)
- Sensitivity: Approx. 10 mV/g (After charge to voltage conversion with the charge amplifier)
- Measurement Range: ±500 g peak
- Frequency Response: 0.5 to 10000 Hz (±5%)
- Operating Temperature Range: -54 to +121 °C

Charge Amplifier:
- Conversion factor (Cf) to produce 1 Vpp from peak charge: Cf ≈ 1000 pF (1 nF)
- Feedback resistor (Rf) for 0.25 Hz low-frequency cutoff: Rf ≈ 620 kΩ
- Operational amplifier: Low-noise, FET input, ultra-low bias current operational amplifier (e.g., AD8628)

Low-Pass Filter:
- Topology: Second-order Butterworth filter
- Cutoff Frequency: 2 Hz
- Components: Standard value resistors (10 kΩ) and capacitors calculated for a 2 Hz cutoff frequency
- Damping Factor: 0.707 (maximally flat Butterworth response)

High-Pass Filter:
- Topology: Second-order Sallen-Key configuration
- Cutoff Frequency: 0.25 Hz
- Components: R1 = R2 = 63.7 kΩ (rounded to standard values), C1 = C2 = 10 μF
- Damping Factor: 0.707 (maximally flat Butterworth response)

Gain Stage:
- Non-Inverting Operational Amplifier Configuration
- Desired Gain: 10 (to achieve 1 Vpp output from a 100 mVpp input)
- Power Supply: Battery-operated, with voltage regulation as needed for the selected op-amp

Anti-Aliasing Filter:
- Topology: Second-order Butterworth active low-pass filter
- Cutoff Frequency: 4.41 kHz (1/10 of the ADC sampling rate)
- Components: Assuming 10 nF capacitor, resistor ≈ 3.6 kΩ
- Operational amplifier should have low noise and bandwidth greater than 10 times the cutoff frequency

Analog to Digital Converter (ADC):
- Resolution: 12 bits (for a voltage resolution of approx. 244 µV per bit)
- Sampling Rate: 16 Hz (minimum based on the Nyquist theorem for a 2 Hz maximum frequency)
- Topology: Successive Approximation Register (SAR) ADC
- Interface Protocol: SPI for fast data transfer
- Power Supply: 3.3 V to 5 V, depending on system design
- Additional Features: Internal voltage reference, differential input capability

Output Block:
- Buffer Amplifier: Operational amplifier in voltage follower configuration (e.g., AD8628)
- Level Shifter: As needed based on required output voltage range
- Voltage Protection: TVS diodes and series resistors for over-voltage protection
- Communication Protocol: I2C or SPI, or Bluetooth module (e.g., HC-05) for wireless communication
- Output Connector: Standard connectors based on analog or digital output requirements
- Isolation: Optional, using optical isolators or transformers

The complete solution is designed to meet the requirements of a portable device capable of measuring low-frequency vibration with an output of 1 Vpp and maintaining a low offset voltage of less than 10 mV. The design is battery-powered, and components are chosen for low power consumption and industrial temperature range operation.