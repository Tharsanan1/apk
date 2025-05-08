# Use official Python image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Copy app files
COPY app.py .

# Install Flask
RUN pip install --no-cache-dir flask

# Expose the port Flask will run on
EXPOSE 8000

# Run the app
CMD ["python", "app.py"]
