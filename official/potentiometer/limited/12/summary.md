Pendulum Angle Measurement System Design

High-Level Architecture Overview
The system is designed for measuring the angle of a pendulum using a potentiometer and consists of several blocks: Input Conditioning, Signal Processing, Output Conditioning, and Data Acquisition. Each block is critical in translating the mechanical angle into a digital signal.

Input Conditioning
- Component: Potentiometer (Bourns 3590S series or Vishay Spectrol Model 534)
  - Type: Precision rotary, 10kΩ resistance, wirewound or conductive plastic
  - Operating Temperature: -55°C to +125°C
  - Power Rating: 0.5W to 1W

- Component: Buffer Amplifier
  - Op-Amp: AD8605 or OPA333 for low offset voltage and drift, Rail-to-Rail output
  - Configuration: Voltage Follower
  - Power Supply: ±15V if dual supply; single supply if required with virtual ground

Signal Processing
- Component: Scaling Amplifier
  - Op-Amp: LM358 (Low power, general-purpose, wide bandwidth)
  - Gain: \( A = \frac{DAQ_{max} - DAQ_{min}}{Pot_{max} - Pot_{min}} = \frac{14V}{10V} = 1.4 \)
  - Resistors: \( R_f = 4kΩ \) (or nearest standard value), \( R_g = 10kΩ \)

- Component: Notch Filter (50 & 60 Hz)
  - Topology: Multiple Feedback Notch Filter
  - Quality Factor (Q): 10
  - Attenuation Depth: At least 20 dB at the notch frequencies
  - Resistors: \( R1 = R2 = 31.8kΩ \) (closest 1% standard value)
  - Capacitors: \( C1 = C2 = 10nF \) (closest 1% standard value)

Output Conditioning
- Component: Offset Adjust
  - Configuration: Summing Amplifier using an operational amplifier
  - Offset Voltage: Adjustable during calibration via a trim potentiometer

- Component: Anti-aliasing Filter
  - Type: 4th-order Butterworth active filter
  - Cutoff Frequency: 400 Hz
  - Op-Amps: TL074 or OP07

Data Acquisition (DAQ) Block
- ADC Resolution: At least 12 bits for precision angle measurement
- Sampling Rate: Matching DAQ rate of 1000 samples per second
- Input Range: ±7V to match the DAQ input range

The entire system is designed to be adjusted and calibrated to ensure accurate conversion of the pendulum's angular position to a digital signal suitable for the DAQ system. The potentiometer is mechanically linked to the pendulum to measure the angle, and the output is conditioned through buffering, scaling, and filtering to fit the DAQ's input range. Noise at 50 and 60 Hz is attenuated using notch filters. The system also includes an anti-aliasing filter to prevent sampling artifacts, and the DAQ converts the conditioned analog signal into digital format at a high resolution to ensure precision.