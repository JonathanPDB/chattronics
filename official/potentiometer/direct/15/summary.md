Pendulum Angle Measurement System Design Summary

1. Potentiometer
   - Type: Single-turn rotary potentiometer
   - Model: Bourns 3590S-2-103L
   - Resistance: 10 kOhms
   - Mechanical Rotation: 360 degrees
   - Operating Temperature Range: -10 to 70 degrees Celsius
   - Sensitivity: Linear with angle, adequate to distinguish small changes in angle
   - Assumptions: Linear relationship between angle and resistance

2. Buffer (Voltage Follower)
   - Topology: Unity-Gain Buffer using an operational amplifier
   - Op-Amp: OPA340 or similar
   - Input Impedance: > 1 MΩ
   - Output Impedance: A few ohms
   - Frequency Response: -3dB point at least 5 Hz or higher

3. Level Shifter
   - Topology: Summing amplifier configuration
   - Op-Amp: LM358
   - Resistors: R1 = 14 kOhms (to inverting input), R2 = 10 kOhms (from inverting input to output), R3 = 10 kOhms (from non-inverting input to ground)
   - Reference Voltage: Precision 7V reference for Vref
   - Output Range Adjustment: From -5V/+5V to -7V/+7V

4. Notch Filter
   - Topology: Twin-T Notch Filter
   - Center Frequency: 55 Hz (to attenuate 50 and 60 Hz noise)
   - Quality Factor (Q): Approximately 5 to 10
   - Stopband Attenuation: Minimum of 20 dB at 50 and 60 Hz
   - Op-Amps: TL072 or equivalent for low noise performance
   - Resistors/Capacitors: Precision 1% tolerance

5. Amplifier
   - Topology: Non-inverting operational amplifier configuration
   - Amplification Gain: 1.4
   - Op-Amps: TL071, LM358, or OP27
   - Power Supply Range: Minimum +/-10V for op-amp operation
   - Assumptions: Linear voltage output from potentiometer ranges from -5V to +5V over pendulum's motion

6. Anti-Aliasing Filter
   - Topology: Active Butterworth Low-Pass Filter
   - Order: Second Order (2-pole)
   - Cutoff Frequency: 100 Hz
   - Attenuation at Nyquist Frequency (500 Hz): At least -40 dB
   - Op-Amps: TL072 or LM358 for low noise performance
   - Resistors/Capacitors: Based on op-amp characteristics and filter performance
   - Power Supply: +5 V to +15 V range suitable for many op-amps

Additional Assumptions and Notes:
- The system is intended to function within a typical indoor environment, with a temperature range of -10°C to 70°C.
- Precision components are selected to ensure stability and accuracy of the system.
- Resistors used in the notch filter and anti-aliasing filter have a 1% tolerance to maintain the designed characteristics.
- The amplifier's gain calculation is based on the assumption that the output range of the potentiometer needs to be adjusted to fit the DAQ's full scale without clipping.