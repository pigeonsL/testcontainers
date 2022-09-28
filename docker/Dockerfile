#Dockerfile
#安装了一个轻量级的模拟器ros-melodic-ros-tutorials
FROM osrf/ros:melodic-desktop-full
RUN apt-get update && apt-get install -y \
    ros-melodic-ros-tutorials \
    ros-melodic-common-tutorials \
    && rm -rf /var/lib/apt/lists/