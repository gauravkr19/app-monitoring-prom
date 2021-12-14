FROM maven:3.6-openjdk-11-slim as builder

# Copy local code to the container image.
WORKDIR /app
COPY pom.xml ./
COPY config.yaml ./
COPY src ./src/

# Build a release artifact.
RUN mvn package -DskipTests

FROM adoptopenjdk/openjdk11:jdk-11.0.10_9-alpine-slim

#RUN wget -q https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.15.0/jmx_prometheus_javaagent-0.15.0.jar
#COPY --from=builder /app/config.yaml /config.yaml

# Copy the jar to the production image from the builder stage.
COPY --from=builder /app/target/helloworld-*.jar /helloworld.jar

# Run the web service on container startup.
CMD ["java", "-Dserver.port=${PORT}","-jar", "/helloworld.jar"]