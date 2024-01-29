Comprehensive Analog Instrumentation System for Machine Monitoring

Pressure Sensor Array:
- Sensors: Piezoresistive strain-gauge pressure sensors, Honeywell Sensotec Model TJE, with a pressure range of 0 to 1000 psi.
- Sensitivity: Assuming 1 to 3 mV/V at full scale.

Amplification Block Pressure:
- Topology: Instrumentation Amplifier.
- Gain Calculation: To amplify the 1 mV max output to a standard ADC input range of 0-5V, a gain of 5000 is required.
- Components: Precision resistors for fixed gain (1kΩ for adjustable gain), low-noise op-amps like INA125P or AD620.

Filtering Block Pressure:
- Topology: Butterworth low-pass filter using Sallen-Key configuration.
- Order: Second-order (two-pole) or fourth-order if necessary.
- Cutoff Frequency: Set at 450 Hz to allow signals up to 400 Hz with minimal attenuation.

Temperature Sensor Array:
- Sensors: Infrared radiation temperature sensors, Melexis MLX90614, capable of measuring -70°C to +380°C.
- Accuracy: ±0.5°C with resolution better than 0.02°C.
- Output: Digital (PWM or SMBus), simplifies signal conditioning requirements.

Linearization Block Temp:
- Method: Software-based using a microcontroller with an ADC. Lookup table or mathematical algorithm to correct nonlinearity.
- Components: High-resolution ADC (12-bit or higher), microcontroller (e.g., STM32 or PIC series).

Amplification Block Temp:
- Topology: Instrumentation Amplifier.
- Gain Calculation: Gain of 50 required to match a linearized sensor output range of 0mV to 100mV to the ADC input range of 0V to 5V.
- Components: AD620 with a 1kΩ gain-setting resistor (Rg) for a gain of 50, and additional RC filter as needed.

Filtering Block Temp:
- Topology: Butterworth low-pass filter with Sallen-Key configuration.
- Order: Second-order (two-pole) filter.
- Cutoff Frequency: 450 Hz to ensure integrity of signals up to 400 Hz.

Multiplexer:
- Type: 8-to-1 Analog Multiplexer.
- Requirements: Low on-resistance (<100 Ohms), high off-isolation, and bandwidth of at least 1 MHz.
- Components: Robust CMOS or BiCMOS analog multiplexer, such as ADG1608 or MAX4051A.

Analog-to-Digital Converter (ADC):
- Type: Successive Approximation Register (SAR) ADC.
- Resolution: 12-bit or higher to ensure 1% accuracy.
- Sampling Rate: At least 2 kHz per channel to capture the maximum signal frequency of 400 Hz without aliasing.
- Input Range: Adjustable to match amplified sensor voltages.

Digital Interface to Computer:
- Protocol: USB for communication, offering ease of use and compatibility with most modern computers.
- Data Rate: USB 2.0 with a transfer rate of 480 Mbps, more than adequate for the data rate required.
- Electrical Isolation: Opto-isolators for electrical isolation.
- Additional Processing: Packetization with error-checking (e.g., CRC) before sending data to the computer.

Overall, this comprehensive system is designed to accurately monitor pressure and surface temperature variations at multiple points on a machine, with careful consideration given to maintaining a 1% accuracy through the use of precise amplifiers, filters, and sensor arrays. The digital data is processed and interfaced with the computer via a robust USB connection, ensuring reliable and efficient analysis.