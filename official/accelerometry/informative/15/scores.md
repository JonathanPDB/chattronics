Score: 2
Explanations: 
1. The feedback resistance (Rf) is 8 MΩ and the feedback capacitance (Cf) is 200 pF. The product Rf*Cf is 8e6 * 200e-12 = 1.6e-3, which is roughly 0.509 (since 1/(2*pi) is approximately 0.159). This does not meet the requirement of being roughly around 2/pi (allowing up to 10% difference).

2. The charge amplifier gain is not provided directly, but the peak charge output is given as 100 pC at 1 g acceleration, and the corresponding peak voltage is 0.5 V. However, we cannot determine if the gain multiplied by 1.61E-12 and divided by Cf is roughly around 1 without the explicit gain value.

3. The peak-to-peak charge output is mentioned as 100 pC at 1 g acceleration. This does not meet the requirement of being around 161 pC.

4. The bias current is not provided explicitly in the summary.

5. 0.01 divided by the feedback resistance (Rf) gives 0.01 / 8e6 = 1.25e-9 A, which is the bias current according to the requirement. However, since the bias current is not explicitly provided, we cannot confirm this requirement.

6. The project does use a charge amplifier to condition the piezoelectric sensor.

7. The project is optimized for an input oscillation of up to 2 Hz, which is mentioned in the context of the low-frequency cutoff of the amplifier. The calculations backing this up are not provided, but the design intent is clear.

8. The output voltage is not explicitly stated to be 1 V peak to peak.

9. The offset voltage is not mentioned; therefore, we cannot confirm if it is kept below 10 mV.

From the provided summary, it is clear that the project uses a charge amplifier to condition the piezoelectric sensor and is optimized for an input oscillation of up to 2 Hz. However, the other requirements are either not met or cannot be confirmed with the provided information.