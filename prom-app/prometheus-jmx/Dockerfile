# Use the official maven/Java 11 image to create a build artifact.
# https://hub.docker.com/_/maven
FROM maven:3.6-openjdk-11-slim as builder

# Copy local code to the container image.
WORKDIR /app
COPY pom.xml ./
COPY config.yaml ./
COPY src ./src/

# Build a release artifact.
RUN mvn package -DskipTests

# Use AdoptOpenJDK for base image.
# It's important to use OpenJDK 8u191 or above that has container support enabled.
# https://hub.docker.com/r/adoptopenjdk/openjdk11
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM adoptopenjdk/openjdk11:jdk-11.0.10_9-alpine-slim

RUN wget -q https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.15.0/jmx_prometheus_javaagent-0.15.0.jar

COPY --from=builder /app/config.yaml /config.yaml

# Copy the jar to the production image from the builder stage.
COPY --from=builder /app/target/helloworld-*.jar /helloworld.jar

# Run the web service on container startup.
CMD ["java", "-javaagent:./jmx_prometheus_javaagent-0.15.0.jar=9404:config.yaml", \
     "-Djava.security.egd=file:/dev/./urandom", \
     "-Dcom.sun.management.jmxremote.ssl=false", \ 
     "-Dcom.sun.management.jmxremote.authenticate=false", \
     "-Dcom.sun.management.jmxremote.port=5555", \
     "-Dserver.port=${PORT}","-jar", \
     "/helloworld.jar"]