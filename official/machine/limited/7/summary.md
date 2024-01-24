Design of an Analog Monitoring System for Pressure and Temperature

In designing an analog monitoring system for machine pressure and temperature, the architecture is segmented into several key blocks: sensor arrays, signal conditioning (including amplification and filtering), sample and hold circuits, a multiplexer, an analog-to-digital converter (ADC), and a data interface. Below is a compilation of the design choices and specifications for each block.

Pressure_Sensor_Array:
- Sensor Type: Industrial-grade strain-gauge pressure sensors
- Pressure Range: 0 to 10 bar (0 to 145 psi)
- Output: Maximum 1 mV, with sensitivity stable across the range
- Operating Temperature Range: -40°C to 125°C
- Example Sensor: Honeywell Model STJE or equivalent

Instrumentation_Amplifier_Pressure:
- Topology: Three-op-amp instrumentation amplifier
- Gain: 5000 (to scale 1 mV sensor output to a 5V ADC input range)
- Power Supply: ±15V or single supply if unipolar signal
- CMRR: Above 80 dB
- Example IC: Analog Devices AD620 or Texas Instruments INA118

Lowpass_Filter_Pressure:
- Filter Type: 2nd-order lowpass Butterworth filter
- Cutoff Frequency: 400 Hz
- Roll-off Rate: 12 dB/octave
- Components: Precision resistors (1% tolerance) and capacitors (5% tolerance)

Sample_and_Hold_Pressure:
- IC: LF398 or equivalent S&H IC
- Hold Capacitor: 0.1µF ceramic capacitor
- Buffer Amplifier: Precision op-amp like OP07
- Power Supply: ±15V or ±5V

Temperature_Sensor_Array:
- Sensor Type: Infrared radiation detectors, non-contact
- Temperature Range: -40°C to +150°C
- Resolution: 0.02°C
- Power Supply: 3.3V to 5V
- Output Signal: I²C-compatible digital interface
- Example Sensor: MLX90614 or similar

Instrumentation_Amplifier_Temperature:
- Topology: Three-op-amp instrumentation amplifier
- Gain: 50 (to scale 100 mV sensor output to a 5V ADC input range)
- Power Supply: 5V single-supply
- CMRR: 80 dB or higher
- Example IC: Analog Devices AD8421 or Texas Instruments INA333

Lowpass_Filter_Temperature:
- Filter Type: 4th-order lowpass Butterworth filter
- Cutoff Frequency: 400 Hz
- Roll-off Rate: 24 dB/octave
- Components: Precision resistors (1% tolerance) and capacitors (5% tolerance)

Sample_and_Hold_Temperature:
- Same specifications as Sample_and_Hold_Pressure

Multiplexer:
- Type: 16:1 differential MUX
- Crosstalk Isolation: -80 dB or higher
- Switching Time: Less than 1 ms
- Control Logic: Compatible with digital IO voltage levels
- Example IC: Analog Devices ADG1606

ADC:
- Resolution: 12-bit or higher
- Sampling Rate: ≥2 kHz per channel
- Voltage Range: 0-5V
- Interface: SPI (Serial Peripheral Interface)
- Operating Temperature Range: Industrial range (-40°C to +85°C)

Data_Interface:
- Communication Protocol: USB 2.0 or 3.0 for plug-and-play functionality
- Data Throughput: 102.4 kbps for 16-bit resolution at 800 sps over 8 channels
- Safety and Data Integrity: CRC or similar error checking, potential for encryption if required
- Hot-Swapping: Supported via USB

Calculation and Design Steps:
- Determined the gain for each instrumentation amplifier based on the sensor's maximum output and the desired ADC input range.
- Chose the filter order and calculated the cutoff frequency to prevent aliasing while allowing relevant frequencies up to 400 Hz.
- Selected components with appropriate tolerances and characteristics, considering temperature stability and noise immunity.
- Ensured power supply compatibility across all stages, with a preference for standard industrial voltage levels.

This solution provides a comprehensive analog monitoring system architecture using industrial standards and best practices for monitoring pressure and temperature on machinery with high accuracy and reliability.