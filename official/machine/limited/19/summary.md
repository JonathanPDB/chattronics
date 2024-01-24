Industrial Machine Pressure and Temperature Monitoring System Design

Pressure Measurement Path:
- Sensors: Strain-gauge pressure sensors with a maximum output of 1 mV and expected sensitivity of approximately 2 mV/V/bar.
- Instrumentation Amplifiers: Dual-stage configuration with a total gain of 4900 to amplify the 1 mV pressure signal up to a 5V ADC input range. Each stage with a gain of 70 using precision op-amps like AD620 or INA118.
- Power Supply: ±15V for the instrumentation amplifiers.

Temperature Measurement Path:
- Sensors: Infrared temperature sensors capable of measuring temperatures in the range of -20°C to 500°C. A response time of less than 500 ms and a resolution better than 0.1°C are desirable.
- Linearization Circuit: A microcontroller-based approach with ADC and lookup tables or polynomial approximation algorithms to linearize the temperature sensor's output.
- Amplifiers: A non-inverting amplifier topology to amplify the temperature sensor output with a gain of 50, to match the 5V ADC input range.

Shared Signal Path:
- Multiplexer: An 8-to-1 analog multiplexer with a low on-resistance and fast switching capabilities for an ADC sampling rate of at least 1 kSPS per channel.
- Anti-Aliasing Filter: A 4th-order Butterworth filter with a cutoff frequency of 350 Hz to prevent aliasing of frequencies above the Nyquist frequency.

ADC:
- Resolution: 16-bit resolution to achieve the 1% accuracy requirement, which corresponds to a resolution of 0.0015% of the full-scale range.
- Sampling Rate: Minimum of 1 kSPS per channel to capture signals up to 400 Hz according to the Nyquist theorem.
- Type: Successive Approximation Register (SAR) ADC for a balance between speed, power consumption, and resolution.
- Input Range: Adjustable or with a PGA to match the conditioned signals from pressure and temperature sensors.

Data Interface:
- Communication Protocol: Universal communication protocol such as USB or Ethernet for compatibility with a broad range of computer systems and networks.
- Power Supply: Typically 5V or 3.3V DC, in line with common digital circuit requirements.

Component Selection and Calculation Details:
1. Pressure Sensor Amplification Calculation:
   Gain needed = ADC Input Range / Sensor Output Range
               = 5V / 1mV
               = 5000

   Assuming two-stage amplification, each stage needs a gain of 70, for a total combined gain of 4900.

2. Temperature Sensor Amplification Calculation:
   Gain needed = ADC Input Range / Linearized Sensor Output Range
               = 5V / 100mV
               = 50

   This can be achieved with a single amplification stage.

3. Anti-Aliasing Filter Design:
   Order: 4th-order
   Cutoff Frequency: 350 Hz (slightly below the Nyquist frequency of 400 Hz)
   Topology: Butterworth for maximally flat passband response

Overall, the system is designed to amplify, filter, and digitize the signals from the pressure and temperature sensors, ensuring signals are accurately captured for analysis by a computer. The design includes provisions for a flexible and robust Data Interface capable of adapting to various user needs and scenarios. All component values and system parameters are chosen based on typical industrial application standards and the project's specifications to ensure compatibility and reliable operation.