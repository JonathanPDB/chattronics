Analog Instrumentation System for Pressure and Temperature Monitoring

Pressure Measurement Subsystem:
- Strain-Gauge Sensor: Full-bridge strain-gauge, Honeywell Sensotec Model TJE or similar, with a pressure range of 0 to 10 MPa and sensitivity of 0.1 μV/Pa.
- Instrumentation Amplifier: Three-op-amp configuration (e.g., AD620 or INA118) with a gain of 5000 to amplify the 1 mV signal from the strain-gauge to a suitable level for the ADC. The gain can be set using precision resistors and a potentiometer for calibration.
- Low Pass Filter: Second-order Sallen-Key Butterworth filter designed to have a cutoff frequency of 350 Hz to eliminate frequencies above 400 Hz. Calculated component values for R = 10 kΩ and C = 4.7 nF.
- Sample & Hold Circuit: Using a monolithic Sample & Hold IC such as the LF398, with a hold capacitor value of 0.01 μF to 0.1 μF, polypropylene or Teflon for low leakage.

Temperature Measurement Subsystem:
- Infrared Detector: Non-contact infrared thermopile sensor, e.g., Heimann Sensor HTM1735, capable of measuring temperatures from -20°C to +1000°C.
- Logarithmic Amplifier: AD8307 or similar, with a scale factor K chosen so that the full temperature range corresponds to the ADC input range (e.g., 0V to 5V).
- Low Pass Filter: Similar second-order Sallen-Key Butterworth active filter as in the pressure subsystem, with component values designed for a cutoff frequency just above the signal frequency of interest (e.g., 450 Hz).
- Sample & Hold Circuit: Similar configuration as in the pressure subsystem, ensuring fast acquisition and low droop rate during the hold phase.

Analog Multiplexer:
- 16-Channel Analog Multiplexer: CD74HC4067 or equivalent, capable of handling the voltage levels from both the pressure and temperature sensor conditioning circuits with fast enough switching to sample each channel sequentially at the desired ADC rate.

Analog-to-Digital Converter (ADC):
- 12-bit SAR ADC, capable of at least 32 kSPS total throughput to sample all 16 channels within a 1-second interval. Selected for its resolution finer than 1 mV (least significant bit) to resolve 1% changes in sensor outputs. Low-power, small-footprint, and SPI interface for communication with the computer.

Output to Computer:
- Interface: USB 3.0 or Ethernet for data transfer to computer.
- Data Format and Protocol: Raw binary or ASCII encoding, with a packet structure to indicate data type (pressure or temperature) and sensor index.

General Design Considerations:
- All electronic components selected should be suitable for the industrial temperature range of -40°C to 85°C.
- The system should incorporate proper shielding and grounding to minimize electrical noise and interference.
- Power supply voltages for the analog circuitry are assumed to be ±15 V, with consideration for low power consumption and efficient operation.
- The PCB material should be FR4 with a ground plane for noise immunity and thermal stability.