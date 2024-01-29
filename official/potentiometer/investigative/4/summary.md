Pendulum Angle Measurement System Design

This system is designed to calculate the angle of a pendulum using a linear potentiometer, with the signal processed through various stages before being read by a DAQ system. The DAQ is limited to a +/- 7V input with a sampling rate of 1000 samples per second.

1. Potentiometer:
A Bourns 3590S-2-103L 10k Ohm potentiometer is suggested due to its precision and reliability. The potentiometer's output will vary linearly with the angle of the pendulum within a range of 45 to 135 degrees, with 90 degrees being the neutral position. The 10-turn span of this potentiometer provides a high resolution, and the mechanical limits should be designed to prevent the shaft from rotating beyond the specified angle range.

2. Buffer:
An operational amplifier (op-amp) in a voltage follower configuration acts as the buffer stage. A general-purpose op-amp like the LM324 is suitable, operating from a dual supply of +/- 12V. This configuration provides a high input impedance, which prevents loading of the potentiometer, and a low output impedance suitable for driving the next stage.

3. Low-Pass Filter:
A second-order active low-pass Butterworth filter is implemented to attenuate 50 Hz and 60 Hz noise. A Sallen-Key topology is used with a cutoff frequency of 40 Hz. Resistor and capacitor values will need to be calculated based on the selected op-amp characteristics. The components should be selected to achieve a minimum of 20 dB attenuation at the noise frequencies.

4. Voltage Divider:
To scale the potentiometer's -10V to +10V output range to the DAQ's -7V to +7V range, a resistive voltage divider is employed. The chosen ratio is R1 = 30kΩ and R2 = 70kΩ, providing the necessary scaling while maintaining a high enough impedance to avoid loading effects.

5. DAQ:
While the exact ADC resolution within the DAQ is unknown, the design assumes a resolution sufficient to capture the detail provided by the potentiometer and buffer stage. The ADC must accept the scaled ±7V input range and convert the analog signal at a sampling rate of 1000 samples per second. An external anti-aliasing filter should be implemented before the ADC, with a cutoff frequency just below 500 Hz, to ensure no aliasing effects occur.

Component Summary:
- Potentiometer: Bourns 3590S-2-103L 10k Ohm
- Buffer Op-Amp: LM324 (or equivalent rail-to-rail op-amp)
- Low-Pass Filter Resistors & Capacitors: Values to be calculated for a cutoff of 40 Hz
- Voltage Divider: R1 = 30kΩ, R2 = 70kΩ
- DAQ: Assumed to have suitable ADC resolution and sampling rate

This system design provides a simple, effective approach to measuring the pendulum's angle and conditioning the signal for the DAQ system while ensuring noise rejection and proper scaling.