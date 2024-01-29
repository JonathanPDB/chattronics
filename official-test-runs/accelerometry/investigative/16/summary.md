Low-Frequency Vibration Measurement Device Design

The designed device is a portable instrument for measuring low-frequency vibrations up to 2 Hz with a piezoelectric accelerometer. Below is the compilation of the entire solution provided for each block in the system architecture.

Accelerometer Block:
- Sensor: A piezoelectric accelerometer with a sensitivity of 100 pC/g.
- Suggested Model: PCB Piezotronics model 352A24 or equivalent.
- Measurement Range: ±1 g (assumed for general vibration measurement purposes).
- Operating Temperature: -40°C to +85°C (typical industrial range).
- Power Requirements: ±5V to ±15V, current draw typically under 10 mA.

Charge Amplifier Block:
- Configuration: Inverting operational amplifier to convert charge to voltage.
- Gain (Aq): Designed to convert the maximum charge output at 1g (100 pC) to a voltage signal. If aiming for a maximum output of 4V peak, Aq = 40,000 V/Coulomb.
- Feedback Capacitor (Cf): Cf = 25 pF (calculated from Aq = 1/Cf).
- Feedback Resistor (Rf): Rf = 25.3 MΩ (calculated based on a -3 dB point at 0.25 Hz).
- Input Resistor (Rin): Rin = 10 * Rf (to discharge the input capacitance of the accelerometer).

Low-Pass Filter Block:
- Topology: Second-order Sallen-Key or Butterworth active filter.
- Cutoff Frequency (fc): 3 Hz to ensure minimal attenuation within the passband.
- Quality Factor (Q): 0.707 for Butterworth response for a maximally flat passband.
- Component Values: Calculated based on the desired Q factor and fc.

High-Pass Filter Block:
- Topology: Second-order Butterworth high-pass filter.
- Cutoff Frequency (fc): 0.25 Hz (-3 dB point).
- Damping Factor (ζ): 0.707 for Butterworth response.
- Component Values: For a 0.25 Hz cutoff and a standard capacitor value of 680 nF, the series resistor is calculated using the formula R = 1 / (2 * π * fc * C), resulting in a value close to the chosen capacitor.

Buffer Block:
- Configuration: Unity Gain Buffer (Voltage gain, Av = 1) to maintain signal amplitude and provide impedance matching.
- Frequency Response: Maintains the 3 dB point at 0.25 Hz.
- Op-Amp Choice: Low-noise operational amplifier, such as OPA2134 or equivalent.
- High-Pass Filtering: A 680 nF capacitor is suggested to achieve the desired cutoff frequency at the input.

Output Stage Block:
- Configuration: Operational Amplifier Voltage Follower.
- Gain (Av): 1 (to match the 1 V peak-to-peak output from the charge amplifier).
- Op-Amp Choice: Low-noise, rail-to-rail output op-amp suitable for battery operation, such as OPA344 for 5V systems or LMV321 for 3.3V systems.

Overall architecture ensures low-frequency signal integrity from the sensor through to the output stage, with a focus on maintaining a flat frequency response across the bandwidth of interest, minimal noise introduction, and appropriate impedance matching.