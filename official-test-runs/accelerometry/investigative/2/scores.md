Score: 3
Explanations: 
The project summary covers the following requirements:

1. The feedback resistance (Rf) of 150 kΩ and the feedback capacitance (Cf) of 2 nF give a product of Rf * Cf = 300 μs. The required value is approximately 2/pi (which is about 0.637), and 300 μs is not within 10% of this value. Therefore, this requirement is not met.

2. The gain of the charge amplifier is not explicitly stated, so we cannot calculate G x 1.61p / Cf to check if it is roughly around 1. Therefore, this requirement is not met.

3. The max charge (Q_max) is given as 1000 pC, which is much higher than the required peak-to-peak charge output of 161 pC. This requirement is not met.

4. Bias current provision is not mentioned in the summary. Therefore, this requirement is not met.

5. The bias current is not provided, and thus we cannot verify if 0.01 divided by the feedback resistance equals the bias current. This requirement is not met.

6. The project uses a charge amplifier to condition the piezoelectric sensor. This requirement is met.

7. The design is optimized for capturing frequencies below 2 Hz, with a cutoff frequency of 0.1 Hz on the charge amplifier and 10 Hz on the low-pass filter. While detailed calculations are not provided, the given frequencies indicate that the design is likely optimized for the required input oscillation. This requirement is met.

8. The desired output voltage is stated as 1 V peak-to-peak. This requirement is met.

9. The offset voltage is not mentioned in the summary, so we cannot determine if it is kept below 10 mV. Therefore, this requirement is not met.

In summary, requirements 6, 7, and 8 are met, resulting in a score of 3.