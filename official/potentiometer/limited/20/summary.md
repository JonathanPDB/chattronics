Pendulum Angle Measurement System Design

The project involves the design and implementation of an analog system to calculate the angle of a pendulum using a potentiometer as the primary sensor and a DAQ for digital data acquisition. The system architecture consists of the following blocks: Potentiometer, Scaling Amplifier, Band-Stop Filter, Anti-Aliasing Filter, and DAQ.

**Potentiometer Sensor Block**
- Type: Precision wirewound or conductive plastic potentiometer.
- Model Recommendation: Bourns 3590S series.
- Value: 10 kOhms linear.
- Voltage Range: -10V to +10V.
- Temperature Range: Standard operation (-55°C to 125°C).
- Lifetime: Mechanical life of about 1 million cycles (extendable with a higher cycle life potentiometer or a non-contacting Hall effect sensor).
- Resolution and Sensitivity: Infinite resolution (analog device), sensitivity approximately 0.222 volts per degree.
- Mounting: Accommodate mechanical linkage to the pendulum and the wooden structure.
- Electrical Connection: Three terminals for power supply (-10V and +10V) and wiper output.

**Scaling Amplifier Block**
- Topology: Non-inverting amplifier using an operational amplifier.
- Op-Amp Recommendation: LM358.
- Gain Setting: 0.7 (to scale the potentiometer's voltage from -10V/+10V to -7V/+7V).
- Power Supply: +15V or ±15V.
- Resistor Values: Ri = 10kΩ, Rf = 4.3kΩ (to achieve the desired gain).
- Gain Calculation: G = 1 + (Rf / Ri), where for G = 0.7, Ri = 10kΩ and Rf = 4.3kΩ.

**Band-Stop Filter Block**
- Topology Suggestion 1: Passive Twin-T Notch Filter for 50 Hz and 60 Hz.
- Resistor Values: 50 Hz - 3.2 kOhms; 60 Hz - 2.7 kOhms.
- Capacitor Values: 50 Hz and 60 Hz - 1 uF.
- Topology Suggestion 2: Active Notch Filter (using op-amps) for higher Q factor.
- Filter Order: Second-order (single-stage) for each frequency.

**Anti-Aliasing Filter Block**
- Topology: Active low-pass filter.
- Filter Order: 4th (Butterworth response for maximally flat passband).
- Cutoff Frequency: 100 Hz.
- Filter Design: Sallen-Key or Multiple Feedback.
- Op-amp Recommendation: LM324 or TL084.
- Component Values: Capacitors: C1 = C2 = 10nF, Resistors: R1 = R2 = 159 kOhms (calculated based on fc = 100Hz and C = 10nF).
- Power Supply: Appropriate for the op-amp, typically between +5V to +15V.

**DAQ (Data Acquisition System) Block**
- Sampling Rate: 1000 samples per second.
- ADC Type: Successive Approximation Register (SAR) ADC.
- ADC Resolution: 12 bits or higher.
- ADC Accuracy: ±1 LSB or better.
- ADC Input Range: Bipolar, at least ±7V.
- Additional Requirement: External anti-aliasing filter due to the lack of information on built-in capabilities.

This analog system design takes into account the project's voltage ranges, frequency attenuation requirements, and the DAQ's specifications to provide an effective solution for accurate pendulum angle measurements.