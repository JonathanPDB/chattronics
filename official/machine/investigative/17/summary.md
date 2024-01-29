Design of an Analog Instrumentation System for Machine Pressure and Temperature Monitoring

Design Overview:
The system is designed to monitor pressure and surface temperature variations at eight points on a machine, interfacing with an ADC for subsequent computer analysis. The system includes arrays of sensors, signal conditioning units with amplification and filtering, multiplexing, and conversion from analog to digital signals.

Pressure Sensor Array:
- Sensor Type: Strain gauge-based pressure sensors (exact model unspecified)
- Sensor Characteristics: 1 mV maximum output, 1% accuracy, sensitivity assumed to be 0.1 mV/bar for a 10 bar range
- Operating Temperature Range: Assume -40°C to 85°C for industrial applications
- Power Supply: Typically 10 to 30 VDC with built-in voltage regulation (exact values unspecified)

Temperature Sensor Array:
- Sensor Type: Infrared radiation detectors (MLX90614 suggested)
- Sensor Characteristics: 0.02°C resolution, wide measurement range (-70°C to +380°C), factory calibrated
- Power Supply: 3.3V to 5V, low power consumption

Instrumentation Amplifier (for Pressure):
- Amplifier Type: Programmable gain instrumentation amplifier (AD8237 for single supply or INA118 for dual supplies suggested)
- Gain Calculation: Assuming ADC range of 0-5V, Gain = 5000; for ±10V, Gain = 20000
- Bandwidth: >2 kHz to support up to 400 Hz signal frequency
- CMRR: >80 dB for high noise rejection
- Input Impedance: >10 MΩ to avoid loading the sensor
- Noise: Low-noise design to maintain 1% accuracy

Instrumentation Amplifier (for Temperature):
- Amplifier Type: Standard instrumentation amplifier with adjustable gain (~50) and integrated active low-pass filter
- Bandwidth: At least 4 kHz to accommodate a 400 Hz maximum signal frequency
- Filter Characteristics: Second-order Butterworth with a cutoff frequency of about 1 kHz and a damping factor (ζ) of approximately 0.707

8-to-1 Multiplexer:
- IC Suggestion: ADG708/ADG709 for low on-resistance, fast switching, and broad bandwidth (>100 MHz)
- Bandwidth: Must exceed 4 kHz, based on signal frequency requirements
- Control Interface: Flexible to support standard protocols (SPI, I2C, parallel, etc.)
- Features: Include over-voltage protection if necessary

Anti-Aliasing Filter:
- Filter Type: Sallen-Key or multiple feedback low-pass filter
- Order and Cutoff Frequency: Second-order (12 dB/octave) with a cutoff frequency of 1 kHz
- Characteristics: Butterworth response for flat passband, -3 dB gain at 1 kHz, attenuation of at least -40 dB for frequencies well above 1 kHz

Analog to Digital Converter (ADC):
- Resolution: 16-bit (or higher) for 1% accuracy
- Sampling Rate: Minimum 2 kHz per channel, aggregate sampling rate of 16 kHz for all eight channels
- Interface: Flexible digital interface (SPI, I2C)
- Input Configuration: Differential inputs recommended for common-mode noise rejection

Computer (Data Analysis):
- Data Acquisition: Resolution and oversampling to improve resolution and reduce noise
- Signal Processing: Digital low-pass filtering, linearization algorithms for temperature data
- Data Handling and Interface: Robust communication protocol (SPI, I2C, USB, Ethernet)
- Data Analysis: Use of software like MATLAB, LabVIEW, or Python with appropriate libraries for analysis and visualization
- Environmental Tolerance and Protection: Industrial-grade computers with EMI/RFI shielding if necessary
- User Interface: User-friendly GUI for monitoring and control

This summary compiles the suggested components and their characteristics for the design of an instrumentation system to monitor pressure and surface temperature variations on a machine, with the capability to interface with an ADC for computer analysis. Component values and types are proposed based on industry standards and reasonable assumptions, where specific project details were not provided.