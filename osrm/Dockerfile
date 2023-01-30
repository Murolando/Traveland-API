FROM osrm/osrm-backend:latest

COPY ./data /data
WORKDIR /data
EXPOSE 5000
RUN osrm-extract -p "/opt/car.lua" "/data/north-caucasus-fed-district-latest.osm.pbf"
RUN osrm-contract "/data/north-caucasus-fed-district-latest.osrm"
CMD osrm-routed "/data/north-caucasus-fed-district-latest.osrm"



