Pendulum Angle Measurement System Design

This design comprises several key components organized into functional blocks: the potentiometer, buffer, anti-alias filter, and ADC. Below is the summary of each block with the suggested components and configurations.

Potentiometer:
- Type: Linear rotary potentiometer
- Resistance Range: 10 kOhms
- Rotational Range: Must match or exceed the pendulum's angular range (45 to 135 degrees)
- Model Suggestion: Bourns 3590S-2-103L or similar high-resolution potentiometer
- Power Source: -10 to +10 V, ensuring the output voltage swing matches the DAQ input range
- Mounting: Must be aligned with the pendulum's pivot to avoid mechanical play or misalignment

Buffer (Voltage Follower):
- Function: Impedance matching without amplification (Gain = 1)
- Op-Amp Characteristics: Rail-to-rail output, low offset voltage, low noise
- Power Supply: Assuming single-supply, use standard voltage like +15V or slightly above the highest expected output from the potentiometer
- Op-Amp Recommendations: TL072, LM324, or similar

Anti-Alias Filter:
- Topology: Active Low-Pass Filter, Sallen-Key or Multiple Feedback (MFB)
- Order: 2nd-order or higher depending on required noise attenuation
- Cutoff Frequency: 150 Hz recommended to stay well below Nyquist frequency (500 Hz)
- Filter Type: Butterworth for flat passband response, Chebyshev for steeper roll-off if more critical than passband flatness
- Component Recommendations: Low-noise op-amps like TL072, OPA2134; 1% tolerance metal film resistors; stable capacitors like MLCC or film capacitors
- Special Attention: Attenuate 50 and 60 Hz interference

ADC (Analog-to-Digital Converter):
- Resolution: 12-bit (4096 discrete levels) or higher
- Sampling Rate: 1000 samples per second
- Input Voltage Range: Bipolar ADC accommodating at least +/-7 V, or unipolar with level shifting and scaling
- Input Impedance: >= 100 kOhms
- Digital Interface: SPI or I2C, depending on DAQ/microcontroller interface capabilities

Calculations for each block are based on standard practices and expectations for a pendulum swing measurement application. Specific component values for the filter can be calculated using online tools or filter design software once the final op-amp model and desired cutoff frequency are chosen.

The design should be compatible with a DAQ system that accepts a voltage range of +/- 7V and has a sampling rate of 1000 sps. Environmental conditions and operational temperatures were not specified; if the pendulum will be used in harsh conditions, the selection of components with appropriate temperature ratings or IP ratings will be required.