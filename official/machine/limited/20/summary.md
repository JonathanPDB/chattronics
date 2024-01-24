Analog Data Acquisition System for Pressure and Temperature Monitoring

System Overview:
The proposed system consists of an array of pressure and temperature sensors interfaced with analog signal conditioning blocks, an analog-to-digital converter (ADC), and a data interface for computer analysis. The system monitors pressure and surface temperature variations at eight points on a machine, with a 1% accuracy requirement for both sensors.

Pressure Sensor Array:
- Sensor Type: Strain gauge pressure sensor
- Model Suggestion: Honeywell Model 31 Series or equivalent
- Pressure Range: 0-1000 psi (assumed)
- Accuracy: 0.25% to 0.5% full-scale to allow for signal conditioning errors
- Sensitivity: Assuming 2 mV/V sensitivity
- Overpressure Limit: Typically 150% of the rated pressure
- Output: Max 1 mV, in line with the specified maximum output
- Signal Bandwidth: Capable of capturing up to 400 Hz frequencies

Instrumentation Amplifier Array:
- Configuration: Each strain gauge sensor connected to an individual instrumentation amplifier
- Gain (A_v): 5000, calculated to amplify the 1 mV sensor output to 5V for the ADC
- Power Supply (V_s): ±15V assumed for ample dynamic range
- CMRR: ≥80 dB to mitigate common-mode noise
- Quiescent Current: Microampere range for low power consumption
- Input Impedance: >1 MΩ
- Gain Setting: Precision resistors with 0.1% tolerance or better
- Example of Gain Setting Resistors: If Rin is 10kΩ, Rf should be 490kΩ for a gain of approximately 5000

Temperature Sensor Array:
- Sensor Type: Infrared (IR) radiation detector with adjustable emissivity
- Model Suggestion: MLX90614 from Melexis
- Temperature Range: -20 to 500°C (assumed)
- Supply Voltage: 2.6V to 3.6V
- Current Consumption: ~1.5mA
- Resolution: 0.02°C (exceeds 1% accuracy requirement)
- Output: Digital (PWM or SMBus)

Amplifier Array (for Temperature Sensors):
- Topology: Non-inverting operational amplifier
- Gain (A_v): 50, to map the linearized 0-100 mV output to a 0-5V ADC input range
- Power Supply (V_s): ≥5V for rail-to-rail capability
- GBWP: ≥20 kHz considering the 400 Hz measurement frequency
- Input Impedance: High, typically >1 MΩ
- Component Selection: Op-amps like OPAx340 series or AD8605 with rail-to-rail input/output
- Gain Setting: For a gain of 50, with Rin as 10kΩ, Rf would be 490kΩ

Low-Pass Filter Array:
- Type: 4th order Butterworth for both pressure and temperature signal paths
- Cutoff Frequency: 500 Hz, providing a sharp roll-off without ripple
- Attenuation Rate: -24 dB/octave beyond cutoff
- Power Supply: Matched with the amplifier array
- Component Values: Based on 10 nF capacitors and 40 kΩ resistors for a single stage

Multiplexer:
- Type: 8-to-1 analog multiplexer, such as ADG1606 from Analog Devices
- Switching Time: Microsecond range or lower
- Control Interface: SPI for efficiency
- Supply Voltage: 3.3V or 5V, matching the system
- Channel Isolation: High to prevent signal bleed-over

ADC:
- Resolution: 10 bits or higher, aligning with 1% accuracy requirement
- Sampling Rate: 8 kHz overall, 1 kHz per channel
- Input Voltage Range: 0-5V to match the output range of the signal conditioning
- Interface: SPI for high-speed data transfer
- Example ADC: ADS7950 from Texas Instruments with integrated multiplexer

Data Interface:
- Interface Type: USB 2.0 or higher for compatibility and sufficient data rates
- Power Supply: Through USB connection or external if required
- Data Protocol: Packets with headers, footers, and simple checksum or CRC for error detection

With this architecture and component selection, the data acquisition system can accurately capture the required pressure and temperature data and interface with a computer for analysis.