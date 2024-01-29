Pendulum Angle Measurement System Design

POTENTIOMETER SELECTION
- A 10 kOhm linear potentiometer with a rotational angle of at least 270 degrees to cover a measurement range of 45 to 135 degrees.
- The potentiometer should have a robust mechanical design, a wide operating temperature range, and at least 0.5W power rating.
- It is assumed that the potentiometer will output a voltage range from -10V at 45 degrees to +10V at 135 degrees.

BUFFER AMPLIFICATION STAGE
- A non-inverting operational amplifier (op-amp) configuration with fractional gain (attenuation) is required.
- The necessary gain to adjust the potentiometer's output voltage to the DAQ input range of +/-7V is calculated as 0.7.
- Assuming ideal op-amps, R1 could be 10kΩ and R2 could be 3.3kΩ based on the formula V_out = (1 + (R2/R1)) * V_in.
- The op-amp should have a symmetric power supply, which can be +/-12V or +/-15V.

LOW-PASS FILTER
- An active low-pass Sallen-Key filter design is proposed with a cutoff frequency (fc) of 50 Hz.
- A second-order (two-pole) filter with a -12 dB/octave roll-off beyond the fc is recommended.
- A damping factor (ζ) between 0.5 and 1 (critical damping) is appropriate. For a Butterworth response, ζ = 0.707.
- The operational amplifier for the filter stage should be low-noise, low-offset, and have a sufficient bandwidth, such as the TL072.

NOTCH FILTER
- A Twin-T or Multiple Feedback Band-pass filter topology is suggested for creating notches at 50 Hz and 60 Hz.
- The Q factor for the notch filters should be around 10 to 30 to create a narrow stopband without affecting adjacent frequencies.
- The depth of the notch should be at least 20 dB.
- The power supply voltage levels for the notch filter's operational amplifiers should be compatible with the available +/- 10 Volts.

DAQ SYSTEM
- A minimum sampling rate of 50 Hz is recommended based on the pendulum motion and the DAQ's specified rate of 1000 samples per second.
- A 12-bit ADC like the ADS8319 with a bipolar input range of +/- Vref and SPI or I2C digital output is recommended, which can be set for the desired range.
- The DAQ system should accommodate the buffered potentiometer signal range of +/-7V.

ADDITIONAL CONSIDERATIONS
- All active components require power supply voltages that match the available source and provide sufficient headroom.
- The design assumes common environmental conditions; if extremes are expected, components must be rated accordingly.
- The system architecture comprises four main blocks: Potentiometer, Buffer, Low-Pass Filter, and Notch Filter, which interface with the DAQ system to provide digital angle measurements.
- If more specific requirements are provided later, such as mechanical constraints for the potentiometer or specific noise and bandwidth requirements for the filters, the design can be refined.