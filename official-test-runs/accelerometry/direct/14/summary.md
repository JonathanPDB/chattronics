Portable Low-Frequency Vibration Measurement Device Design

This summary consolidates the key points of the proposed design for a portable device capable of measuring low-frequency vibrations, following an analog signal processing chain from the sensor to the digital signal processing output.

Piezoelectric Accelerometer:
- Sensor Model: PCB Piezotronics Model 352C33 with a sensitivity of 100 pC/g.
- Expected Operating Conditions: Wide temperature range of -54°C to +121°C.
- Maximum Acceleration Measurement: ±0.1g (approximated based on a 5 cm peak oscillation at 2 Hz).

Charge Amplifier:
- Maximum Output Charge (Q_max): 160 pC (calculated from the 100 pC/g sensitivity and a peak acceleration of approximately 1.6g).
- Feedback Capacitor (C_f): 160 pF, chosen to generate a 1 V peak-to-peak output for Q_max.
- Feedback Resistor (R_f): 3.9 MΩ, selected to establish a -3 dB point at 0.25 Hz for the low-frequency response.
- Operational Amplifier: Low-noise op-amp with low input bias current to minimize the offset voltage.

Low-pass Filter (0.25 Hz -3dB):
- Filter Topology: Multiple Feedback Low-Pass (MFB LP) Filter.
- Filter Type: Second-order Butterworth for a maximally flat passband.
- Component Values: Assuming a 10 μF capacitor, the resistors required would be approximately 6.37 MΩ. Due to practical considerations, resistor values may need to be adjusted, or an active filter design with scaled components may be implemented.

Gain Stage:
- Expected Input Voltage Range: Assumed 10 mV peak-to-peak from the charge amplifier.
- Gain (G): 100, to amplify the assumed input to a 1 V peak-to-peak output.
- Topology: Non-inverting operational amplifier configuration.
- Op-Amp Example: ADA4077-2 or OPA2277.
- Precision Resistors: 99 kΩ for R1 and 1 kΩ for R2 for fixed gain.

Buffer Amplifier:
- Topology: Voltage Follower (Unity Gain Buffer).
- Operational Amplifier: Low-noise, precision op-amp with FET inputs, such as the OPA134 series.
- Gain: Unity (A_v = 1).

Anti-Aliasing Filter (2 Hz cut-off):
- Filter Type: Second-order Butterworth low-pass filter.
- Component Values: R1 = R2 ≈ 39 kΩ, C1 = C2 ≈ 2 nF, based on the standard Butterworth design for a 2 Hz cutoff.

Analog-to-Digital Converter (ADC):
- Resolution: 12-bit or higher to provide a resolution well below the 10 mV offset requirement.
- Sampling Rate: At least 20 Hz to capture the vibration signal adequately.

Digital Signal Processing & Output:
- DSP Requirements: Sampling rate of 10 to 20 Hz, FFT analysis for frequency content if needed.
- Data Storage: Non-volatile memory (e.g., SD card) for data logging.
- Data Output: Potentially wireless transmission (e.g., Wi-Fi, Bluetooth) or direct display on an LCD.

This architecture and component selection provide a comprehensive starting point for the design of a portable vibration measurement system. Fine-tuning of component values and additional system-level validation will be required during the development process to meet all project specifications and performance targets.