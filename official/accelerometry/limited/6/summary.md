Portable Low-Frequency Vibration Measurement Device Design

The design of a portable low-frequency vibration measurement device involves several key components and calculation steps to create a functional system that meets the project requirements.

Piezoelectric Accelerometer:
- Recommended Sensor: PCB Piezotronics 352C33, with a sensitivity of 100 pC/g.
- Calculated charge for 5 cm peak-to-peak acceleration: Q = sensitivity * acceleration = 100 pC/g * 0.05 g = 5 pC peak-to-peak.

Charge Amplifier:
- Op-Amp: AD8628 or OPA128 for low bias current and low noise.
- Feedback Capacitor (C_f): Assuming a charge amplifier output of 10 mV peak-to-peak, C_f is chosen to be 1 pF for a theoretical 1 V peak-to-peak output.
- Feedback Resistor (R_f): Selected to set the low-frequency cutoff at 0.25 Hz, a standard value of 620 MΩ is used, resulting in a corner frequency slightly above the desired 0.25 Hz.

High-Pass Filter:
- Type: Active Bessel high-pass filter for flat group delay to minimize phase distortion.
- Order: 2nd order to provide a reasonable trade-off between complexity and performance.
- Components: Resistors with low temperature coefficients and low leakage capacitors.
- Resistor (R): The standard value of 620 kΩ is chosen to achieve a cutoff frequency near the desired 0.25 Hz with C = 1 μF.

Amplifier:
- Op-Amp: Low-noise operational amplifier such as the AD8628 or OPAx177.
- Gain: Non-inverting configuration with a gain of 100, using standard resistor values of 1 kΩ for R2 and 100 kΩ for R1.
- Power Supply: A single supply voltage, typically around 5 V, suitable for portable devices.
- Bandwidth: Chosen to comfortably exceed the frequency requirements with a GBP much higher than 200 Hz.

Anti-Aliasing Filter:
- Topology: Passive RC low-pass filter followed by an active second-order Sallen-Key low-pass filter.
- Cutoff Frequency: Set at 10 Hz to prevent aliasing, well below the Nyquist frequency based on a 100 SPS ADC.
- Components for Sallen-Key Filter: Standard resistor values of 10 kΩ, and capacitors of 1.59 µF for a Butterworth response.

ADC:
- Type: Delta-Sigma for high resolution and low-frequency measurements.
- Resolution: 24-bit to ensure fine quantization and high dynamic range.
- Sampling Rate: At least 10 times the highest frequency of interest; 100 SPS is suggested.
- SNR: At least 90 dB to ensure accurate low-level signal detection.

Overall, the device encompasses an accelerometer to capture the vibrations, a charge amplifier to convert the charges to voltages, a high-pass filter to eliminate low-frequency noise and DC offsets, an amplification stage to scale the signal to the desired level, an anti-aliasing filter to clean the signal before sampling, and an ADC to digitize the signal for further processing or display.