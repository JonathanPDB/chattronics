Design of a Pressure and Temperature Monitoring System for Machinery

System Overview:
The system is designed to monitor pressure and temperature at eight points on a machine, with the data to be analyzed by a computer. The architecture includes an array of pressure and temperature sensors, signal conditioning circuits, multiplexing, anti-aliasing filtering, analog-to-digital conversion, and a data interface to the computer.

Pressure Sensor Array:
- Sensor Type: Piezoresistive strain-gauge pressure sensor.
- Suggested Model: Honeywell Model 19C or similar industrial-grade sensors.
- Pressure Range: 0 to 10 bar (0 to 145 psi) as a standard range, with the capability to select different ranges if necessary.

Temperature Sensor Array:
- Sensor Type: Infrared radiation detectors.
- Suggested Model: MLX90614 from Melexis, with a measurable temperature range of -70°C to +380°C and various FOV options.
- Resolution: 0.01°C for industrial applications.

Pressure Signal Conditioning:
- Instrumentation Amplifier: INA114 or AD620, with external resistors for gain configuration.
- Gain: Calculated based on the ADC input range; for a 5V range, a gain of 5000 (R_G ~ 10Ω) and for a 3.3V range, a gain of 3300 (R_G ~ 15Ω) is proposed.
- Low-Pass Filter: Sallen-Key or Multiple Feedback topology with a 500 Hz cutoff frequency using operational amplifiers like TL072 or LT1057.
- Component values for LPF: Assuming C = 0.01 µF, R ≈ 3.18 kΩ (closest standard value is 3.2 kΩ).

Temperature Signal Conditioning:
- Amplification to scale the 100 mV sensor signal to the ADC's 0-5V input range.
- Gain: A gain of 50 is needed to achieve full-scale output, determined by the feedback and input resistor ratio (Rf = 49*Rin).
- Low-Pass Filter: Butterworth with a 500 Hz cutoff frequency to eliminate frequencies above 400 Hz.
- Linearization Network: Custom designed based on specific sensor output characteristics, potentially including precision resistors or active circuitry.

8-to-1 Analog Multiplexer:
- Recommended Model: ADG1606, which supports 16 channels, has low on-resistance, and wide bandwidth.
- Bandwidth Requirement: At least 4 kHz to ensure minimal phase and amplitude distortion.
- Settling Time: Less than 100 microseconds to ensure accurate sampling.

Anti-Aliasing Filter:
- Filter Type: Low-pass Butterworth filter.
- Order: 4th order to balance performance and complexity.
- Cutoff Frequency: Between 100 Hz to 200 Hz, depending on the ADC sampling rate.

Analog-to-Digital Converter (ADC):
- Resolution: 16-bit to maintain sensor accuracy and a resolution of approximately 76.3 µV per bit for a 5V reference.
- Sampling Rate: At least 1 kS/s per channel to capture the expected signal dynamics.

Data Interface to Computer:
- Interface Protocol: USB is recommended for its support and bandwidth.
- Implementation: Use a UART to USB bridge if the ADC outputs serial data; implement checksum or CRC for data integrity.

The above components and suggested values are based on standard industrial practices and typical specifications. Adjustments may be needed to match specific project requirements, which should be confirmed with detailed system specifications and sensor datasheets.