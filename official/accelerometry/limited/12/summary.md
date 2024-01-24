Portable Low-Frequency Vibration Measurement Device Design

This design summary outlines the key components and calculations for creating a portable device capable of measuring low-frequency vibrations using a piezoelectric accelerometer. The device comprises several stages: signal acquisition by the accelerometer, charge amplification, signal filtering, voltage amplification, anti-aliasing filtering (if needed), and digital conversion (if needed).

Piezoelectric Accelerometer
- Chosen Model: PCB Piezotronics Model 352A60 (or equivalent with required sensitivity and robustness)
- Sensitivity: 100 pC/g
- Operating Temperature Range: -40 to 85°C
- Vibration and Shock Resistance: At least 20 g and 5000 g, respectively
- Mounting: Compact and stud-mounted for portability and ease of installation

Charge Amplifier
- Op-Amp: LMC662 or equivalent with low input bias current
- Feedback Capacitor (Cf): 240 pF (standard value close to calculated 247.5 pF)
- Feedback Resistor (Rf): 2.7 MΩ (standard value close to calculated 2.65 MΩ)
- Input Resistor (Rin): 2.7 MΩ (to set the high-pass cutoff frequency at 0.25 Hz)
- Power Supply: ±5 V or ±12 V (dual supply for bipolar output swing)
- Gain Calculation:
  - Peak acceleration \( A = (2\pi \cdot 2\ \text{Hz})^2 \times 0.05\ \text{m} \approx 0.3951\ \text{m/s}^2 \)
  - Output charge at peak acceleration \( Q \approx 0.3951\ \text{m/s}^2 \times 981\ \text{pC/(m/s}^2) \approx 387.63\ \text{pC} \)
  - Gain \( G = \frac{V_{out-pp}}{Q} \approx \frac{1\ \text{Vpp}}{387.63\ \text{pC}} \approx 2.58\ \text{MV/C} \)
  - \( C_f \) chosen to achieve desired gain \( G \) for a cutoff frequency \( f_c \) of 0.25 Hz

High-Pass Filter
- Type: Active Bessel High-Pass Filter
- Order: Second-order for a balance between complexity and performance
- Cutoff Frequency (fc): 0.25 Hz (to comply with the 3 dB low-frequency requirement)
- Configuration: Sallen-Key or Multiple Feedback (MFB) for easy implementation
- Component Values: To be calculated based on the standard formulae for Bessel filters

Voltage Amplifier
- Op-Amp Topology: Non-inverting for high input impedance
- Gain (G): 10 (20 dB) to amplify the signal to 1 Vpp
- Supply Voltage: ±5 V or single supply of 5 V (rail-to-rail op-amp recommended)
- Op-Amp Selection: Low-noise, rail-to-rail input/output, precision op-amp for accurate signal amplification
- Example Op-Amp Models: Texas Instruments OPAx340 series, Analog Devices AD8605 series, Maxim Integrated MAX44245 series

Anti-Aliasing Filter
- Topology: Sallen-Key Low-Pass Filter
- Type: Butterworth for a maximally flat amplitude response
- Order: Second-order for sufficient roll-off
- Cutoff Frequency (fc): 2 Hz (chosen to prevent aliasing for signals up to 2 Hz)
- Component Values: Equal-value resistors R1 = R2 = 10 kΩ, equal-value capacitors C1 = C2 ≈ 1.59 μF (standard value chosen to be close to calculation)

Analog-to-Digital Converter (ADC) (If Required)
- Type: Sigma-Delta or SAR ADC based on resolution and power consumption requirements
- Resolution: At least 12 bits for precise measurement
- Sampling Rate: Minimum of 10 Hz to capture the waveform accurately
- Voltage Reference: Precise and stable to match the 1 Vpp range
- Power Supply: Compatible with portable device constraints; low power consumption preferred
- Digital Interface: SPI or I2C for simple connectivity
- Component Selection: Based on accuracy, speed, and power supply constraints

Note: The actual selection of components and final design should consider component tolerances, availability, and system-level constraints such as size, cost, and power consumption. Precision components and op-amps should be used to minimize noise and offset voltages. Component values calculated above are based on typical industry standards and idealized conditions; practical implementation may require fine-tuning.