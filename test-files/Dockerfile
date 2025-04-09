# 🚩 Using latest tag (unpredictable builds)
FROM ubuntu:latest

# 🚩 Exposing a secret in an environment variable
ENV API_KEY=supersecretapikey

# 🚩 Running as root (no USER instruction)
# No USER line means container runs as root by default

# 🚩 No healthcheck defined
# No HEALTHCHECK instruction present

# 🚩 Installing unnecessary tools & missing apt-get upgrade
RUN apt-get update && apt-get install -y \
    curl \
    nano \
    net-tools

# 🚩 Using ADD instead of COPY
ADD . /app

# 🚩 No .dockerignore present (assumed)
# Sensitive files might get copied if .dockerignore is missing

# 🚩 No multi-stage builds
# This increases final image size unnecessarily

CMD ["bash"]
