Portable Low-Frequency Vibration Measurement Device Design

This design summary outlines the components and calculations for a portable device that measures low-frequency vibrations using a piezoelectric accelerometer. The device architecture includes a piezoelectric accelerometer, charge amplifier, low-pass filter, anti-aliasing filter, analog-to-digital converter (ADC), and digital display.

Piezoelectric Accelerometer:
- Type: Charge mode piezoelectric accelerometer
- Sensitivity: 100 pC/g
- Operating Temperature Range: -40°C to +85°C
- Example Model: Endevco 2221F (or similar models from PCB Piezotronics or Kistler)

Charge Amplifier:
- Operational Amplifier: Low input bias current and low offset voltage op-amp (e.g., OPAx277 series)
- Feedback Capacitor (C_f): 150 pF
- Feedback Resistor (R_f): 4.7 MΩ
- Calculation for Gain: Given the sensitivity of 100 pC/g and maximum expected acceleration of 1g, the charge developed is Q = 100 pC. The gain (G) is 1 V / 100 pC = 10^4 V/C. Using \(C_f = Q / V_{out}\), we calculate \(C_f\) as 100 pC / 1 V = 100 pF. A standard value of 150 pF is selected. To determine \(R_f\), we use \(R_f = 1 / (2\pi f_c C_f)\), where \(f_c = 0.25 \text{Hz}\). This results in \(R_f\) ≈ 4.24 MΩ, and we choose the nearest standard value of 4.7 MΩ.

Low-pass Filter (0.25 Hz cutoff):
- Filter Type: Second-order Butterworth active low-pass filter
- Operational Amplifiers: Low-noise (e.g., OPAx177 or AD8676)
- Capacitors: \(C1 = C2 = 10 \mu F\) (standard value)
- Resistors: \(R1 = R2 = 62 k\Omega\) (nearest standard value to calculated 63.7 kΩ)
- Roll-off Rate: -12 dB/octave beyond 0.25 Hz cutoff frequency

Anti-aliasing Filter (2 Hz cutoff):
- Filter Type: Butterworth
- Order: Second-order
- Cutoff Frequency: 2 Hz
- Component Values: R = 82 kΩ, C = 1 µF
- Configuration: Single-stage RC low-pass filter

Analog-to-Digital Converter (ADC):
- Topology: Sigma-Delta (Σ-Δ)
- Resolution: 16 bits
- Sampling Rate: ≥20 SPS (samples per second)
- Input Voltage Range: ±1.5 V or wider
- Input Type: Differential
- Power Consumption: Low-power design (battery operation consideration)
- Interface: SPI or I2C

Digital Display Processing:
- Display Technology: LCD (for low power consumption and versatility)
- Resolution and Size: Moderate resolution to display numeric values and simplified graphs
- User Interface: Basic buttons for navigation and configuration
- Refresh Rate: A few updates per second (for numerical values)
- Power Conservation: Low-power modes with automatic backlight dimming and sleep functionality
- Software Interface: Firmware developed in C or C++ on an embedded microcontroller platform

The overall design is focused on portability, accuracy, and user-friendliness while ensuring low power consumption for extended battery life. Component values and specific models are chosen based on common standards for the application of low-frequency vibration measurement.