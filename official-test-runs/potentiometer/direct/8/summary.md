Pendulum Angle Measurement System Design

Overall System Architecture and Component Selection

Potentiometer: Bourns 3590S-2-103L (10 kΩ)
- This potentiometer converts the pendulum's angular position into a variable resistance and corresponding voltage signal.
- Chosen for its precision, availability, and durability suitable for angle measurements.

Buffer Amplifier: Operational Amplifier (Op-Amp) Voltage Follower
- Utilizes a low-noise precision op-amp like the AD8605 or Texas Instruments OPAx340 series.
- No external components are required; it provides high impedance loading to the potentiometer and low impedance driving capabilities for subsequent stages.

Level Shifter: Summing Amplifier Configuration
- Op-Amp: LM358 or an equivalent rail-to-rail Op-Amp.
- Resistors: R1 = R2 = 10 kΩ for voltage division; Ri = 10 kΩ, Rf = 20 kΩ for feedback network to provide a gain of 2.
- For shifting the 45 to 135 degrees voltage range from -3.33 V to +3.33 V to within the DAQ's ±7 V input range.

Notch Filters: Active Twin-T Notch Filter Configuration
- For 50 Hz Notch: Resistors R = 10 kΩ, Capacitors C = 3.18 nF, Feedback Resistor R3 = 318 kΩ.
- For 60 Hz Notch: Resistors R = 10 kΩ, Capacitors C = 265 nF, Feedback Resistor R3 = 318 kΩ.
- Op-Amp: General-purpose op-amp like TL071 or LM358.

Anti-aliasing Filter: Second-order Butterworth Low-pass Filter
- Op-Amps: AD797 or OPA2134 for low noise and precision.
- Resistors and Capacitors: Based on a cutoff frequency just below 500 Hz (exact values depend on the specific design and components chosen).

Gain Stage: Non-inverting Amplifier with Low-pass RC Filter
- Op-Amp: TL071 or equivalent, operating on a dual supply of ±10 V.
- Gain: Set to 0.65 using resistors R1 = 10 kΩ, R2 = 6.5 kΩ.
- Anti-aliasing RC Filter: R = 10 kΩ, C = 3.3 μF for cutoff frequency just below 500 Hz.

Data Acquisition System (DAQ): Successive Approximation Register (SAR) ADC
- Resolution: 12-bit or 16-bit for higher resolution if needed.
- Sampling Rate: ≥ 1000 samples per second.
- Input Voltage Range: Scaled to DAQ's ±7 V input.

Calculations and Design Considerations:

Level Shifter Calculation:
Vout_shifted = (Vin + 10V) * (Rf / (Ri + Rf)) - 10V * (Rf / (Ri + Rf))
Targeting a gain of 0.7 (±7 V from DAQ range over ±10 V from the potentiometer), we chose a gain of 0.65 for safety. With Ri = 10 kΩ, we get Rf = 6.5 kΩ.

Notch Filter Component Selection:
For 50 Hz, we use the formula C = 1 / (2 * π * f * R), with R = 10 kΩ and f = 50 Hz, yielding C ≈ 3.18 nF. Similarly, for 60 Hz, with R = 10 kΩ, we get C ≈ 265 nF.

Anti-aliasing Filter Design:
With the cutoff frequency (fc) set just below 500 Hz, we use R = 10 kΩ and calculate C using C = 1 / (2 * π * fc * R), resulting in C ≈ 3.3 μF as the nearest standard value.

With the information provided and under the assumption of an indoor application without extreme environmental conditions, these selections and calculations form the basis for a robust pendulum angle measurement system suited for general use.