Pendulum Angle Measurement System Design

This summary encompasses the entire solution, including the high-level architecture and detailed implementation for each block of the pendulum angle measurement system.

1. Potentiometer:
   - Type: Single-turn linear potentiometer
   - Model: Bourns 3590S-2-103L 10K Ohm Potentiometer
   - Operational Temperature Range: -10°C to +85°C
   - Power Rating: Using standard potentiometers with 0.5W or 1W is sufficient.
   - Sensitivity: 0.2222V/degree, calculated based on the potentiometer's linear response from -10V at 45 degrees to +10V at 135 degrees.

2. Buffer (Voltage Follower):
   - Configuration: Voltage follower using a rail-to-rail operational amplifier
   - Op-Amp: Analog Devices AD8628 or Texas Instruments OPA2277
    - Gain: Unity (1)
    - Bandwidth: Minimum of 10 kHz
    - Power Supply: Matched to the -10 V to +10 V range used for the potentiometer
    - Protection: Include input protection such as series resistors (1kΩ) and clamp diodes (e.g., 1N4148)

3. Scale and Shift:
   - Gain (Av): 1 (unity gain)
   - Reference Voltage (Vref): 0 V if the DAQ expects 0 V for the 90-degree position or adjusted to meet the DAQ's expected zero reference
   - Op-Amp: LM358 in a non-inverting configuration
   - Resistors: Rf = Ri = 10kΩ (1% tolerance for precision)

4. Band-Stop Filter (Twin-T Notch Filter):
   - Attenuation: At least 20 dB at 50 Hz and 60 Hz
   - Q Factor: Around 10
   - Insertion Loss: Less than 3 dB
   - Component Values for 50 Hz Notch:
     - Capacitors (C1 and C2): 3.18 uF
     - Resistors (R1 and R2): 318 Ohms
     - Resistor (R3): 31.8 kOhms
   - Adjusted Values for 60 Hz Notch:
     - Capacitors (C1 and C2): 2.65 uF
     - Resistors (R1 and R2): 265 Ohms
     - Resistor (R3): 26.5 kOhms

5. Anti-Aliasing Filter (Second-Order Butterworth Low-Pass Filter):
   - Cutoff Frequency: 50 Hz
   - Configuration: Passive filter
   - Resistor (R): 3.2 kOhms
   - Capacitor (C): 1 uF
   - Quality Factor (Q): 1/√2 for a 2-pole Butterworth filter

6. DAQ System (ADC):
   - Type: Successive Approximation Register (SAR) ADC
   - Resolution: 12-bit (4096 discrete values)
   - Sampling Rate: 1000 samples per second (1 kSPS)
   - Input Voltage Range: Configurable to +/-7 V
   - Reference Voltage: Equal to or just above the maximum expected input voltage (7 V)
   - ADC Driver: Required if impedance matching is necessary

Calculation Examples:
- Sensitivity of Potentiometer: The output voltage range is -5V to +5V for pendulum angles of 45 to 135 degrees. So, sensitivity is 10V over a 90-degree range, resulting in 0.2222V/degree.
- Anti-Aliasing Filter Cutoff Frequency: With a desired cutoff frequency (fc) of 50 Hz and choosing a resistor value (R) of 3.2 kOhms, capacitor (C) value is calculated using fc = 1 / (2 * π * √(R * C)), leading to C ≈ 1 uF.

All component values and configurations are given based on standard values and typical requirements for such a measurement system. The exact values should be fine-tuned based on the final environment, DAQ input characteristics, and component availability.