FROM scratch
COPY echo-server / 
ENTRYPOINT ["/echo-server"]
