Score: 2
Explanations: 
1. The feedback resistance (Rf) is 10 MΩ and the feedback capacitance (Cf) is 62.81 nF. The product Rf*Cf is 10e6 Ω * 62.81e-9 F = 0.6281 s, which is not close to the required 2/pi ≈ 0.6366 s (within 10% tolerance).
2. The charge amplifier gain is not explicitly given, thus we cannot calculate G x 1.61p / Cf to see if it is roughly around 1 (within 10% tolerance).
3. The peak-to-peak charge output is calculated to be around 628.3 pC which is not "around 161 pC" as required.
4. The bias current is not provided explicitly.
5. The feedback resistance (Rf) is 10 MΩ, which would imply a bias current of 0.01 / 10e6 = 1e-9 A or 1 nA. However, without the explicit value of the bias current, we cannot confirm this requirement.
6. The project uses a charge amplifier to condition the signal from the piezoelectric sensor.
7. The project is optimised for an input oscillation of 2 Hz and 5 cm amplitude and has calculations to back it up.
8. The output voltage range for the ADC is 0 to +1 V, which can be interpreted as 1 V peak-to-peak, although it's not explicitly stated that the charge amplifier's output itself is 1 V peak-to-peak.
9. The offset adjustment is mentioned through the use of a parallel trimmer in the feedback resistor, but there is no explicit statement that the offset is kept below 10 mV.

In conclusion, only items 6 and 7 are clearly reached. Item 8 is somewhat implied but not explicitly stated for the charge amplifier's output, and item 9 is partially addressed but not quantitatively confirmed. Items 1, 2, 3, 4, and 5 are either not met or cannot be confirmed with the given information.