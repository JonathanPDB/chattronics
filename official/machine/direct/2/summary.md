Design of an Analog Instrumentation System for Monitoring Pressure and Surface Temperature

Pressure Sensor Array:
- Pressure Range Assumption: 0-1000 psi
- Sensor Model Suggestion: Honeywell Model FP2000 or equivalent
- Sensitivity: Typically 2 mV/V/psi with a 10 V supply, yielding 20 mV/psi
- Accuracy: ±1% Full Scale (FS)
- Operating Temperature Range: -40°C to 85°C

Instrumentation Amplifier (for Pressure Sensors):
- Topology: Three-op-amp configuration
- Required Gain: To scale the sensor's maximum output of 1 mV to the ADC's input range (0-5V), a gain of 5000 is needed (G = 5V / 0.001V).
- Power Supply: ±15 V
- Op-Amp Recommendations: AD620 or INA118 for precision and low noise
- Bandwidth: >2 kHz for accommodating signal frequencies up to 400 Hz

Temperature Sensor Array:
- Temperature Range Assumption: -70°C to 380°C
- Sensor Recommendation: Melexis MLX90614ESF-BAA-000-TU-ND or equivalent with digital output for easier integration
- Resolution: 0.1°C to ensure the 1% accuracy
- Power Supply: 3.3V to 5V
- Current Consumption: <10mA per sensor

Nonlinear Correction (for Temperature Sensors):
- Method: Piecewise-linear approximation using precision resistors and op-amps
- Op-Amps: LM358 or equivalent, used for multiple segments linearization
- Resistors: High-precision metal film (0.1% tolerance)
- Voltage Reference: Stable reference like LM4040 for op-amps

Multiplexer:
- Capacity: 16 channels to accommodate all sensor signals
- Bandwidth: >2 kHz, at least 5 times the highest frequency of interest (400 Hz)
- Suggested Model: ADG1606 from Analog Devices or equivalent with low on-resistance
- Supply Voltage: ±5V to ±15V
- Settling Time: <250 ns

Anti-Aliasing Filter:
- Topology: 4th order Butterworth low-pass filter
- Cutoff Frequency: 400 Hz with 60 dB of attenuation starting at twice the cutoff frequency (800 Hz)
- Op-Amp Recommendations: OPAx134 series for low noise and stability
- Components: Metal film resistors (1% tolerance), ceramic or film capacitors (5% tolerance)

ADC:
- Topology: Successive Approximation Register (SAR)
- Resolution: 12 bits for 1% accuracy across a 0-5V range
- Sampling Rate: At least 2 kHz per channel to satisfy Nyquist for a 400 Hz signal
- Input Range: 0-5V, accommodating the amplified sensor signal levels
- Interface: SPI for high-speed data transfer
- Power Supply: Low to moderate power consumption, compatible with ±15V or 5V systems

Data Acquisition System:
- Data Throughput: Minimum of 12 bits * 16 channels * 800 samples/second = 153.6 kbps
- Interface: High-speed serial or parallel (SPI, I2C) or USB for easy connectivity to computers, with Ethernet or RS-485 for long distances
- Processing: Capable of real-time data logging and necessary signal processing (if needed)
- Operating Environment: Industrial-grade components with wide temperature range and proper shielding
- Data Security: Encryption and integrity measures such as checksums or CRC for sensitive data

Environmental Considerations:
- Ensure proper shielding, grounding, and protection against environmental factors such as dust, water, and corrosive substances.
- Temperature effects on the sensors and instrumentation amplifiers should be compensated for to maintain accuracy.
- The sensors and components should have IP ratings appropriate for the operating environment.

Note:
- The design includes an analog front-end to condition, linearize, and multiplex the signals from the pressure and temperature sensors before digitization by the ADC.
- The values and models provided are based on general industrial standards and should be adjusted once specific sensor characteristics are provided.
- The system's final design should be verified with detailed analysis including thermal effects, power consumption, and overall system integration considerations.