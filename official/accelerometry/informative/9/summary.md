Low-Frequency Vibration Measurement Device Design

SENSOR SELECTION:
- A piezoelectric accelerometer is chosen for vibration measurement with a sensitivity of 100 pC/g. The device is intended to measure vibrations up to a frequency of 2 Hz.

CHARGE AMPLIFIER DESIGN:
- With a sensitivity of 100 pC/g and a peak acceleration of 0.05 g, the expected charge output is 5 pC peak, resulting in 10 pC peak to peak.
- Desired voltage output is 1 V peak-to-peak.
- Feedback capacitor (Cf) calculated as 10 pC / 0.5 V = 20 pF for a 1 V peak-to-peak output.
- Feedback resistor (Rf) for the charge amplifier selected based on a low-frequency cutoff, with Rf calculated for a cutoff frequency just above the signal frequency of interest (2 Hz).
- A more realistic value for Rf, such as 10 MΩ, could be chosen, leading to a recalculated Cf for practical considerations of 3.18 nF.

LOW-PASS FILTER DESIGN:
- A first-order passive RC low-pass filter is proposed with a cutoff frequency around 5 Hz to minimize signal attenuation and phase shift at the target frequency of 2 Hz.
- The calculated component values for the low-pass filter are C = 10 nF and R ≈ 3.3 kΩ (using a standard resistor value).

BUFFER STAGE:
- A unity-gain operational amplifier configuration is employed for the buffer stage to ensure signal integrity without loading the preceding filter stage.
- Op-amp selection includes general-purpose, low-offset voltage op-amps, such as the OPA344 or AD8605, suitable for low-voltage, low-frequency applications.
- The op-amp is chosen with a decent gain-bandwidth product, low-offset voltage, CMOS technology for low power consumption, and rail-to-rail input/output capability.

OUTPUT STAGE:
- The output is designed to maintain a low impedance capable of driving a variety of loads without compromising the signal.
- Provision for a standard BNC or SMA output connector is included for easy connection to measurement and monitoring equipment.
- Output protection against over-voltage and ESD is recommended using series resistors, clamping diodes, or transient voltage suppression diodes.
- A variable gain stage and a DC offset adjustment capability are suggested to adapt the output level to match the requirements of various interfacing equipment.
- For future upgrades, provisions for an Analog-to-Digital Converter (ADC) with a suitable sampling rate and resolution are to be included.