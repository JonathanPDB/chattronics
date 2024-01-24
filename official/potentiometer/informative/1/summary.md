Pendulum Angle Measurement System Design

High-Level Architecture:
The system begins with a potentiometer that measures the pendulum's angle. The varying resistance alters the voltage, which is buffered and then scaled down by a voltage divider. The signal passes through low-pass and notch filters to attenuate unwanted frequencies and is then level-shifted to match the DAQ's input range. Finally, the DAQ samples the signal and converts it into a digital format.

Potentiometer:
- Type: Precision conductive plastic potentiometer
- Model Suggestion: Bourns 3590 Precision Potentiometer
- Value: 10 kOhms
- Operating Range: Based on standard conditions (-10°C to +70°C)
- Lifetime: At least 1 million cycles
- Sensitivity: Should provide a linear output with the pendulum's angle

Buffer:
- Function: Impedance matching, unity gain
- Op-Amp Selection: LM358 or equivalent, considering a dual-supply operation of +/- 9V to +/- 12V

Voltage Divider:
- Purpose: Scale the potentiometer's voltage range (-10V to +10V) down to within the DAQ's range (+/-7V)
- Resistors: R1 = 3 kOhms, R2 = 7 kOhms (1% tolerance or better)

Low-Pass Filter:
- Topology: Second-order Sallen-Key low-pass filter
- Cutoff Frequency (f_c): 10 Hz
- Components: Two 100 nF capacitors, two resistors approximating 159.15 kOhms
- Damping: Butterworth response, damping ratio (ζ) of 1/sqrt(2)

Notch Filters:
- Frequency: One notch filter at 50 Hz, another at 60 Hz
- Topology: Twin-T or multiple feedback active filter design
- Q Factor: 10-30 (for a narrow bandwidth)
- Attenuation: >20 dB at the notch frequencies

Attenuator/Level Shifter:
- Purpose: Further scale the voltage if necessary and adjust DC offset
- Configuration: Resistive voltage divider using precision resistors

DAQ:
- ADC Type: 12-bit SAR ADC
- Sampling Rate: At least 2000 samples per second
- SNR: Minimum 60 dB
- Input Configuration: Single-ended or differential based on environmental noise
- Communication Protocol: SPI or I2C

Overall System Characteristics:
- The system follows a linear signal chain with each block serving a distinct purpose to condition the signal for accurate digital conversion.
- Component values should be fine-tuned during implementation to account for real-world variance.
- The design allows for flexibility in the choice of components and topologies based on the requirements and constraints of the specific application.