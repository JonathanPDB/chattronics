Score: 6
Explanations: 
The project summary provided addresses several of the requirements listed, but not all of them:

1. The potentiometer is indeed used as a voltage divider, as it provides variable resistance corresponding to the pendulum's position.

2. The voltage applied to the potentiometer is not explicitly mentioned as +/- 10 V. However, the calculation for the attenuator's attenuation factor mentions a 10V input, which implies that the potentiometer might be operating at +/- 10 V. This is not clear enough to confirm compliance without assumption.

3. The architecture includes a voltage divider, a buffer, and a low-pass filter, which is relatively simple and meets the requirement.

4. The DAQ input voltage is centered at 0, with a range of +/- 7V, as specified.

5. The maximum voltage applied to the DAQ is 7V, as ensured by the voltage divider design (which produces a maximum output voltage of about 6.98V).

6. There is a low pass filter mentioned, the Sallen-Key Low-Pass Filter, which will help avoid aliasing.

7. There is a Twin-T Notch Filter designed to attenuate frequencies around 50 and 60 Hz.

8. The requirement for the filter to have a gain of at least -20 dB at 500 Hz is not explicitly mentioned. The cutoff frequency for the low-pass filter is stated to be around 10 Hz, but without knowing the order of the filter or the specifics of the design, we cannot confirm if it meets the -20 dB at 500 Hz requirement.

Therefore, the project meets requirements 1, 3, 4, 5, 6, and 7, but due to lack of explicit information or confirmation, requirements 2 and 8 are not verifiably met.