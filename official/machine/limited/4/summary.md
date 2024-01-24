Design of an Analog Instrumentation System for Pressure and Temperature Monitoring

System Architecture Overview:
The system is designed to monitor pressure and surface temperature variations at eight points on a machine. It consists of an array of strain-gauge pressure sensors and infrared radiation temperature sensors whose outputs are conditioned by respective signal conditioning blocks. The conditioned signals are then fed into a multiplexer, filtered by an anti-aliasing filter, digitized by an ADC, and finally interfaced with a computer for data analysis.

Pressure Sensor Array:
- Sensor Type: Strain-gauge pressure sensors with identical sensitivity.
- Output: Maximum of 1 mV.
- Suggested Model: Honeywell HSC Series or NXP MPX Series.
- Accuracy: 1% accuracy requirement met with sensors of appropriate caliber.

Temperature Sensor Array:
- Sensor Type: Infrared radiation detectors for surface temperature measurement.
- Output: Maximum of 100 mV with nonlinearity; requires linearization.
- Suggested Model: MLX90614 from Melexis, or equivalent with analog output or an added DAC.
- Accuracy: 1% accuracy requirement with a resolution of 0.1°C or better.

Instrumentation Amplifiers (for Pressure):
- Gain: Required gain of 5000 calculated from the 1 mV sensor output to match a 5V ADC input range.
- Suggested Op-Amp: AD8421 or INA118.
- Power Supply: Dual ±15V or single 5V, chosen based on op-amp specifications.
- Filtering: Low-pass filter with a cutoff frequency above 400 Hz, implemented to reduce noise.

Instrumentation Amplifiers (for Temperature):
- Gain: 50, to amplify the linearized temperature signal from 100 mV to a usable range up to 5V for the ADC.
- Suggested Op-Amp: AD620 or INA333 for their precision and low noise characteristics.
- Power Supply: Same as pressure amplification stage.
- Filtering: Similar low-pass filter as in the pressure signal path.

Linearization Circuitry (for Temperature):
- Methods: Analog piecewise linearization using op-amps or digital linearization via microcontroller.
- Suggested Components: Precision resistors and op-amps for analog or microcontroller with integrated ADC for digital.

16-to-1 Analog Multiplexer:
- Model: CD74HC4067 or equivalent CMOS-based analog multiplexer.
- Control: Sequentially selects signals from each of the 16 amplified channels (8 pressure and 8 temperature).
- Speed: Switch time significantly shorter than the ADC sampling interval.

Anti-Aliasing Filter:
- Type: Active low-pass filter using a Sallen-Key topology.
- Cutoff Frequency: Set at 450 Hz to attenuate frequencies higher than the maximum frequency of interest (400 Hz).
- Order: Fourth-order implemented as two cascaded second-order stages.
- Components: Assuming capacitor values of 10 nF, resistor values are calculated to be approximately 35.7 kΩ for the desired cutoff frequency.

Analog-to-Digital Converter (ADC):
- Type: Successive Approximation Register (SAR) ADC.
- Resolution: 12-bit or higher to ensure 1% accuracy.
- Sampling Rate: Minimum 2 kHz per channel, considering Nyquist criteria and filter roll-off.
- Interface: Recommended SPI for high-speed data transfer.

Data Interface to Computer:
- Communication Protocol: USB or Ethernet, depending on the real-time data transfer requirements.
- Data Rate: 204.8 kbps calculated for 16 channels at a 16-bit ADC with a sampling rate of 800 samples per second.
- Integrity: Intrinsic error detection and correction mechanisms of chosen protocols, with additional application-level checksum or CRC if needed.
- Isolation: Use of opto-isolators or magnetic isolation for galvanic isolation if needed.

The design choices presented are based on industry standards and reasonable assumptions to create a robust and accurate data acquisition system capable of meeting the project's requirements. Component selection, power supply considerations, and interfacing details are suggested to ensure a functional and reliable system.