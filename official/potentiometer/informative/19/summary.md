Pendulum Angle Measurement System Design

This design summary outlines the proposed architecture and component selection for a system that calculates the angle of a pendulum using a linear potentiometer. The analog output is conditioned to be compatible with a Data Acquisition System (DAQ) that accepts a ±7V input.

Potentiometer:
A high-precision linear potentiometer is recommended for measuring the pendulum angle. It will provide a variable resistance corresponding to the pendulum's position within a 45 to 135-degree range. A suggested model is the Bourns PTA4543-2015DP-A103, which offers a 10K ohm resistance and is designed for robust, reliable operation.

Buffer (Unity Gain Amplifier):
The buffer stage will consist of a non-inverting unity gain configuration using a general-purpose, rail-to-rail output operational amplifier such as the LM324 or TLV2372. This buffer isolates the potentiometer from the rest of the signal chain, ensuring that the varying impedance of the potentiometer does not affect subsequent stages. The power supply voltage for the op-amp will be selected based on the available values, ideally +/- 15V or +/-12V.

Attenuator (Voltage Divider):
To match the output of the potentiometer to the DAQ's input range, a passive attenuator using a voltage divider is proposed. The required attenuation factor (AF) is calculated as follows:

AF = V_in_max / V_out_max = 10V / 7V ≈ 1.43

Selecting standard resistor values for the voltage divider, we calculate:

R1 = 4.32 kΩ (closest standard value to achieve the desired attenuation)
R2 = 10 kΩ

This configuration gives a real attenuation factor of approximately 1.432, producing a maximum output voltage of about 6.98V when the potentiometer is at 10V, slightly below the ±7V limit.

Low-Pass Filter:
A Twin-T Notch Filter will handle the attenuation of 50 and 60 Hz frequencies, while a Sallen-Key Low-Pass Filter will condition the overall signal. The notch filter will be designed for a quality factor (Q) between 2 to 5, targeting the specific frequencies. For the low-pass filter, a second-order Butterworth configuration is suggested, with a cut-off frequency (f_c) of around 10 Hz, which is sufficiently above the pendulum's maximum swing frequency and provides a flat passband response.

ADC (Analog-to-Digital Converter):
A 12-bit ADC with a sampling rate greater than 1 kHz is recommended for digitizing the conditioned signal. An ADC with a high input impedance (>1 MΩ) and an SPI or I2C digital interface is ideal. It should have an internal or external reference voltage slightly above ±7V for accurate and stable measurements.

The system is designed with general industrial conditions in mind; if specific environmental or power considerations arise, components should be selected to match those requirements.