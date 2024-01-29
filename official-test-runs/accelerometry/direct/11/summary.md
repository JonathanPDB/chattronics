Design of a Portable Low-Frequency Vibration Measurement Device

The proposed architecture for the portable vibration measurement device consists of several key blocks: the piezoelectric accelerometer, charge amplifier, low-pass filter, instrumentation amplifier, ADC, anti-aliasing filter, and digital processing with data output. Each block is designed to work in harmony with the others to ensure accurate and reliable vibration measurements.

Accelerometer Block:
- Selected Sensor: Endevco model 2221F or equivalent with 100 pC/g sensitivity.
- Measurement Range: ±1g to measure peak accelerations around 0.05g.
- Frequency Response: Flat from 0.25 Hz to 2 Hz.
- Resolution: Capable of discerning changes as small as 1 mg.
- Temperature Range: Standard industrial range from -40°C to 85°C.
- Power Supply: 5V with a current draw of less than 10 mA.

Charge Amplifier Block:
- Gain: 100 MV/C to produce 1 V peak-to-peak output for 5 pC peak charge.
- Feedback Capacitor (C_f): 1 nF.
- Feedback Resistor (R_f): 620 kΩ for a 3 dB point at approximately 0.25 Hz.
- Bias Current (I_bias_max): Less than 16 pA to keep the voltage offset below 10 mV.
- Operational Amplifier: LMC6001 or equivalent with ultra-low bias current.

Low-pass Filter Block:
- Topology: Multiple Feedback (MFB) active low-pass filter.
- Order: Second-order (two-pole) for a 12 dB/octave roll-off.
- Cutoff Frequency (f_c): 3 Hz to ensure low-frequency response is 3 dB down at 0.25 Hz.
- Damping Ratio (ζ): 1/sqrt(2) for a Butterworth response (Q = 0.707).
- Resistors (R1, R2): 10 kΩ.
- Capacitor (C1): Approximately 5.3 nF (use standard values of 4.7 nF or 5.6 nF or a combination).

Instrumentation Amplifier Block:
- Topology: Three-op-amp instrumentation amplifier configuration.
- Power Supply: +5V single supply.
- Gain: Set to 1000 (60 dB) to amplify the signal from millivolts to volts.
- Noise Performance: Low noise with input noise density below 10 nV/√Hz.
- Bandwidth: At least 20 Hz.
- CMRR: At least 80 dB.
- Gain Resistor (Rg): Approximately 49 ohms for a gain of 1000.
- Suggested Models: AD620 or INA128.

ADC Block:
- Type: SAR ADC.
- Resolution: 16-bit for high-fidelity measurements.
- Sampling Rate: 10 Hz to capture the waveform accurately.
- Input Voltage Range: At least ±1 V or 0 to 1 V for single-ended input.
- Accuracy: INL and DNL less than 1 LSB at 16-bit resolution.
- Input Type: Differential input for improved noise rejection.
- Interface: SPI or I2C for minimal wiring and portability.
- Power Consumption: Low power for battery efficiency.
- Anti-Aliasing Filter: Cutoff just above 2 Hz, implemented with a simple first-order or second-order low-pass filter.

Anti-Aliasing Filter Block:
- Topology: Active low-pass Butterworth filter.
- Order: Second-order (two-pole).
- Cutoff Frequency (f_c): 8 Hz to prevent aliasing while allowing up to 2 Hz signals.
- Component Values: Standard values of 2 kΩ for resistors (R) and 10 nF for capacitors (C) in a Sallen-Key configuration.

Digital Processing and Data Output:
- Functionality: Include digital filtering and FFT for frequency analysis.
- Data Storage: MicroSD card slot for logging data.
- Connectivity: USB interface for data transfer, optional Bluetooth or Wi-Fi for wireless communication.
- Data Format: CSV or JSON for easy integration with analysis software.
- Power Supply: Low power consumption design with a focus on minimal power drain during data output.