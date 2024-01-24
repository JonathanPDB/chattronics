Design of a Portable Low-Frequency Vibration Measurement Device

The proposed portable device for measuring low-frequency vibrations consists of several key stages, including a piezoelectric sensor, charge amplifier, low-pass filter, buffer, and an optional analog-to-digital converter (ADC).

Piezoelectric Sensor:
A generic piezoelectric accelerometer with a sensitivity of 100 pC/g is suggested since no specific model was provided. A commonly available option suitable for low-frequency measurements up to 2 Hz is the PCB Piezotronics 352C series. The sensor must be matched with the subsequent charge amplifier to ensure proper signal conversion and amplification.

Charge Amplifier:
The charge amplifier is designed to convert the charge signal from the piezoelectric sensor into a useful voltage signal. Assuming a maximum expected acceleration of 1 g and a sensor sensitivity of 100 pC/g, the amplifier's gain is calculated to output a 1 V peak-to-peak signal as follows:

Q = 100 pC/g at 1 g = 100 pC
Desired V_out = 1 V peak-to-peak
Gain (G) = V_out / Q = 1 V / 100 pC = 10^10 V/C

The feedback capacitor (Cf) and the feedback resistor (Rf) of the charge amplifier are chosen to provide the calculated gain and to set the time constant for the filter. However, using standard resistor values and practical component considerations, a lower Cf and a smaller Rf that yields an acceptable gain might be necessary. Low-noise components should be selected.

Low-pass Filter:
A fourth-order Butterworth low-pass filter is recommended for a clear, flat passband and a steep roll-off beyond the cutoff frequency. The cutoff frequency is set slightly above the highest frequency of interest at 2.5 Hz. The characteristics are as follows:

- Cutoff Frequency (fc): 2.5 Hz
- Filter Order: Fourth-order (24 dB/octave)
- Passband Ripple: 0 dB (maximally flat)
- Output Impedance: Compatible with the buffer stage

Component values for the filter will depend on the chosen op-amp characteristics and must be calculated to achieve the desired cutoff frequency.

Buffer:
A voltage follower (unity-gain buffer) topology using an op-amp is suggested for the buffer stage. This provides a high input impedance and low output impedance without amplification of the signal. A suitable op-amp for this purpose is the Texas Instruments OPAx2333 series or Analog Devices AD8605, which are low noise and have low power consumption. The buffer must accommodate the expected voltage swing from the filter and maintain signal integrity.

Analog-to-Digital Converter (ADC):
An optional ADC block can be included if digital processing or interface is required. A Delta-Sigma ADC with a resolution of at least 16 bits is suggested for high precision. The sampling rate should be at least 10 samples per second (sps), higher than twice the highest frequency component (2 Hz), ensuring accurate digitization. The selected ADC should have high SNR (> 90 dB) and low THD (< -100 dB). Communication interface (SPI or I2C) depends on the digital system requirements.

Each stage of the device must be designed with careful consideration of the signal chain requirements, including impedance matching, noise performance, and power supply management. The final specifications of each stage will need to be fine-tuned based on the detailed requirements and constraints of the application at hand.