Score: 2
Explanations: 
1. The feedback resistance (Rf) is 15 MΩ and the feedback capacitance (Cf) is 50 nF. The product of these two is Rf x Cf = 15e6 Ω x 50e-9 F = 0.75 s. This value should be around 2/pi (approximately 0.637 s), but it is not within 10% of this value (10% of 0.637 is approximately 0.0637, so the range would be 0.5733 to 0.7007 s), hence this requirement is not met.

2. The charge amplifier gain (G) is not provided, so we cannot calculate G x 1.61p / Cf. However, if we had G, we would verify that G x 1.61e-12 / 50e-9 should be roughly around 1. Since G is not provided, this requirement cannot be confirmed as met.

3. The peak-to-peak charge output is not specified in the project summary, so we cannot confirm that it is around 161 pC.

4. The bias current is implied to be provided by the charge amplifier, as the power supply requirements mention a current of 2 to 20 mA. However, the exact bias current value is not specified, so we cannot confirm it directly.

5. 0.01 divided by the feedback resistance (Rf = 15 MΩ) should equal the bias current. 0.01 / 15e6 Ω = 666.67e-12 A (or 666.67 pA), but we don't have the exact bias current value to verify this, so we cannot confirm that this requirement is met.

6. The project uses a charge amplifier to condition the piezoelectric sensor, which is explicitly mentioned.

7. The project is optimized for an input oscillation of 2 Hz and 5 cm amplitude, as mentioned in the Piezoelectric Sensor Block section.

8. The output voltage is not specified to be 1 V peak to peak in the project summary.

9. The offset voltage is not mentioned, so we cannot confirm whether it is kept below 10 mV or not.