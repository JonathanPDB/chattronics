Analog Water Temperature Measurement System Design

This summary presents an analog water temperature measurement system using a Vishay NTCLE100E3 thermistor to measure the temperature of water in a beaker, with the corresponding voltage output readable by a multimeter in the 0-20V range.

Thermistor Bridge & Linearization:
- Thermistor: Vishay NTCLE100E3, assumed to have a resistance of 10kΩ at 25°C and a B-constant around 3500K.
- Linearization Resistor: Calculated to match the resistance of the thermistor at the midpoint temperature of 50°C. Using the formula R(T) = R25 * exp(B * (1/T - 1/T25)), the value is roughly estimated and then fine-tuned based on standard resistor values.

Signal Conditioning:
- Signal Path: The Wheatstone bridge output is fed into an Instrumentation Amplifier, followed by a Low-Pass Filter, an Adjustable Gain Stage, and finally a Level Shifter to produce a signal within the required voltage range for the multimeter.
- Instrumentation Amplifier: Employing an op-amp like AD620 or INA114 with a gain set to amplify the bridge's millivolt-level output signal to a range that utilizes the full dynamic range of the subsequent stages.
- Low-Pass Filter: A 2nd-order Butterworth filter with a cutoff frequency of 10 Hz is used to remove high-frequency noise. Resistors (R1 = R2) of 10 kΩ and capacitors (C1 = C2) of 1.5 µF form the filter.
- Adjustable Gain Stage: A non-inverting op-amp configuration with a gain range of 100 to 1000, adjustable via a multi-turn precision potentiometer to match the voltage output to the desired range.
- Level Shifter: An op-amp configured as an adder to add a DC offset to the signal and ensure the output voltage starts from 0V. Resistors for the voltage divider and offset circuit are chosen based on the desired shift and supplied voltage.

Output Stage:
- Buffer Amplifier: A voltage follower op-amp circuit is used to provide a low impedance output to the multimeter. An op-amp like LM324 or MCP602 with a rail-to-rail output is used to ensure the full 0-20V output range.

Multimeter:
- High-impedance multimeter with at least 10 MΩ input impedance is recommended to ensure accurate measurement without loading the buffer amplifier. The multimeter should have a voltage resolution and accuracy within 1% of the reading or better.

Thermal Management:
- Self-heating of the thermistor is minimized by designing the circuit to limit the power dissipation below 10mW and keeping the biasing current through the thermistor below 1mA.

Calibration and Testing:
- The system is calibrated using a reference thermometer to ensure the voltage output accurately corresponds to the temperature reading.

Safety and Protection:
- Overvoltage protection is included in the circuit design, and a stable power supply is ensured. The electronics are enclosed to protect against environmental factors and electrical hazards.