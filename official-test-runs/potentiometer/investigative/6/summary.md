Pendulum Angle Measurement System Design

POTENTIOMETER SELECTION
- Component: Precision Potentiometer
- Specifications: 10 kOhm, linear response
- Operating Range: 45 to 135 degrees corresponding to -5 V to 5 V output (with a -10 V to 10 V supply)
- Suggested Model: Bourns 3590S series or similar for precision and reliability
- Considerations: Ensure the potentiometer has a mechanical angle in excess of the required electrical angle and is rated for the environment in which it will be used.

BUFFER STAGE
- Topology: Unity Gain Voltage Follower (Op-Amp Buffer)
- Op-Amp: General-purpose operational amplifier such as TL071 (low noise, suitable supply voltage range)
- Power Supply: Assumed to be shared with DAQ system, typical +/-7V to +/-15V
- Characteristics: High input impedance, low output impedance

LOW-PASS FILTER
- Type: 2nd order active Butterworth filter with a notch filter for 50/60 Hz noise
- Cutoff Frequency: 100 Hz
- Notch Filter Frequencies: 50 Hz and 60 Hz
- Operational Amplifiers: Selected for adequate bandwidth and low noise characteristics
- Components: Resistors and capacitors chosen for tight tolerances and stability

GAIN STAGE
- Topology: Non-inverting operational amplifier configuration
- Gain Calculation: (DAQ input range) / (Potentiometer output range) = (7 V) / (5 V) = 1.4 (or use unity gain for simplicity)
- Op-Amp Selection: General-purpose such as LM358 or TL081 for higher performance
- Power Supply: Assumed to be shared with DAQ system, details unspecified
- Additional Considerations: Include voltage supply stabilization (bypass capacitors) and consider a single supply with rail-to-rail op-amps if necessary.

ANTI-ALIASING FILTER
- Type: 2nd order Butterworth Low-Pass Filter
- Cutoff Frequency: 50 Hz (one-tenth of Nyquist frequency)
- Stopband Attenuation: >40 dB beyond half the sampling rate
- Components: Precision resistors and capacitors, low-noise operational amplifier

DATA ACQUISITION SYSTEM (DAQ)
- ADC Type: Successive Approximation Register (SAR) ADC
- Sampling Rate: 1000 samples per second
- Resolution: 12-bit recommended (higher resolution possible if needed)
- Input Range: Configured to match the DAQ system's ±7 V input range
- Accuracy: Dependent on resolution and quality of the ADC reference voltage
- Communication Interfaces: SPI or I2C for single-channel measurements (assumed)

Additional notes:
- The potentiometer should be securely mounted to minimize mechanical stress and inaccuracies.
- The buffer stage should be designed with consideration for DAQ input impedance, although specific details were not provided.
- The gain and anti-aliasing stages should ensure signal fidelity and prevent aliasing without introducing significant noise or distortion.
- The ADC selection is based on a balance between resolution, sampling rate, and expected signal bandwidth.
- All operational amplifiers should have bypass capacitors near their power supply pins for noise reduction.
- The entire system design accounts for potential noise sources, such as 50/60 Hz, and includes appropriate filtering to mitigate them.