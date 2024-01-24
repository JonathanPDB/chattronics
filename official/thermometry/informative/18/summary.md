Water Temperature Measurement System with Analog Output

GENERAL ARCHITECTURE:

The project is designed to measure the temperature of water using an NTC thermistor and output the measured temperature as a voltage in the range of 0 to 20 Volts, suitable for reading by a multimeter.

1. NTC Thermistor:
   - Type: Vishay NTCLE100E3 (expected value of 10kΩ at 25°C with a B-value around 3977 K, but actual values should be verified with the datasheet).
   - Waterproofing: Epoxy or glass coating, or a stainless steel/PTFE housing for submersion.
   - Self-heating: Minimized by using a high-value series resistor or constant-current source to limit current.
   - Stability: Typical drift < 0.1% per year, calibration may be required for precision applications.
   - Response Time: A bare thermistor bead responds within a second; encapsulation increases the time constant.

2. Linearization Circuit:
   - Technique: A parallel resistor (Rp) is used to linearize the response of the NTC thermistor around the midpoint of the temperature range.
   - Estimated Rp: Approximately 5kΩ, subject to fine-tuning based on actual R50 value from the thermistor's datasheet.
   - Supply Voltage for Voltage Divider: 24V (chosen to achieve a 0-20V range after scaling).

3. Buffer Amplifier:
   - Topology: Voltage follower using an operational amplifier with FET input for high input impedance.
   - Gain: 1 (unity gain).
   - Input Impedance: >1 MΩ.
   - Bandwidth: >1 Hz.
   - Supply Voltage: Adequately higher than 20V, e.g., ±15V or a single supply of at least 24V.

4. Voltage to Temperature Converter:
   - Configuration: Wheatstone bridge and Op-Amp based linearization.
   - Gain: Variable, set to map the voltage from the Wheatstone bridge to the desired output range.
   - Calibration: Adjustable gain via trimpot or similar component.
   - Supply Voltage: ±15V (assuming standard lab power supplies).

5. Output Scaling Amplifier:
   - Topology: Adjustable gain instrumentation amplifier.
   - Gain: Variable, based on input range from the voltage to temperature converter. Hypothetical gain of 20 is used for a 0-1V input range mapping to 0-20V output.
   - Amplifier Model: AD623 or similar for rail-to-rail input/output.
   - Additional Features: Overvoltage protection, filtering, and offset adjustment as needed.
   - Supply Voltage: ±15V or created from a single-ended supply with a rail splitter.

ADDITIONAL DETAILS:

- Protection and Stability: Input protection diodes, decoupling capacitors, and other basic protection features are included as good engineering practice.
- Temperature Compensation: Not specified, but a temperature compensation network can be added if necessary.

Overall, the system is designed with standard components and characteristics based on common practices for temperature measurement applications. The system includes a linearization circuit for the NTC thermistor, a buffer amplifier, a voltage to temperature converter, and an output scaling amplifier to provide a precise and linear 0-20V output corresponding to the water temperature. Component values and specific configurations are suggested based on standard practices and should be fine-tuned based on actual component specifications and user requirements.