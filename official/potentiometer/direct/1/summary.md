Pendulum Angle Measurement System

This summary consolidates the proposed architecture and design for an analog pendulum angle measurement system using a potentiometer, signal conditioning, filtering, and data acquisition unit (DAQ).

**Potentiometer Sensor Block**
- Type: Precision linear potentiometer
- Selected Model: Bourns 3590 Precision Potentiometer (or equivalent)
- Operating Temperature Range: -40°C to +85°C
- Power Supply Requirements: Regulated power supply with low ripple, providing a stable -10V to +10V
- Electrical Characteristics: 
  - Resistance: 10 kOhm
  - Resolution: Better than 0.1%
  - Linearity: ±0.5% or better
- Mechanical Stress Considerations: 
  - Mechanical limit to the potentiometer's sweep range for 45 to 135 degrees of the pendulum's movement

**Offset Adjustment Block**
- Topology: Differential amplifier using an operational amplifier (op-amp)
- Parameters: 
  - Op-amp: Low-offset voltage and drift op-amp (e.g., AD8628, OPA227)
  - Resistors for voltage divider: 1% metal film resistors, 10 kOhm each
  - Potentiometer for fine adjustment: 10 kOhm trimmer potentiometer
  - Feedback and input resistors for unity gain: 10 kOhm, 1%

**Amplification Block**
- Topology: Non-inverting operational amplifier configuration
- Gain (Av): 1.4 (calculated based on desired DAQ input range of +/-7V)
- Op-amp: Single-supply op-amp with rail-to-rail input/output capabilities (e.g., LM324)
- Feedback Resistor (Rf): 14 kOhm (calculated for the desired gain)
- Input Resistor (Ri): 10 kOhm (for unity gain)

**LowPassFilter1 Block**
- Topology: Multiple Feedback Low-Pass Filter (MFLPF)
- Cutoff Frequency: 40 Hz (to attenuate 50 Hz interference)
- Order: 2nd-order (12 dB/octave roll-off)
- Passband Gain: 0 dB (unity gain)
- Damping Factor: Q factor of 0.707 for a Butterworth response
- Resistors: 390 kOhm (nearest standard value to calculated 398 kOhm based on 10 nF capacitors)

**LowPassFilter2 Block**
- Topology: Sallen-Key low-pass filter
- Cutoff Frequency (f_c): 100 Hz (to attenuate 60 Hz interference and prevent aliasing)
- Order: 2nd order (40 dB/decade roll-off)
- Passband Ripple: Maximally flat Butterworth response
- Component Values: 
  - Resistors: R1 = R2 = 10 kOhm (standard value)
  - Capacitors: C1 = C2 = 160 nF or closest available (rounded from 159 nF)

**AntiAliasingFilter Block**
- Topology: Sallen-Key low-pass filter with a Butterworth response
- Cutoff Frequency: 100 Hz (to prevent aliasing)
- Order: Second-order (two-pole) with a 12 dB/octave roll-off
- Resistors: 10 kOhm (standard value)
- Capacitors: 160 nF or closest available (rounded from 159 nF)
- Op-amp: TL072 or equivalent with good frequency response

**ADC Block within DAQ**
- Topology: Successive Approximation Register (SAR) ADC
- Resolution: 12-bit (4096 distinct values)
- Sampling Rate: 1000 samples per second as required
- Input Voltage Range: +/-7V to match the full-scale output of the signal conditioning stage
- Interface: SPI (Serial Peripheral Interface) or I2C as per integration requirements
- Operating Conditions: Industrial temperature ranges if necessary

**Overall System Considerations**
The system provides a complete solution from the physical measurement using the potentiometer through signal conditioning, filtering, and conversion to a digital signal by the ADC. The design accounts for environmental factors, power supply stability, and ensures that the DAQ's input voltage range and sampling rate are fully utilized while suppressing unwanted noise and preventing signal aliasing.