Analog Temperature Measurement System Design

OVERVIEW
This design entails an analog temperature measurement system using a Vishay NTCLE100E3 NTC thermistor to measure the temperature of water in a beaker, with the output given as a voltage range of 0 to 20 Volts, readable by a multimeter.

NTC THERMISTOR
- Sensor: Vishay NTCLE100E3 NTC thermistor.
- Measurement Range: 10 to 90 degrees Celsius.
- The thermistor should be adequately protected from the water to prevent electrical shorts and corrosion.

BUFFER STAGE
- Operational Amplifier: LM324 or similar high input impedance op-amp.
- Configuration: Voltage follower to prevent loading the thermistor.
- Additional: If needed, an RC low-pass filter for EMI rejection with R = 10 kΩ and C = 1 µF (cutoff frequency approximately 16 Hz).

AMPLIFICATION STAGE
- Instrumentation Amplifier: Choosing between Analog Devices AD623 (Single-Supply, Rail-to-Rail) or Texas Instruments INA128/INA129 (Requires Dual-Supply).
- Gain Settings: Variable, set by a precision potentiometer post-empirical testing to match 0-20V out for the 10-90°C temperature range.

LINEARIZATION STAGE
- Method: A parallel resistor with the NTC thermistor to linearize the response at 50°C.
- Thermistor Beta Value: Assuming approximately 3950K.
- Thermistor Resistance at 25°C (R25): 10kΩ.
- Estimated Thermistor Resistance at 50°C (R50°C): 
  R(50°C) = 10kΩ * exp[3950 * (1/323.15K - 1/298.15K)] ≈ 9259Ω.
- Parallel Resistor (Rp): 9.2kΩ to linearize response around 50°C.
- Power Supply (Vsupply): 24V or as per availability.
- Gain of Non-Inverting Op-Amp Stage: Calculated based on Vsupply and the desired output voltage range. An example is given below.

CALCULATIONS FOR LINEARIZATION STAGE
- Assuming Rmin and Rmax are the thermistor's resistance at 10°C and 90°C, respectively.
- Vmax across the thermistor in the voltage divider would be calculated as:
  Vmax = Vsupply * (Rmax / (Rmax + Rp)) - Vsupply * (Rmin / (Rmin + Rp)).
- With Vsupply selected, the gain required for the op-amp (Av) would be:
  Av = 20V / Vmax.

ADDITIONAL NOTES
- The system should be calibrated using the potentiometer to ensure the output voltage range matches the measured temperature range.
- The multimeter used must have an accuracy of at least 0.1% and have a voltage range setting that can accommodate the 0-20V output.
- The instrumentation amplifier's gain and the linearization circuit should be fine-tuned based on empirical measurements for the desired accuracy.

CONCLUSION
The proposed design provides a linearized voltage output in response to varying temperatures measured by an NTC thermistor. The system utilizes analog signal conditioning techniques, ensuring the output is within a measurable range for standard multimeters.