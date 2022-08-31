FROM apache/airflow:2.3.0
USER root
RUN rm -rf  /etc/apt/sources.list.d/*
RUN sed -i "s@http://\(deb\|security\).debian.org@https://mirrors.aliyun.com@g" /etc/apt/sources.list
RUN apt update && apt-get install -y gcc g++ python-dev libsasl2-dev
USER airflow
RUN pip install apache-airflow-providers-cncf-kubernetes apache-airflow-providers-apache-spark apache-airflow-providers-apache-hive

