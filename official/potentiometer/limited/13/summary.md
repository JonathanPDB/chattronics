Pendulum Angle Measurement System Design

The project involves developing a system to calculate the angle of a pendulum using a potentiometer. The system must measure angles from 45 to 135 degrees, with 90 degrees corresponding to the steady position. The potentiometer's output will be conditioned and sent to a Data Acquisition (DAQ) system that samples at 1000 samples per second. The DAQ's maximum accepted input voltage is +/- 7V, and frequencies around 50 and 60 Hz should be attenuated.

1. Potentiometer:
   - A precision linear potentiometer, such as the Bourns 3590S series, is recommended.
   - The potentiometer is rated for 10 kOhms and should handle a power supply range from -10V to +10V.
   - The sensitivity required is 14V span over 90 degrees, equating to 0.1556 V/°.
   - Angular resolution, based on a 12-bit DAQ, is approximately 0.022°.
   - Operating temperature range is selected to be -40°C to 125°C to ensure stability across various conditions.

2. Offset Adjustment:
   - A non-inverting op-amp level shifter circuit is used to center the voltage swing within the DAQ's input voltage range.
   - The op-amp chosen could be the OPA2337, which operates within a supply voltage of -10V to +10V.
   - Resistors in the circuit should have a 0.1% tolerance for precision.

3. Amplifier:
   - A non-inverting amplifier configuration is proposed to utilize the full range of the DAQ input.
   - The gain required is calculated to be 1.4, to scale the potentiometer's output voltage of ±5V for the extreme angles to within the DAQ's acceptable input range of ±7V.
   - Resistor values should be chosen to achieve the desired gain, with precision resistors of 0.1% tolerance.

4. Low-Pass Filter:
   - A 2nd order Butterworth filter is suggested with a cutoff frequency at 250 Hz.
   - Assuming standard capacitor values of 0.1 uF, the resistors would be calculated as approximately 318.3 kΩ, with a standard value of 320 kΩ being used.
   - The filter is designed to maintain signal integrity up to the cutoff frequency while attenuating higher frequencies.

5. Notch Filters:
   - Twin-T Notch Filter designs are chosen for 50 Hz and 60 Hz to attenuate power line interferences.
   - For a 50 Hz notch filter, the component values calculated for resistors and capacitors are based on a desired attenuation and quality factor, with resistor values set to 600 ohms and capacitors to approximately 5.3 μF.
   - A similar approach is used for the 60 Hz notch filter, with component values adjusted for the target frequency.

6. Anti-Aliasing Filter:
   - A 2nd order active low-pass Butterworth filter is used with a cutoff frequency just below 500 Hz to prevent aliasing.
   - The filter's design will involve an operational amplifier with low noise and low offset voltage, such as the AD8628.

7. DAQ:
   - The ADC chosen is a Successive Approximation Register (SAR) type with a resolution of at least 12 bits and a sampling rate equal to or greater than 1000 samples per second.
   - It should have a single-ended input configuration, suitable for connecting to the potentiometer output.
   - The ADC's input buffer will ensure compatibility with the DAQ's input voltage range.

The suggested components and configurations provide a complete signal conditioning path from the potentiometer to the DAQ, ensuring accurate measurement of the pendulum's angle within the specified voltage range and sampling specifications.