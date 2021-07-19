FROM python:alpine

RUN python3 -m venv /opt/venv

# Install dependencies:
COPY requirements.txt .
RUN /opt/venv/bin/pip install -r requirements.txt

# Run the application:
COPY server.py .
ENTRYPOINT ["/opt/venv/bin/python", "server.py"]