Temperature Measurement and Signal Conditioning for NTC Thermistor

Overview:
This solution outlines a complete signal conditioning circuit for measuring the temperature of water with a thermistor and providing an output voltage suitable for measurement with a multimeter.

NTC Thermistor:
The primary sensor used is a Vishay NTCLE100E3 NTC thermistor. It operates within the temperature range of 10 to 90 degrees Celsius, and the self-heating effect must be maintained below 1%.
- A typical R25 value for this thermistor is around 10kΩ.
- The expected resolution for temperature measurement should be 0.1°C or better.

Parallel Resistor for Linearization:
To linearize the NTC's response at 50 degrees Celsius (midpoint), a parallel fixed resistor is utilized. The resistance of the parallel resistor (Rp) is chosen to match the thermistor's resistance at the linearization point.
- Assuming a standard R25 of 10kΩ and using the thermistor's beta value from the datasheet, calculate the resistance at 50°C (R50).
- Rp = R50 (with a 1% tolerance, and a temperature coefficient of 50 ppm/°C or better).

Buffer Amplifier:
A unity-gain buffer amplifier provides high input impedance to prevent loading the thermistor circuit and low output impedance.
- An op-amp like TL071 or LF356 is suitable for this stage.
- The input impedance of the buffer should be at least 10 times higher than the highest resistance of the thermistor.

Gain Stage:
An amplification stage boosts the signal to a level suitable for the desired 0-20 V output range.
- Assuming an input range of 0-100 mV across the thermistor after linearization, calculate the gain (Av) needed to reach the output range: Av = 20V / 100mV = 200.
- Use a non-inverting amplifier configuration with a precision op-amp.
- Select R1 (feedback resistor) to be 200kΩ and R2 (input resistor) to be 1kΩ for the desired gain.

Level Shifter:
The level shifter adjusts the DC level of the amplified signal to match the 0-20 V range of the multimeter.
- Use an op-amp in a summing amplifier configuration with a variable resistor for fine-tuning.
- Rail-to-rail output capability is necessary if the power supply is close to the maximum output voltage.
- Precision resistors should be used for stability.

Anti-Aliasing Filter:
A low-pass filter removes high-frequency noise that could affect the multimeter reading.
- A Sallen-Key second-order low-pass filter with a cutoff frequency of 30 Hz is proposed.
- The filter provides a maximally flat response (Butterworth) with a roll-off rate of 40 dB/decade.
- Use standard off-the-shelf values for resistors and capacitors based on the desired cutoff frequency.

Output Buffer:
The output buffer ensures the signal can drive the multimeter with a high input impedance without significant voltage drop.
- Use a voltage follower configuration with a unity gain op-amp.
- The op-amp should support the output current range of 10-20 mA and provide rail-to-rail output if needed.
- Bandwidth of the buffer should exceed the signal bandwidth, typically a few kHz is sufficient.

The final circuit design should be assembled and tested to adjust buffer, gain, and level shifter stages in line with actual component values and performance. Fine-tuning can be performed based on experimental data, and component values can be adjusted as necessary during prototyping.