Analog Instrumentation System for Monitoring Pressure and Surface Temperature

Pressure Sensor Array
- Selection: Strain-gauge pressure sensors with sensitivity presumed at 2 mV/V/psi, for an industrial range of 0-100 psi.
- Model Suggestion: Honeywell Sensotec Model TJE, providing an output of 2 mV/V with recommended excitation of 10 VDC.

AmplifierPressure
- Configuration: Instrumentation amplifier with a gain calculated to map the sensor output to the ADC input range.
- Gain Calculation: To convert the maximum sensor output (1 mV at full scale) to a 5V ADC input, a gain of 5000 is required.
- Component Suggestion: INA118 or AD620 with a power supply of ±15V for industrial applications.
- Gain Setting: Gain = 5V / 1mV = 5000. Using INA118, Rg is set to approximately 10Ω for the required gain.

FilterPressure
- Type: 4th-order low-pass Butterworth filter.
- Cutoff Frequency: 450 Hz to ensure minimal attenuation of signals up to 400 Hz.
- Order: Fourth-order, yielding an attenuation rate of -24 dB/octave.

Temperature Sensor Array
- Type and Model Suggestion: MLX90614ESF-BAA by Melexis for non-contact infrared temperature measurement.
- Temperature Range: -70°C to +380°C to cover a wide range of temperatures.

PreAmplifierTemp
- Topology: Non-inverting amplifier configuration using operational amplifiers.
- Gain: Factor of 50 to scale the temperature sensor output from 0-100 mV to 0-5V.
- Component Suggestion: LT1013 or AD precision series for low-offset voltage and low temperature drift.

Linearizer
- Topology: Piecewise linearization using operational amplifiers.
- Components: LM324 or AD623 op-amps, 10 kΩ precision resistors, 10 kΩ multi-turn potentiometers, LM4040 precision voltage reference.
- Technique: Adjust breakpoints of the piecewise linearization using potentiometers and summing amplifier.

FilterTemp
- Type: Second-order low-pass Butterworth filter.
- Cutoff Frequency: 500 Hz to accommodate the full range of interest.
- Component Values: 10 kΩ resistors with 33 nF capacitors resulting in a slightly lower cutoff frequency of approximately 482 Hz.

Multiplexer
- Selection: 8-to-1 analog multiplexer for selecting one of the eight sensor signals.
- Model Suggestion: CD4051B or DG408 designed for high-performance data acquisition systems.
- Features: Low on-resistance, low channel-to-channel crosstalk, fast switching speed in the microseconds range.

ADC
- Resolution: 12-bit to ensure fine-grained voltage detection and meet the 1% accuracy requirement.
- Sampling Rate: At least 2 kSPS to capture the waveform without aliasing.
- Digital Output Format: SPI for higher data rates and compatibility with most computer systems.
- Power Supply: 3.3V or 5V with low-noise characteristics.

Computer
- Specifications: Mid-range multi-core processor, at least 8 GB RAM, and SSD storage.
- Interfaces: USB or Ethernet for receiving data from the ADC.
- Software: Data logging and visualization software with real-time or batch processing capabilities.
- Environment: Industrial-grade hardware if the computer operates in harsh conditions.