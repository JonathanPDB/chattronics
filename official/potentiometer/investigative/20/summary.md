Pendulum Angle Measurement System Design

System Overview:
The project consists of a linear potentiometer to measure the angle of a pendulum, signal conditioning circuits to prepare the analog signal, and a DAQ system to digitize and analyze the measurement. Below is a detailed description of each component and its required specifications.

Linear Potentiometer:
- Type: Linear, 10 kOhms resistance.
- Model Suggestion: Vishay 157 series or Bourns similar models.
- Expected Angular Range: 45 to 135 degrees, with the steady position at 90 degrees.
- Linearity: ±0.5% or better.
- Resolution/Sensitivity: Infinite (analog device), limited by the noise level and DAQ's resolution.
- Operating Temperature Range: Typically -40°C to +85°C for industrial applications.
- The mechanical travel should match the angular range of the pendulum's swing.

Buffer Amplifier:
- Type: Voltage follower using an operational amplifier (op-amp).
- Op-amp Suggestion: LM324 or OPAx340 series for rail-to-rail input/output capability.
- Input Impedance: > 1 MΩ to prevent signal loading.
- Output Impedance: < 100 Ω to avoid signal loss to the subsequent stage.
- Bandwidth: Greater than the highest frequency component of the pendulum's oscillation.

Scaling Amplifier:
- Configuration: Non-inverting amplifier.
- Gain (Scaling Factor): 0.7, to scale the -10 to +10 V range from the potentiometer to the +/- 7 V input range of the DAQ.
- Op-amp Suggestion: Precision op-amp like LM324.
- Gain Setting Resistors: R1 = 10 kOhm, R2 = 4.33 kOhm, to achieve the desired gain.
- Feedback Capacitor (Cf): 10 pF to 100 pF, in parallel with R2 for stability.

Filtering (Notch and Anti-Aliasing Filters):
- Notch Filters: Active Twin-T or Multiple Feedback Band-Pass configured as notch for 50 and 60 Hz.
- Notch Filter Attenuation: >20 dB at notch frequencies.
- Anti-Aliasing Filter: Active second-order Butterworth low-pass filter with a cutoff frequency of around 100 Hz.
- Anti-Aliasing Filter Configuration: Op-amp-based with appropriate capacitors and resistors for setting the cutoff frequency.

DAQ System:
- ADC Type: SAR ADC (suggested due to balance of speed, resolution, and cost).
- ADC Resolution: 12-bit, providing 4096 discrete levels.
- Sampling Rate: At least double the highest frequency of interest, with 100 Hz as a recommended value.
- Communication Protocol: SPI or I2C, assuming standard compatibility.
- Power Supply: Powered within the operating voltage range, typically 5V or 3.3V.

The above components and their values are suggested based on standard practices and the assumption of typical requirements. Adjustments might be needed based on further details provided about the DAQ system, the environmental conditions, and the specific behavior of the pendulum.