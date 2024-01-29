Design of a Pressure and Temperature Monitoring System for Industrial Machinery

The design for an analog instrumentation system to monitor pressure and temperature at eight points on a machine includes the following components and specifications:

Pressure Sensor Array:
- Pressure Sensors: 8x Strain-gauge pressure sensors
- Pressure Range: 0-100 psi (assumed, adjustable based on application)
- Sensitivity: 1 mV/psi
- Accuracy: 1% Full Scale
- Operating Temperature Range: -40°C to 85°C (industrial environment assumption)
- Safety Ratings: IP65 or equivalent (assumed, for industrial environment)
- Excitation Voltage: 5 V (assumed, typical for strain gauges)
- Suggested Model: Honeywell Sensing and Productivity Solutions SSC Series or equivalent

Temperature Sensor Array:
- Temperature Sensors: 8x Infrared radiation detectors
- Temperature Range: -70°C to +380°C (based on Melexis MLX90614)
- Response Time: Approximately 10 ms to 100 ms
- Thermal Resolution: 0.02°C at room temperature
- Output: Digital signal (PWM or I2C)
- Accuracy: ±0.5°C around room temperatures
- Suggested Sensor: Melexis MLX90614 or equivalent

Pressure Signal Conditioning:
- Gain: Minimum gain of 5000 to scale 1 mV maximum pressure sensor output to at least 0-5 V
- Low-Pass Filter: Cutoff frequency slightly above 400 Hz (to filter frequencies higher than 400 Hz)
- Power Supply: ±15 V for dual supply or +5 V for a rail-to-rail single supply instrumentation amplifier
- Excitation Voltage Source: 5 V stable reference voltage

Temperature Signal Conditioning (Simplified):
- Buffer Amplifier: High-input impedance, unity-gain, to prevent loading of the sensor signal
- Low-Pass Filter: Assuming 50 Hz cutoff frequency with components R = 3.2 kΩ and C = 1 µF (to filter out high-frequency noise)

8-to-1 Multiplexer:
- Voltage Range for Input Channels: 1 mV to 100 mV
- Bandwidth: At least 800 Hz to handle signal frequencies up to 400 Hz
- Switching Time: Significantly less than 2.5 ms (period of the highest frequency of interest)
- Suggested Models: DG408, ADG1606 or equivalent

Analog-to-Digital Converter (ADC):
- Resolution: 16-bit (to provide a fine quantization level for the accuracy of 1% FS)
- Sampling Rate: 1 kHz (to capture the full dynamics of the pressure and temperature signals)
- Communication Interface: SPI/I2C (based on compatibility with computer systems)
- Power Supply Voltage: 3.3V to 5V
- Operating Temperature Range: -40°C to 85°C (industrial grade)
- Suggested ADC: ADS1115 or equivalent

Data Interface to Computer:
- Interface: USB for short distances or Ethernet for networked sensors (assumed, for general compatibility)
- Communication Protocol: USB HID for USB or TCP/IP for Ethernet (assumed, for ease of integration)
- Data Rate: To accommodate the ADC's sampling rate, resolution, and number of sensors (calculated based on specific details)
- Robustness: Additional protective measures such as opto-isolation (assumed, for industrial environments)
- Compliance: CE and FCC standards (assumed, for regulatory compliance)

Note: The specifics of the system such as exact pressure range, operating conditions, power supply specifications, and environmental requirements were not provided. The above summary is based on reasonable assumptions and common practices in industrial monitoring system applications. Final component selection and design should be validated through simulation and/or prototyping to ensure it meets the performance requirements of the actual application.