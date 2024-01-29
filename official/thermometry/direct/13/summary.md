Analog Temperature Measurement System for Water

Overview:
This solution provides an analog-based temperature measurement system for water temperatures between 10°C and 90°C using a Vishay NTCLE100E3 thermistor. The system outputs a voltage linearly corresponding to the measured temperature range, with an output range of 0 to 20V suitable for reading by a multimeter.

Sensor Block:
- Selected Sensor: NTC thermistor, Vishay NTCLE100E3 series, assumed typical resistance at 25°C is 10k ohms with a Beta value around 3950K.
- The sensor block requires waterproofing if the sensor is to be submerged and should minimize self-heating to less than 1% of the temperature reading.

Linearization Network:
- A resistor in parallel with the NTC thermistor to linearize its response around 50°C is used.
- Assuming the thermistor's resistance at 50°C is approximately 10k ohms, the linearization resistor, R_parallel, is chosen to be 10k ohms as well.

Buffer Amplifier:
- Configuration: Non-inverting voltage follower (unity gain).
- Op-amp Selection: A CMOS op-amp such as the TLC2272 or MCP602, selected for its low input bias current and offset voltage, and capable of single-supply operation.
- Power Supply: Assumed to be +5V or +12V, depending on availability.
- Gain: 1 (unity gain), to ensure a high input impedance and low output impedance.
- Decoupling capacitors: 0.1 μF ceramic capacitors for noise reduction.

Signal Conditioning Amplifier:
- A non-inverting amplifier configuration is used to amplify the signal from the thermistor.
- The gain should be adjustable to accommodate the actual thermistor characteristics.
- Gain Calculation: Assuming the linearized network outputs 0.5V at 10°C and 2.5V at 90°C, the gain required is 10 (calculated using Gain = (20V - 0V) / (2.5V - 0.5V)).
- Op-amp Selection: Instrumentation amplifiers like INA126 or AD623 are suitable for their precise gain setting capability and low offset voltage.
- Additional features: Trim potentiometer to allow fine-tuning of gain.

Output Stage Amplifier:
- Configuration: Non-inverting amplifier with a gain of 10, assuming a signal conditioning output range of 0-2V needs to be mapped to a 0-20V range.
- Gain: 10 (calculated using Gain = 20V / 2V).
- Op-amp Selection: A rail-to-rail output op-amp such as the OPAx177 or AD8476 is recommended.

Multimeter:
- Measurement Mode: DC voltage with a high input impedance (>10MΩ).
- Selection: True RMS digital multimeter capable of measuring voltages up to 20V.
- Resolution and Accuracy: As high as possible within budget to ensure precision in temperature readings.

Additional Considerations:
- Self-heating of the thermistor is managed by keeping the operating current low.
- The linearization resistor should have a low temperature coefficient and be rated for low power.
- Environmental considerations like moisture protection are important if the circuitry is exposed to water.
- Calibration and adjustment mechanisms are integrated for component tolerances and to maintain accuracy.