Portable Low-Frequency Vibration Measurement Device Design

SENSOR BLOCK
For the sensor, a piezoelectric accelerometer with a sensitivity of 100 pC/g is chosen. The model PCB Piezotronics 352C33 is suggested as a starting point, with an expectation to calibrate it to match the required sensitivity. The sensor must handle a maximum charge output of 500 pC and operate within a temperature range typical for industrial applications (-40°C to +85°C). The accelerometer's sensitivity and maximum expected charge output are used to calculate the gain required in the charge amplifier stage.

CHARGE AMPLIFIER BLOCK
A charge amplifier is designed to convert the charge output from the piezoelectric sensor to the required voltage level. The design parameters for the charge amplifier are:
- Feedback Capacitor (Cf): 1 nF
- Feedback Resistor (Rf): Considering a load capacitance (Cl) of 100 pF and a cutoff frequency (fc) of 0.1 Hz (for a stable low-frequency response), the feedback resistor is calculated as Rf ≈ 1.59 MΩ (rounded to the nearest standard resistor value).
- Operational Amplifier: Recommended are low-noise op-amps such as the AD797 or LT1028.

LOW-PASS FILTER BLOCK
Two options are provided for the filter topology:
Option 1: Butterworth Active Low-Pass Filter
- Order: 2nd order
- Cutoff Frequency: 0.25 Hz for a 3 dB down point
Option 2: Chebyshev Active Low-Pass Filter
- Order: 2nd order
- Cutoff Frequency: Slightly higher than 0.25 Hz with an acceptable passband ripple

BUFFER BLOCK
The buffer stage employs a voltage follower using an operational amplifier to provide a low output impedance and high input impedance. Component selection includes:
- Operational Amplifier: A rail-to-rail input/output op-amp such as the OPAx340 series.
- Bypass Capacitors: 0.1 uF ceramic capacitors for power supply decoupling.

ADC BLOCK
The ADC specifications are chosen based on general requirements for low-frequency vibration measurement:
- ADC Topology: Successive Approximation Register (SAR) ADC
- Sampling Rate: 10 Hz
- Resolution: 12 bits
- Input Range: At least 1 V peak-to-peak
- Interface: SPI or I2C

Each block in the design is chosen to ensure compatibility with the others and to meet the overall requirements of measuring low-frequency vibrations with a portable device. The use of a piezoelectric accelerometer, a charge amplifier with low-noise characteristics, a low-pass filter with a sub-Hz cutoff frequency, a buffer stage with a rail-to-rail op-amp, and a SAR ADC with moderate resolution but low power consumption creates a cohesive system that balances performance, power, and portability.