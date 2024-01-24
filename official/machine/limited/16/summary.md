Analog Instrumentation System for Pressure and Temperature Monitoring

This summary outlines the design for an analog instrumentation system to monitor pressure and temperature at eight points on a machine, with data acquisition for computer analysis.

Pressure Sensor Array:
- Strain-Gauge Sensor: Omega PX309 Series, 0-100 psi, 1 mV/V sensitivity, -40°C to 85°C operating range.
- Instrumentation Amplifier: AD620, AD623 or equivalent, Gain of 5000 for scaling 1 mV to 5 V, configured for the gain using a gain resistor (Rg) calculated based on the amplifier's datasheet.
- Pressure Multiplexer: CD4051B CMOS 8-to-1 analog multiplexer, or ADG1208 differential multiplexer for handling multiple sensor channels.
- Pressure Signal Conditioning: Incorporates low-pass filter and voltage reference, with an op-amp buffer stage for signal integrity.
- Pressure Anti-Aliasing Filter: 2nd order active Sallen-Key Low-Pass Filter, Cutoff frequency at 400 Hz, using resistors of 3.3 kΩ and capacitors of 0.1 µF.

Temperature Sensor Array:
- IR Detector: Suggested model Melexis MLX90614ESF-BAA, -70°C to +380°C temperature range, 10mV/°C linear analog output. Temperature range and resolution accommodate the 1% accuracy requirement.
- Temp Nonlinear Correction: Analog linearization using operational amplifiers configured to match the sensor's nonlinearity, with precision resistors for gain setting and capacitors for stability.
- Temperature Multiplexer: Analog Devices ADG708 or similar 8-to-1 multiplexer to select one temperature sensor output at a time for ADC conversion.
- Temperature Signal Conditioning: Gain stage to scale sensor output, with low-pass filtering using a cutoff frequency of 500 Hz, for resistor values approximating 3.3 kΩ and capacitor values of 0.1 µF.

Analog-to-Digital Converter (ADC):
- ADC: 12-bit resolution to maintain 1% accuracy, with a sampling rate of at least 2 kSPS per channel to prevent aliasing, considering Nyquist theorem. Suggested ADC should support desired data rate and have high input impedance.

Data Acquisition System Interface:
- Communication Interface: USB for easy connectivity or Ethernet for network-based monitoring, with data rates sufficient for the application (USB 2.0 or Fast Ethernet).
- Data Formatting: Binary format with structured metadata for synchronization and error checking.
- Power Consumption and Latency: Low-power design with minimized latency through efficient protocol design and high-speed interfaces.

Overall, the design focuses on accuracy, stability, and compatibility with industrial environments. Component selections and circuit configurations are based on standard practices and the information available, with flexibility to adjust based on further requirements or specific application details.