Pressure and Temperature Monitoring System Design Summary

This summary compiles the details of an analog instrumentation system designed to monitor pressure and temperature at eight points on a machine, with the goal of digital analysis.

Pressure Sensor Array:
- Sensor Model: Industrial grade strain-gauge pressure sensor with a range of 0 to 1000 psi for versatility.
- Sensitivity: Typically around 2 mV/V/psi, scaled to a maximum output of 1 mV.
- Resolution and Accuracy: 0.1 psi for a 1000 psi range, exceeding 1% accuracy.
- Operating Temperature Range: -40 to 85°C.
- Output Signal: mV/V relative to a given excitation voltage.
- Environmental Resistance: IP65 or higher for protection against dust, water, and industrial fluids.
- Mounting: Threaded process connections for versatility.
- Electrical Connection: Shielded cable or robust connectors for industrial settings.
- Example Sensor: Honeywell Model 317 or similar.

Temperature Sensor Array:
- Sensor Model: MLX90614ESF-BAA Infrared Temperature Sensor by Melexis, offering non-contact measurement with digital output.
- Temperature Range: -70°C to +380°C.
- Output: Digital via I2C, factory-calibrated.
- Response Time: Approximately 80 milliseconds.
- Environmental Conditions: TO-39 package for resilience, may require additional protection based on application.
- Sensitivity and Resolution: Factory-set resolution of 0.02°C.
- Power Requirements: Operates at 3.3V or 5V with 1.5mA current consumption.

Instrumentation Amplifier_P:
- Gain: Set to 5000 to match the pressure sensors' 1 mV full-scale output to a typical 5V ADC input.
- Op-Amp: Analog Devices AD623 or similar, with adjustable gain.
- Gain Setting Resistors: Precision resistors with 0.1% tolerance for accurate gain setting.

Instrumentation Amplifier_T:
- Gain: Set to magnify the 100 mV full-scale output of the temperature sensors to the ADC's input range.
- Op-Amp: Precision op-amps with a low offset voltage, such as the OP07 or LT1057.
- Configuration: Three-op-amp instrumentation amplifier topology selected for high input impedance and CMRR.

Low-Pass_Filter_P:
- Type: Active filter topology, Butterworth response for flat passband.
- Order: 4th-order filter recommended for -24 dB/octave roll-off.
- Cutoff Frequency: Set at 450 Hz to cover the signal range of interest without significant attenuation.

Low-Pass_Filter_T:
- Topology: Sallen-Key active low-pass filter.
- Order: 4th-order Butterworth response.
- Cutoff Frequency: Set at 450 Hz to minimize signal distortion.

Gain_Stage_P:
- Gain: Fixed gain of 5000 to scale the pressure sensor's output.
- Op-Amp: Precision op-amp with low noise and offset voltage.
- Feedback Resistors: Precision resistors with a tolerance of 0.1% or better.

Gain_Stage_T:
- Gain: Approximately 42 for a 0-5 V ADC input range with a 100 mV sensor output.
- Op-Amp: Low-noise, precision op-amps like the AD8628 or OPA2188.

Anti_Aliasing_Filter:
- Type: 4th-order active Butterworth low-pass filter.
- Cutoff Frequency: Set at 450 Hz.
- Op-Amps: Low noise, rail-to-rail operational amplifiers.

Multiplexer:
- Model: ADG1606 from Analog Devices or equivalent.
- Capability: 16:1 multiplexing with SPI control, low on-resistance, and low charge injection.
- Control Interface: SPI for ease of use with microcontrollers.

ADC:
- Type: 16-bit SAR ADC, with a sampling rate of at least 1 kHz per channel (8 kHz aggregate).
- Interface: SPI for high-speed data transfer.
- Operating Temperature: Industrial temperature range of -40 to 85°C.
- Voltage Reference: Precision voltage reference for accuracy over temperature variance.

Linearization_T:
- Method: Logarithmic amplifier using an op-amp and diode for sensors with exponential response.
- Op-Amp: Precision op-amp, e.g., AD829 or LT1057.
- Diode: Silicon diode like 1N4148 or temperature-compensated version for log applications.
- Second Stage: Inverting amplifier for slope adjustment, using the same op-amp type.

The provided values and models are suggested based on general-purpose industrial applications and may need to be adapted to the specific sensors and conditions of the user's project. All the suggested components should be carefully evaluated, and their datasheets reviewed for compatibility with the user's system.