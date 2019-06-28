FROM centos

WORKDIR /var/local/tools

# Copy the directory contents into the container at /var/local
COPY . /var/local/tools


EXPOSE 8080

# Define environment variable
ENV NAME TOOLS

CMD ["sh", "./bin/restart.sh"]