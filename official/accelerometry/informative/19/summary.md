Design of a Portable Low-Frequency Vibration Measurement Device

System Overview:
The design comprises a piezoelectric accelerometer, a charge amplifier, a low-pass filter, a buffer, and an ADC. The objective is to measure low-frequency vibrations up to 2 Hz with a peak amplification of 5 cm.

Sensor:
- A piezoelectric accelerometer is chosen for its sensitivity and dynamic range. Suggested model: TE Connectivity Model 834M1-0500 or similar with 100 pC/g sensitivity.
- The accelerometer's specifications such as capacitance and any built-in conditioning were not specified.

Charge Amplifier:
- Sensitivity of the accelerometer: 100 pC/g.
- The maximum expected acceleration (a_max) ≈ 0.016 g.
- Maximum expected charge (Q_max) = 1.6 pC.
- Desired output voltage = 1 V peak to peak (V_out = 0.5 V half of peak-to-peak).
- Feedback capacitance (C_f) assumed 10 nF (assuming sensor capacitance ~ 1 nF).
- Feedback resistor (R_f) calculated to be approximately 3.125 MΩ.
- Calculations for R_f: 
  R_f = V_out * C_f / Q_max
      = (0.5 V * 10 nF) / 1.6 pC
      ≈ 3.125 MΩ
- Low-frequency cutoff (f_c) ≈ 5 Hz.
- Operational Amplifier: AD8628 or LTC6244.

Low-Pass Filter:
- Active second-order Sallen-Key filter design.
- Cutoff frequency set slightly above the highest frequency of interest, chosen as 5 Hz.
- Resistor (R) = 10kΩ.
- Capacitor (C) ≈ 3.3 µF (standard value close to calculated value of 3.18 µF).
- Damping factor (ζ) between 0.6 to 0.8 for a slightly underdamped system.

Buffer:
- Topologies: Voltage Follower (unity gain) or Non-Inverting Amplifier (if gain needed).
- Gain: Unity gain (gain = 1) or a modest gain between 2 to 5.
- Power Supply Voltage: Assuming 3.3V or 5V for a single supply, or ±5V for a dual supply.
- Bandwidth: At least 20 Hz to ensure accurate signal reproduction.
- Slew Rate: At least 0.5 V/µs.
- Operational Amplifiers: LM324 or OPA2134.

ADC:
- Type: Delta-Sigma ADC.
- Resolution: 24-bit.
- Sampling Rate: 10 to 100 SPS (configurable).
- Digital Output Interface: SPI (default), I2C optional.
- SNR: > 90 dB.
- THD: < -100 dB.
- Includes Integrated Anti-Aliasing Filter with cutoff slightly above 2 Hz.

Please note that these specifications are based on general assumptions due to the lack of specific input on certain parameters. The final design should be validated with actual sensor data and system requirements.