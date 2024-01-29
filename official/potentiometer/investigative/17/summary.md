Pendulum Angle Measurement System Design Summary

POTENTIOMETER BLOCK:
- Type and Model: Conductive plastic potentiometer, suggested model Bourns 3590S-2-103L, 10 kOhm.
- Voltage Rating: Greater than 10 V, suggested 25 V.
- Temperature Range: -55°C to +125°C (typical for robust operation).
- Resolution: Infinite (analog), effective resolution determined by DAQ system.
- Sensitivity: (10V/10kOhm)/(90 degrees) = 1.11 mV/degree.
- Durability: >1 million cycles suggested.

BUFFER:
- Op-Amp: General-purpose op-amp with high input impedance (>1 MOhm) and sufficient bandwidth (>1 kHz). Suggested models include LM358 or TL081.
- Configuration: Voltage follower.
- Power Supply: Assumed +/-15V or a single supply if using LM358.
- Gain: Unity.

GAIN_BLOCK (ATTENUATION STAGE):
- Topology: Two-stage inverting amplifier.
- Op-Amp: Precision op-amp with rail-to-rail output. Examples: OP07, AD823, LT1013.
- Gain: First stage -0.7 (Attenuation), Second stage -1 (Phase correction).
- Power Supply for op-amp: +/-15V (Typical for analog circuitry).
- Feedback Resistors: R1 = 3 kΩ, R2 = 10 kΩ for each stage.

LEVEL SHIFTER:
- Op-Amp: Rail-to-rail operational amplifier like LM324.
- Power Supply: Above 7 V, such as 9 V or 12 V.
- Vbias Calculation: Vbias = (R2/(R1+R2))*Vcc + (Desired Shift), where R1 and R2 form a voltage divider for virtual ground.
- Resistors: Precision resistors (1% tolerance or better) for setting DC offset.

ANTI_ALIASING_FILTER:
- Topology: Active Sallen-Key Low-Pass Filter with additional Notch Filter at 50/60 Hz if required.
- Order: 2nd order.
- Cutoff Frequency: 250 Hz for a buffer below the Nyquist frequency.
- Op-Amp: Low-noise Op-Amp like TL072.
- Quality Factor (Q): 0.707 for Butterworth response.
- Precision Resistors and Capacitors for construction.
- Sampling Rate of DAQ: 1000 samples per second.

DAQ (ADC) INTERFACE:
- ADC Resolution: 12-bit or 16-bit recommended.
- Sampling Rate: 1000 samples per second.
- Communication Interface: SPI or I2C, determined by DAQ system capabilities.
- SNR: >60 dB.
- Power Supply Voltage: Standard 3.3V or 5V.
- Input Range: Matched with the output of the signal conditioning stage.

Please note that specific component values and configurations need to be adjusted based on the final design details, supply voltages, and DAQ system specifications.