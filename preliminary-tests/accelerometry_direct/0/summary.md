Portable Low-Frequency Vibration Measurement Device Design

This summary compiles the details of a portable device designed to measure low-frequency vibrations using a piezoelectric accelerometer, complete with charge amplifier, low-pass filter, gain stage, and output buffer.

Piezoelectric Accelerometer:
- Model: PCB Piezotronics model 352C33 or similar with 100 pC/g sensitivity
- Measurement Range: Must handle >0.04 g to measure 5 cm peak at 2 Hz without saturating
- Frequency Range: Flat response down to at least 0.25 Hz
- Temperature Range: Operable within -40 to +85°C
- The charge generated for peak acceleration:
  a_peak_g = (2 * π * 2)^2 * 0.05 m / 9.80665 m/s² ≈ 0.04 g
  Q_peak = 100 pC/g * 0.04 g = 4 pC

Charge Amplifier:
- Feedback Capacitor (Cf): 8 nF (to achieve 0.5 V output for 4 pC charge)
- Feedback Resistor (Rf): 79.6 MΩ (calculated for a cutoff at 0.25 Hz)
- Op-Amp: Low bias current FET input op-amp (e.g., TL071)
- Bias current: Typically 1 pA
- Power Supply: Assumed ±15V
- Output offset voltage: Designed to be <10 mV

Low-Pass Filter:
- Topology: Active Butterworth filter
- Order: Second-order (two-pole)
- Cutoff Frequency: -3 dB at 0.25 Hz
- Resistors: 10 kΩ (precision resistors with 1% tolerance)
- Capacitors: 68 μF (NP0/C0G for low temperature coefficient)
- Op-Amp for Sallen-Key configuration: OPAx233 series or similar
- Bandwidth: Filter's GBW should be sufficiently above 2Hz (at least 10 times)

Gain Stage:
- Configuration: Non-inverting operational amplifier
- Gain: 1 (assuming charge amplifier is designed to output 0.5 V for 4 pC charge)
- Op-Amp: Low-noise, low-offset, precision op-amp with sufficient GBWP (e.g., OPAx333 series)
- Power Supply: Single supply of 5V or 3.3V

Output Buffer:
- Configuration: Voltage follower (unity gain buffer)
- Op-Amp: Low-offset, low-noise, rail-to-rail output (e.g., OPAx333 series)
- Load Impedance: Capable of driving 10 kΩ to 100 kΩ
- Bypass Capacitors: 0.1 µF ceramic capacitors
- Power Supply: 5V single supply

With all the components and configurations outlined, the device can measure vibrations at 2 Hz with a peak displacement of 5 cm and output a 1 V peak-to-peak voltage. The offset voltage will be kept below 10 mV, and the low-frequency response will be 3 dB down at 0.25 Hz.