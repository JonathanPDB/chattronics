Pendulum Angle Measurement System Design

POTENTIOMETER:
- Sensor: 10 kOhm linear potentiometer, model Bourns 3590S-2-103L or equivalent
- The potentiometer acts as a variable voltage divider, where the output voltage is proportional to the pendulum's angle.

ATTENUATOR:
- Attenuation Ratio: 0.7 (to scale -10V to +10V range to -7V to +7V)
- Resistors: R1 = 4.3 kOhm (closest standard value), R2 = 10 kOhm
- Voltage Divider: The attenuator consists of two resistors in series, with the junction providing the attenuated output.

NOTCH FILTER:
- Twin-T Notch Filter topology, designed to attenuate both 50 Hz and 60 Hz.
- Quality Factor (Q Factor): 10
- Attenuation: -40 dB at 50 Hz and 60 Hz frequencies
- Component Values (calculated using standard formulas):
  - For 50 Hz: Capacitors (C) 3.18 µF, Resistors (R) 318 Ohms
  - For 60 Hz: Capacitors (C) 2.65 µF, Resistors (R) 265 Ohms

ANTI-ALIAS FILTER:
- Active Sallen-Key Low-Pass Filter topology
- Filter Order: 2 (second-order, 12 dB/octave roll-off)
- Cutoff Frequency: 400 Hz
- Operational Amplifier: TL072 (or equivalent) for low noise
- Resistors: R1 = R2 = 39.2 kΩ (1% tolerance)
- Capacitors: C1 = C2 = 1 nF (5% tolerance)

BUFFER:
- Op-Amp Voltage Follower (Unity-Gain Buffer) topology
- Operational Amplifier: Texas Instruments OPAx340 series or Analog Devices AD8605
- Gain: 1 (unity)
- Power Supply: Slightly above the operational voltage range to accommodate rail-to-rail output

DAQ (DATA ACQUISITION SYSTEM):
- ADC Topology: Successive Approximation Register (SAR)
- Resolution: 12-bit, yielding a resolution of 0.00488 V per bit
- Sampling Rate: At least 1000 samples per second
- Input Voltage Range: At least +/- 7 V
- Input Impedance: High, to prevent loading from the buffer stage
- Anti-Aliasing: External, with a cutoff frequency just below 500 Hz
- Notch Filter: External, for 50 Hz and 60 Hz attenuation
- Accuracy: INL and DNL of less than ±1 LSB

Additional Considerations:
- The system's power supply should match the voltage levels used in the circuit design.
- The op-amps should be powered with a supply voltage that prevents clipping at the buffer stage.
- Protection diodes and bypass capacitors should be used where appropriate for signal integrity and component protection.
- Component values have been selected based on standard industry practices, and may need to be adjusted or fine-tuned based on actual testing and environmental conditions.