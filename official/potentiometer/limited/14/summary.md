PENDULUM ANGLE MEASUREMENT SYSTEM DESIGN

POTENTIOMETER SELECTION:
- Type and Model: Bourns 3590S 10 kOhm linear potentiometer
- Operating Temperature Range: 0 to 40 degrees Celsius
- Power Rating: 0.5 W or greater
- Resolution: Infinite voltage resolution (limited by the ADC and noise level)
- Sensitivity: Approximately 111 ohms per degree
- Mechanical Constraints: Standard 6.35 mm (1/4 inch) shaft diameter for panel mounting

BUFFER STAGE:
- Op-Amp: LM324 or TL074 for buffering with a high input impedance and low output impedance
- Power Supply: Dual power supply of +/-10V
- Configuration: Unity-gain voltage follower with decoupling capacitors (100 nF near power supply pins)

DIFFERENTIAL AMPLIFIER:
- Gain (Av): 1.4 (to map -5V to +5V potentiometer output to -7V to +7V DAQ input)
- Resistor Values: R1 = 10 kOhm, R2 = 4 kOhm (R2 is a parallel combination of standard value resistors)
- Op-Amp: Texas Instruments OPAx340 series or Analog Devices AD8605
- Power Supply: Single-supply or dual-supply based on the system design
- Frequency Response and Temperature Range: As per the environmental conditions and system requirements

BAND-STOP FILTER (NOTCH FILTER):
- Filter Type: Active Twin-T Notch Filter for 50 Hz and 60 Hz
- Attenuation Level: -40 dB
- Quality Factor (Q): 10
- Bandwidth: Approximately 5 Hz around 50 Hz and 60 Hz
- Op-Amp: TL072 or the LM358
- Resistor/Capacitor Values: Calculated based on desired notch frequencies and Quality Factor

ANTI-ALIASING FILTER:
- Filter Type: Second-order Sallen-Key Low-Pass Filter
- Cutoff Frequency: 200 Hz (-3 dB point)
- Op-amp: Texas Instruments OPAx134 series
- Resistors: Metal film resistors with a 1% tolerance
- Capacitors: Ceramic capacitors with a 10% tolerance, in series pairs
- Calculation for Cutoff Frequency: C1 = C2 = 3.9 nF, R1 = R2 = 20.3 kOhms

LEVEL SHIFTER:
- Op-Amp: LM358
- Voltage Divider Resistors: R1 = R2 = 10kΩ to create a 0V reference from -10V to +10V power supply
- Summing Input Resistor: R3 = 10kΩ
- Feedback Resistor: R4 = 10kΩ to maintain unity gain and introduce a DC shift

ADDITIONAL CONSIDERATIONS:
- All op-amps used should offer rail-to-rail input and output swing, low offset voltage, suitable bandwidth, and slew rate.
- Precision resistors for the differential amplifier, notch filter, and anti-aliasing filter should be chosen to maintain system accuracy.
- Decoupling capacitors should be used near op-amp power supply pins to filter out noise.
- Power supply arrangements for the op-amps should be consistent with the overall system design (single-supply or dual-supply).
- The entire system should be designed to handle the environmental conditions in which it will operate.

By implementing the above specifications, the pendulum angle measurement system will provide accurate angle measurements within the specified range and environmental conditions.