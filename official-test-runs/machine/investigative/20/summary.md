Analog Data Acquisition System for Machine Pressure and Temperature Monitoring

Overview:
The system is designed to monitor pressure and temperature at eight points on a machine. The pressure is measured using strain-gauge sensors with a max output of 1 mV, while the temperature is measured using infrared radiation detectors with a max output of 100 mV. The system includes signal conditioning, multiplexing, analog-to-digital conversion, and digital interfacing with a computer for data analysis.

Pressure Sensor Array:
- Type: Strain-gauge pressure sensor
- Requirements: ±1% accuracy, temperature compensation (if needed), sensitivity sufficient to produce max 1 mV output under full-scale pressure.

Instrumentation Amplifier (Pressure):
- Topology: Three-op-amp instrumentation amplifier
- Gain: Assuming a 0-5 V ADC range, gain = 5000 V/V for a unipolar ADC or 2500 V/V for a bipolar ADC.
- Bandwidth: Minimum of 400 Hz
- Op-amp selection: Low noise, high CMRR (e.g., AD620, INA125)
- Power supply: Assume standard ±15 V or single supply as per amplifier requirements.

Low-Pass Filter (Pressure):
- Topology: Sallen-Key or multiple feedback (MFB) low-pass filter
- Cutoff frequency (fc): Set to 800 Hz to accommodate signal frequencies up to 400 Hz
- Order: Second-order for a roll-off of -12 dB/octave or fourth-order for -24 dB/octave, if necessary.
- Component values (for second-order): Resistors (R1 = R2) = 10 kΩ, Capacitors (C1 = C2) = 19.9 nF for fc = 800 Hz using standard op-amps.

Temperature Sensor Array:
- Suggestion: MLX90614 infrared thermometer sensor for its digital output and wide temperature measurement range.

Instrumentation Amplifier (Temperature):
- Topology: Three-op-amp instrumentation amplifier
- Gain: 50 V/V to scale 100 mV sensor output to a 5 V ADC range.
- Bandwidth: Minimum of 400 Hz
- Op-amp selection: Low noise, high CMRR, suitable GBWP (e.g., AD623, INA128/INA129)
- Power supply: As per amplifier requirements, with single supply operation possible.

Low-Pass Filter (Temperature):
- Topology: Second-order Butterworth filter using an op-amp in the Sallen-Key configuration.
- Cutoff frequency (fc): 10 Hz to ensure minimal phase delay in temperature readings.
- Order: Second-order for a reasonable trade-off between complexity and performance.
- Component values: Resistors R1=R2=10 kOhm, Capacitors C1=C2=1.59 µF.

Multiplexer:
- Type: CMOS analog multiplexer, such as CD74HC4067
- Channels: Minimum 16 single-ended channels to accommodate all sensors.
- Features: High impedance inputs, low on-resistance, and low crosstalk.

Analog-to-Digital Converter (ADC):
- Type: 16-bit Successive Approximation Register (SAR) ADC
- Sampling rate: At least 2 kSPS per channel to satisfy Nyquist criterion
- Input voltage range: 0 to 100 mV to match sensor outputs
- Channels: Single ADC channel with time-division multiplexing via MUX
- Additional features: Internal MUX, sample-and-hold circuit.

Computer:
- Interface: Capable of handling common communication protocols (USB, Ethernet, SPI, I2C).
- Data buffering: FIFO buffers to manage peak data loads.
- Processing: Efficient algorithms optimized for the computer's capabilities.
- Robustness: Error checking, fault tolerance, and recovery mechanisms.
- Scalability: Modular software design for easy expansion or integration.

Please note that the above recommendations are based on the limited information provided and common industrial practices. Specific component models and values should be further refined based on detailed specifications and requirements when available.