Pendulum Angle Measurement System Design

This summary encompasses the design considerations and recommendations for a pendulum angle measurement system using a potentiometer, signal conditioning circuitry, and a DAQ system.

**Potentiometer Sensor Block**
- Type: Linear Rotary Potentiometer
- Model: Bourns 3590 Precision Potentiometer Series or similar
- Resistance Value: 10 kΩ
- Operating Temperature Range: -40°C to 85°C (assumed industrial standard)
- Voltage Rating: 10 V across terminals, but typically operated at 5V
- Power Rating: ≥10 mW, recommended 0.5 W to 1 W for safety margin
- Mechanical Endurance: Rated for at least 1 million cycles (assumed)
- Tolerance: ±5%
- Temperature Coefficient: ±250 ppm/°C (assumed industrial standard)
- Environmental Conditions: Sealed or industrial-grade for dust and humidity resistance
- Electrical Connection: Three-terminal with conductive plastic element for smooth output

**Buffer Stage**
- Topology: Unity-gain voltage follower using an op-amp
- Op-Amp Suggestion: TL081 or similar with a wide bandwidth
- Power Supply: Typically +/- 15V

**Scaling Amplifier**
- Topology Options: Non-inverting amplifier or Differential amplifier
- Gain: Variable, to be tuned during implementation. Example calculation assuming 2.5V range for pendulum angles:
  - G = (DAQ_max_voltage - DAQ_min_voltage) / (Pot_max_voltage - Pot_min_voltage)
  - G = (7V - (-7V)) / (3.75V - 1.25V) = 5.6 (assuming)
- Components: Operational amplifier chip, precision resistors, and a potentiometer or trimmer for gain adjustment

**Low Pass Filter**
- Filter Type: Low Pass Butterworth
- Order: 2nd order
- Cutoff Frequency: 100 Hz
- Component Values: To be determined, typically kilohm-range resistors and nano to microfarad-range capacitors

**Notch Filters**
- Notch Filter for 50Hz:
  - Topology: Active Twin-T Notch Filter
  - Center Frequency (f0): 50 Hz
  - Q Factor: ≈10 for sharp attenuation
  - Attenuation: ≥20 dB at 50 Hz
  - Op-Amp: TL072 or NE5532
  - Resistor Values: 10 kΩ
  - Capacitor Values: 330 nF (adjusted for availability)
- Notch Filter for 60Hz:
  - Topology: Active Twin-T Notch Filter
  - Center Frequency (f0): 60 Hz
  - Q Factor: 10 to 15
  - Attenuation: ≥20 dB at 60 Hz
  - Op-Amp: TL071 or LM358
  - Resistor Values: Based on the chosen Q factor and standard values
  - Capacitor Values: Standard value closest to calculated requirement

**Data Acquisition System (DAQ)**
- ADC Topology: Successive Approximation Register (SAR)
- Sampling Rate: ≥1000 samples per second
- Resolution: 12-bit to 16-bit (assumed suitable)
- Input Range: Configurable to ±7 V
- Input Impedance: High (MΩ range)
- Op-Amp Power Supply: Compatible with DAQ voltage range, such as +/- 9V

The system is designed with flexibility to accommodate fine-tuning during implementation, as specific requirements for some parameters were not provided. The selection of components and their values are based on typical industry standards and assumptions made from the information given. The final design and component values will need validation during prototyping and testing.