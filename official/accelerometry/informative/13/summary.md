Design of a Portable Low-frequency Vibration Measurement Device

Sensor Block: Piezoelectric Accelerometer
- Sensor Type: Piezoelectric accelerometer
- Sensitivity: 100 pC/g
- Model Suggestion: Endevco Model 2271A or similar
- Operating Conditions: Industrial or commercial grade, with a general operating temperature range from 0°C to 70°C, or wider if the device is used in harsher environments.

Charge Amplifier Block:
- Gain Calculation: Assuming a peak acceleration of 0.5 g, the output charge of the accelerometer will be 50 pC. To achieve a 1 V peak-to-peak output, the gain should be 20 MV/C.
- Amplifier Type: Low-noise operational amplifier
- Power Supply: Standard ±5V or ±12V, unless specific limits are provided
- Passive Components: High input impedance relative to the sensor capacitance

Low-pass Filter Block:
- Filter Type: Active Butterworth filter for maximally flat passband
- Order: Fourth-order (four-pole) design for a -24 dB/octave roll-off
- Cutoff Frequency: Approximately 5 Hz to allow signals at 2 Hz to pass through with minimal attenuation
- Components: Standard resistor value of 10 kΩ, and capacitor values calculated using fc = 1/(2π√(R1R2C1C2)) with C1 = C2 = 1 µF to set fc at 5 Hz
- Power Supply: Assuming a standard ±15V supply for active filter designs
- Operational Amplifiers: Low-noise Op-Amps with bandwidth exceeding the highest frequency of interest (e.g., TL074 or similar)

Analog-to-Digital Converter (ADC) Block:
- ADC Type: Sigma-Delta ADC recommended for high resolution and low noise
- Resolution: At least 16 bits to ensure fine granularity of the measured vibrations
- Sampling Rate: 10 to 20 samples per second (10-20 Hz) to capture the maximum frequency of interest with an adequate margin
- Signal Bandwidth: Slightly above 2 Hz, with a low-pass filter cutoff frequency of approximately 5 Hz
- Power Consumption: Low-power ADC for portable device considerations
- Interface/Digital Output Format: SPI or I2C interface depending on the microcontroller or data processing unit used; single-ended or differential output format based on system requirements