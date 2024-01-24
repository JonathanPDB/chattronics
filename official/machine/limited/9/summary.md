Design of an Analog Instrumentation System for Pressure and Temperature Monitoring

Pressure Sensor Array:
- Proposed Sensor: Honeywell S Series, or similar, with a range of 0-100 psi.
- Sensitivity: Typically around 2 mV/V, with a 10 V excitation voltage leading to a 20 mV full-scale output.
- Accuracy: At least 1% Full Scale Output (FSO) or better.
- Nonlinearity: Less than 1% FSO.
- Operating Temperature Range: -40 to 85 °C.
- Power Supply: 10 VDC excitation with current consumption typically under 10 mA.

Instrumentation Amplifier (for Pressure Sensors):
- Amplification: Gain of 5000 required to amplify 1 mV to 5V (G = 5V / 1mV).
- Topology: Instrumentation Amplifier with high CMRR (e.g., AD623).
- Power Supply: Dual +/-15V for ample headroom.
- Gain Setting: Precision resistors for gain-setting, Rg calculated using the formula G = 1 + (50 kΩ / Rg).
- Noise: Low noise characteristics for high-precision measurements.

Low Pass Filter (400 Hz):
- Type: Butterworth for maximally flat frequency response.
- Order: 4th-order for -24 dB/octave roll-off beyond 400 Hz.
- Cutoff frequency: 400 Hz.
- Passband ripple: 0 dB.
- Stopband rejection: At least -48 dB at twice the cutoff frequency (800 Hz).

Temperature Sensor Array:
- Suggested Sensor: MLX90614 from Melexis or equivalent, operating within the industrial temperature range.
- Range: -70°C to +380°C.
- Resolution: 0.02°C, which exceeds the requirement for 1% accuracy.
- Digital Interface: I2C for direct microcontroller interfacing.

Non-linear Correction Block (for Temperature Sensors):
- Approach: Digital correction using a microcontroller with ADC and computational capabilities.
- Microcontroller: ARM Cortex-M series or equivalent.
- ADC Resolution: 12-bit or higher resolution.

Amplifier (for Temperature Sensors):
- Topology: Non-Inverting Amplifier Post Non-linear Correction with a gain of 50.
- Power Supply: Dual +/-15V or single 5V supply based on available power sources.
- Gain Adjustments: Trimming capabilities for fine-tuning during system calibration.

8-to-1 Multiplexer:
- Model: CD4051 series or equivalent solid-state analog multiplexer.
- Voltage Handling: Capable of handling the maximum post-amplification sensor outputs.
- Control Interface: 3-bit digital control compatible with microcontroller logic.
- EMI Considerations: Shielded cables and metallic enclosure for EMI reduction.

Analog-to-Digital Converter:
- Resolution: 12-bit or higher to ensure 1% accuracy level.
- Sampling Rate: Minimum of 6400 Hz to accommodate 800 Hz per channel for eight channels.
- Topology: Successive Approximation Register (SAR) for a balance between speed and precision.
- Interface Protocol: Serial communication over USB or I2C for easy interfacing with computers or microcontrollers.
- Operating Temperature Range: Industrial-grade ADC with a wide temperature range.

Data Acquisition System:
- Data Throughput: Must handle at least 76.8 kbps (800 samples/sec/channel * 12 bits/sample * 8 channels).
- Real-time: Minimal latency with the potential for real-time processing.
- Computer Interface: USB or Ethernet for data transfer.
- Data Format: Standard formats like CSV or JSON.
- Signal Processing: Raw data output with correct scaling and calibration.

Additional System Considerations:
- All instrumentation should be rated for industrial temperature ranges.
- Calibration procedures should be implemented for all sensors in conjunction with their respective signal conditioning stages.
- Protection against environmental factors such as humidity, vibration, and EMI should be included in the design.
- Component tolerances should be selected to ensure filter characteristics and system accuracy remain within specifications.