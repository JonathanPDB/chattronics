Pendulum Angle Measurement System Design

**Potentiometer Selection**
- Type: Linear potentiometer, suitable for mapping the pendulum's 45 to 135 degrees range of motion.
- Suggested Model: Bourns PTA series, PTA4543-2015DP-A103 or equivalent, with 10 kOhm resistance.
- Expected Power Source: -10V to +10V with a corresponding output voltage range.
- Implementation: Mechanical linkage is required to ensure the potentiometer's travel correlates with the pendulum's angular range.

**Buffer Stage**
- Function: Isolation of potentiometer from downstream loading and gain adjustment.
- Design: Non-inverting operational amplifier configuration.
- Gain Setting: A gain of 0.7 is needed to scale down the potentiometer's -10V to +10V output to match the DAQ's +/-7V input range.
- Resistor Values: Based on the gain formula (Av = 1 + (R2/R1)), choose R1 = 10kΩ and R2 = 3kΩ.
- Op-amp Model: Precision op-amps like AD8628 or OPA2277, which have low offset voltage and low noise.
- Bandwidth: At least 1kHz to capture the pendulum's motion without aliasing.

**Low-Pass Filter**
- Purpose: Removal of high-frequency noise and prevention of aliasing.
- Cutoff Frequency: Between 150 Hz and 200 Hz to ensure capturing the pendulum motion.
- Filter Order: First-order or second-order Butterworth filter for flat passband or second-order Bessel filter for linear phase characteristics.
- Topology: Active filter using operational amplifiers for better performance and control.

**Notch Filter**
- Function: Attenuation of 50 Hz and 60 Hz power line noise.
- Depth of Attenuation: Aim for at least 20 dB at notch frequencies.
- Filter Order: First-order Twin-T notch filter for each frequency. Consider higher-order active designs for deeper notches if needed.
- Resistor and Capacitor Values: For a 50 Hz notch, R = 10 kΩ and C = 318.3 nF; for a 60 Hz notch, R = 10 kΩ and C = 265.3 nF. These values should be fine-tuned.

**Data Acquisition (DAQ) System**
- ADC Type: Successive Approximation Register (SAR) ADC for its balance between speed, accuracy, and cost.
- Resolution: A 12-bit resolution to offer a voltage resolution of approximately 3.42 mV per bit across the ±7 V input range.
- Sampling Rate: 1000 samples per second, with capability slightly higher than the minimum requirement.
- Input Impedance: High, to prevent loading effects on the signal conditioning circuitry.
- Anti-Aliasing Filter: Required before the ADC with a cutoff below 500 Hz (Nyquist frequency).

**Overall Considerations**
- The potentiometer's mechanical linkage, buffer stage gain, low-pass filter cutoff, and notch filter values need to be tested and possibly adjusted during prototyping.
- Component choices should focus on stability and precision to ensure the reliability of angle measurements.
- The system must be validated to ensure that the analog signal accurately represents the pendulum's angle and that noise is sufficiently filtered before analog-to-digital conversion.