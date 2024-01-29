PENDULUM ANGLE MEASUREMENT SYSTEM DESIGN

POTENTIOMETER SELECTION
- Type: Rotary potentiometer
- Model Suggestion: Bourns 3540S precision potentiometer
- Total Resistance: 10 kOhms
- Resolution: Better than 9 degrees per step
- Sensitivity: Approximately 55.6 Ohms/degree
- Operating Temperature Range: -40°C to +125°C (if sealed, based on environmental conditions)
- Rotational Life: Greater than 1 million cycles
- Shaft Material: Metal
- Mounting: Servo mount style

BUFFER STAGE DESIGN
- Op-amp: General-purpose precision op-amp such as TL071
- Input Impedance: > 1 MΩ
- Output Impedance: < 100 Ω
- Bandwidth: > 3 MHz (TL071)
- Slew Rate: > 13 V/µs (TL071)
- Power Supply: Powered by a dual supply voltage of +/- 10 V
- Bypass Capacitors: 0.1 µF and 10 µF across each supply rail to ground

FILTER DESIGN
- Topology: Active Twin-T Notch Filter
- Attenuation: Minimum 40 dB at 50 Hz and 60 Hz
- Resistors & Capacitors: Calculated to create the 50 Hz and 60 Hz notch frequencies
- Op-amp: Low-noise op-amp with sufficient bandwidth, such as the TL071
- Passband Ripple: Less than 0.5 dB
- Phase Shift: Minimal in the passband

AMPLIFIER DESIGN
- Topology: Non-inverting amplifier
- Gain: 0.7 (to map potentiometer maximum from 10V to DAQ maximum of 7V)
- Feedback Resistor (Rf): 3k Ohms
- Ground Resistor (Rg): 10k Ohms
- Power Supply: Single 15V for the op-amp

PROTECTION CIRCUIT DESIGN
- TVS Diode: SMBJ7.5CA (bidirectional, surface mount, standoff voltage of 7.5 V)
- Series Resistor: 100 Ohms, 1/2 Watt (limits current to 70 mA at 7 V)

ANGLE MEASUREMENT ANALOG-TO-DIGITAL CONVERSION (ADC)
- Type of ADC: Successive Approximation Register (SAR)
- Resolution: 12 bits or higher
- Sampling Rate: 1000 samples per second
- Input Voltage Range: Configurable to accept +/- 7V
- Anti-Aliasing Filter: Low-pass with cutoff frequency just below 500 Hz
- Interface: Standard DAQ communication protocols
- Accuracy: ±1 least significant bit (LSB)

CALCULATIONS & DESIGN NOTES:
- The potentiometer's sensitivity is based on a linear relationship between the angle and resistance, with a full-scale range of 180 degrees equating to the potentiometer's maximum resistance value of 10 kOhms.
- The buffer stage's op-amp is chosen for its high input impedance, ensuring that the potentiometer is not loaded, and its low output impedance, allowing it to drive the filter stage effectively.
- The gain of the amplifier stage is calculated as follows:
  Gain (G) = 1 + (Rf/Rg)
  To achieve G = 0.7, the ratio of Rf to Rg is set to -0.3, which can be achieved with Rf = 3k Ohms and Rg = 10k Ohms.
- The response time of the protection circuit is designed to be in the order of microseconds to protect against high-frequency transients, considering the DAQ's sampling rate.
- The ADC's resolution ensures that the quantization error is minimized, allowing for precise measurement of the pendulum angle.
- Overall, components and design strategies were chosen based on best engineering practices to achieve a reliable and accurate measurement system for the specified application.