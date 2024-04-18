FROM fedora:39
COPY ./counter /counter
RUN chmod +x /counter
ENTRYPOINT [ "/counter" ]