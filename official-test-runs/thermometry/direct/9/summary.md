Analog Temperature Measurement System Design

The system is designed to measure the temperature of water in a beaker using a Vishay NTCLE100E3 NTC thermistor and output a voltage between 0 and 20 Volts corresponding to a temperature range between 10°C and 90°C. The analog domain is used for signal conditioning.

NTC Thermistor Sensor Block:
- Sensor: Vishay NTCLE100E3 with a typical resistance of 10kΩ at 25°C and a beta value of approximately 3977K.
- Operating current: Kept low to prevent self-heating, typically under 100 µA.
- Environmental considerations: The sensor will be waterproof and chemically resistant.

Linearization Resistor Block:
- Linearization Resistor (Rp): Estimated at approximately 6.5kΩ to match thermistor resistance at 50°C.
- Calculation: Rp = R25 * exp(-beta * ((1/(50+273.15)) - (1/298.15))) where R25 = 10kΩ and beta = 3977K.
- High-precision metal film resistor with 1% tolerance and a temperature coefficient of 50 ppm/°C or better.

Buffer:
- Op-Amp: TLV2371 or equivalent voltage follower configuration.
- Power Supply: +5V or +12V (single supply).

Amplification:
- Topology: Non-inverting op-amp configuration.
- Gain: Between 7 to 10, adjustable via feedback resistor values.
- Op-Amp: General-purpose LM324 or precision op-amp like AD8605 for low offset and noise.

Level Shifting:
- Op-Amp Adder Circuit: To add a DC offset to the amplified signal.
- Resistors: Rf = 100kΩ, Rin = 20kΩ for an estimated gain of 5. Rref also 100kΩ for offset addition.
- Power Supply: Single or dual supply, depending on availability.

Voltage Scaling:
- Topology: Non-inverting amplifier configuration.
- Gain: 4, with resistor values R1 = 10kΩ and R2 = 30kΩ for the non-inverting op-amp.
- Op-Amp: LM324 or TL074.

Anti-Aliasing Filter:
- Configuration: Second-order Sallen-Key active low-pass Butterworth filter.
- Cutoff Frequency: 1 Hz.
- Components: R1 = R2 ≈ 1.59 MΩ, C1 = C2 = 0.1 μF.

Output Conditioning and Measurement:
- Load Impedance: Assumed high impedance typical of digital multimeters (around 10 MΩ).
- Equipment: High-precision digital multimeter with sufficient resolution and accuracy for the 0-20V range.
- Calibration: Adjust the scaling and offset resistors to calibrate the output range to match the temperature accurately.

Overall, the system provides a reliable and accurate analog solution for temperature measurement using a common lab bench power supply. Precision components and careful design considerations ensure minimal self-heating, stability, and accuracy. The multi-stage conditioning process linearizes, amplifies, shifts, scales, and filters the signal from the NTC thermistor for a direct, measurable voltage output.