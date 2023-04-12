# firewatch-sensor
This repository contains the code for a firewatch module(all components of a sensor put together). The code initializes the co2 sensor, which can also measure humidity and temperature, starts and secures a lorawan connection to the things network, and creates and sends a payload to the things network. The code is in tinygo, which I have had no previous experience in until I started this project.

## Aim of the Project
The goal of this project is to create a prototype for what would be a cluster of sensors in remote areas to detect wildfires. Nowadays there are ever increasing wildfire occurring throughout the world causing billions of dollars in damages. Hopefully this project will be able to notify respective authorities of wildfires and save resources. 

For this firewatch module I am using the following parts:
    - Lora E5 Devboard
        I am using this board since it has a lora RF and mcu chip installed. And for the sake of creating a prototype sensor makes everything easier to have it all in one. It was also the only board I could get my hands on.

    -  SCD40 Co2, humidity, temperature sensor.
        The SCD40 is an amazing co2 sensor and is super accurrate. It actually measures co2 whereas other sensors approximate it from VOC gas concentration. A good co2 sensor is key if the firewatch module is going to be placed outside.
    - GPS






to run use:

make