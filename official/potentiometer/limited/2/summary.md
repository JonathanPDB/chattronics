Pendulum Angle Measurement System Design

This summary provides a comprehensive overview of the proposed solution for measuring the angle of a pendulum using a potentiometer and various analog signal conditioning blocks, followed by digital acquisition (DAQ).

**Sensor Block (Potentiometer):**
- Type: Linear taper potentiometer
- Value: 10 kOhms
- Angle Measurement Range: 45 to 135 degrees
- Suggested Model: Bourns 3590S-2-103L
- Environmental Conditions: Assuming standard indoor conditions
- Temperature Range: -55°C to +125°C
- Power Rating: Potentiometer is well within its power rating with a maximum voltage of 10V

**Buffer Stage:**
- Type: Unity-Gain Buffer (Voltage Follower)
- Op-Amp: TL071 or equivalent FET input Op-Amp
- Power Supply: ±10V or single supply with virtual ground
- Input/Output Capacitors: 0.1 µF ceramic capacitors for power supply stability

**Level Shifter:**
- Configuration: Non-inverting amplifier with adjustable DC offset
- Op-Amp: General-purpose op-amp (e.g., LM358) powered by ±15V
- Voltage Reference (Vref): 2.5V precision voltage reference IC or stable voltage divider
- Resistors: R1 = 7 kOhms, R2 = 3 kOhms for scaling, Roffset and Rgain = 10 kOhms each for fine-tuning

**Amplification Block:**
- Configuration: Non-inverting op-amp
- Op-Amp: Low-noise, precision op-amp (e.g., OPAx177 or AD8676)
- Power Supply: Voltage to suit the available power and rail-to-rail capability if single supply
- Gain Calculation: If scaling is needed, gain A_V = desired output range / level shifter output range
- Resistors: R_f and R_g chosen based on required gain, with R_g typically around 10 kOhms

**Notch Filter:**
- Topology: Twin-T Notch Filter
- Frequencies: 50 Hz and 60 Hz with at least 40 dB attenuation
- Quality Factor (Q): Greater than 10 for narrow notches
- Component Values: Calculated based on:
\[ R = \frac{1}{2 \pi f C} \]
\[ R_p = 2R \]
\[ C_p = \frac{C}{2} \]
where R is resistance, C is capacitance, R_p is parallel resistance, C_p is parallel capacitance, and f is the notch frequency.

**Anti-Aliasing Filter:**
- Type: Second-order active low-pass Butterworth filter
- Cutoff Frequency: Approximately 150 Hz to preserve signal components <100 Hz and attenuate >500 Hz
- Attenuation: -40 dB or more at the Nyquist frequency
- Op-Amp: General-purpose op-amp suitable for filter design

**Data Acquisition System (DAQ):**
- ADC Type: Successive Approximation Register (SAR)
- Resolution: 12 bits or higher for adequate quantization
- Input Voltage Range: +/- 7V to match DAQ specifications
- Sampling Rate: At least 2 ksps (kilo-samples per second) for overhead beyond the DAQ's 1 ksps
- Communication Protocol: SPI or I2C if not specified, to ensure compatibility
- Anti-Aliasing: If no built-in AA filter, use an external filter with a -40 dB attenuation at 500 Hz

With these components and specifications, the pendulum angle measurement system is expected to provide accurate and reliable data for input into the DAQ. The analog signal conditioning stages will ensure that the potentiometer's voltage is properly buffered, scaled, shifted, and filtered before being converted to a digital signal by the ADC.