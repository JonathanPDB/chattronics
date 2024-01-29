Pendulum Angle Measurement System Design Summary

1. Pendulum_Potentiometer
   - Type: Linear taper potentiometer
   - Resistance: 10 kOhm
   - Mechanical rotation range: At least 270 degrees to cover the pendulum's 45 to 135-degree range
   - Suggested model: Bourns 3590S-2-103L or similar
   - Voltage Sensitivity: 0.037V/degree, calculated based on a 10V range over 270 degrees
   - Operating temperature range: Typically -55°C to +125°C

2. Signal_Buffer
   - Topology: Operational amplifier in a voltage follower configuration
   - Op-amp selection: FET-input or CMOS op-amp, e.g., TL071 or OPAx134 series
   - Gain: Unity (1x)
   - Power Supply: Dual power supply configuration matching the potentiometer power source (e.g., +/-12 V or +/-15 V)

3. Low_Pass_Filter
   - Filter type: Active Butterworth Low-Pass Filter
   - Order: 4th
   - Cutoff Frequency: 5 Hz
   - Components: Operational Amplifiers (e.g., TL072 or LM358), Precision Resistors (1% tolerance), Capacitors (1% tolerance)
   - Attenuation: More than 40 dB at 50 Hz and 60 Hz

4. Voltage_Scaler
   - Scaling factor: 7/10 to convert the -10 to +10 volts power supply to within the DAQ's +/- 7V range
   - Voltage Divider: R1 = 3kOhm, R2 = 7kOhm precision resistors
   - Buffer: A buffer with a FET input for high input impedance

5. Anti_Aliasing_Filter
   - Topology: Second-order active low-pass Butterworth filter
   - Cutoff Frequency: 100 Hz
   - Op-Amp: TL072 or similar for low noise characteristics
   - Components: Precision Resistors and Capacitors (1% tolerance)

6. DAQ_Input (ADC block)
   - Sampling Rate: 1000 samples per second
   - Resolution: Suggested 12-bit or 16-bit depending on the required detail
   - Interface: Assumed to be standard (SPI or I2C)
   - Anti-aliasing: Included to ensure maximum pendulum motion frequency is below 500 Hz

Calculations and Notes:
- For the voltage scaler, using a voltage divider formula Vout = Vin * (R2 / (R1 + R2)), where Vout/Vin should be 7/10, R2 = 7kOhm and R1 = 3kOhm are calculated.
- For the low-pass filter, a 4th order Butterworth filter with a cutoff frequency of 5 Hz is chosen to ensure a flat passband and significant attenuation of 50 Hz and 60 Hz noise.
- The ADC should match or exceed the sampling rate of 1000 samples per second with a suggested resolution of 12 or 16 bits based on precision requirements.

This solution provides a complete architecture for the pendulum angle measurement system, ensuring accurate analog-to-digital conversion of the pendulum's angle while minimizing interference and signal distortion.