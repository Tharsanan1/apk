# Use official Python runtime as base image
FROM python:3.9-slim

# Set working directory in container
WORKDIR /app

# Copy the Python script to the container
COPY server.py .

# Install Flask
RUN pip install --no-cache-dir flask

# Make port 5000 available to the world outside this container
EXPOSE 5000

# Define environment variable to ensure Python output is sent straight to terminal
ENV PYTHONUNBUFFERED=1

# Run server.py when the container launches
CMD ["python", "server.py"]