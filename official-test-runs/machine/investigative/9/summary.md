Design of an Analog Instrumentation System for Monitoring Pressure and Surface Temperature

Pressure Sensor Array:
- Sensor Type: Strain-gauge based pressure transducer
- Operating Range: Typically, 0-10 bar can be assumed for industrial applications
- Sensitivity: Assuming 1 mV full-scale output, the sensitivity would be 0.1 µV/Pa
- Accuracy: ±1% full scale
- Frequency Response: Must be at least 400 Hz to capture relevant signals
- Output: mV level from the strain gauge
- Suggested Sensor Model: Consider industrial-grade options from manufacturers like Honeywell or Sensata

Instrumentation Amplifier Array:
- Gain: To map 1 mV sensor output to a 0-5V ADC input range, a gain of 5000 is required
- Signal Bandwidth: At least 2 kHz to accommodate signals up to 400 Hz
- Power Supply: Assuming a single-sided 5V, use rail-to-rail output op-amps

Temperature Sensor Array:
- Sensor Type: Infrared radiation temperature sensors, suitable for non-contact temperature measurement
- Temperature Range: Assumed industrial range of -20°C to 500°C
- Resolution: Typically around 0.1°C
- Suggested Sensor Model: MLX90614 from Melexis, which has a temperature range of -70°C to +380°C and offers an accuracy of ±0.5°C in the -40°C to +125°C range

Non-Linear Signal Conditioning:
- Topologies:
  - Analog circuitry using operational amplifiers and precision components for piecewise linearization
  - Digital processing with a microcontroller using a lookup table or mathematical model for linearization
- ADC Compatibility: Signal to be scaled to match the ADC’s voltage range (0-5V or 0-10V)

Low Pass Filter Array:
- Cutoff Frequency: 500 Hz to allow signals up to 400 Hz and attenuate higher frequencies
- Filter Order: Second-order low pass Butterworth filter with a roll-off rate of 40 dB per decade
- Damping: Q factor of 0.7071 for a Butterworth response
- Op-Amps: Low-noise, wide bandwidth op-amps like OPA2134 or AD823

Multiplexer:
- Configuration: 8-to-1 analog multiplexer
- Bandwidth: Greater than 1 kHz
- On-Resistance: Typically below 100 ohms
- Example Parts: ADG408/ADG409, CD74HC4051

Analog-to-Digital Converter (ADC):
- Resolution: 12-bit or higher to exceed the 1% sensor accuracy
- Sampling Rate: Greater than 1 kHz per channel, considering a minimum requirement of 800 Hz as per the Nyquist criterion
- Input Voltage Range: Should match the output range of the conditioned sensor signals
- Interface: SPI is recommended for high-speed data transfer
- Power Supply: 3.3 to 5 V, typical for industrial applications
- Example ADC: Successive-approximation register (SAR) ADC with 12-bit resolution, >1 kHz sampling rate, integrated 8-channel multiplexer, and SPI interface

Note: All suggested components and specifications are based on typical industrial applications and standard practices. They should be further refined based on the detailed requirements of the specific machine and application in which the monitoring system will be used.