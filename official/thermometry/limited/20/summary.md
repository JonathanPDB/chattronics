Water Temperature Measurement System Design

The design revolves around an NTC thermistor-based temperature sensing system that outputs a voltage between 0 and 20 Volts corresponding to a temperature range of 10 to 90 degrees Celsius. The entire system avoids digital conversion and is suitable for direct reading through a multimeter.

Thermistor and Linearization:
- Sensor: Vishay NTCLE100E3 NTC thermistor.
- Linearization Resistor: A fixed resistor parallel to the thermistor, chosen based on the resistance at the midpoint (50°C), calculated from the datasheet's beta value (assuming 3950K) and resistance at 25°C (typically 10kΩ).
- Linearization will improve the accuracy around the midpoint temperature.

Wheatstone Bridge:
- The bridge is configured with the thermistor and a linearization resistor forming one leg and precision resistors forming the other.
- Supply Voltage (Vsupply): 5V to minimize self-heating in the thermistor.
- Resistors R1, R2, and R3: 10kΩ (matched to the thermistor's resistance at 50°C).

Instrumentation Amplifier:
- An instrumentation amplifier (like the Analog Devices AD620) is used to amplify the differential voltage from the Wheatstone bridge.
- The amplifier has a specified gain (G) set to map the bridge's voltage range to the desired output range after level shifting.

Variable Gain Stage:
- A precision operational amplifier (like the OPA277) configured with a potentiometer in the feedback loop for adjustable gain.
- Gain (G) is initially estimated and then fine-tuned during calibration.

Level Shifter:
- Assumed power supply voltages for level shifting: ±15V.
- Op-Amp Level Shifter Circuit: A summing amplifier configuration with an op-amp to shift the voltage level.

Output Buffer:
- A unity-gain voltage follower using an operational amplifier (like the LM324) to provide a low impedance drive to the multimeter.
- Important for maintaining the signal integrity and preventing loading the output.

Anti-aliasing Filter (Optional):
- Topology: Third-order active Butterworth low-pass filter.
- Cutoff Frequency: 10 Hz to eliminate high-frequency noise and ensure a clean signal for the multimeter.
- Components: Precision op-amps, metal film resistors, and ceramic capacitors.

Multimeter:
- Measurement device: Any digital multimeter with at least 10 MΩ input impedance and suitable voltage measurement resolution (1 mV or better for precision).
- External ADCs are not used, and the entire signal conditioning is performed in the analog domain.

Self-Heating Consideration:
- The thermistor's excitation current is minimized to keep self-heating below 1%.

Calibration:
- The system should be calibrated against a known temperature reference, adjusting the variable gain to correlate the output voltage to the temperature readings.

Resistor Calculations:
For the level shifter, assuming the instrumentation amplifier outputs a 0-5V range for the temperature span with a ±15V supply:
- Gain (A) of level shifter: 20V / 5V = 4.
- If the non-inverting input is grounded, Rf/Rin = 3 for a non-inverting amplifier configuration.
- If Rf = 30kΩ, Rin = 10kΩ.

Filter Calculations:
For the third-order Butterworth filter:
- Op-amp choice should have a bandwidth exceeding the cutoff frequency. Component values will be determined using standard filter design equations based on the chosen op-amp characteristics.

The proposed solution accounts for all critical aspects of the analog signal chain required to measure the temperature of water in a beaker with a thermistor and to output a voltage readable by a standard multimeter.