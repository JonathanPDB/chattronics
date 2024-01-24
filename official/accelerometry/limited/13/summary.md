Portable Low-Frequency Vibration Measurement Device

The device architecture is designed to accurately measure low-frequency vibrations using a piezoelectric accelerometer with a sensitivity of 100 pC/g. The system converts mechanical vibrations into an electrical charge, which is then processed to output a digital signal representing the vibration.

1. Piezoelectric Accelerometer (PiezoSensor):
   - A suitable accelerometer is the Endevco model 2271A.
   - Sensitivity: 100 pC/g
   - Estimated max acceleration from 5 cm peak oscillation at 2 Hz: 0.79g
   - Charge output at peak acceleration: 79 pC
   - Operating Temperature Range: -40°C to +85°C

2. Charge Amplifier:
   - Converts the charge from the accelerometer to a voltage signal.
   - Feedback Capacitor (Cf): 1 nF for stability.
   - Feedback Resistor (Rf): 10 MΩ to set low-frequency cutoff at 0.25 Hz.
   - Operational Amplifier: AD8628 for ultra-low offset voltage and bias current.
   - Additional gain stage needed with a gain of approximately 12.66 (for example, using a non-inverting amplifier configuration with gain set by resistors Rg1 = 1 kΩ and Rg2 = 12 kΩ).

3. High-Pass Filter (HighPassFilter):
   - Sallen-Key High-Pass Filter topology.
   - Second-order Butterworth response for maximally flat passband.
   - Cutoff Frequency (f_c): 0.25 Hz to be 3 dB down at the lowest frequency of interest.
   - Component values: C1 = C2 = 10 μF and R1 = R2 = 62 kΩ (standard value close to calculated 63.66 kΩ).

4. Amplifier (Amplifier):
   - Instrumentation amplifier topology selected for precision.
   - Required Gain (Av): 100 to amplify charge amplifier output to 1 Vpp.
   - Power supply: ±15 V chosen for a wide dynamic range.
   - Operational amplifier choice: AD620 or similar for low noise and precision.

5. Anti-Aliasing Filter (AntiAliasingFilter):
   - Butterworth Low-Pass Filter with a cutoff frequency slightly higher than twice the highest frequency of interest (2 Hz), set to 5 Hz.
   - Second-Order Filter to provide sufficient attenuation for frequencies above 10 Hz.
   - Component values: R = 10 kΩ, C = 3.18 μF (chosen for practical stability, actual value subject to availability, and further refinement during prototyping).

6. Analog-to-Digital Converter (ADC):
   - Sigma-Delta (Σ-Δ) ADC topology recommended for high resolution and noise immunity.
   - Resolution: At least 16 bits to maintain fine granularity below the 10 mV offset requirement.
   - Sampling Rate: At least 64 Hz to accurately capture the waveform and allow for digital filtering.
   - Interface: Serial interface such as SPI or I2C for ease of integration with microcontrollers or digital processors.
   - ADC should be operable within a single-supply voltage range of 3.3V to 5V, typical for battery-operated devices.

Notes:
- The detailed component values and specific part numbers should be refined based on prototyping and additional system requirements such as size, power consumption, and environmental constraints.
- The op-amp characteristics, filter component tolerances, and ADC resolution should be considered in a real-world scenario and might require adjustment during prototyping.
- The actual mass of the object to which the accelerometer is attached and environmental conditions will affect the final component selection and system calibration.