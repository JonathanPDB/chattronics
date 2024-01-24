Pendulum Angle Measurement System Design

The project entails developing an electronic system to calculate the angle of a pendulum using a potentiometer, with the output being sent to a Data Acquisition (DAQ) system. Below is a summary of the system's architecture and component selection based on the provided constraints and requirements.

1. Potentiometer:
   - Type and Model: Precision linear conductive plastic potentiometer.
   - Example Model: Bourns 3549 series.
   - Resistance Value: 10 kOhm.
   - Voltage Range: -10V to +10V (corresponding to a pendulum angle range of 45° to 135°).
   - Sensitivity: 0.222V/degree.

2. Offset Adjustment:
   - Circuit Topology: Non-inverting summing amplifier.
   - Operational Amplifier: Precision rail-to-rail op-amp, e.g., ADA4077-2.
   - Resistors (Precision, low temperature coefficient):
       - R1: 10 kOhm (Input resistor).
       - R2: 14 kOhm (Feedback resistor).
       - R3: 14 kOhm (Grounding resistor).
       - R4 & R5: 20 kOhm each (Voltage divider to set 0 V reference).
       - RV1: 10 kOhm trim pot (Optional, for fine adjustment).

3. Gain Stage:
   - Non-inverting amplifier with a gain of 1.4 to scale -5V to 5V signal to -7V to 7V.
   - Operational Amplifier: Low noise, rail-to-rail op-amp like LM324 or TL084.
   - Resistors:
       - Rin: 10 kOhm (Input resistor).
       - Rf: 4 kOhm (Feedback resistor).

4. LowPassFilter (LPF):
   - Filter Type: Active second-order Butterworth low-pass filter.
   - Cutoff Frequency: 100 Hz.
   - Operational Amplifier: Low noise op-amp, e.g., TL072.
   - Resistors and Capacitors (calculated for 100 Hz using standard equations):
       - R1 = R2: 16 kOhm.
       - C1 = C2: 100 nF.

5. NotchFilter:
   - Filter Type: Active Multiple Feedback Notch Filter.
   - Center Frequencies: 50 Hz and 60 Hz.
   - Quality Factor (Q): Between 10 and 30.
   - Operational Amplifier: Low-noise, precision op-amp like LM358 or OP07.
   - Resistors & Capacitors: Values calculated using standard notch filter design equations for each center frequency.

6. AntiAliasingFilter:
   - Filter Type: Active Butterworth Low-Pass Filter.
   - Order: Second-Order.
   - Cutoff Frequency: 250 Hz.
   - Attenuation at Nyquist Frequency: At least -24 dB.
   - Operational Amplifier: Low-noise, rail-to-rail op-amp with a high slew rate.
   - Resistors & Capacitors: Precision components with low temperature coefficients.

7. DAQ:
   - ADC Resolution: 12-bit for adequate resolution.
   - Sampling Rate: At least 1000 samples per second to match DAQ rate.
   - Input Voltage Range: ±7V to match the conditioned signal.
   - ADC Type: Successive Approximation Register (SAR) ADC topology.
   - Interface: Compatible with DAQ system, e.g., SPI or I2C.

The values and models provided are based on the common requirements for such a system and are meant to act as a starting point. The final values should be refined through simulation and prototyping. The components were chosen considering the need for precision, stability, and low noise to accurately capture the pendulum's motion and convert it into an electrical signal for the DAQ system.