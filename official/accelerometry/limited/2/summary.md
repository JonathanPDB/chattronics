Design of a Portable Low-Frequency Vibration Measurement Device

Piezoelectric Sensor Block:
- Sensor Type: Piezoelectric accelerometer
- Recommended Model: Endevco model 2271A/AM20 or similar
- Sensitivity: 100 pC/g
- Measurement Range: Assuming a peak acceleration of 0.8 g based on a 5 cm peak amplitude at 2 Hz frequency
- Operational Temperature Range: Typically -55°C to +125°C for industrial-grade sensors

Charge Amplifier Block:
- Peak Charge Calculation: Q_peak = Sensitivity × A_peak = 100 pC/g × 0.8 g = 80 pC
- Feedback Capacitor (C_f): Chosen to be 10 pF for practical purposes
- Feedback Resistor (R_f): Calculated for a -3 dB cutoff at 0.25 Hz
  R_f = 1 / (2π × 0.25 Hz × 10 pF) ≈ 63.7 MΩ (standard value approximation)
- Op-amp Suggestion: Low input bias current, low offset voltage, e.g., AD8628 or LMC662

Low-Pass Filter Block:
- Filter Type: Second-order Sallen-Key Butterworth low-pass filter
- Cutoff Frequency: 0.25 Hz
- Component Values (Example): R1 = R2 = 100 kΩ, C1 = C2 = 10 μF, using TL072 op-amp or similar

High-Pass Filter Block:
- Filter Type: Second-order Sallen-Key Butterworth high-pass filter
- Cutoff Frequency: 0.25 Hz to achieve a 3 dB down point
- Component Values (Example): R1 = R2 = 100 kΩ, C1 = C2 ≈ 10 μF, using TL072 op-amp or similar

Voltage Amplifier Block:
- Topology: Non-inverting operational amplifier configuration
- Expected Gain (A_v): 10 (20 dB), assuming 100 mV peak to peak from charge amplifier
- Power Supply: A single supply configuration, with the supply voltage a few volts higher than the output swing required
- Gain Setting Resistors (Example): R1 = 10kΩ, R2 = 100kΩ
- Op-amp Suggestion: Low offset voltage, low drift, e.g., OP07 or AD8628

Anti-Aliasing Filter Block:
- Filter Type: Second-order Butterworth low-pass filter
- Chosen Cutoff Frequency: 10 Hz based on a conservative ADC sampling rate assumption
- Component Values (Example): C = 0.1 µF, R ≈ 159.15 kΩ

Analog to Digital Converter (ADC) Block:
- ADC Type: Sigma-Delta (Σ-Δ)
- Resolution: 16-bit or higher
- Sampling Rate: At least 10 samples per second (10 Hz)
- Input Range: 1 V peak-to-peak
- Reference Voltage: Internal for portability
- Power Consumption: Low-power variant
- Interface: Serial (SPI or I2C)

General Design Considerations:
- All op-amps and active components chosen should have specifications that support low noise, low offset voltage, and stability over temperature variations.
- The device should use high-quality components with low temperature coefficients for stability over varying temperatures.
- The anti-aliasing filter and ADC sampling rate should be chosen to effectively prevent aliasing while minimizing power consumption for portability.
- The overall design should consider practical aspects such as power efficiency, the physical size of components, integration ease, and cost-effectiveness.