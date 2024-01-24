Pendulum Angle Measurement System Design

The project involves designing an analog system to measure the angle of a pendulum using a 10 kOhm linear potentiometer with a power source varying between -10 and 10 Volts. The DAQ system samples at 1000 samples per second and accepts a maximum input voltage of +/- 7 V.

Potentiometer:
- Model: Bourns 3590S-2-103L or equivalent.
- Characteristics: 10 kOhm linear taper, 2 W power rating at 70°C, operating temperature range of -55°C to +125°C.

Offset Adjustment:
- Configuration: Differential Amplifier using an operational amplifier (Op-Amp).
- Gain (Av): Set to 7/10 to match the potentiometer output to DAQ input range.
- Resistor Values: Assuming R1 = 10 kOhms, R2 should be 7 kOhms.

Amplifier:
- Configuration: Non-inverting operational amplifier.
- Gain (A): Approximate gain of 1.4 to scale the potentiometer voltage to the DAQ's input range.
- Resistor Values: R1 = 10 kOhms, R2 = 3.9 kOhms (chosen from standard values for a gain slightly less than 1.4).

Filter_50Hz and Filter_60Hz (Assumed similar requirements for both filters):
- Type: Active Band-Stop (Notch) Filter.
- Center Frequency (f0): 50 Hz and 60 Hz respectively.
- Quality Factor (Q): 10 for 50 Hz filter, 30 for 60 Hz filter.
- Attenuation Level: At least -20 dB for 50 Hz and -30 dB for 60 Hz at their respective center frequencies.
- Op-Amp: TL072 or similar low-noise operational amplifier.

Anti-Aliasing Filter:
- Type: Active Low-Pass Sallen-Key Filter.
- Order: Second-order.
- Cut-off Frequency: 150 Hz.
- Op-Amp: TL072 or similar.
- Quality Factor (Q): 0.707 for a Butterworth response.
- Resistors/Capacitors: Values calculated based on standard Sallen-Key formulas to achieve the desired cut-off frequency and response.

Overall Architecture:
1. The Potentiometer acts as a transducer, measuring the pendulum's angle and converting it into a voltage signal.
2. The Offset Adjustment block scales the potentiometer's output to fit the DAQ's input voltage range.
3. The Amplifier block adjusts the signal amplitude, ensuring it does not exceed the DAQ's maximum voltage input.
4. The Filter_50Hz and Filter_60Hz blocks remove power line noise from the signal.
5. The Anti-Aliasing Filter prevents high-frequency noise from creating artifacts in the digital signal post-ADC conversion.
6. The DAQ samples the conditioned analog signal at 1000 samples per second, converting it to a digital representation for analysis.