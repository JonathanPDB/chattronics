Design of a Data Acquisition System for Pressure and Temperature Monitoring

System Architecture Overview:
The data acquisition system is composed of several key blocks: pressure sensors, temperature sensors, instrumentation amplifiers, a multiplexer, a sample and hold circuit, an anti-aliasing filter, an analog-to-digital converter (ADC), and a microcontroller. Each block is designed to ensure accurate capture of pressure and temperature data at multiple points on a machine.

Pressure Sensors:
- Type: Strain gauge pressure sensors
- Output: Maximum of 1 mV
- Suggested Range: 0-10 bar
- Sensitivity: Assumed 0.1 mV/bar if the maximum pressure is 10 bar
- Suggested Model: Honeywell Sensing and Productivity Solutions' Model 19C series
- Accuracy: 1% full-scale

Temperature Sensors:
- Type: Infrared radiation detectors for non-contact surface temperature measurement
- Output: Maximum of 100 mV
- Suggested Model: MLX90614 with a typical temperature measurement range of -70°C to +380°C
- Sensitivity and Resolution: Factory-calibrated resolution of 0.02°C and an accuracy of ±0.5°C

Instrumentation Amplifiers:
- Gain (Pressure): 5000 V/V to amplify the 1 mV sensor signal to a standard ADC input range of 5V
- Gain (Temperature): 50 V/V to amplify the 100 mV sensor signal to the same ADC input range
- Suggested Op-Amp: Precision type such as AD620 or INA128
- Other: Use of precision resistors for accurate gain setting

Multiplexer:
- Type: 16:1 analog CMOS multiplexer
- Bandwidth: >4 kHz to handle the maximum signal frequency of 400 Hz
- Suggested Model: ADG1606 from Analog Devices
- On-Resistance: Less than 100 ohms
- Crosstalk: -80 dB or better
- Switching Time: Less than one microsecond

Sample & Hold Circuit:
- Op-Amp: High-speed precision op-amp, like AD8031 or OPA657
- Sampling Capacitor: Around 0.1 to 1 μF, polypropylene or Teflon type
- Buffer: A high-speed buffer to drive the ADC input

Anti-Aliasing Filter:
- Type: Active Low-Pass
- Configuration: Sallen-Key
- Order: 2nd (attenuation rate of -40 dB/decade)
- Cutoff Frequency: 600 Hz
- Components: Low-noise op-amp with >2 MHz bandwidth, 1% metal film resistors, low-tolerance capacitors

Analog-to-Digital Converter (ADC):
- Type: SAR ADC
- Resolution: At least 12-bit to achieve the desired 1% accuracy
- Interface: Assume SPI for high data rates
- Input Range: Assumes a unipolar range compatible with the amplified sensor outputs

Microcontroller/Processor:
- Capabilities: Must handle real-time digital filtering, control the multiplexer, and manage data acquisition
- Flexible communication interfaces such as USB, SPI, I2C, UART, or Ethernet
- Adequate memory to buffer data if needed

General Assumptions and Recommendations:
- The exact specifications of the pressure and temperature ranges, ADC sampling rate, and other numerical values need to be defined based on the system's application.
- A conservative design approach has been taken, assuming typical industrial conditions and standard component ranges.
- The system is designed to meet a general accuracy requirement of 1% for both pressure and temperature measurements.
- All suggested models and types are subject to change based on the final application requirements and component availability.