Pendulum Angle Measurement System Design

POTENTIOMETER SELECTION:
- Component: Single-turn linear potentiometer (e.g., Bourns PTV09A series)
- Electrical Rotation: >90 degrees to cover pendulum range from 45 to 135 degrees
- Resistance: 10 kOhms
- Voltage: Up to 7 Volts
- Current: Maximum 0.7 mA (calculated from I = V/R = 7V / 10kOhms)

BUFFER AMPLIFIER:
- Topology: Voltage follower (unity gain buffer)
- Op-Amp: Rail-to-rail input/output (e.g., MCP6021)
- Power Supply: +/-15V (if dual supply op-amp) or 10V (if single supply)
- Decoupling: 0.1 uF capacitor near op-amp power supply pins

LOW-PASS FILTER WITH NOTCH (TO REMOVE 50/60 HZ NOISE):
- Notch Filters: Twin-T network for 50 Hz and 60 Hz (switchable or band-stop design)
- Low-Pass Filter: 2nd or 3rd order Butterworth with cutoff at 400 Hz
- Resistors: Precision with 1% tolerance
- Capacitors: Precision with 5% tolerance

GAIN STAGE:
- Topology: Non-inverting amplifier with adjustable gain
- Op-Amp: Precision with low offset (e.g., OPA2277 or AD823)
- Gain: Adjustable up to 1.4 (set with trim potentiometer)
- Power Supply: Based on op-amp requirements

ANTI-ALIASING FILTER:
- Topology: 2nd order active low-pass filter (e.g., Sallen-Key or Multiple Feedback)
- Cutoff Frequency: 250 Hz
- Filter Type: Butterworth for maximally flat passband
- Resistors (R): 10 kOhms
- Capacitors (C): Chosen for cutoff frequency (f_c = 1 / (2 * π * R * C)), approximately 68 nF (nearest standard value)

DATA ACQUISITION SYSTEM (DAQ):
- Sampling Rate: 1000 samples per second
- Suggested ADC: Successive Approximation Register (SAR) ADC
- Resolution: 12 to 16 bits (for precision angle measurement)
- Input Type: To be determined (single-ended or differential)
- Communication Interface: To be determined (SPI, I2C, USB, etc.)

The above components and values are based on the information provided and standard engineering practices. The design can be fine-tuned with further details about environmental conditions, power consumption, and specific signal characteristics.