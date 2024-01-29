Analog Temperature Measurement System for Water Using a Thermistor

The system architecture is designed to measure the temperature of water in a beaker using an NTC thermistor and convert this measurement into a voltage output ranging from 0 to 20 Volts. The proposed system is purely analog and avoids digital-to-analog conversions outside the use of a standard multimeter for reading the output voltage. Below is the compilation of the entire solution organized by architecture blocks.

1. Thermistor Sensor Block (NTC Thermistor - Vishay NTCLE100E3):
   - NTCLE100E3 to measure temperatures between 10°C and 90°C.

2. Linearization Resistor Block:
   - Assuming a beta value of 3950 K, thermistor resistance at 50°C is estimated and matched with a parallel resistor to linearize the response at the midpoint of the desired temperature range.

3. Buffer Amplifier Block:
   - Operational Amplifier (Op-Amp): Precision low-noise rail-to-rail op-amp (e.g., LT1028 or OP27) to minimize loading on the thermistor.
   - Configuration: Non-inverting unity-gain.
   - Power Supply: Assumed ±15V for dual supply or a higher single supply voltage if necessary.

4. Gain Stage Block:
   - Op-Amp: Precision low-noise rail-to-rail op-amp.
   - Gain: Calculated as 20 to amplify a 0-1V signal to a 0-20V range.
   - Resistors: R1 = 1kΩ (feedback), R2 = 51Ω (grounded).

5. Level Shifter Block:
   - Op-Amp: Precision low-noise op-amp (e.g., LT1469).
   - Resistor Values: R1 = 10kΩ, R2 = 100kΩ for gain, R3 = 160kΩ for offset.
   - Power Supply: Assumed ±15V dual supply for maximum dynamic range.

6. Low-Pass Filter Block:
   - Topology: Active Low-Pass Filter (Sallen-Key), Second-order Butterworth.
   - Cutoff Frequency (fc): 1 Hz to reduce high-frequency noise.
   - Resistors: R1, R2 = 10kΩ for equal values.
   - Capacitors: C1, C2, use 15 µF or the closest standard value.
   - Op-Amp: Low-noise op-amp (e.g., TL071).

7. Output Amplifier Block:
   - Op-Amp: Rail-to-rail output op-amp capable of 0-20V output within the given power supply constraints.
   - Gain: Unity (1x), configured to drive the multimeter with the required voltage level and impedance.

Calculations for Thermistor and Linearization Resistor:
R(T) = R25 * exp(Beta * (1/T - 1/T25)), where R(T) is the resistance at temperature T, R25 is 10kΩ for the NTCLE100E3 at 25°C, Beta is approximately 3950 K, and T is in Kelvin.

Calculations for Gain Stage:
Assuming the voltage after the thermistor and linearization network ranges from 0 to 1V for 10°C to 90°C, the gain required to scale this to 0-20V is 20.
Gain (Av) = 1 + (Rf / Ri), with chosen values of Rf = 1kΩ and Ri = 51Ω, we achieve a gain of 20.

Overall, each block in the system has been designed with consideration of the constraints provided and typical values where specific data was not available. The selected components and configurations ensure a robust and accurate analog temperature measurement system suitable for the described application.