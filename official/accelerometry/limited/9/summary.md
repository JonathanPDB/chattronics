Design of a Portable Low-Frequency Vibration Measurement Device

This solution encompasses the design of a portable device that measures low-frequency vibration using a piezoelectric accelerometer and includes signal conditioning and output amplification.

Piezoelectric Accelerometer:
- Selected Sensor: TE Connectivity Model 832M1-0500
- Sensitivity: 100 pC/g
- Measurement Range: ±50 g
- Operating Temperature Range: -55°C to +125°C
- Supply Voltage Requirement: 18 to 30 VDC (additional voltage regulation circuit may be required)

Charge Amplifier:
- Assumed Peak Acceleration: 0.05 g (from 5 cm peak oscillation)
- Maximum Charge (Qmax): 5 pC (calculated using 100 pC/g sensitivity)
- Feedback Capacitor (C_f): 5 pF
- Feedback Resistor (R_f): Approximately 127 Megaohms (standard value: closest is 120 Megaohms)
- Operational Amplifier: Low-noise JFET input op-amp, e.g., TL071 or OPAx134 series
- Supply Voltage: ±5V to ±15V (selected based on power supply availability)
- Low-Frequency Cutoff: 0.25 Hz (3 dB down at this frequency)
- Amplifier Output Voltage Formula: V_out = Q/C_f

High-Pass Filter:
- Topology: 1st order Sallen-Key High-Pass Filter
- Cutoff Frequency (f_c): 0.25 Hz (roll-off to minimize phase distortion)
- Resistor (R): 620 kΩ (standard value close to the calculated 636.62 kΩ requirement)
- Capacitor (C): 1 µF
- Power Supply: ±5V to ±15V (based on portable device constraints)

Low-Pass Filter:
- Option 1: 4th-order Active Butterworth Low-Pass Filter
  - Cutoff Frequency: 5 Hz
  - Passband Ripple: None (Butterworth characteristic)
  - Attenuation: 24 dB per octave beyond the cutoff frequency
  - Operational Amplifier: e.g., TL072
  - Power Supply: ±15V

- Option 2: 2nd-order or 3rd-order Active Chebyshev Low-Pass Filter
  - Cutoff Frequency: 5 Hz
  - Passband Ripple: 0.5 dB
  - Attenuation: 36 dB per octave for a 3rd-order filter beyond the cutoff frequency
  - Operational Amplifier: e.g., TL072
  - Power Supply: ±15V

Output Amplifier:
- Topology: Non-inverting Operational Amplifier
- Gain (A_v): 10x
- Desired Output: 1 V peak-to-peak
- Supply Voltage: +5 V (single-supply operation assumed for portability)
- Feedback Resistor (R_f): 90 kΩ
- Ground-Reference Resistor (R_g): 10 kΩ
- Load Impedance: ≥ 10 kΩ
- Operational Amplifier: Low-noise, rail-to-rail output op-amp capable of single 5 V supply operation, e.g., OPAx340 series or AD8605

The device's design is organized into distinct blocks following the signal path from the accelerometer to the final output. All components were selected based on typical values and requirements for portable instrumentation. Component values were calculated to ensure the desired behavior of each stage, considering the overall device objectives of measuring low-frequency vibrations with a 1 V peak-to-peak output while minimizing offset and maximizing signal integrity.