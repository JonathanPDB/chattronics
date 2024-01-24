Comprehensive Design of an Analog Instrumentation System for Machine Monitoring

The project entails designing an analog instrumentation system to monitor pressure and surface temperature variations at eight points on a machine, with the data to be analyzed by a computer. The system includes an array of pressure and temperature sensors, signal conditioning circuits, a multiplexer, a low pass filter, and an analog-to-digital converter (ADC).

Pressure Sensor Array:
- Sensor Type: Strain-gauge pressure sensors with a 1 mV maximum output.
- Suggested Model: Honeywell Sensotec Model TJE series, sensitivity matched to the expected pressure range.
- Accuracy: 1%
- Instrumentation Amplifier Gain: Assuming a 0 to 5 V ADC range, a gain of 5000 is needed. This can be achieved with a two-stage approach: a first stage with a gain of 100 and a second stage with a gain of 50.
- Precision Resistor for Gain Setting: Using a 0.1% tolerance resistor in the gain network to achieve the 1% accuracy requirement.

Temperature Sensor Array:
- Sensor Type: Infrared radiation detectors for non-contact measurement with a nonlinear scale and a maximum output of 100 mV.
- Suggested Models: Texas Instruments TMP006 or Melexis MLX90614.
- Accuracy: 1%
- Instrumentation Amplifier Gain: To bring the maximum output to a 0 to 5 V ADC range, a gain of 50 is calculated.
- Gain Resistor for AD620 (Example): RG = (49.4 kOhm) / (50 - 1) = 1.02 kOhm, selecting the nearest standard value of 1 kOhm.

Multiplexer:
- Type: Analog Multiplexer IC (e.g., ADG508F or CD74HC4067) or FPGA-based for higher performance requirements.
- Signal Frequencies: Up to 400 Hz expected from both pressure and temperature sensors.
- Sampling Rate: A conservative rate of at least 2 kHz per channel is recommended.

Low Pass Filter:
- Type: Butterworth (third-order) using a Sallen-Key topology.
- Cutoff Frequency (fc): 500 Hz to provide a buffer above the maximum signal frequency.
- Component Values: Calculated based on operational amplifier characteristics and power supply voltages.

Analog to Digital Converter (ADC):
- Type: Successive Approximation Register (SAR) ADC.
- Sampling Rate: At least 1 kHz per channel.
- Resolution: 16 bits to balance cost and performance.
- Input Type: Differential inputs to increase noise immunity.
- Communication Interface: SPI for speed and reliability.
- Operating Voltage: Compatible with TTL logic levels (3.3V or 5V).

General Considerations:
- Precision resistors and standard value capacitors should be used to maintain the performance of the amplifiers and filters.
- Power supply choices should be considered based on the availability and compatibility with other system components.
- Calibration procedures and environmental protection for the sensors may be necessary based on the final application specifics.

The design provides a robust framework for the data acquisition system. However, specifics regarding pressure ranges, ADC input ranges, and environmental conditions will need to be defined to finalize component values and calibrations.