Pendulum Angle Measurement System Design

This summary outlines the key components and design considerations for a pendulum angle measurement system using a linear potentiometer and data acquisition (DAQ) system.

Potentiometer Block:
- Type: 10 kOhm linear potentiometer, suggested model could be from the Bourns PTA series.
- Voltage Output: Linearly changes from -5 V to 5 V for 45 to 135-degree pendulum movement.
- Resolution: Assuming 1000 samples per second from the DAQ and 90-degree range, the angular resolution is 0.09 degrees per sample.
- Mounting: Must ensure full range of pendulum motion is within the potentiometer's range.
- Environmental Conditions: Standard industrial range, typically -10°C to +70°C.

Amplifier Block:
- Configuration: Unity-gain buffer amplifier to match impedances and buffer the signal.
- Op-amp Selection: Precision low-offset, low noise, unity-gain stable op-amp (e.g., AD8628).
- Power Supply: At least ±15 V to ensure op-amp can produce voltages close to supply rails.

Low-Pass Filter Block:
- Type: Active second-order Butterworth low-pass filter.
- Cutoff Frequency: 200 Hz to prevent noise and high-frequency interference.
- Components: Operational amplifiers, resistors, and capacitors with 1% tolerance for accurate filter characteristics.

Anti-Aliasing Filter Block:
- Type: Active Butterworth Low-Pass Filter, possibly 4th order if additional attenuation is needed at 50 and 60 Hz.
- Cutoff Frequency: 250 Hz, to be below the Nyquist frequency of 500 Hz given the 1000 samples per second DAQ rate.
- Components: Operational amplifiers with suitable bandwidth and low noise characteristics, resistors, and capacitors with low tolerance values.

Data Acquisition (DAQ) Block:
- Sampling Rate: 1000 samples per second.
- ADC Type: Successive Approximation Register (SAR) ADC is recommended due to its balance between speed, cost, and performance.
- Resolution: Suggested 12-bit resolution for fine granularity in angle measurement, which offers 4096 discrete levels.
- Input Impedance: Should be high to avoid loading the potentiometer circuit, typically from 10 kΩ to 1 MΩ.
- Voltage Reference: Requires an accurate reference with a tolerance within ±0.1% or better.

Overall Considerations:
- Signal integrity is paramount, with careful attention to the linearity, resolution, and noise immunity of all components in the signal path.
- The system must be calibrated to ensure that the voltage output of the potentiometer accurately correlates with the mechanical angle of the pendulum.
- The design should account for environmental factors such as temperature and humidity, ensuring that all components are rated for the expected operating conditions.
- The final system design should be verified through testing to ensure that all components work together cohesively and that the system meets the desired specifications.