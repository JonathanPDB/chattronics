Design and Implementation of an Analog Instrumentation System for Machine Monitoring

Pressure Sensors:
- Industrial-grade strain-gauge pressure sensors.
- Pressure range and accuracy: Typically 0-100 psi, 0-500 psi, or 0-1000 psi with at least 1% accuracy.
- Sensitivity: Assuming a 10 V excitation voltage, a sensitivity of 0.1 mV/V/psi might be typical.
- Operating Temperature Range: -40°C to 85°C or as required by the application.
- Supply Voltage: Commonly 5 V, 10 V, or 24 V DC.

Instrumentation Amplifiers for Pressure Sensors (x8):
- Gain: Adjustable up to 5000 to match the ADC’s input range.
- Input offset voltage: As low as possible, preferably in the microvolt range.
- CMRR and DMRR: Above 80 dB.
- Supply: Dual ±15 V or as available.
- Suggested models: AD620 for low gain, INA125 or AD8421 for higher gain applications.

Low-pass Filters (x8) for Pressure:
- Active Butterworth Low-Pass Filter topology recommended.
- Second-order (two-pole) filter with a roll-off rate of -12 dB/octave.
- Cutoff frequency: 400 Hz to pass desired frequencies up to 400 Hz and attenuate higher frequencies.
- Components: Resistors R1 = R2 = 10 kΩ, Capacitors C1 = C2 ≈ 1/(2πRfc) with C calculated to achieve the desired cutoff frequency.
- Damping ratio (ζ): 0.707 for a maximally flat passband.

Temperature Sensors:
- Model suggestions: Melexis MLX90614 (digital output), Texas Instruments TMP006 or TMP007 (analog output).
- Temperature range: Typically -70°C to +380°C.
- Sensitivity and resolution: As per sensor specifications, with a resolution fine enough to meet the 1% accuracy requirement.

Nonlinear Amplifiers for Temperature Sensors (x8):
- Amplifier stage with adjustable gain to match the ADC’s input range.
- Linearization mechanism, analog or digital.
- Low-pass filtering with a cutoff frequency above 400 Hz.

Low-pass Filters (x8) for Temperature:
- Active Butterworth or Bessel Low-Pass Filter topology, depending on phase linearity requirements.
- Second-order filter with a cutoff frequency of 500 Hz to prevent aliasing.
- Components: Similar to the pressure filters, adjusted for temperature sensor characteristics.

16-to-1 Analog Mux:
- Fast switching to support a minimum sampling rate of 800 Hz per channel.
- Bandwidth: A few kHz to pass the highest frequencies of interest without attenuation.
- Suggested ICs: ADG1606 or DG406 for high signal integrity and fast switching.

ADC:
- Type: Multiplexed successive approximation register (SAR) ADC.
- Sampling rate: At least 1 kHz per channel.
- Resolution: 12-bit or higher to meet the 1% accuracy requirement.
- Input range: 0-5V or 0-3.3V based on the signal levels post amplification.
- Interface: SPI for higher data rates, compatible with the computer's capabilities.

Computer and Digital Processing:
- Software: LabVIEW or MATLAB for data acquisition and analysis.
- Signal Processing: Digital filtering, signal conditioning, and analysis algorithms.
- Data Storage and Management: Database management if long-term storage is required.
- Visualization: Real-time graphical interfaces with plots and gauges.
- Computational Requirements: Multi-core processor with high clock speeds and at least 8 GB RAM.
- Reliability and Redundancy: Industrial-grade hardware or failover systems for high reliability.

This summary outlines the main components and their specifications required to design and implement the analog instrumentation system for monitoring pressure and temperature in a machine environment. Specific component values and models are suggested based on typical requirements and should be fine-tuned based on further analysis or testing.