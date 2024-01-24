Pendulum Angle Measurement System Design

This summary provides an overview of the proposed architecture for an analog pendulum angle measurement system, complete with component values and design considerations.

Potentiometer:
- Type: Linear Taper Potentiometer
- Model Suggestion: Bourns 3590S Precision Potentiometer
- Resistance: 10 kOhms
- Voltage Span: -10V to +10V
- Temperature Range: -55 °C to +125 °C
- Sensitivity: 20V/90° = 0.222V/°
- Resolution: Assuming a 12-bit DAQ, 14V/4096 ≈ 3.4mV per step, equivalent to 0.0153 degrees per step (3.4mV / 0.222V/°)

Amplifier:
- Topology: Non-inverting operational amplifier
- Op-Amp Selection: Analog Devices AD8605 or similar low-offset voltage op-amp
- Gain (Av): 1.4 (7V out for 5V in)
- Resistors: R1 = 10 kOhms, R2 = 4 kOhms
- Power Supply: ±15V

Filter50Hz:
- Topology: Multiple Feedback (MFB) Active Notch Filter
- Notch Frequency (f₀): 50 Hz
- Quality Factor (Q): 15
- Attenuation at 50 Hz: 20 dB
- Resistors: R1 = R2 = 330 kOhms (chosen for Q and f₀)
- Capacitors: C = 9.65 nF (adjusted from 10 nF for precise f₀)
- Operational Amplifier: TL072 or LM358

Filter60Hz:
- Component values and design approach similar to the Filter50Hz block but tuned for 60 Hz.

Anti-Aliasing Filter:
- Topology: Sallen-Key Butterworth Low-Pass Filter
- Order: 2nd order
- Cutoff Frequency (fc): 150 Hz (to ensure a safety margin below the Nyquist frequency of 500 Hz)
- Resistors: R1 = R2 = 10 kOhms (for a given C value)
- Capacitors: C ≈ 106 nF (calculated for the chosen R and fc)

Data Acquisition System (DAQ):
- ADC Resolution: 12-bit
- Sampling Rate: 2 ksps (kilosamples per second)
- Integrated Sample and Hold: Required
- Interface: SPI or parallel (to be compatible with the DAQ system)
- Power Consumption: Low to moderate (standard ADC)

All component values provided are initial estimates and may need to be fine-tuned during the prototyping phase. Additionally, mechanical considerations such as the mounting of the potentiometer and the construction of the wooden structure are to be determined based on the specific application requirements.