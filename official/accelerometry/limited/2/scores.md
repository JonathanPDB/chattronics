Score: 3
Explanations: 
1. The feedback resistor (Rf) is calculated to be approximately 63.7 MΩ. The feedback capacitance (Cf) is given as 10 pF. The product Rf * Cf = 63.7 MΩ * 10 pF = 637 MΩ·pF which is not close to the required 2/π (approximately 0.637 MΩ·pF). The value is off by a factor of 1000 and thus does not meet the requirement.
2. The gain from the voltage amplifier block is given as 10, but there is no clear indication of the charge amplifier gain, which is necessary to evaluate this requirement. Since the charge amplifier's output is given as 100 mVpp for an assumed peak charge of 80 pC, we cannot directly infer the gain without additional information. Therefore, this requirement cannot be confirmed as met without further data.
3. The peak charge output calculated from the piezoelectric sensor is 80 pC, which is not close to the 161 pC required.
4. There is no explicit mention of the bias current being provided, so this requirement is not met.
5. Since the bias current is not provided, we cannot confirm that 0.01 divided by the feedback resistance equals the bias current.
6. The project uses a charge amplifier to condition the signal from a piezoelectric sensor, so this requirement is met.
7. The project is optimized for an input oscillation of 2 Hz and 5 cm amplitude with calculations to back it up, so this requirement is met.
8. The output voltage from the voltage amplifier is expected to be 1 Vpp, which is in line with the requirement for the ADC input range, so this requirement is met.
9. There is no information provided about the offset voltage being below 10 mV, other than suggestions for op-amps with low offset voltage. However, without specific measurements or design calculations, we cannot confirm that the offset is kept below 10 mV.

In summary, the project explicitly meets requirements 6, 7, and 8. The other requirements are either not met or cannot be confirmed with the given information.