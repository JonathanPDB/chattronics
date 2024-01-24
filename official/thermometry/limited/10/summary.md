Analog Temperature Measurement System Architecture

The following is a systematic compilation of the analog temperature measurement system designed to measure the temperature of water within a range of 10 to 90 degrees Celsius using a Vishay NTCLE100E3 thermistor and output a voltage of 0 to 20 Volts suitable for a multimeter.

Thermistor Sensor:
- Sensor Type: NTC Thermistor (Vishay NTCLE100E3) with a nominal resistance of 10kΩ at 25°C.

Wheatstone Bridge:
- Configuration: One leg with the NTC thermistor and a parallel linearization resistor, and another leg with two precision fixed resistors.
- Resistor Values: R1 = R2 = 10kΩ (Fixed resistors), Rp (Parallel resistor for linearization at 50°C midpoint) ≈ 5.92kΩ.
- Excitation Voltage: 10V for Wheatstone Bridge supply.

Linearization Circuit:
- Linearization Resistor: Rp = 6.2kΩ precision metal film resistor, 1% tolerance, chosen to linearize the thermistor response at the 50°C midpoint.

Voltage Reference:
- Series voltage reference using the LM4040 with a fixed output voltage suitable for the Wheatstone bridge (5V or 10V).
- Supply Voltage: ±15V dual supply.

Instrumentation Amplifier (In-Amp):
- Selected Model: AD623 or INA118 with gain set by an external resistor.
- Gain Calculation: Adjustable, to accommodate the voltage output from the Wheatstone bridge.
- Power Supply: Compatible with ±15V dual supply or a single supply if needed, with a low offset voltage and noise.

Buffer Amplifier:
- Topology: Voltage follower (unity-gain buffer) using a precision op-amp.
- Operational Amplifier Selection: OPA227, LT1013 for dual supply or LM324, TLC2272 for single supply.
- Gain: 1 (unity gain), with high input impedance and low output impedance to prevent loading effects.

Gain Stage:
- Topology: Non-inverting amplifier configuration.
- Operational Amplifier: LM324 or TL081, selected based on the power supply and performance requirements.
- Gain Calculation: Assuming linearization output range of 0.1V to 0.9V, the gain required is \( G = \frac{20V}{0.9V - 0.1V} = 22.22 \approx 22 \).

Output Stage:
- The op-amp in the Gain Stage will amplify the signal to the required 0-20V output range.
- Protection: Clamp diodes for transient protection.

Multimeter:
- Recommended Resolution: At least 4.5 to 6.5 digits for precise voltage measurement.
- Input Impedance: High impedance, typically 10 MΩ or more.
- Voltage Measurement Range: Capable of measuring 0-20 volts DC with auto-ranging capability.
- Features: Data logging and connectivity options, if needed.
- Portability: Handheld or benchtop based on application requirements.
- Recommended Models: Fluke 87V Industrial Multimeter or Agilent U1233A for handheld; Fluke 8846A Precision Multimeter or Keysight 34461A for benchtop.

General Considerations:
- Allow for warm-up and stabilization before taking measurements.
- Test and calibrate the setup with known temperature standards.
- Adjust gain and component values during calibration for accuracy.
- Use proper decoupling and layout techniques to minimize noise and interference.

This architecture provides a robust base for the analog temperature measurement system, with the flexibility to adjust and fine-tune based on further specifications or testing.