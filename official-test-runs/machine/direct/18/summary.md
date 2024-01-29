Pressure and Temperature Monitoring System Design

Architecture Overview:
The design encompasses an array of eight strain-gauge pressure sensors and eight infrared radiation temperature sensors, signal conditioning blocks (including amplification and filtering), a multiplexer (MUX), an analog-to-digital converter (ADC), and a data processing unit. The system is tailored for accurate monitoring of pressure and temperature variations at multiple points on a machine, with the processed data ready for computer analysis.

Pressure Sensor Array:
- Sensor Type: General-purpose industrial strain-gauge pressure sensors. A good starting point would be the Honeywell S Series (e.g., Model SPT Series).
- Pressure Range: 0-100 psi, assuming typical industrial applications.
- Temperature Range: -40°C to 85°C.
- Sensitivity: Assuming a sensitivity of 2 mV/V/psi.
- Output: 1 mV maximum output.
- Material Compatibility: Stainless steel diaphragm for fluid compatibility and corrosion resistance.
- Resolution: With a 16-bit ADC and a 1 mV full-scale range, the resolution is 0.0153 mV per bit.
- Certification: Sensors conform to common industry standards like CE and RoHS.

Instrumentation Amplifier Array:
- Topology: Three-op-amp instrumentation amplifier for each pressure sensor.
- Gain: Approximately 25 (to convert the maximum output voltage of 1 mV from the sensor to a level suitable for the ADC).
- Components: INA118 or AD620 as typical instrumentation amplifier ICs.
- Voltage Supply: ±15 V supply is assumed for op-amps.
- Offset Voltage and Drift: <1 mV, <5 µV/°C.
- CMRR: > 100 dB.

Temperature Sensor Array:
- Sensor Type: Infrared radiation detectors; suggested model MLX90614 from Melexis.
- Temperature Range: -20°C to 500°C.
- Resolution: 0.1°C, providing better than required 1% accuracy.

Non-Linear Correction Block:
- Implementation using a microcontroller with a look-up table or polynomial correction based on a generic non-linear curve.
- Components: Microcontroller with internal ADC (e.g., STM32, ATmega328, or ESP32).
- ADC Resolution: 12-bit or higher to discern 0.0244 mV per bit against a 100 mV full-scale range.

Amplifier Array (for Temperature Sensors):
- Topology: Non-inverting amplifiers for each temperature sensor.
- Gain: Approximately 50 (to amplify the 100 mV sensor output to match the 0-5 V ADC range).
- Op-Amps: Low-noise precision op-amps such as OPAx177 or AD8615.
- Power Supply: Standard ±15 V supply assumed for op-amps.

Anti-Aliasing Filter:
- Filter Type: 4th-order active Butterworth filter for its flat passband characteristic.
- Cutoff Frequency: 500 Hz.
- Filter Order: 4th-order (24 dB/octave roll-off).
- Components: Operational amplifiers like OPAx134 or AD823, with precision resistors and capacitors.

Multiplexer (MUX):
- MUX Channels: 16-channel analog multiplexer/demultiplexer IC such as CD74HC4067 to accommodate all 16 sensor channels.
- Control Interface: SPI or parallel control, managed by a microcontroller.
- Voltage Levels: TTL or CMOS compatible logic (3.3 V or 5 V).

Analog-to-Digital Converter (ADC):
- Type: Successive approximation register (SAR) ADC suitable for medium to high-resolution applications.
- Resolution: 16-bit to accommodate the full dynamic range of the sensor outputs and maintain the 1% accuracy.
- Sampling Rate: At least 2 kHz per channel to prevent aliasing and ensure accurate representation of the signal.
- Interface: SPI or I2C for communication with the data processing unit.
- Input Voltage Range: Adjustable to accommodate 0-100 mV from temperature sensors after amplification and linearization.
- Additional Features: Internal voltage reference, onboard PGA for signal range matching.

Data Processing Unit:
- Processing: Calibration, scale linearization, unit conversion, and simple averaging with optional real-time digital filtering.
- Software: LabVIEW, MATLAB/Simulink, or Python with NumPy and SciPy libraries for implementing data processing algorithms.
- Communication Interface: USB or Ethernet for data transmission to computer analysis software.
- Data Format: Structured binary or ASCII format suitable for various analysis software.

Overall, the system focuses on maintaining the accuracy of measurements and ensuring robust performance in an industrial environment. Each component and block is selected or designed to be flexible and adaptable to a broad range of potential machine monitoring applications. The end-to-end solution is capable of capturing high-fidelity data that is ready for computer analysis to inform maintenance, performance tuning, or other operational insights.