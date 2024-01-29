Design of an Analog Instrumentation System for Monitoring Pressure and Surface Temperature Variations

Pressure Sensor Array:
- Suggested model: Honeywell Trustability HSC Series or equivalent industrial strain-gauge pressure sensors.
- Expected measurable pressure range: To be specified, but sensors typically offer ranges from 1 psi to 150 psi.
- Accuracy of pressure measurement: ±0.25% FSS (Full-Scale Span) typical, exceeding the required 1% accuracy.
- Operating temperature range: Typically -40°C to +85°C, with built-in temperature compensation to maintain accuracy.
- Power Supply Requirements: Typically 3.3V to 5V; current consumption details to be obtained from the chosen sensor's datasheet.

Temperature Sensor Array:
- Suggested model: MLX90614 (Melexis) or similar infrared thermometer sensors.
- Temperature range: -70°C to +380°C (to be selected based on operational requirements).
- Output: Digital (e.g., I2C communication interface).
- Accuracy: ±0.5°C around room temperatures, generally within the 1% requirement.
- Resolution: 0.02°C.
- Power Supply: 3.3V or 5V, depending on the model variant.
- Environmental considerations: Standard industrial-grade sensors typically accommodate a wide range of environmental conditions.

Signal Conditioning:
Instrumentation Amplifier for Pressure Sensors:
- Gain Calculation: To amplify the 1 mV maximum output to 5V full range for the ADC.
- Gain (pressure) = 5V / 1mV = 5000.

Instrumentation Amplifier for Temperature Sensors:
- Gain Calculation: To amplify the 100 mV maximum output to 5V full range for the ADC.
- Gain (temperature) = 5V / 100mV = 50.
- Amplifiers should have high input impedance and CMRR.

Linearization Circuit for Temperature Sensors:
- Techniques: Look-Up Table or Polynomial Approximation.
- Implementation: Can be achieved in software within the microcontroller.
- ADC resolution should be at least 10-bit, 12-bit or higher recommended.

Low-Pass Filter:
- Suggested Order: Second-order Butterworth filter for a flat frequency response.
- Cutoff Frequency: Set at or slightly below 400 Hz to ensure attenuation of higher frequencies.
- Passive Filter Components: Using standard values, C = 0.01 µF, R ≈ 3.9 kΩ for RC filter configuration.

8:1 Multiplexer:
- Suggested model: ADG1608 (Analog Devices) or equivalent with a suitable bandwidth (>1 kHz) and low crosstalk.
- Input Voltage Range: 0 to 100 mV to match the higher output voltage of the temperature sensors.
- Switching Speed: <100 µs to support a sampling rate of at least 1 kSps per channel.

Microcontroller:
- Features: Should support multiple communication interfaces (USB, UART, SPI, I2C).
- Capabilities: DMA, sufficient processing speed and memory, and compatibility with industrial protocols.
- Power Supply: Wide operating voltage range, industrial temperature grade.

Data Interface:
- Interface Options: USB, Ethernet, RS-232/485, Wireless (Wi-Fi/Bluetooth).
- Isolation: Implement galvanic isolation for robustness.
- Protection: Include surge protection measures such as TVS diodes.
- Environmental Considerations: Use industrial-grade connectors, shielding, and grounding techniques.

The above components and techniques have been suggested based on a balance between general industrial standards and the specific requirements provided throughout the conversation. Final component selection, calibration, and environmental protection measures will depend on the detailed operational parameters of the machine and environment where the system will be deployed.

Note: The exact gain values for the amplifiers are subject to change based on the ADC's specific input range and resolution, which were not provided. The multiplexer model and filter component values are also only suggestions and may need to be adjusted based on further system specifications and empirical testing.