Water Temperature Measurement System Design

This system employs a Vishay NTCLE100E3 thermistor to measure the temperature of water within a beaker and outputs a voltage readable by a multimeter.

NTC Thermistor Sensor:
- Vishay NTCLE100E3 series NTC thermistor is proposed for temperature sensing.
- Nominal resistance at 25°C (R25) is typically 10 kΩ.
- Temperature measurement range is 10°C to 90°C.

Linearization Resistor:
- A parallel resistor is used to linearize the thermistor response at the midpoint (50°C). The resistance value (R_L) should match the NTC's resistance at 50°C, which needs to be calculated or obtained from the datasheet.
- The resistance at any temperature can be estimated using the Beta parameter equation:
\[ R_{T} = R_{25} \cdot e^{B\left(\frac{1}{T} - \frac{1}{T_{25}}\right)} \]
  where \( T \) is the temperature in Kelvin, \( R_{T} \) is the resistance at temperature \( T \), and \( B \) is the Beta value.

Signal Conditioning:
- A Wheatstone bridge and an operational amplifier circuit are employed to amplify the thermistor signal.
- Assuming linear behavior, the amplifier gain is set to map the thermistor resistance change to a voltage range suitable for the ADC.
- Operational amplifier example: AD623 or equivalent.
- Power supply: Standard 5V, assuming no specific constraints provided.

Anti-Aliasing Filter:
- A second-order Butterworth low-pass filter with a cutoff frequency of 10 Hz is proposed to ensure signal integrity before ADC conversion, assuming a typical ADC sampling rate of 1 kHz.
- The filter provides a flat passband with a 12 dB/octave roll-off beyond the cutoff frequency.

ADC (Analog to Digital Converter):
- A SAR ADC with a 12-bit resolution is suggested, such as Microchip MCP3201, to digitize the conditioned analog signal.
- A sampling rate between 1 to 10 samples per second (1 to 10 Hz) is adequate for capturing temperature changes in the beaker.

Output Voltage Amplifier:
- A non-inverting operational amplifier configuration is recommended to achieve an output voltage range of 0 to 20 Volts.
- Assuming the signal conditioning stage outputs a 0 to 1 Volt signal, an amplifier gain of 20 is required to obtain the 0 to 20 Volts output range.
- Gain calculation: Av = Vout_max/Vin_max = 20V/1V = 20
- Feedback resistor (Rf) calculation assuming Ri = 10kΩ: Rf = Av × Ri - Ri = 19 × 10kΩ = 190kΩ.
- Precision op-amp with rail-to-rail output capability, such as AD8628 or equivalent.

Voltage Output:
- The output voltage range that correlates with the temperature range of 10 to 90 degrees Celsius is 0 to 20 Volts.
- A precision multimeter capable of reading voltages in the 0 to 20 Volts range is used for direct measurement.
- A voltage follower circuit may be included to prevent loading effects.

Calibration:
- The output voltage is calibrated against known temperature points (e.g., ice bath for 0°C, boiling water for 100°C) to confirm accuracy within the 10-90°C range.
- Adjustments to the signal conditioning circuit may be necessary to align with calibration points.

Note: The provided details are based on general assumptions and typical values. The actual component values, system parameters, and calibrations may vary and should be adjusted according to the specific characteristics of the chosen thermistor and system requirements.