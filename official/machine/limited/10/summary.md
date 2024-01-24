ANALOG INSTRUMENTATION DESIGN FOR PRESSURE AND TEMPERATURE MONITORING SYSTEM

Pressure Sensors:
- Sensor type: Piezoresistive strain-gauge pressure sensor
- Suggested model: Honeywell Trustability HSC Series
- Pressure range: 0-10 bar
- Accuracy: ≤ 0.25% FS
- Sensitivity: Sufficient for required resolution
- Operating temperature range: -40°C to 85°C
- Supply voltage: 3.3 to 5 V DC
- Output: 1 mV maximum, Wheatstone bridge configuration

Temperature Sensors:
- Sensor type: Infrared radiation detector
- Suggested model: MLX90614 by Melexis or OS-MINIUSB by Omega
- Temperature range: -20°C to 500°C (MLX90614), up to 1030°C (OS-MINIUSB)
- Resolution: At least 0.5°C to meet 1% accuracy
- Output: Nonlinear, up to 100 mV maximum

Instrumentation Amplifiers:
- Topology: Three-op-amp instrumentation amplifier configuration
- Op-Amps: AD8421 or INA118
- Gain for pressure sensors (Gp): 5000 to scale 1 mV to 5 V output
- Gain for temperature sensors (Gt): 50 to scale 100 mV to 5 V output
- Features: Adjustable gain, low offset voltage, high CMRR

16-to-1 Analog Multiplexer:
- Model: CD4051B or similar CMOS analog multiplexer
- Input voltage range: 0-5V (to match sensor amplifications)
- Bandwidth: ≥1 kHz
- Input impedance: >1 MΩ
- Crosstalk: Ideally < -80 dB

Sample & Hold Circuit:
- Topology: Track-and-hold or closed-loop S&H amplifier
- Hold Capacitor (CH): 0.01 µF to 0.1 µF, low leakage
- Analog Switch: CMOS with low on-resistance and low charge injection
- Acquisition Time: ≤2 µs
- Hold Time: ≥ ADC conversion time + multiplexer switching time
- Buffer Amplifier: Unity gain, high slew rate (>5 V/µs), low offset (<1 mV)

Anti-Aliasing Filter:
- Topology: 4th-order Butterworth low-pass filter
- Cutoff frequency: 500 Hz
- Attenuation rate: 80 dB/decade

Analog-to-Digital Converter:
- Resolution: 12 bits or higher
- Sampling Rate: ≥2500 samples per second per channel
- Input Voltage Range: 0-5V
- SNR: >90 dB
- THD: < -90 dB
- Features: Multiplexer, Sample-and-Hold, Digital Filtering

Temperature Linearization:
- Method: Digital linearization using a microcontroller or analog circuit
- Op-Amps for Analog Linearization: LM358 or TLV2372
- Microcontroller for Digital Linearization: Arduino Nano, STM32F103
- Components: 0.1% metal film resistors, ceramic or film capacitors

Digital Interface to Computer:
- Interface: USB 2.0 or USB 3.0 for high-speed data transfer
- Data Formatting: CSV or plain text
- Data Transfer: Continuous streaming or batch processing, depending on real-time needs