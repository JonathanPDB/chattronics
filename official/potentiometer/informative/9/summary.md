Pendulum Angle Measurement System Design Summary

This summary presents a comprehensive solution for the pendulum angle measurement system, which includes the potentiometer, buffer, anti-alias filter, voltage divider, and DAQ system. 

Potentiometer:
- Type: Linear
- Resistance: 10 kOhms
- Suggested Model: Bourns 3590S-2-103L Wirewound Potentiometer
- Operating Range: Expected to be from -3.33V to 3.33V for the angular range of 45 to 135 degrees

Buffer Amplifier:
- Topology: Non-inverting Unity Gain Buffer
- Gain: 1 (Unity Gain)
- Power Supply Range: -10V to +10V
- Suggested Op-Amp: TL071 or TL081 (standard), OPA340 or MCP602 (rail-to-rail)

Anti-Alias Filter:
- Type: 2nd-order Active Low-pass Filter (Butterworth)
- Cutoff Frequency (fc): 100 Hz
- Roll-off: 12 dB/octave
- Attenuation at 50/60 Hz: >20 dB (to be confirmed by user)
- Op-Amp: TL072 (dual supply) or TLV2372 (single supply)
- Resistors/Capacitors: Calculated based on fc and selected op-amp

Voltage Divider:
- Configuration: Adjustable with a trimpot for fine-tuning
- Resistors: R1 = 10 kOhm (fixed), R2 = 10 kOhm trimpot
- Purpose: Scale the signal to DAQ's input voltage range and set the midpoint at 0 V for the 90-degree steady position

Data Acquisition System (DAQ):
- Resolution: At least 10 bits (1024 levels) for sufficient angular resolution
- Sampling Rate: 1000 samples per second or higher
- Input Voltage Range: Bipolar +/- 7 V or equivalent unipolar range

Calculations and Considerations:
1. Anti-Alias Filter: The cutoff frequency was chosen to be 100 Hz, below the Nyquist frequency (500 Hz) and above the highest expected pendulum frequency. The filter order was selected as second-order for simplicity, and standard component values will be utilized for implementation.

2. Voltage Divider: Given that the desired voltage at the DAQ's input corresponding to the pendulum's steady position at 90 degrees is 0 V, R1 and R2 were chosen to allow for calibration and scalability. The divider will scale down the maximum potentiometer output voltage to fit the DAQ's input voltage range.

3. Buffer: The non-inverting unity gain buffer was chosen to maintain signal integrity without amplification. The op-amp selection was based on the operating voltage range of the buffer and whether rail-to-rail capabilities are necessary.

4. Potentiometer: The selected potentiometer is a precision rotary potentiometer with linear taper and a mechanical angle sufficient to cover the pendulum's angular range. It is chosen based on general-purpose requirements.

5. DAQ: The specifications for the ADC within the DAQ were chosen to ensure compatibility with the signal conditioning stages and to maintain the overall system resolution and sampling rate requirements.

This design provides a reliable system for measuring the pendulum angle with the necessary signal conditioning to interface with a DAQ system that cannot be altered or tuned, while taking into account the angular range and steady-state conditions.