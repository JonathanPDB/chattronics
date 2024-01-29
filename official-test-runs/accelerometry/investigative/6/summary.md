Design of a Portable Low-Frequency Vibration Measurement Device

Piezoelectric Sensor:
- Sensor Type: Piezoelectric accelerometer
- Suggested Model: PCB Piezotronics 352A60
- Sensitivity: 100 pC/g (adjusted in charge amplifier)
- Frequency Range: 0.06 to 6000 Hz (with the relevant range being up to 2 Hz)
- Operating Temperature Range: -54 to +121 °C
- Electrical Isolation: Base isolated
- Size and Weight: Small and lightweight for portability

Charge Amplifier:
- Gain (G) = V_output_peak_to_peak / Q_max = 1V / 100 pC = 10^7 V/C
- Practical Feedback Resistor (Rf) = 10 MΩ
- Feedback Capacitor (Cf) calculated for a lower 3 dB point at 0.25 Hz, Cf = τ / Rf ≈ 0.64 s / 10 MΩ = 64 nF
- Selected Operational Amplifier: Low-noise precision op-amp with high input impedance and low bias current

Low Pass Filter:
- Filter Type: Active Butterworth low-pass filter
- Cutoff Frequency: Slightly above 2 Hz, adjusted to around 10 Hz for design purposes
- Filter Order: 2 (12 dB/octave or 40 dB/decade)
- Component values (assuming a standard op-amp): R = 1.5 MΩ, C = 10 nF
- Damping Factor (ζ) for Butterworth: ζ = 1/√2

Amplifier:
- Topology: Instrumentation Amplifier
- Gain (G) assuming full-scale ADC range match: G = 3.3 (if ADC range is 0-3.3V and sensor output is 1V peak to peak)
- Operational Amplifiers: AD620, INA114, or similar precision op-amps
- Power Supply: Dual supply (e.g., ±5V) or single supply with rail-to-rail op-amps
- Additional Features: Programmable gain for different sensitivities or measurement ranges

Anti-Aliasing Filter:
- Filter Type: Active Butterworth low-pass filter
- Filter Order: 2 (second-order)
- Cutoff Frequency (f_c): Slightly higher than 2 Hz, adjusted to around 5 Hz
- Topology: Active Sallen-Key
- Component values for f_c: Resistors (R1 = R2) = 10 kΩ, Capacitors (C1 = C2) = 3.3 uF (closest standard value to calculated 3.18 uF)
- Operational Amplifier: Selected based on bandwidth, slew rate, and noise characteristics

Additional Design Considerations:
- Ensure high input impedance across the signal path to avoid loading the piezoelectric sensor.
- Use low-noise design techniques, including shielding and proper PCB layout, to minimize EMI and noise pickup.
- Implement power-saving features for portability, such as power-down modes in the ADC and low-power op-amps.
- Consider environmental and mechanical constraints in the final enclosure design for durability and compliance with standards.
- Fine-tuning of component values may be required during prototyping due to component tolerances and actual behavior.