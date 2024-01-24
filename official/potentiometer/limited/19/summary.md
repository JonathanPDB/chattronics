Pendulum Angle Measurement System Design

This summary consolidates the design details for the proposed pendulum angle measurement system which utilizes a linear potentiometer to infer the pendulum's angle. The output is conditioned to be compatible with a Data Acquisition System (DAQ) that has a maximum input voltage of ±7V and a sampling rate of 1000 samples per second.

Potentiometer:
- Model: Bourns PTV09A-4020F-B103 or equivalent
- Type: 10 kOhm, linear taper
- Operating Temperature Range: -10°C to +70°C
- Sensitivity: Approximately 0.222 V/degree based on a voltage range of 20V and an angular range of 90 degrees.
- Mechanical Mounting: Panel-mount design with a 6.35mm shaft diameter.

Offset Adjustment:
- The offset adjustment block requires a non-inverting summing amplifier to shift the baseline voltage so that the 90 degrees position corresponds to 0V.
- Op-Amp: AD712 or similar for low offset voltage and wide power supply range.
- Resistors for summing circuit: 
  - R1 and R2 set at 10 kOhm to match the potentiometer's resistance.
  - If Rg is the ground resistor connected to the inverting input and Rf the feedback resistor, then Rf/Rg = 0.7.
  - Choose Rg = 10 kOhm, which sets Rf at approximately 7 kOhm (nearest standard value: 6.8 kOhm).
- Voltage reference for offset (Vref): 2.5V.
- Power Supply: Dual, ±15V to accommodate input and output voltage ranges.

Amplifier:
- An amplifier block is used to scale the voltage signal from the offset adjustment to fit the DAQ's input range.
- Gain (Av): The amplifier gain is set to 0.7, calculated as 7V / 10V.
- Op-Amp: Texas Instruments OPA2277 or equivalent, rail-to-rail output, high CMRR.
- Resistors for adjustable gain: 
  - Ri (input resistor) = 10kΩ.
  - Rf (feedback resistor) = 7kΩ, adjustable using a 10kΩ potentiometer.

Filter (Band-Stop):
- A band-stop filter is designed to attenuate 50 Hz and 60 Hz noise frequencies.
- Center frequency (f0): 55 Hz, designed to attenuate both 50 Hz and 60 Hz.
- Q factor: Between 10 to 30 for sharp attenuation.
- Attenuation: >40 dB at 50 Hz and 60 Hz.
- Passband ripple: Less than 0.5 dB.
- Stopband attenuation: At least 40 dB.
- Order: Second-order filter for each notch, with higher orders if necessary.

Anti-Aliasing Filter (Low-Pass):
- A second-order low-pass filter is implemented as an anti-aliasing filter before the signal is digitized by the DAQ.
- Cutoff Frequency: 450 Hz, to remove frequencies above the Nyquist frequency.
- Damping: Q = 0.707 for a Butterworth response.
- Op-Amp: Low-noise, low-offset voltage op-amp.

Data Acquisition System (DAQ):
- Sampling Rate: 1000 samples per second.
- Input Voltage Range: ±7V.
- Resolution: A 12-bit resolution is recommended for the ADC to offer 4096 discrete levels.
- Maximum Frequency Component: Signals up to 500 Hz should be effectively handled by the ADC.
- Anti-Aliasing Filter: Simple low-pass with a cutoff frequency close to 500 Hz.

The above design provides a robust framework for the pendulum angle measurement system. The final implementation may require tuning of component values and further calibration based on the actual behavior of the system during the commissioning phase.