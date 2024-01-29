DESIGN OF A PORTABLE LOW-FREQUENCY VIBRATION MEASURING DEVICE

The designed system is composed of a piezoelectric accelerometer, charge amplifier, low-pass filter, high-pass filter, amplifier stage, anti-aliasing filter, and an ADC. Here is the detailed architecture with calculations and component choices:

**Piezoelectric Accelerometer:**
- Model: PCB Piezotronics 352C33
- Sensitivity: 100 pC/g
- Frequency Response: Up to 2 Hz with a low-frequency limit of 0.25 Hz
- Expected Acceleration: a = (2 * π * f)^2 * d ≈ 0.198 m/s² (~0.02 g) with f = 2 Hz and d = 0.05 m
- Calculated Charge: Q = Sensitivity * a ≈ 628.3 pC

**Charge Amplifier:**
- Op-amp: Low-noise JFET or CMOS op-amp (e.g., AD8628)
- Feedback Capacitor (C_f): 62.81 nF
- Feedback Resistor (R_f): 10 MΩ with a parallel trimmer for offset adjustment

**Low-Pass Filter:**
- Topology: Passive second-order low-pass RC filter
- Cutoff Frequency (f_c): 3 Hz
- Resistance (R1 = R2): 47 kΩ
- Capacitance (C1 = C2): 1.2 µF

**High-Pass Filter:**
- Topology: First-Order Active High-Pass Filter
- Cutoff Frequency (f_c): 0.25 Hz
- Resistance (R): 100 kΩ
- Capacitance (C): ≈ 10 μF
- Op-amp: TL071 or AD8610 for low noise

**Amplifier Stage:**
- Topology: Non-Inverting Operational Amplifier Configuration
- Op-amp: AD8605 or similar low-noise, rail-to-rail op-amp
- Gain Setting Resistors: R1 = 1 kΩ and Rf = 99 kΩ for a gain of 100
- Feedback Capacitor (Cf): 10 pF to 100 pF for stability
- Power Supply: Single +5 V operation

**Anti-Aliasing Filter:**
- Topology: Butterworth low-pass filter
- Order: Second-order (Two-pole)
- Cutoff Frequency: 5 Hz
- Component Selection: Resistors of 1% tolerance and capacitors with low temperature coefficients

**ADC:**
- Resolution: 16-bit
- Sampling Rate: Minimum 10 Samples per second (S/s)
- SNR: ≥ 90 dB
- Communication Interface: SPI or I2C
- Power Consumption: Low-power mode with < 1 mA current draw
- ADC Type: SAR ADC
- Input Voltage Range: 0 to +1 V
- Accuracy and Linearity: INL/DNL within ±1 LSB

The overall design ensures a portable, low-power, and high-accuracy system capable of measuring low-frequency vibrations with minimal noise and offset. It includes provisions for temperature compensation of the sensor and uses precision components to maintain stability and accuracy over a range of operating conditions.

Component values and op-amp choices were made to balance performance, power consumption, and cost, making the design suitable for a wide range of portable vibration measurement applications.