Pendulum Angle Measurement System Design

This summary encompasses the design and specifications for an analog pendulum angle measurement system. The system incorporates a potentiometer for angular position sensing, signal conditioning, attenuation, and interfacing with a DAQ system.

Potentiometer:
- Type: Precision conductive plastic potentiometer.
- Electrical Range: 10 kOhms.
- Mechanical Angle Range: Ideally, a 270-degree mechanical rotation potentiometer, allowing the electrical range to be distributed evenly over the pendulum's 45 to 135-degree mechanical swing.
- Model Suggestion: Bourns Model 3549 or equivalent, with a lifecycle of 3 million cycles.

Buffer (Voltage Follower):
- Purpose: To prevent loading the potentiometer and to maintain signal integrity.
- Op-Amp: OPAx140 or equivalent, chosen for its high input impedance (>1 MΩ), low output impedance (<100 Ω), and low noise (<10 nV/√Hz).
- Power Supply: Single-supply operation within the -10V to +10V range.

Low Pass Filter:
- Topology: Active Filter, Fourth-order Butterworth for maximally flat passband.
- Cutoff Frequency: Assumed 10 Hz to accommodate the slow swinging motion of the pendulum.
- Op-Amps: TL072 or NE5532 for low noise performance.
- Component Values: To be calculated based on the filter topology selected.

Notch Filter (50 Hz and 60 Hz):
- Topology: Active Twin-T or Multiple Feedback (MFB) with an operational amplifier for high Q factor and narrow bandwidth.
- Q Factor: Between 10 and 30, to attenuate noise without affecting adjacent frequencies.
- Notch Depth: Between -20 to -40 dB to suppress power line noise.
- Component Values: To be calculated based on the chosen Q factor and desired notch depth.
- Op-Amp: Low noise and low offset characteristics, such as the TL072 or NE5532.

Attenuator:
- Purpose: To scale the output voltage from the potentiometer to the acceptable input voltage range of the DAQ.
- Attenuation Ratio (AR): Calculated as 0.7 based on the DAQ's maximum input voltage of 7V and potentiometer's voltage range of 10V.
- Resistors: R1 = 3kΩ (approx.), R2 = 7kΩ, chosen to provide the necessary voltage scaling. The resistor values are standardized to reflect the attenuation ratio directly.

DAQ System:
- Sampling Rate: 1000 samples per second.
- Additional Specifications: High enough resolution, ideally at least 12-bit, to distinguish minor changes in angle.
- Input Configuration: To be determined based on the specifics of the DAQ model, which affects the requirement for single-ended or differential inputs, input impedance, and potential need for isolation.