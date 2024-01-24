Design Summary for a Machine Monitoring System

Analog Instrumentation and Data Acquisition System for Pressure and Temperature Monitoring

1. Pressure Sensors:
   - Suggested Model: Honeywell Sensotec Model TJE series or equivalent strain-gauge pressure sensors.
   - Output Signal: 1 mV (max).
   - Recommended Gain for Amplification: 1000x to convert 1 mV to 1 V for ADC input.

2. Temperature Sensors:
   - Sensor Type: Non-contact infrared radiation detectors.
   - Suggested Model: MLX90614 from Melexis or equivalent.
   - Output Signal: 100 mV (max).
   - Recommended Gain for Amplification: 10x to convert 100 mV to 1 V for ADC input.

3. Instrumentation Amplifiers:
   - Models: AD620, INA118, or INA128/INA129 series.
   - Gain Configuration: 
     - Pressure Sensors: Gain of 1000x.
     - Temperature Sensors: Gain of 10x.
   - Supply Voltage: ±10 V (typical for industrial applications).
   - Input Impedance: Megaohms to gigaohms range.
   - Bandwidth: Minimum of 2 kHz to accommodate signal frequencies up to 400 Hz.

4. Analog Multiplexer:
   - Models: ADG1606 or CD74HC4067.
   - Configuration: 16-to-1 analog multiplexer.
   - On-Resistance: Low, e.g., 4.5 ohms for ADG1606.
   - Switching Time: Fast, e.g., 175 ns for ADG1606.
   - Protection Resistors: 100 ohms at each input of the multiplexer.
   - Bypass Capacitors: 100 nF close to the power supply pins.

5. Anti-Aliasing Filter:
   - Topology: Active low-pass Butterworth filter.
   - Order: 2nd order or 4th order if steeper roll-off is required.
   - Cutoff Frequency (f_c): 500 Hz to prevent aliasing while allowing signal frequencies up to 400 Hz.
   - Op-amps: Low-noise, high-speed, e.g., TL072, AD823.
   - Resistors: 1% metal film for precision.
   - Capacitors: Ceramic or film for stability.

6. Analog-to-Digital Converter (ADC):
   - Topology: Successive Approximation Register (SAR) ADC.
   - Resolution: 12-bit (or 16-bit for greater precision).
   - Sampling Rate: Minimum 12.8 ksamples/s aggregate (800 sps per channel, 16 channels).
   - Interface: SPI for higher throughput.
   - Accuracy: ±1 LSB.
   - Integral and Differential Nonlinearity (INL/DNL): Within ±1 LSB.

7. Computer (Digital Data Processing):
   - Data Acquisition Software: LabVIEW or MATLAB for data acquisition and processing capability.
   - Processing: Capable of both real-time and batch processing.
   - Data Handling: Include signal conditioning, data logging, algorithm development, and visualization tools.
   - ADC Interfacing: Ensure compatibility with standard protocols (SPI, I2C, USB, or Ethernet).