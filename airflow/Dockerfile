FROM apache/airflow:2.3.0
USER root
RUN apt-update && apt-get install -y gcc g++
USER airflow
RUN pip install apache-airflow-providers-cncf-kubernetes apache-airflow-providers-apache-spark apache-airflow-providers-apache-hive

