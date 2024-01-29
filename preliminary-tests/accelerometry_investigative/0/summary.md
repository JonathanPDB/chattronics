Design of a Portable Low-Frequency Vibration Measurement Device

The following summary outlines the architecture and component values for a portable device capable of measuring low-frequency vibrations using a piezoelectric accelerometer and associated signal conditioning stages.

Piezoelectric Accelerometer:
- Sensor Type: Piezoelectric accelerometer
- Sensitivity: 100 pC/g
- Suggested Model: PCB Piezotronics 352C22 or equivalent that operates within the frequency range of 0.25 Hz to at least 2 Hz

Charge Amplifier:
- Feedback Capacitor (C_f): 820 pF (chosen to convert the maximum expected charge to desired voltage)
- Feedback Resistor (R_f): 1 MΩ (to set low-frequency cutoff below 0.25 Hz)
- Operational Amplifier: JFET or MOSFET input op-amp with high input impedance and low noise
- The design uses the charge sensitivity of the accelerometer and expected vibration amplitude to calculate the maximum expected charge:
  ap = (2π * f)^2 * x ≈ 39.478 m/s^2 → ap_g ≈ 4.025 g → Qmax = S * ap_g ≈ 402.5 pC

Low-Pass Filter:
- Cutoff Frequency (f_lp): 5 Hz (to ensure minimal attenuation at 2 Hz)
- Filter Type: Second-order Bessel/Butterworth (user preference)
- Operational Amplifier: Low-noise, FET input operational amplifier
- Resistors and Capacitors: Will be calculated using standard filter design equations to achieve the desired cutoff frequency

High-Pass Filter:
- Cutoff Frequency (f_hp): 0.25 Hz (3 dB down point)
- Filter Type: Second-order Butterworth for flat passband response
- Operational Amplifier: Low-noise precision Op-Amp (e.g., LT1007/LT1037 series)
- Resistors and Capacitors: Chosen for a second-order high-pass Butterworth filter using standard values (e.g., 10 uF capacitors, 62 kΩ resistors)

Gain Stage:
- Voltage Gain (Av): Approx. 625 (calculated to amplify the maximum accelerometer output to 1 Vpp)
- Operational Amplifier: Low-noise precision Op-Amp (e.g., chosen for a fixed-gain design)
- Feedback Resistors: Calculated for desired gain (e.g., 62 kΩ input resistor, 624 kΩ feedback resistor)

Output Interface:
- Operational Amplifier: Rail-to-rail output op-amp (e.g., OPA344 or AD8605), powered by a single supply voltage close to 5V
- Resistor (Rload): Small series resistor (e.g., 10 Ω to 50 Ω)
- Capacitor (Cload): Decoupling capacitor (e.g., 100 nF)

The device design encompasses the full signal chain from the accelerometer through to the output interface, including the conditioning and amplification required to produce a standardized output suitable for data acquisition systems or other recording equipment. Component values and suggested models are based on typical industry standards and the limited data provided. The actual component selection may require adjustments based on more detailed specifications and testing.