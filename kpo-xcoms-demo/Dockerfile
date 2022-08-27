FROM python

WORKDIR /

# creating the file to write XComs to
RUN mkdir -p airflow/xcom
RUN echo "" > airflow/xcom/return.json

COPY multiply_by_23.py ./

CMD ["python", "./multiply_by_23.py"]