FROM apache/airflow:2.3.0
USER root
RUN rm -rf  /etc/apt/sources.list.d/*
RUN sed -i "s@http://\(deb\|security\).debian.org@https://mirrors.aliyun.com@g" /etc/apt/sources.list
RUN apt update && apt-get install -y gcc g++ python-dev libsasl2-dev
ARG USERNAME=airflow
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    # && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
USER airflow
RUN pip install apache-airflow-providers-cncf-kubernetes apache-airflow-providers-apache-spark apache-airflow-providers-apache-hive -i https://mirror.baidu.com/pypi/simple

