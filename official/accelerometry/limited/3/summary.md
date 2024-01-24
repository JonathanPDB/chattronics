Design of a Portable Low-Frequency Vibration Measurement Device

Piezoelectric Accelerometer:
- Sensor Type: Ceramic-based piezoelectric accelerometer.
- Sensitivity: 100 pC/g.
- Operating Temperature Range: -40°C to 85°C.
- Resolution: Fine enough to detect changes well below 0.02 g (0.198 m/s²).
- Supply Voltage: Typically ±5V to ±15V.
- Suggested Model: PCB Piezotronics 352C33 or similar with proper charge sensitivity conversion.

Charge Amplifier:
- Topology: Operational amplifier in a feedback configuration with a capacitive feedback element.
- Sensitivity of Accelerometer: 100 pC/g.
- Peak acceleration: 0.064 g (0.628 m/s² at 5 cm peak amplitude and 2 Hz).
- Expected peak charge: 6.4 pC.
- Desired peak-to-peak voltage output: 1 V.
- Feedback Capacitance (Cf): 12.8 pF calculated from Cf = Q / Vpp (6.4 pC / 0.5 V).
- Gain (G): 78.125 MV/pC (1 / Cf).
- Op-Amp Selection: Low input bias current (< 1 nA) and low noise (< 1 µV/√Hz).
- Power Supply: Single-supply, within 3.3V to 5V.
- High-Pass Filter: To achieve a 3 dB point at 0.25 Hz.

Low-Pass Filter:
- Topology: Second-order Sallen-Key Butterworth.
- Cutoff Frequency: 0.25 Hz.
- Order: Second-order (12 dB/octave).
- Damping Factor (Q): 1/√2 for Butterworth response.
- Passband Ripple: Zero (maximally flat).
- Components: Assuming R = 10 kΩ, C ≈ 63.66 µF (use standard value or combination).

Offset Compensation:
- Topology: Precision op-amp with trimmer potentiometer for manual adjustment.
- Maximum Offset Voltage: Below 10 mV.
- Op-Amp: Low input bias current and low offset voltage, such as ADA4528-1 or OPA188.
- Trimmer Potentiometer: 10kΩ multi-turn.

Output Buffer:
- Topology: Unity Gain Voltage Follower.
- Load Impedance: Capable of driving >1 kΩ to 10 kΩ loads.
- Gain: Unity (Gain = 1).
- Supply Voltage: 5V (Single supply).
- Output Swing: Rail-to-rail or close to rail-to-rail.
- Bandwidth: >20 Hz for margin.
- Slew Rate: >0.5 V/µs.
- Example Op-Amps: Texas Instruments OPA340 or Analog Devices AD8605.

ADC:
- Type of ADC: Sigma-Delta (Σ-Δ).
- Resolution: 24-bit.
- Sampling Rate: 256 Hz (oversampling for the 2 Hz signal).
- Interface: SPI.
- Power Consumption: Low-power for portability.
- SNR and Dynamic Range: High to capture low-level signals.
- Voltage Reference: Precision voltage reference for accuracy.

Anti-Aliasing Filter:
- Cutoff Frequency: Approximately 8 Hz.
- Filter Order: Second-order (two-pole) Butterworth.
- Capacitors (C): Approximately 20 nF (calculated based on a 1 kΩ resistor).
- Resistors (R): Standard 1 kΩ (or adjusted for practical capacitor values).