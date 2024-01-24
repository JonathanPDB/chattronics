Pendulum Angle Measurement System Design

This summary compiles all the details and specifications for the design and implementation of a pendulum angle measurement system using a potentiometer, signal conditioning, filtering, and a data acquisition (DAQ) system.

**Potentiometer (Sensor)**
- Type: Linear rotary potentiometer.
- Value: 10 kOhms.
- Model Suggestion: Bourns 3590 Precision Potentiometer.
- Mounting: Securely attached to the wooden structure with the shaft coupled to the pendulum beam.
- Voltage Range: Expected to vary linearly with the angle of the pendulum, estimated between -5V and +5V for 45° to 135°.

**Buffer (Signal Conditioning)**
- Topology: Voltage Follower Configuration using an Operational Amplifier (Op-Amp).
- Op-Amp Suggestion: Texas Instruments OPA340 or Analog Devices AD8605.
- Power Supply: ±10V (matching the potentiometer's power supply).
- Input and Output Resistances: High input impedance (>1 MΩ) and low output impedance (<100 Ω) to prevent loading and ensure signal integrity.

**Gain Stage (Signal Conditioning)**
- Topology: Non-Inverting Op-Amp Configuration.
- Required Gain (A): \( A = \frac{V_{out,max} - V_{out,min}}{V_{in,max} - V_{in,min}} = \frac{7V - (-7V)}{5V - (-5V)} = 1.4 \).
- Resistor Values: R1 = 4kΩ, R2 = 10kΩ (precision resistors with a 1% tolerance).
- Power Supply: ±10V (dual supply).

**Band-Pass Filter (Signal Conditioning)**
- Active Twin-T Notch Filter Topology for 50 Hz and 60 Hz attenuation.
- Notch Filter Component Values: R = 31.8 kΩ, C = 100 nF (for 50 Hz), R = 26.5 kΩ, C = 100 nF (for 60 Hz).
- Followed by a Sallen-Key Low-Pass Filter with a Butterworth response.
- Low-Pass Filter Component Values: R = 3.9 kΩ, C = 1 nF (for a cutoff frequency of 200 Hz).

**Anti-Aliasing Filter (Signal Conditioning)**
- Sallen-Key Low-Pass Filter Topology.
- Filter Order: Second-order (two-pole) with -12 dB/octave roll-off.
- Cutoff Frequency: Approximately 210 Hz (below half the DAQ's sampling rate).
- Damping Factor: ζ = 0.707 (Butterworth response for flat passband).

**Data Acquisition System (DAQ)**
- ADC Topology: Successive Approximation Register (SAR).
- ADC Resolution: 12 bits (4096 discrete levels).
- SNR: 60-70 dB (corresponding to a dynamic range of about 10-11 bits).
- Input Capacitance and Leakage Current: Low to prevent loading effects.
- Power Supply Voltage: Typically 3.3 V to 5 V for digital circuitry.
- Sampling Rate: 1000 samples per second.
- Maximum Input Voltage: ±7 V.

**General Considerations**
- All Op-Amps should have decoupling capacitors to reduce noise.
- Protection diodes may be used at the buffer output to protect the DAQ from voltage spikes.
- A compact design using SMD components is preferred for modern applications, but through-hole components may be used for prototyping ease.
- The entire system must be calibrated and tested to ensure accurate angle measurement within the operating voltage and angle range.

This comprehensive solution ensures that the pendulum angle is accurately measured and conditioned for analog-to-digital conversion without exceeding the DAQ's input voltage limits or introducing significant noise or distortion to the signal.