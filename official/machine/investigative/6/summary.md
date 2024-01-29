Comprehensive Design Overview for Pressure and Temperature Monitoring System

Pressure Sensor Array:
- Sensor Type: Industrial-grade strain-gauge pressure sensors (e.g., Honeywell Model 31 series or similar).
- Pressure Range: Broad range, common is 0-100 psi (specific range to be defined based on application).
- Accuracy: At least 1% of full scale.
- Environmental Sealing: Appropriate IP rating for protection against moisture and dust, EMI shielding.

Amplifier_Pressure:
- Gain: With a 1 mV full-scale output from the pressure sensor and a 5V ADC input range, the required gain is calculated as 5000 (G = 5V / 1mV).
- Topology: Instrumentation amplifier recommended for its precision, low offset, and low drift (e.g., AD620, INA118).

Filter_Pressure:
- Topology: Active low-pass filter with a Sallen-Key or Multiple Feedback (MFB) configuration.
- Order: 4th order Butterworth filter to provide a flat frequency response up to 400 Hz.
- Cutoff Frequency: Set at approximately 500 Hz to ensure a flat response up to 400 Hz without significant attenuation.

Temperature Sensor Array:
- Sensor Type: Non-contact infrared temperature sensors (e.g., MLX90614 by Melexis) with a standard range of -70°C to +380°C.
- Accuracy: ±0.5°C around room temperatures and a resolution of 0.02°C.
- Field of View: Appropriate for the monitored surface area.

Amplifier_Temperature:
- Gain: Assuming a maximum sensor output of 100 mV and a 5V ADC range, a gain of 50 is calculated (G = 5V / 100mV).
- Topology: Fixed Gain Precision Amplifier or Programmable Gain Amplifier based on the need for flexibility.

Filter_Temperature:
- Topology: Active low-pass filter with a Butterworth response for a maximally flat passband.
- Order: 2nd order is suggested for balancing complexity and performance.
- Cutoff Frequency: A standard starting point around 1 Hz since temperature changes are slow.

Multiplexer:
- Type: High-Speed 16:1 Analog Multiplexer capable of switching speeds much faster than 800 Hz.
- Control Signals: Compatibility with standard logic levels (3.3V or 5V).

ADC:
- Sampling Rate: Minimum overall sampling rate of 12.8 kHz is needed for sequential sampling of 16 channels at 800 Hz each.
- Resolution: At least 12 bits, but 16 bits is recommended for greater accuracy.
- Topology: Successive-approximation register (SAR) ADC or sigma-delta ADC for higher resolution and filtering capabilities.

Computer:
- Processing: Scalable, flexible computing system capable of buffering incoming data, processing, and storage.
- Communication: Standard interfaces such as USB, SPI, I2C, or Ethernet should be supported.
- Software: Modular programming approach, cross-platform libraries, real-time monitoring, and visualization capabilities.
- Environmental Considerations: Industrial-grade components for a wide temperature range, humidity and EMI protection.

Please note that the exact component selection, numerical values for filter components, and further system specifics will need to be validated and adjusted based on detailed project requirements and real-world performance assessments.