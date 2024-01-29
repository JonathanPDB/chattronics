Pendulum Angle Measurement System Design

The design of the pendulum angle measurement system entails several stages, including sensor selection, signal conditioning, analog-to-digital conversion, and data acquisition (DAQ). Below is a compilation of the proposed solution organized by architecture blocks.

**POTENTIOMETER SELECTION**
- Type: Linear Rotary Potentiometer
- Resistance: 10 kOhms
- Power Rating: 0.5W to 1W (Typical for sensor applications)
- Tolerance: ±5% or better
- Linearity: ±0.5% or better
- Temperature Coefficient: 50 ppm/°C or better
- Mechanical Angle of Rotation: At least 270 degrees
- Rotational Life: 1 million cycles or more
- Suggested Model: Bourns 3590S-2-103L

**BUFFER AMPLIFIER**
- Gain: 1 (0 dB)
- Input Impedance: >100 kOhms
- Power Supply Rejection Ratio (PSRR): High to reject 50/60 Hz noise
- Offset Voltage: <1 mV
- Bias Current: <1 nA
- Suggested Op-Amp: OPA227 or OPA128

**SCALING AMPLIFIER**
- Desired Gain: Adjustable from 0.5 to 0.9 to scale the potentiometer's output to DAQ's input range
- Cutoff Frequency for Notch Filter: 50 Hz and 60 Hz for AC noise
- Suggested Topology: Non-inverting configuration with a gain-adjustable via trim pot
- Suggested Op-Amp: OPAx177 or AD823A

**ANTI-ALIASING FILTER**
- Recommended First-Order Cutoff Frequency: 50 Hz
- Recommended Second-Order Cutoff Frequency: 50 Hz (Butterworth for flat passband)
- Filter Order: 1 or 2
- Suggested Topology: Sallen-Key or Multiple Feedback for second-order
- Suggested Op-Amp: Rail-to-rail operational amplifier with a bandwidth at least five times the cutoff frequency

**DATA ACQUISITION SYSTEM (DAQ)**
- ADC Resolution: At least 12 bits for 4096 distinct positions
- ADC Sampling Rate: 2 ksps or higher
- Input Voltage Range: Bipolar, at least +/- 7 V
- Suggested ADC Type: Successive approximation register (SAR)
- Communication Interface: Compatible with DAQ's input (e.g., SPI, I2C)

By implementing the design recommendations above, the system will be capable of accurately measuring the angle of a pendulum with the predefined requirements. The potentiometer will convert the pendulum's angle into a variable resistance, which is then translated into a voltage signal. This signal will be buffered and scaled to fit the DAQ's input range. It will also be filtered to remove unwanted noise, especially from power lines at 50 and 60 Hz, as well as to prevent aliasing of the signal before being digitized by the DAQ's ADC for data processing and analysis.