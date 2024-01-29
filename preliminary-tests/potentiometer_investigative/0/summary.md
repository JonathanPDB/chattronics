Pendulum Angle Measurement System Design

System Overview:
The system is designed to calculate the angle of a pendulum using a potentiometer, with the output being displayed in LabVIEW through a DAQ system. The architecture consists of a potentiometer for angle sensing, a buffer stage to maintain signal integrity, a low-pass filter to eliminate power line noise, an attenuator to scale the voltage to DAQ-compatible levels, an ADC for digital conversion, and the DAQ system for interfacing with LabVIEW.

Potentiometer:
- Type: Linear Taper Potentiometer
- Resistance: 10 kOhm
- Rotational Angle: 180 degrees
- Model Suggestion: Bourns 3590 Precision Potentiometer
- Note: High resolution, tight tolerance for accuracy, high rotational life for durability

Buffer Stage:
- Topology: Non-inverting op-amp configuration
- Input Impedance: >1 MΩ to prevent loading the potentiometer
- Gain: Unity (Gain = 1) for impedance matching
- Op-Amp Suggestions: TL071 for low noise, LM324 for cost-effectiveness
- Power Supply: ±12V to prevent signal clipping

Low-Pass Filter:
- Type: Active low-pass filter
- Topology: Sallen-Key
- Filter order: Second-order
- Cutoff frequency: 40 Hz
- Component values: Assuming unity-gain, C1 = C2 = 0.1 µF, R1 = R2 ≈ 40 kΩ
- Op-Amp: Chosen for low noise and low offset voltage

Attenuator:
- Voltage Divider using two resistors (R1 and R2)
- R1 = 4.3 kOhms, R2 = 10 kOhms
- Configuration: R1 connected in series with R2; the output taken between them
- Calculations: R1 = R2 * [(Vin/Vout) - 1], where Vin = 10V (potentiometer), Vout = 7V (DAQ input)

Analog-to-Digital Converter (ADC):
- Resolution: 12-bit or higher
- Sampling rate: 1 kHz (minimum)
- Input type: Single-ended
- Communication interface: SPI or I2C
- Linearity: INL and DNL < 1 LSB
- Input voltage range: Configurable to match the +/- 7 V range after attenuation

DAQ System:
- Compatibility with LabVIEW for data display and logging
- Digital resolution and sampling rate suitable for pendulum dynamics
- Interface with host computer: USB, PCI, or Ethernet for low-latency data transfer
- Digital processing: Software-based filtering and moving average in LabVIEW for noise reduction and stability
- Data presentation: Real-time graphing, zoom, pan, and logging features in LabVIEW

The system design ensures accurate angle measurement by maintaining signal integrity through buffering, noise elimination through filtering, correct scaling for the DAQ, and appropriate digital conversion and processing for visualization in LabVIEW. Component values and models provided are based on typical requirements and should be refined based on detailed specifications of the DAQ system once known.