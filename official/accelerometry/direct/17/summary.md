Portable Low-Frequency Vibration Measurement Device Design

Analog Front-End

1. Sensor Type and Model:
   - Charge mode piezoelectric accelerometer, suggested model PCB Piezotronics 352C33 with a sensitivity of 100 pC/g.
   - Operating temperature range from -40°C to +85°C.
   - Measurement range up to at least 1g considering the peak amplitude of 5 cm at 2 Hz.

2. Charge Amplifier (ChargeAmp):
   - Feedback components: C_f = 10 nF, R_f = 1 MΩ (with additional low-pass filtering to achieve the desired low-frequency response).
   - Gain (G) to convert 6.4 pC peak charge to 1 Vpp output: G = 1 Vpp / 6.4 pC = 156.25 MV/C.
   - JFET input operational amplifier with low input bias current and low noise.

3. Low-Pass Filter (LPF):
   - Active second-order Butterworth low-pass filter with a cutoff frequency (f_c) of 0.25 Hz.
   - Component values: R1 = R2 = 62 kΩ (or closest standard value), C1 = C2 = 10 µF.
   - Op-amp: Low noise op-amp, e.g., AD8605 or equivalent.

Output Stage

1. OutputStage Amplification:
   - Voltage follower (buffer) configuration with unity gain (1) using a rail-to-rail operational amplifier capable of single-supply operation (e.g., Texas Instruments OPA344).
   - Input impedance greater than 1 MΩ.

Digital Processing

1. Anti-Aliasing Filter for ADC:
   - Sallen-Key second-order low-pass filter with a cutoff frequency of 40 Hz.
   - Component values: R1 = R2 = 10 kΩ, C1 = C2 = 39 nF (or closest standard value).

2. ADC Characteristics:
   - Sigma-Delta ADC topology with a resolution of 24-bit and an assumed sampling rate of 100 Hz.
   - Input range to accommodate the 1 Vpp output; high input impedance (10 kΩ to 1 MΩ).

3. Microcontroller:
   - Low-power microcontroller with data storage capabilities (SD card or internal memory).
   - Sampling rate for the microcontroller's ADC: 10 Hz or higher.
   - Real-time analysis capability if required (FFT computations).
   - Communication: USB and/or wireless options like Bluetooth or Wi-Fi.

The design considerations provided are based on general assumptions and standard practices suitable for a portable vibration measurement device intended to measure low-frequency vibrations. Specific component models and values have been suggested based on these assumptions, and further adjustments may be necessary based on detailed project requirements.