Analog Instrumentation System for Machine Monitoring

Overview:
The system is designed to monitor pressure and surface temperature at eight points of a machine. The architecture comprises sensor arrays for pressure and temperature, signal conditioning blocks for both sensor types, a multiplexer, an anti-aliasing filter, an ADC, a data interface, and a computer for data analysis.

Pressure Sensor Array:
- Sensor Type: Strain-gauge pressure sensor
- Maximum Output: 1 mV
- Suggested Model: Honeywell Model TJE
- Expected Operating Range: 0 to 100 psi (assuming typical industrial range)
- Sensitivity: Approximately 500 µV/V/psi (derived from maximum output)
- Resolution: Dependent on ADC resolution, for a 12-bit ADC, minimum detectable change would be 0.024 psi.

Temperature Sensor Array:
- Sensor Type: Infrared radiation detectors for surface temperature measurement
- Maximum Output: 100 mV
- Accuracy: 1% of reading
- Nonlinear Scale: Requires digital linearization post-ADC or analog correction via DAC
- Suggested Model: Melexis MLX90614 Infrared Thermometer

Signal Conditioning Block - Pressure:
- Amplification: Using an instrumentation amplifier like the AD620 with a gain of 5000 to scale 1 mV to 5V.
- Filtering: RC low-pass filter with a cutoff frequency of 500 Hz (R = 1 kΩ, C ≈ 330 nF).

Signal Conditioning Block - Temperature:
- Amplification: Assuming a 5V ADC range, a gain of 50 is needed (5V/100mV).
- Filtering: RC low-pass filter with a cutoff frequency of 500 Hz (R = 10 kΩ, C ≈ 330 nF).
- Nonlinearity Correction: Addressed digitally post-ADC or with a DAC controlled by a microcontroller for analog correction.

Multiplexer:
- Type: 16-to-1 CMOS analog multiplexer
- Example Part: CD74HC4067
- Considerations: Fast switching speed (>4 kHz), minimal signal-to-noise ratio (SNR ≥ 60 dB).

Anti-Aliasing Filter:
- Type: Active Butterworth low-pass filter
- Order: Fourth-order, providing 80 dB/decade roll-off
- Cutoff Frequency: 500 Hz
- Op-Amps: Low-noise, precision op-amps (e.g., OPAx140 series)
- Components: High-stability, low-tolerance resistors and capacitors.

ADC:
- Type: Successive Approximation Register (SAR)
- Resolution: 16-bit
- Sampling Rate: At least 2 kSPS
- Input Range: 0 to 100 mV
- Interface: Compatible with computer system (e.g., SPI, I2C)
- Features: Integrated PGA (optional), low power consumption.

Data Interface:
- Communication Interface: USB or Ethernet, depending on data throughput requirements
- Data Throughput: At least 1 Mbps to handle the aggregated data from all sensors
- Data Buffering: FIFO buffer to manage data flow and temporary storage
- Environmental Protection: Robust protection against industrial environmental factors.

Computer:
- Specifications: Capable of handling the data throughput and processing requirements
- Software: Flexible to accommodate various data formats and protocols (e.g., MATLAB, Python, LabVIEW)
- Storage: Local hard drive/SSD or cloud-based storage solution
- Visualization: Real-time data visualization tools or dashboards for monitoring.

Calculations and Assumptions:
- Gain for Pressure Signal Conditioning: G = 5V / 1mV = 5000
- Gain for Temperature Signal Conditioning: G = 5V / 100mV = 50
- Low-Pass Filter Cutoff Frequency for Pressure: f_c = 1/(2πRC) with R = 1 kΩ, C ≈ 318 nF (rounded to 330 nF)
- Low-Pass Filter Cutoff Frequency for Temperature: f_c = 1/(2πRC) with R = 10 kΩ, C ≈ 318 nF (rounded to 330 nF)
- ADC Sampling Rate: Minimum of 2 times the highest signal frequency, selected 2 kSPS for better accuracy.

Note: The above component values and recommendations are based on standard industrial ranges and typical specifications, and should be fine-tuned based on the final sensor selections and system requirements.