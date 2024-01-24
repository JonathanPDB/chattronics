Analog Temperature Measurement System for Water Beaker

The proposed system architecture consists of the following blocks, designed to measure the temperature of water in a beaker using an NTC thermistor and to output a voltage easily measured by a multimeter.

Thermistor Sensor (NTCLE100E3):
- Vishay NTCLE100E3 thermistor with a typical resistance of 10kΩ at 25°C.
- Operating temperature range: 10°C to 90°C.
- Self-heating to be minimized by limiting the current through the device.

Linearization Resistor:
- A parallel resistor is used to linearize the thermistor's response at the midpoint temperature of 50°C.
- Resistance value calculated based on the thermistor's resistance at 50°C using the parallel resistance formula and the thermistor's beta value.

Wheatstone Bridge:
- A Wheatstone bridge configuration with one arm being the thermistor and the linearization resistor.
- Bridge powered by a stable 5V voltage reference to minimize self-heating.
- Output voltage range to be linear with temperature change, assuming a few millivolts per degree Celsius change.

Instrumentation Amplifier:
- To amplify the differential signal from the Wheatstone bridge.
- High CMRR and low offset voltage for accurate measurements.
- Gain calculated based on the output voltage of the Wheatstone bridge and the desired 0-20V output range.
- Assuming a 5V to 1V bridge output, a gain of 20V/V is required.

Voltage Reference:
- ADR4540BRZ series voltage reference from Analog Devices.
- Output voltage: 2.5V.
- High accuracy (±0.02%) and low temperature coefficient (5 ppm/°C).

Scaling Amplifier:
- Non-inverting operational amplifier configuration.
- Gain set to scale the instrumentation amplifier output to the 0-20V multimeter range.
- Gain calculation based on an assumed bridge output of 100mV (10°C) to 1V (90°C), resulting in a gain of 20V/V (20V output range over 1V input range).
- Precision metal film resistors were chosen for setting gain, with a tolerance of 1% or better.

Low Pass Filter:
- 2nd order Butterworth filter with a cutoff frequency of 1 Hz.
- Designed to filter high-frequency noise and preserve the temperature signal fidelity.
- Operational Amplifier: TL072 used for low noise.

Output Buffer:
- Unity-gain voltage follower using a rail-to-rail output operational amplifier such as OPA197, OPA192, or OPA2187.
- Low offset voltage and low output impedance to drive the high-impedance input of the multimeter.

Multimeter:
- Should have a voltage measurement range up to 20V DC with a resolution of at least 10mV.
- An input impedance of 10MΩ to ensure accurate measurements without loading effects.

The entire system is designed to function within a temperature range of 10°C to 90°C, with an output voltage range of 0V to 20V, representing the temperature range. The design accounts for self-heating effects to be less than 1% and all signal conditioning is performed in the analog domain. The detailed component selection and calculations ensure that the system will provide accurate and stable temperature measurements.