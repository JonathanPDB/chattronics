Pendulum Angle Measurement System

The Pendulum Angle Measurement System uses a linear potentiometer, signal conditioning circuits including a buffer, low-pass filters, a gain stage (if necessary), and an analog-to-digital converter (ADC) interfaced with a Data Acquisition System (DAQ).

Potentiometer:
- Model: Bourns 3540S or equivalent precision potentiometer.
- Resistance: 10 kOhms.
- Linear taper, suitable for the 45 to 135 degrees pendulum angle range.
- Rotational range: At least 90 degrees.
- Linearity: Within 0.5% to 1% desired for improved accuracy.
- Power supply: Between -10V and +10V.
- Temperature Coefficient: 50 to 100 ppm/°C.

Buffer:
- Op-Amp: Rail-to-rail input/output capable of operating from the -10V to +10V supply.
- Configuration: Voltage follower (unity-gain buffer) to maintain the signal level.
- Example Op-Amp: Texas Instruments OPAx340 or Analog Devices AD8605.
- High input impedance: Typically in the megaohms range, to prevent loading the potentiometer.
- Decoupling Capacitor: 0.1 µF ceramic capacitor close to the Op-Amp's power supply pins to filter out noise.

Low-Pass Filter (for 50 and 60 Hz noise attenuation):
- Filter Type: 2nd-order (12 dB/octave roll-off) or 4th-order (24 dB/octave roll-off) Sallen-Key low-pass filter.
- Cutoff Frequency: Approximately 100 Hz.
- Resistors: R = 1.59 kOhm.
- Capacitors: C = 1 uF.
- Operational Amplifier: A general-purpose op-amp with a gain-bandwidth product greater than 1 MHz.

Anti-Aliasing Filter:
- Filter Type: Active Second-Order Butterworth Low-Pass Filter.
- Cutoff Frequency: 100 Hz.
- Passband Ripple: < 0.1 dB.
- Stopband Attenuation: > 60 dB beyond 500 Hz.
- Component values for a 100 Hz cutoff with 0.01 µF capacitors:
  - Resistors: R = 159.15 kOhm (calculated as R = 1 / (2π × 100 Hz × 0.01 µF)).

Gain Stage:
- Assuming the potentiometer and DAQ voltage ranges match, no gain stage is necessary.
- If a gain stage is needed in the future:
  - Op-Amp: TL071 or a similar general-purpose Op-Amp.
  - Configuration: Non-inverting amplifier with a gain of 1 or adjustable gain using a potentiometer in the feedback loop.

ADC (Analog-to-Digital Converter):
- Topology: Successive Approximation Register (SAR) ADC.
- Resolution: 12 bits (providing 4096 discrete levels for the +/- 7V input range).
- Sampling Rate: 1 kHz (to capture higher-order harmonics).
- Communication Interface: SPI or I2C.

This architecture ensures a robust and accurate measurement of the pendulum's angle by conditioning the signal from the potentiometer and preparing it for digital conversion with minimal noise and distortion.